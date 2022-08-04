from sys import path

from diem.utils import account_address, private_key_bytes
from violas_client.client import Client
from violas_client.vtypes.local_account import LocalAccount

from flask import current_app, g

SERVER = "http://124.251.110.220:50001";
ROOT_PRIVATE_KEY = "7b5a50c2caec5921b7b268cc08c2ce754921e3572aa2f68eddd5824888958b5b";
VASP_PRIVATE_KEY = "346de128de4a6b69bd281ffbd19157fe19272a2d8608ef64708026580aeab11a";

class WalletClient():
    def __init__(self):
        self.cli = Client(server_url=SERVER, root_private=ROOT_PRIVATE_KEY);
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
