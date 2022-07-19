from .base import Base
from .rust_optional import RustOptional

class Optional(RustOptional):
    _type = None

    @classmethod
    def from_type(cls, type):
        ret = cls
        ret._type = type
        return ret
