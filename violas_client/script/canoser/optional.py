# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

from canoser.base import Base
from canoser.rust_optional import RustOptional

class Optional(RustOptional):
    _type = None

    @classmethod
    def from_type(cls, type):
        ret = cls
        ret._type = type
        return ret
