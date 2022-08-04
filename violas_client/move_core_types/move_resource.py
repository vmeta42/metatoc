from .language_storage import StructTag
from .access_path import AccessPath, Accesses
from .account_address import AccountAddress

class MoveResource():
    MODULE_NAME: str
    STRUCT_NAME: str

    @classmethod
    def module_identifier(cls):
        return cls.MODULE_NAME

    @classmethod
    def struct_identifier(cls):
        return cls.STRUCT_NAME

    @classmethod
    def type_params(cls):
        return []

    @classmethod
    def struct_tag(cls, *type_params, module_address=None):
        module_address = AccountAddress.normalize_to_bytes(module_address)
        return StructTag(
            module_address,
            cls.module_identifier(),
            cls.struct_identifier(),
            list(type_params)
        )

    @classmethod
    def resource_path(cls, module_address=None):
        return cls.struct_tag(module_address=module_address).access_vector()

    @classmethod
    def resource_path_for(cls, *type_params, module_address=None):
        return AccessPath.resource_access_vec(cls.struct_tag(*type_params, module_address=module_address), Accesses.empty())

    @classmethod
    def access_path_for(cls, *currency_typetag, module_address=None):
        return cls.struct_tag(*currency_typetag, module_address=module_address).access_vector()
