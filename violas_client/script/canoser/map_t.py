# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from canoser.base import Base
from canoser.int_type import Uint32

class MapT(Base):

    def __init__(self, ktype, vtype):
        self.ktype = ktype
        self.vtype = vtype

    def encode(self, kvs):
        output = b""
        output += Uint32.serialize_uint32_as_uleb128(len(kvs))
        odict = {}
        for k, v in kvs.items():
            odict[self.ktype.encode(k)] = self.vtype.encode(v)
        for name in sorted(odict.keys()):
            output += name
            output += odict[name]
        return output

    def decode(self, cursor):
        kvs = {}
        size = Uint32.parse_uint32_from_uleb128(cursor)
        for _ in range(size):
            k = self.ktype.decode(cursor)
            v = self.vtype.decode(cursor)
            if isinstance(k, list) and isinstance(k[0], int):
                #python doesn't support list as key in dict, so we change list to bytes
                kvs[bytes(k)] = v
            else:
                kvs[k] = v
        #TODO: check the key order of kvs, because lcs has order when serialize map.
        return kvs

    def check_value(self, kvs):
        if not isinstance(kvs, dict):
            raise TypeError(f"{kvs} is not a dict.")
        for k, v in kvs.items():
            if isinstance(self.ktype, list) or \
                (hasattr(self.ktype, 'delegate_type') and isinstance(self.ktype.delegate_type, list)):
                from canoser.types import BytesT
                BytesT().check_value(k)
            else:
                self.ktype.check_value(k)
            self.vtype.check_value(v)

    def __eq__(self, other):
        if not isinstance(other, MapT):
            return False
        return self.ktype == other.ktype and self.vtype == other.vtype

    def to_json_serializable(cls, obj):
        amap = {}
        for k,v in obj.items():
            kk = cls.ktype.to_json_serializable(k)
            vv = cls.vtype.to_json_serializable(v)
            amap[kk] = vv
        return amap


