from canoser import Struct, Uint64, DelegateT, BytesT
from move_core_types.account_address import AccountAddress

class EventKey(DelegateT):
    LENGTH = AccountAddress.LENGTH + 8
    delegate_type = BytesT(LENGTH)

    @staticmethod
    def get_creator_address(key):
        return key[EventKey.LENGTH - AccountAddress.LENGTH:].hex()

    @classmethod
    def new_from_address(cls, addr, salt: int):
        addr = AccountAddress.normalize_to_bytes(addr)
        return salt.to_bytes(8, "little") + addr

class EventHandle(Struct):
    _fields = [
        ("count", Uint64),
        ("key", EventKey)
    ]

    def get_count(self):
        return self.count

    def get_key(self):
        return self.key.hex()

    def get_creator_address(self):
        return EventKey.get_creator_address(self.key)

