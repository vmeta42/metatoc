server_url: str = "http://172.22.50.186:8080"
root_private="5138780620b4b0216fb2ce7962e95260d0ecd527a51cdcdea825a33517ecf6f9"


VLS: str = "VLS"
CORE_CODE_ADDRESS = b"\x02".rjust(16, b"\x00")

from enum import IntEnum

class Meta42CodeType(IntEnum):
    META42_INITIALIZE = 0
    META42_ACCEPT = 1
    META42_MINT_TOKEN = 2
    META42_SHARE_TOKEN_BY_ID = 3
    META42_SHARE_TOKEN_BY_INDEX = 4

script_hashs={
    "meta42_mint_token": "4676ac7785c883cc75973baceb5f446a32f869d5dd094b0b941e43f102241f68",
    "meta42_initialize": "e3589161c8f4be989f267d5221eac3e990c26785f13da0d96bc221a694560280",
    "meta42_accept": "9003e43d1ee2bfb53dffdb4940a65b68e71e7244ea7ecf13516bc8c3d5a8a044",
    "meta42_share_token_by_id": "01014209683c5d0f20e9de02308eae5209c8b5bc5442722665b66514754315fb",
    "meta42_share_token_by_index": "62cf77924c2c0f67afa5c6b6730aa5590bf4f1cb6826b81af3c2a5d27de99dc2"
}

hash_to_type_map = {
    script_hashs["meta42_initialize"]: Meta42CodeType.META42_INITIALIZE,
    script_hashs["meta42_accept"]: Meta42CodeType.META42_ACCEPT,
    script_hashs["meta42_mint_token"]: Meta42CodeType.META42_MINT_TOKEN,
    script_hashs["meta42_share_token_by_id"]: Meta42CodeType.META42_SHARE_TOKEN_BY_ID,
    script_hashs["meta42_share_token_by_index"]: Meta42CodeType.META42_SHARE_TOKEN_BY_INDEX
}

name_to_type_map = {
    "meta42_initialize": Meta42CodeType.META42_INITIALIZE,
    "meta42_accept": Meta42CodeType.META42_ACCEPT,
    "meta42_mint_token": Meta42CodeType.META42_MINT_TOKEN,
    "meta42_share_token_by_id": Meta42CodeType.META42_SHARE_TOKEN_BY_ID,
    "meta42_share_token_by_index": Meta42CodeType.META42_SHARE_TOKEN_BY_INDEX
}

minted_events_key = "0400000000000000458d623300e797451b3e794a45b41065"
shared_events_key = "0500000000000000458d623300e797451b3e794a45b41065"
