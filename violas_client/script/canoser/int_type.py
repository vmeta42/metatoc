# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

import struct
from random import randint
from canoser.base import Base
from struct import pack, unpack
from_bytes = int.from_bytes

class IntType(Base):

    @classmethod
    def to_json_serializable(cls, value):
        return value

    @classmethod
    def encode(cls, value):
        return pack(cls.pack_str, value)

    @classmethod
    def encode_slow(cls, value):
        return value.to_bytes(cls.byte_lens, byteorder="little", signed=cls.signed)

    @classmethod
    def decode_bytes_slow(cls, bytes):
        return unpack(cls.pack_str, bytes)[0]

    @classmethod
    def decode_bytes(cls, bytes):
        return from_bytes(bytes, byteorder='little', signed=cls.signed)

    @classmethod
    def decode(cls, cursor):
        bytes = cursor.read_bytes(cls.byte_lens)
        return cls.decode_bytes(bytes)

    @classmethod
    def from_value(cls, value):
        return value

    @classmethod
    def from_proto(cls, proto):
        return proto

    @classmethod
    def int_unsafe(cls, s):
        ret = int(s)
        cls.check_value(ret)
        return ret

    @classmethod
    def int_safe(cls, s):
        """
        Only allow safe str and valid int to be coerced to destination IntType
        """
        if isinstance(s, bool):
            raise TypeError(f"{s} is not a integer")
        if isinstance(s, int):
            cls.check_value(s)
            return s
        if not isinstance(s, str):
            raise TypeError(f"{s} is not instance of <str>.")
        if len(s) < 1:
            raise TypeError(f"'{s}' is empty.")
        len_min = len(str(cls.min_value))
        len_max = len(str(cls.max_value))
        if len(s) > max(len_min, len_max):
            raise TypeError(f"Length of {s} is larger than max:{max(len_min, len_max)}.")
        ret = int(s)
        cls.check_value(ret)
        return ret


    @classmethod
    def check_value(cls, value):
        if value is None:
            return
        if isinstance(value, bool):
            raise TypeError(f"{value} is not a integer")
        if not isinstance(value, int):
            raise TypeError(f"{value} is not instance of <int>.")
        min, max = cls.min_value, cls.max_value
        if value < min or value > max:
            raise TypeError('value {} not in range {}-{}'.format(value, min, max))

    @classmethod
    def checked_add(cls, v1, v2):
        #rust style api
        cls.check_value(v1)
        cls.check_value(v2)
        try:
            ret = v1+v2
            cls.check_value(ret)
            return ret
        except TypeError:
            return None

    @classmethod
    def random(cls):
        return randint(cls.min_value, cls.max_value)

class Int8(IntType):
    pack_str = "<b"
    byte_lens = 1
    max_value = 127
    min_value = -128
    signed = True

class Int16(IntType):
    pack_str = "<h"
    byte_lens = 2
    max_value = 32767
    min_value = -32768
    signed = True

class Int32(IntType):
    pack_str = "<l"
    byte_lens = 4
    max_value = 2147483647
    min_value = -2147483648
    signed = True

class Int64(IntType):
    pack_str = "<q"
    byte_lens = 8
    max_value = 9223372036854775807
    min_value = -9223372036854775808
    signed = True


class Uint8(IntType):
    pack_str = "<B"
    byte_lens = 1
    max_value = 255
    min_value = 0
    signed = False

class Uint16(IntType):
    pack_str = "<H"
    byte_lens = 2
    max_value = 65535
    min_value = 0
    signed = False

class Uint32(IntType):
    pack_str = "<L"
    byte_lens = 4
    max_value = 4294967295
    min_value = 0
    signed = False

    @classmethod
    def serialize_uint32_as_uleb128(cls, value) -> bytes:
        ret = bytearray()
        while value >= 0x80:
            # Write 7 (lowest) bits of data and set the 8th bit to 1.
            byte = (value & 0x7f)
            ret.append(byte | 0x80)
            value >>= 7

        # Write the remaining bits of data and set the highest bit to 0.
        ret.append(value)
        return bytes(ret)

    @classmethod
    def parse_uint32_from_uleb128(cls, cursor):
        max_shift = 28
        value = 0
        shift = 0
        while not cursor.is_finished():
            byte = cursor.read_u8()
            val = byte & 0x7f
            value |= (val << shift)
            if val == byte:
                return value
            shift += 7
            if shift > max_shift:
                break
        raise ValueError(f"invalid ULEB128 representation for Uint32")


class Uint64(IntType):
    pack_str = "<Q"
    byte_lens = 8
    max_value = 18446744073709551615
    min_value = 0
    signed = False


class Int128(IntType):
    byte_lens = 16
    max_value = 170141183460469231731687303715884105727
    min_value = -170141183460469231731687303715884105728
    signed = True

    @classmethod
    def encode(cls, value):
        return value.to_bytes(16, byteorder="little", signed=True)



class Uint128(IntType):
    byte_lens = 16
    max_value = 340282366920938463463374607431768211455
    min_value = 0
    signed = False

    @classmethod
    def encode(cls, value):
        return value.to_bytes(16, byteorder="little", signed=False)

