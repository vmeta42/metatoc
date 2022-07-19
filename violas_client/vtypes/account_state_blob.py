from ..canoser import Struct

class AccountStateBlobView(Struct):
    _fields = [
        ("blob", bytes)
    ]
