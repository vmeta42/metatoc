import os

from flask import Flask

def create_app(test_config = None):
    app = Flask(__name__, instance_relative_config=True);
    app.config.from_mapping(
        SECRET_KEY = "def",
        DATABASE = os.path.join(app.instance_path, "flaskr.sqlite")
    );

    if test_config is None:
        app.config.from_pyfile("config.py", silent=True);
    else:
        app.config.from_mapping(test_config);

    try:
        os.makedirs(app.instance_path);
    except OSError:
        pass;

    from . import db
    db.init_app(app);

    from . import violas_client_wrapper
    violas_client_wrapper.init_app(app);

    from . import wallet
    app.register_blueprint(wallet.bp);

    from . import block
    app.register_blueprint(block.bp);

    from . import chain
    app.register_blueprint(chain.bp);

    return app;
