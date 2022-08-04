from ..canoser import Struct
from ..move_core_types.account_address import AccountAddress

class MintedTokenEvent(Struct):
    _fields = [
        ("token_id", bytes),
        ("path", bytes),
        ("minter", AccountAddress)
    ]


class ShareTokenEvent(Struct):
    _fields = [
        ("sender", AccountAddress),
        ("receiver", AccountAddress),
        ("token_id", bytes),
        ("metadata", bytes),
    ]
