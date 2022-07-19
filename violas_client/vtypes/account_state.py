from ..canoser import Struct
from .account_info import AccountInfoResource, GlobalInfoResource
from .contants import CORE_CODE_ADDRESS

class AccountState(Struct):
    _fields = [
        ("ordered_map", {bytes: bytes})
    ]

    def exists(self):
        if isinstance(self.ordered_map, dict):
            return True
        return False

    def get_account_info_resource(self, currency_module_address=CORE_CODE_ADDRESS):
        if self.exists():
            resource = self.get(AccountInfoResource.resource_path(currency_module_address))
            if resource:
                return AccountInfoResource.deserialize(resource)

    def get_global_info_resource(self, currency_module_address=CORE_CODE_ADDRESS):
        if self.exists():
            resource = self.get(GlobalInfoResource.resource_path(currency_module_address))
            if resource:
                return GlobalInfoResource.deserialize(resource)

    def get(self, key):
        if self.exists():
            return self.ordered_map.get(key)
