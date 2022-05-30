from canoser.types import type_mapping
from canoser.base import Base


class DelegateT(Base):
    delegate_type = 'delegate'

    @classmethod
    def dtype(cls):
        return type_mapping(cls.delegate_type)

    @classmethod
    def encode(cls, value):
        return cls.dtype().encode(value)

    @classmethod
    def decode(cls, cursor):
        return cls.dtype().decode(cursor)

    @classmethod
    def from_value(cls, value):
        return cls.dtype().from_value(value)

    @classmethod
    def check_value(cls, value):
        cls.dtype().check_value(value)

    @classmethod
    def to_json_serializable(cls, value):
        #TODO: bad smell
        #if hasattr(cls, "to_json_serializable"):
        if 'to_json_serializable' in cls.__dict__.keys():
            return cls.to_json_serializable(value)
        return cls.dtype().to_json_serializable(value)
