from sys import path

from diem.utils import account_address
from violas_client.client import Client
from violas_client.vtypes.contants import server_url
from violas_client.vtypes.local_account import LocalAccount

from flask import current_app, g

SERVER="http://124.251.110.220:50001";

class WalletClient():
    def __init__(self):
        self.cli = Client(server_url=SERVER);
        try:
            with current_app.open_resource("instance/wallet.key", "r") as f:
                self.p_vasp = LocalAccount.from_private_key_hex(f.read());
        except FileNotFoundError:
            self.p_vasp = LocalAccount.generate();
            self.cli.create_parent_vasp(self.p_vasp.account_address, self.p_vasp.auth_key.prefix(), "VLS", b"metatoc");
            self.cli.meta42_accept(self.p_vasp);

            with open("./instance/wallet.key", "a") as f:
                f.write(self.p_vasp.public_key_bytes.hex());

        print(self.p_vasp.account_address, self.p_vasp.public_key_bytes.hex());

    def GetNewAccount(self):
        account = LocalAccount.generate();
        self.cli.create_child_vasp(self.p_vasp, account.account_address, account.auth_key.prefix());
        self.cli.meta42_accept(account);

        return account.account_address, account.public_key_bytes.hex();

    def MintNewToken(self, private_key, token):
        account = LocalAccount.from_private_key_hex(private_key);
        self.cli.meta42_mint_token(account, token);

    def ShareToken(self, from_private_key, to_address, token):
        from_account = LocalAccount.from_private_key_hex(from_private_key);
        self.cli.meta42_share_token_by_path(from_account, to_address, token);

    def GetTokens(self, address, token):
        paths = self.cli.get_paths(address);

        if token in paths:
            return True;
        else:
            return False;
