# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from .base import Base
from .cursor import Cursor
from .types import type_mapping
from .int_type import Uint32
from .struct import TypedProperty
import json

#TODO: how to support discontinuous index in enum

class RustEnum(Base):
    _enums = []

    @classmethod
    def get_index(cls, name):
        for index, (ename, _) in enumerate(cls._enums):
            if ename == name:
                return index
        raise TypeError(f"name:{name} not in enum {cls}")

    @classmethod
    def new_with_index_value(cls, index, value):
        if not cls._enums:
            raise TypeError(f'{cls} has no _enums defined.')
        if index < 0 or index >= len(cls._enums):
            raise TypeError(f"index{index} out of bound:0-{len(cls._enums)-1}")
        _name, datatype = cls._enums[index]
        ret = cls.__new__(cls)
        ret._init_with_index_value(index, value, datatype)
        return ret

    def _init_with_index_value(self, index, value, datatype):
        self._index = index
        self.value_type = type_mapping(datatype)
        self.value = value

    def __init__(self, name, value=None):
        if not self.__class__._enums:
            raise TypeError(f'{self.__class__} has no _enums defined.')
        index = self.__class__.get_index(name)
        _name, datatype = self._enums[index]
        if name != _name:
            raise AssertionError(f"{name} != {_name}")
        self._init_with_index_value(index, value, datatype)

    #__getattr__ only gets called for attributes that don't actually exist.
    #If you set an attribute directly, referencing that attribute will retrieve it without calling __getattr__.
    #If you need to catch every attribute regardless whether it exists or not, use __getattribute__ instead.
    def __getattr__(self, name):
        if name == '_index':
            return None
        try:
            return self._index == self.__class__.get_index(name)
        except TypeError:
            return super().__getattr__(self, name)

    def __setattr__(self, name, value):
        if name == "value":
            TypedProperty.check_type(self.value_type, value)
            self.__dict__[name] = value
        elif name == "_index" or name == "value_type":
            self.__dict__[name] = value
        else:
            raise TypeError(f"{name} not allowed to modify in {self}.")

    @property
    def index(self):
        return self._index

    @property
    def enum_name(self):
        name, _ = self.__class__._enums[self._index]
        return name

    @classmethod
    def encode(cls, enum):
        ret = Uint32.serialize_uint32_as_uleb128(enum.index)
        if enum.value_type is not None:
            ret += enum.value_type.encode(enum.value)
        return ret

    @classmethod
    def decode(cls, cursor):
        index = Uint32.parse_uint32_from_uleb128(cursor)
        _name, datatype = cls._enums[index]
        if datatype is not None:
            value = type_mapping(datatype).decode(cursor)
            return cls.new_with_index_value(index, value)
        else:
            return cls.new_with_index_value(index, None)

    @classmethod
    def check_value(cls, value):
        if not isinstance(value, cls):
            raise TypeError('value {} is not {} type'.format(value, cls))

    def __eq__(self, other):
        if not isinstance(other, self.__class__):
            return False
        return self.index == other.index and self.value == other.value

    def to_json_serializable(self):
        if self.value_type == None:
            return self.enum_name
        jj = self.value_type.to_json_serializable(self.value)
        return {self.enum_name : jj}

    def __str__(self):
        return self.to_json(indent=2)

    def __repr__(self):
        return self.__class__.__qualname__ + self.to_json(indent=2)

    def to_json(self, sort_keys=False, indent=4):
        amap = self.to_json_serializable()
        return json.dumps(amap, sort_keys=sort_keys, indent=indent)



