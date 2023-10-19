# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

import struct
from .int_type import Uint32
from .base import Base

class BytesT(Base):

    def __init__(self, fixed_len=None, encode_len=True):
        if fixed_len is not None and fixed_len <= 0:
            raise TypeError("byte len must > 0".format(fixed_len))
        if fixed_len is None and not encode_len:
            raise TypeError("variable length sequences must encode len.")
        self.fixed_len = fixed_len
        self.encode_len = encode_len


    def encode(self, value):
        output = b""
        if self.encode_len:
            output += Uint32.serialize_uint32_as_uleb128(len(value))
        output += value
        return output


    def decode(self, cursor):
        if self.encode_len:
            size = Uint32.parse_uint32_from_uleb128(cursor)
            if self.fixed_len is not None and size != self.fixed_len:
                 raise TypeError(f"{size} is not equal to predefined value: {self.fixed_len}")
        else:
            size = self.fixed_len
        return cursor.read_bytes(size)

    def check_value(self, value):
        if not isinstance(value, bytes):
            raise TypeError('value {} is not bytes'.format(value))
        if self.fixed_len is not None and len(value) != self.fixed_len:
            raise TypeError("len not match: {}-{}".format(len(value), self.fixed_len))


    def __eq__(self, other):
        if not isinstance(other, BytesT):
            return False
        return self.fixed_len == other.fixed_len and self.encode_len == other.encode_len


    def to_json_serializable(cls, obj):
        return obj.hex()


class ByteArrayT(Base):

    def encode(self, value):
        output = b""
        output += Uint32.serialize_uint32_as_uleb128(len(value))
        output += bytes(value)
        return output


    def decode(self, cursor):
        size = Uint32.parse_uint32_from_uleb128(cursor)
        return bytearray(cursor.read_bytes(size))

    def check_value(self, value):
        if not isinstance(value, bytearray):
            raise TypeError('value {} is not bytearray'.format(value))


    def __eq__(self, other):
        if not isinstance(other, ByteArrayT):
            return False
        return True

    def to_json_serializable(cls, obj):
        return obj.hex()

