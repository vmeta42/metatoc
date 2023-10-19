# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from .base import Base
from .types import type_mapping
from .bool_t import BoolT
from .struct import TypedProperty
from . import Struct


class RustOptional(Base):
    _type = None

    def __init__(self, value=None):
        if not self.__class__._type:
            raise TypeError(f'{self.__class__} has no _type defined.')
        self.__dict__["value_type"] = type_mapping(self.__class__._type)
        self.value = value

    def __setattr__(self, name, value):
        if name == "value":
            if value is not None:
                TypedProperty.check_type(self.value_type, value)
            self.__dict__[name] = value
        else:
            raise TypeError(f"{name} not allowed to modify in {self}.")


    @classmethod
    def encode(cls, optional):
        if optional.value is not None:
            ret = BoolT.encode(True)
            ret += optional.value_type.encode(optional.value)
            return ret
        else:
            return BoolT.encode(False)

    @classmethod
    def decode(cls, cursor):
        exist = BoolT.decode(cursor)
        if exist:
            value = cls._type.decode(cursor)
            return cls(value)
        else:
            return cls()

    @classmethod
    def check_value(cls, value):
        if not isinstance(value, cls):
            raise TypeError('value {} is not {} type'.format(value, cls))


    def __eq__(self, other):
        if not isinstance(other, self.__class__):
            return False
        return self.value == other.value

    def to_json_serializable(self):
        if self.value is None:
            return None
        return self.value_type.to_json_serializable(self.value)
