import sys
import functools

from flask import (
    Blueprint, flash, g, redirect, render_template, request, url_for
)
from flask.json import jsonify

from .db import get_db
from .response import Response
from .ViolasClientWrapper import ViolasClient

bp = Blueprint("wallet", __name__);

@bp.route("/signup", methods = ["GET"])
def signup():
    resp = Response();

    if request.method == "GET":
        cli = ViolasClient();
        address, private_key = cli.GetNewAccount();
        resp.setData({"address":address, "private_key":private_key});


    return jsonify(str(resp))
