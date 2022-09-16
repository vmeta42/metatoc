from ..canoser import Struct, Int64, Int8
from .move_resource import MoveResource, StructTag
from .event import EventHandle
from .account_address import AccountAddress


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

class ConfigResource(Struct):
    _fields = [
        ("consensus", bytes),
        ("validator_network_addresses", bytes),
        ("fullnode_network_addresses", bytes)
    ]
class ValidatorInfoResource(Struct):
    _fields = [
        ("addr", AccountAddress),
        ("consensus_voting_power", Int64),
        ("config", ConfigResource),
        ("last_config_update_time", Int64)
    ]
class DiemSystemResource(Struct):
    _fields = [
        ("scheme", Int8),
        ("validators", [ValidatorInfoResource])
    ]
class DiemConfigResource(Struct, MoveResource):
    MODULE_NAME = "DiemConfig"
    STRUCT_NAME = "DiemConfig"
    _fields = [
        ("payload", DiemSystemResource)
    ]