from ..canoser import DelegateT, BytesT
from ..canoser.util import int_list_to_hex

class AccountAddress(DelegateT):
    LENGTH = 16
    HEX_LENGTH = LENGTH * 2
    delegate_type = BytesT(LENGTH, encode_len=False)

    @classmethod
    def default(cls):
        return b"\x00" * cls.LENGTH

    @staticmethod
    def authentication_key(public_key):
        #TODO
        pass

    @staticmethod
    def from_public_key(public_key):
        pass

    @classmethod
    def normalize_to_bytes(cls, address):
        if isinstance(address, list):
            address = int_list_to_hex(address)
        if isinstance(address, str):
            return cls.from_hex(address)
        if isinstance(address, bytes):
            if len(address) != cls.LENGTH:
                raise ValueError(f"{address} is not a valid address.")
            return address
        raise TypeError(f"Address: {address} has unknown type.")

    @classmethod
    def from_hex(cls, address: str):
        if address.startswith("0x") or address.startswith("0X"):
            address = address[2:]
        if len(address) < cls.HEX_LENGTH:
            address =address.rjust(cls.HEX_LENGTH, "0")
        if len(address) == cls.HEX_LENGTH:
            return bytes.fromhex(address)
        raise ValueError(f"{address} is not a valid address")
