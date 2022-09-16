from ..canoser import Struct
from .account_info import AccountInfoResource, GlobalInfoResource, DiemConfigResource
from .contants import CORE_CODE_ADDRESS
from .network_address import NetworkAddress

from ipaddress import ip_address


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

    def get_validator_config(self):
        if self.exists():
            resource = self.get(b'\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\nDiemConfig\nDiemConfig\x01\x07\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\nDiemSystem\nDiemSystem\x00')
            if resource:
                return DiemConfigResource.deserialize(resource)

    def get_validator_infos(self):
        config = self.get_validator_config()
        infos = []
        for v in config.payload.validators:
            fullnode_network_addresses = v.config.fullnode_network_addresses
            network_address = NetworkAddress.deserialize(fullnode_network_addresses[2:])
            info = {
                "address": v.addr.hex(),
                "ip":str(ip_address(network_address.Protocol[0].value.to_bytes(4, byteorder="little"))),
                # "port": network_address.Protocol[1].value
                "port": 50001
            }
            infos.append(info)
        return infos


    def get(self, key):
        if self.exists():
            return self.ordered_map.get(key)
