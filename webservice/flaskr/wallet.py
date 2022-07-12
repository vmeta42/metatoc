import sys
import functools

from flask import (
    Blueprint, flash, g, redirect, render_template, request, url_for
)
from flask.json import jsonify

from flaskr.db import get_db
from flaskr.response import Response

bp = Blueprint("wallet", __name__)

@bp.route("/signup", methods = ["GET"])
def signup():
    resp = Response();

    if request.method == "GET":
        resp.setData({"address":"wallet.address", "private_key":"wallet.private.key"});


    return jsonify(str(resp))
