# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from .base import Base
from .cursor import Cursor
from .types import type_mapping
import json

class TypedProperty:
    def __init__(self, name, expected_type):
        self.name = name
        self.expected_type = expected_type

    def __set__(self, instance, value):
        TypedProperty.check_type(self.expected_type, value)
        instance.__dict__[self.name] = value

    @staticmethod
    def check_type(datatype, value):
        if datatype is None:
            if value is not None:
                raise TypeError(f'{datatype} mismatch {value}')
            else:
                return
        check = getattr(datatype, "check_value", None)
        if callable(check):
            check(value)
        else:
            raise TypeError('{} has no check_value method'.format(datatype))

class Struct(Base):
    _fields = []
    _initialized = False

    @classmethod
    def initailize_fields_type(cls):
        if not cls._initialized:
            cls._initialized = True
            for name, atype in cls._fields:
                setattr(cls, name, TypedProperty(name, type_mapping(atype)))

    def __init__(self, *args, **kwargs):
        self.__class__.initailize_fields_type()

        if len(args) > len(self._fields):
            raise TypeError('Expected {} arguments'.format(len(self._fields)))

        # Set all of the positional arguments
        for (name, _type), value in zip(self._fields, args):
            typed = getattr(self, name)
            typed.__set__(self, value)

        # Set the remaining keyword arguments
        for name, _type in self._fields[len(args):]:
            if name in kwargs:
                typed = getattr(self, name)
                typed.__set__(self, kwargs.pop(name))

        # Check for any remaining unknown arguments
        if kwargs:
            raise TypeError('Invalid argument(s): {}'.format(','.join(kwargs)))


    @classmethod
    def encode(cls, obj):
        output = b''
        for name, atype in obj._fields:
            value = getattr(obj, name)
            output += type_mapping(atype).encode(value)
        return output

    @classmethod
    def decode(cls, cursor):
        ret = cls.__new__(cls)
        ret.__init__()
        for name, atype in ret._fields:
            prop = getattr(ret, name)
            mtype = type_mapping(atype)
            assert mtype == prop.expected_type
            value = mtype.decode(cursor)
            prop.__set__(ret, value)
        return ret

    @classmethod
    def from_value(cls, value):
        ret = cls.__new__(cls)
        ret.__init__()
        for name, atype in ret._fields:
            prop = getattr(ret, name)
            mtype = type_mapping(atype)
            assert mtype == prop.expected_type
            v = mtype.from_value(value.get(name))
            prop.__set__(ret, v)
        if len(value) > len(ret.__dict__):
            if not (value.get("type") is not None and len(value) == len(ret.__dict__)+1):
                raise TypeError(f"Type mismatched {len(value)} != {len(ret.__dict__)}\n"
                                f"source:{value},\n "
                                f"dest:{type(ret), ret}")
        return ret

    @classmethod
    def from_proto(cls, proto):
        ret = cls.__new__(cls)
        ret.__init__()
        for name, atype in ret._fields:
            prop = getattr(ret, name)
            mtype = type_mapping(atype)
            assert mtype == prop.expected_type
            v = mtype.from_proto(getattr(proto, name))
            prop.__set__(ret, v)
        return ret


    @classmethod
    def check_value(cls, value):
        if value is None:
            return
        if not isinstance(value, cls) and value is not None:
            raise TypeError('value {} is not {} type'.format(value, cls))

    def __eq__(self, other):
        if type(self) != type(other):
            return False
        for name, atype in self._fields:
            v1 = getattr(self, name)
            v2 = getattr(other, name)
            if v1 != v2:
                return False
        return True

    def to_json_serializable(self):
        amap = {}
        for name, atype in self._fields:
            value = getattr(self, name)
            if isinstance(value, TypedProperty):
                amap[name] = None
            else:
                atype = type_mapping(atype)
                amap[name] = atype.to_json_serializable(value)
        return amap

    def __str__(self):
        return self.to_json(indent=2)

    def __repr__(self):
        return self.__class__.__qualname__ + self.to_json(indent=2)

    def to_json(self, sort_keys=False, indent=4):
        amap = self.to_json_serializable()
        return json.dumps(amap, sort_keys=sort_keys, indent=indent)

