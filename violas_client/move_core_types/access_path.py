from ..canoser import Struct, Uint64, RustEnum, DelegateT
from .account_address import AccountAddress
from .identifier import Identifier

class Field(Identifier):
    pass

class Access(RustEnum):
    _enums = [
        ("Field", Field),
        ("Index", Uint64)
    ]

SEPARATOR = '/'
class Accesses(DelegateT):
    delegate_type = [Access]

    @classmethod
    def empty(cls):
        return []

    @classmethod
    def as_separated_string(cls, accesses):
        path = ""
        for access in accesses:
            if access.Field:
                path += access.value
            elif accesses.Index:
                path += str(access.value)
            else:
                raise AssertionError("Unreachable")
            path += SEPARATOR
        return path

class AccessPath(Struct):
    CODE_TAG = 0
    RESOURCE_TAG = 1

    _fields = [
        ("address", AccountAddress),
        ("path", bytes)
    ]

    @classmethod
    def new(cls, address, path):
        ret = cls()
        ret.address = address
        ret.path = path
        return ret

    @classmethod
    def new_for_account(cls):
        pass

    @classmethod
    def new_for_sent_event(cls, address):
        pass

    @classmethod
    def new_for_received_event(cls, address):
        pass

    @classmethod
    def resource_access_vec(cls, tag, accesses):
        key = cls.RESOURCE_TAG.to_bytes(1, "little")
        key += tag.hash()
        key += str.encode(Accesses.as_separated_string(accesses))
        return key

    @classmethod
    def resource_access_path(cls, key, accesses):
        path = AccessPath.resource_access_vec(key.type_, accesses)
        return AccessPath(key.address, path)

    @classmethod
    def code_access_path_vec(cls, key):
        root = cls.CODE_TAG.to_bytes(1, "little")
        root += key.hash()
        return root



