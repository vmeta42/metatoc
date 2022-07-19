from .cursor import Cursor
from io import StringIO
import json

class Base:
    """
    All types should implment following four methods:

    def encode(cls_or_obj, value)

    def decode(cls_or_obj, cursor)

    def check_value(cls_or_obj, value)

    def to_json_serializable(obj)
    def to_json_serializable(acls, value)
    """


    def serialize(self):
        return self.__class__.encode(self)

    @classmethod
    def deserialize(cls, buffer, check=True):
        cursor = Cursor(buffer)
        ret = cls.decode(cursor)
        if not cursor.is_finished() and check:
            raise IOError("bytes not all consumed:{}, {}".format(
                len(buffer), cursor.offset))
        return ret
    
        

