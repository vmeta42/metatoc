# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from canoser.int_type import Uint8
from canoser.tuple_t import TupleT
from canoser.map_t import MapT
from canoser.str_t import StrT
from canoser.bytes_t import BytesT, ByteArrayT
from canoser.bool_t import BoolT
from canoser.array_t import ArrayT


def my_import(name):
    components = name.split('.')
    mod = __import__(components[0])
    for comp in components[1:]:
        mod = getattr(mod, comp)
    return mod

def type_mapping(field_type):
    """
    Mapping python types to canoser types
    """
    if field_type == str:
        return StrT
    elif field_type == bytes:
        return BytesT()
    elif field_type == bytearray:
        return ByteArrayT()
    elif field_type == bool:
        return BoolT
    elif type(field_type) == list:
        if len(field_type) == 0:
            return ArrayT(Uint8)
        elif len(field_type) == 1:
            item = field_type[0]
            return ArrayT(type_mapping(item))
        elif len(field_type) == 2:
            item = field_type[0]
            size = field_type[1]
            return ArrayT(type_mapping(item), size)
        elif len(field_type) == 3:
            item = field_type[0]
            size = field_type[1]
            encode_len = field_type[2]
            return ArrayT(type_mapping(item), size, encode_len)
        else:
            raise TypeError("Array has one item type, no more.")
        raise AssertionError("unreacheable")
    elif type(field_type) == dict:
        if len(field_type) == 0:
            ktype = BytesT()
            vtype = [Uint8]
        elif len(field_type) == 1:
            ktype = next(iter(field_type.keys()))
            vtype = next(iter(field_type.values()))
        else:
            raise TypeError("Map type has one item mapping key type to value type.")
        return MapT(type_mapping(ktype), type_mapping(vtype))
    elif type(field_type) == tuple:
        arr = []
        for item in field_type:
            arr.append(type_mapping(item))
        return TupleT(*arr)
    elif type(field_type) == str:
        return my_import(field_type)
    else:
        return field_type
