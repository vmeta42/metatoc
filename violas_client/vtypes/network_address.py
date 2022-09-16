from ..canoser import RustEnum, Struct, Int16, Int32, Int8


class Protocol(RustEnum):
    _enums = [
        ("Ip4", Int32),
        ("Ip6",bytes),
        ("Dns",bytes),
        ("Dns4",bytes),
        ("Dns6",bytes),
        ("Tcp", Int16),
        ("Memory",bytes),
        ("NoiseIK",bytes),
        ("Hanshake",Int8)
    ]


class NetworkAddress(Struct):
    _fields = [
        ("Protocol", [Protocol])
    ]


class NetworkAddresses(Struct):
    _fields = [
        ("NetworkAddress", [bytes])
    ]