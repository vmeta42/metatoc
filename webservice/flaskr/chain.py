# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from flask import (
    Blueprint, flash, g, request
)
from flask.json import jsonify

from .violas_client_wrapper import ViolasClient
from .response import Response

bp = Blueprint("chain", __name__);

@bp.route("/nodes", methods = ["GET"])
def GetNodeInfo():
    resp = Response();

    if request.method == "GET":
        cli = ViolasClient();
        infos = cli.GetNodeInfo();
        for i in infos:
            i.pop("address");

        resp.setData(infos);

    return jsonify(str(resp));
