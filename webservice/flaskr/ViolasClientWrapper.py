from os import getenv
from sys import path
from time import sleep

from diem.jsonrpc import NetworkError
from diem.utils import account_address, private_key_bytes
from flask.cli import with_appcontext
from violas_client.client import Client
from violas_client.vtypes.local_account import LocalAccount

from flask import current_app, g
import click

SERVER = getenv("CHAIN_ENV", "http://validator:8080");
VASP_PRIVATE_KEY = "346de128de4a6b69bd281ffbd19157fe19272a2d8608ef64708026580aeab11a";

class ViolasClient():
    def __init__(self):
        self.root_private_key = self.__GetPrivateKey();
        self.cli = Client(server_url=SERVER, root_private=self.root_private_key);
        self.p_vasp = LocalAccount.from_private_key_hex(VASP_PRIVATE_KEY);

    def GetNewAccount(self):
        account = LocalAccount.generate();
        self.cli.create_child_vasp(self.p_vasp, account.account_address, account.auth_key.prefix());
        self.cli.meta42_accept(account);

        return account.account_address.to_hex(), private_key_bytes(account.private_key).hex();

    def MintNewToken(self, private_key, address, token):
        account = LocalAccount.from_private_key_hex(private_key);

        self.cli.meta42_mint_token(account, token);

    def ShareToken(self, from_private_key, to_address, token):
        from_account = LocalAccount.from_private_key_hex(from_private_key);
        self.cli.meta42_share_token_by_path(from_account, account_address(to_address), token);

    def GetTokens(self, address):
        return self.cli.get_paths(address)

    def HasToken(self, address, token):
        paths = self.cli.get_paths(address);

        if token in paths:
            return True;
        else:
            return False;

    def GetAddressOfAccount(self, private_key):
        account = LocalAccount.from_private_key_hex(private_key);
        return account.account_address.to_hex();

    def InitializeChain(self):
        self.cli.allow_publishing_module(True);
        self.cli.publish_compare();
        self.cli.publish_meta42();
        self.cli.allow_custom_script();
        self.cli.create_parent_vasp(self.p_vasp.account_address, self.p_vasp.auth_key.prefix(), "VLS", b"metatoc");
        self.cli.meta42_initialize(self.p_vasp);

    def __GetPrivateKey(self):
        with open("./diem-share/mint.key", "rb") as f:
            key = f.read();
        return key[1:33].hex();

@click.command("init-chain")
@with_appcontext
def init_chain_command():
    while(True):
        try:
            cli = ViolasClient();
            cli.InitializeChain();
            click.echo("Initialized the Violas Chain.");
            break;
        except NetworkError as e:
            print(e);
            print("Wait 5 seconds for Violas start.....");
            sleep(5);
            continue

def init_app(app):
    app.cli.add_command(init_chain_command);
