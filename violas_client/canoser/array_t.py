# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from .base import Base
from .int_type import Uint32, Uint8
import struct


class ArrayT(Base):

    def __init__(self, atype, fixed_len=None, encode_len=True):
        self.atype = atype
        if fixed_len is not None and fixed_len <= 0:
            raise TypeError("arr len must > 0".format(fixed_len))
        if fixed_len is None and not encode_len:
            raise TypeError("variable length sequences must encode len.")
        self.fixed_len = fixed_len
        self.encode_len = encode_len

    def encode(self, arr):
        if self.fixed_len is not None and len(arr) != self.fixed_len:
             raise TypeError(f"{len(arr)} is not equal to predefined value: {self.fixed_len}")
        output = b""
        if self.encode_len:
            output += Uint32.serialize_uint32_as_uleb128(len(arr))
        for item in arr:
            output += self.atype.encode(item)
        return output

    def decode(self, cursor):
        arr = []
        if self.encode_len:
            size = Uint32.parse_uint32_from_uleb128(cursor)
            if self.fixed_len is not None and size != self.fixed_len:
                 raise TypeError(f"{size} is not equal to predefined value: {self.fixed_len}")
        else:
            size = self.fixed_len
        for _ in range(size):
            arr.append(self.atype.decode(cursor))
        return arr

    def from_value(self, value):
        if value is None:
            return []
        return [self.atype.from_value(v) for v in value]

    def from_proto(self, proto):
        return [self.atype.from_proto(v) for v in proto]

    def check_value(self, arr):
        if self.fixed_len is not None and len(arr) != self.fixed_len:
            raise TypeError("arr len not match: {}-{}".format(len(arr), self.fixed_len))
        if not isinstance(arr, list):
            raise TypeError(f"{arr} is not a list.")
        for item in arr:
            self.atype.check_value(item)

    def __eq__(self, other):
        if not isinstance(other, ArrayT):
            return False
        return self.atype == other.atype and self.fixed_len == other.fixed_len

    def to_json_serializable(cls, obj):
        if cls.atype == Uint8:
            return struct.pack("<{}B".format(len(obj)), *obj).hex()
        ret = []
        for _, item in enumerate(obj):
            data = cls.atype.to_json_serializable(item)
            ret.append(data)
        return ret

