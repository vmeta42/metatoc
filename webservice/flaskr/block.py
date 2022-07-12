import string

from flask import (
    Blueprint, flash, g, redirect, request, url_for, jsonify
)

from flaskr.db import get_db
from flaskr.response import Response

bp = Blueprint("block", __name__);

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

        db.execute(
            "INSERT INTO transactions (transaction_type, from_address, token_name) "
            "VALUES (?, ?, ?)",
            ("CREATE", address, path)
        );
        db.commit();

    elif request.method == "GET":
        address = request.args.get("address");
        offset = request.args.get("offset", 0, type=int);
        limit = request.args.get("limit", 10, type=int);

        db = get_db();
        if address is None:
            tokens = db.execute(
                "SELECT token_name FROM transactions "
                "WHERE transaction_type='CREATE' "
                "LIMIT ? OFFSET ?",
                (limit, offset)
            ).fetchall();
        else:
            tokens = db.execute(
                "SELECT token_name FROM transactions "
                "WHERE transaction_type='CREATE' "
                "AND (from_address=? OR to_address=?) "
                "LIMIT ? OFFSET ?",
                (address, address, limit, offset)
            ).fetchall();

        paths = [];
        for token in tokens:
            paths.append(token["token_name"]);

        resp.setData({"paths": paths});

    elif request.method == "PUT":
        db = get_db();
        params = request.get_json();
        private_key = params["private_key"];
        from_address = params["from_address"];
        to_address = params["to_address"];
        token_name = params["token_name"];

        db.execute(
            "INSERT INTO transactions (transaction_type, from_address, to_address, token_name) "
            "VALUES (?, ?, ?, ?)",
            ("TRANSFER", from_address, to_address, token_name)
        );
        db.commit();

    return jsonify(str(resp));

@bp.route("/paths/<path:hdfs_path>", methods = ["GET"])
def GetPathByAddress(hdfs_path):
    resp = Response();
    address = request.args.get("address");

    db = get_db();

    token = db.execute(
        "SELECT * "
        "FROM transactions "
        "WHERE (from_address=? or to_address=?) "
        "AND token_name=?",
        (address, address, hdfs_path)
    ).fetchall();

    if token is None:
        resp.setData({"data_path": ""});
    else:
        resp.setData({"data_path": "http://hdfs/file/url"});

    return jsonify(str(resp));
