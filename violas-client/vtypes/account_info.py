from canoser import Struct
from vtypes.move_resource import MoveResource, StructTag
from vtypes.event import EventHandle


class TokenResource(Struct):
    MODULE_NAME = "Meta42"
    STRUCT_NAME = "Token"

    _fields = [
        ("path", bytes),
    ]

class TokenResource(Struct):
    MODULE_NAME = "Meta42"
    STRUCT_NAME = "Token"

    _fields = [
        ("path", bytes),
    ]

class AccountInfoResource(Struct, MoveResource):
    MODULE_NAME = "Meta42"
    STRUCT_NAME = "AccountInfo"

    _fields = [
        ("tokens", [TokenResource]),
        ("minted_events", EventHandle),
        ("sent_events", EventHandle),
        ("received_events", EventHandle)
    ]


class GlobalInfoResource(Struct, MoveResource):
    MODULE_NAME = "Meta42"
    STRUCT_NAME = "GlobalInfo"

    _fields = [
        ("minted_events", EventHandle),
        ("shared_events", EventHandle)
    ]