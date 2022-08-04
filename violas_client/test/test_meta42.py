from ..client import Client
from ..vtypes.contants import server_url
from ..vtypes.local_account import LocalAccount
from ..vtypes.contants import Meta42CodeType

cli = Client(server_url=server_url)
p_vasp = LocalAccount.from_private_key_hex("346de128de4a6b69bd281ffbd19157fe19272a2d8608ef64708026580aeab11a")

def test_init():
    cli.allow_publishing_module(True)
    cli.publish_compare()
    cli.publish_meta42()
    cli.allow_custom_script()
    cli.meta42_initialize(p_vasp)
    assert_type(p_vasp.account_address, Meta42CodeType.META42_INITIALIZE)

def test_mint_token():
    child_account = LocalAccount.generate()
    cli.create_child_vasp(p_vasp, child_account.account_address, child_account.auth_key.prefix())
    path = "/" + child_account.account_address.to_hex()
    cli.meta42_mint_token(child_account,path)
    assert_type(child_account.account_address, Meta42CodeType.META42_MINT_TOKEN)
    paths = cli.get_paths(child_account.account_address)
    assert paths[0] == path

def test_share_token():
    src_account = LocalAccount.generate()
    cli.create_child_vasp(p_vasp, src_account.account_address, src_account.auth_key.prefix())
    path = "/" + src_account.account_address.to_hex()
    cli.meta42_mint_token(src_account,path)
    assert_type(src_account.account_address, Meta42CodeType.META42_MINT_TOKEN)
    shared_account = LocalAccount.generate()
    cli.create_child_vasp(p_vasp, shared_account.account_address, shared_account.auth_key.prefix())
    cli.meta42_accept(shared_account)
    assert_type(shared_account.account_address, Meta42CodeType.META42_ACCEPT)
    cli.meta42_share_token_by_path(src_account, shared_account.account_address, path)
    assert_type(src_account.account_address, Meta42CodeType.META42_SHARE_TOKEN_BY_ID)
    paths = cli.get_paths(shared_account.account_address)
    assert paths[0] == path


def assert_type(account_address, type):
    seq = cli.get_account_sequence(account_address)
    tx = cli.get_account_transaction(account_address, seq-1)
    assert tx.get_tx_type() == type
