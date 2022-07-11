import sys
import functools

from flask import (
    Blueprint, flash, g, redirect, render_template, request, url_for
)

from flaskr.db import get_db

bp = Blueprint("wallet", __name__)

@bp.route("/signup", methods = ["GET"])
def signup():
    if request.method == "GET":
        return;

    return;
