import string
import os
import uuid

from flask import (
    Blueprint, flash, g, redirect, request, url_for, jsonify
)

from .db import get_db
from .response import Response
from .violas_client_wrap import ViolasClient
from .ResultCode import ResultCode

bp = Blueprint("block", __name__);

FILE_PATH = "data";
FILE_DOWNLOAD_URL = "https://"

@bp.route("/paths", methods = ["POST", "GET", "PUT"])
def paths():
    resp = Response();

    if request.method == "POST":
        db = get_db();
        params = request.get_json();
        private_key = params["private_key"];
        address = params["address"];
        path = params["path"];
        content = params["content"];

        token = db.execute(
            "SELECT * FROM transactions "
            "WHERE token_name=?",
            (path,)
        ).fetchone();

        if token is not None:
            resp.setStatus(ResultCode.DATA_ALERADY_EXISTED);
            return jsonify(str(resp));

        cli = ViolasClient();
        cli.MintNewToken(private_key, address, path);

        fileName = uuid.uuid4().hex;
        db.execute(
            "INSERT INTO transactions (transaction_type, from_address, token_name, file_name) "
            "VALUES (?, ?, ?, ?)",
            ("CREATE", address, path, fileName)
        );
        db.commit();

        dataPath = os.path.join(os.path.abspath(os.path.curdir), FILE_PATH);
        if not os.path.exists(dataPath):
            os.makedirs(dataPath);

        with open(os.path.join(dataPath, fileName), mode = 'x') as f:
            f.write(content)

    elif request.method == "GET":
        address = request.args.get("address");
        offset = request.args.get("offset", 0, type=int);
        limit = request.args.get("limit", 10, type=int);

        db = get_db();
        paths = [];
        if address is None:
            tokens = db.execute(
                "SELECT token_name FROM transactions "
                "WHERE transaction_type='CREATE' "
                "LIMIT ? OFFSET ?",
                (limit, offset)
            ).fetchall();

            for token in tokens:
                paths.append(token["token_name"]);
        else:
            # tokens = db.execute(
            #     "SELECT token_name FROM transactions "
            #     "WHERE from_address=? OR to_address=? "
            #     "LIMIT ? OFFSET ?",
            #     (address, address, limit, offset)
            # ).fetchall();
            cli = ViolasClient();
            tokens = cli.GetTokens(address);

            for token in tokens:
                paths.append(token);

        resp.setData({"paths": paths});

    elif request.method == "PUT":
        db = get_db();
        params = request.get_json();
        private_key = params["private_key"];
        from_address = params["from_address"];
        to_address = params["to_address"];
        token_name = params["token_name"];

        token = db.execute(
            "SELECT * FROM transactions "
            "WHERE token_name=? "
            "AND (from_address=? OR to_address=?)",
            (token_name, from_address, from_address)
        ).fetchone();

        if token is None:
            resp.setStatus(ResultCode.RESULT_DATA_NONE);
        else:
            cli = ViolasClient();
            cli.ShareToken(private_key, to_address, token_name);

            db.execute(
                "INSERT INTO transactions (transaction_type, from_address, to_address, token_name, file_name) "
                "VALUES (?, ?, ?, ?, ?)",
                ("TRANSFER", from_address, to_address, token_name, token["file_name"])
            );
            db.commit();

    return jsonify(str(resp));

@bp.route("/paths/<path:hdfs_path>", methods = ["GET"])
def GetPathByAddress(hdfs_path):
    resp = Response();

    hdfs_path = "/" + hdfs_path;
    params = request.get_json();
    private_key = params["private_key"];

    cli = ViolasClient();
    address = cli.GetAddressOfAccount(private_key);

    if not cli.HasToken(address, hdfs_path):
        resp.setStatus(ResultCode.PERMISSION_NO_ACCESS);

        return jsonify(str(resp));

    db = get_db();

    token = db.execute(
        "SELECT * "
        "FROM transactions "
        "WHERE (from_address=? or to_address=?) "
        "AND token_name=?",
        (address, address, hdfs_path)
    ).fetchone();

    if token is None:
        resp.setStatus(ResultCode.RESULT_DATA_NONE);
    else:
        dataPath = os.path.join(os.path.abspath(os.path.curdir), FILE_PATH);
        with open(os.path.join(dataPath, token["file_name"]), mode='r') as f:
            data = f.read();

        resp.setData({"data": data});

    return jsonify(str(resp));
