from dataclasses import dataclass
import typing
from diem import serde_types as st
from diem import diem_types
from diem.stdlib import (
    ScriptCall,
    ScriptFunctionCall,
    decode_u8vector_argument,
    decode_u64_argument,
    decode_address_argument,
    bcs_deserialize

)

from diem.diem_types import (
    Script,
    ScriptFunction,
    TransactionPayload,
    TransactionPayload__ScriptFunction,
    Identifier,
    ModuleId,
    TypeTag,
    AccountAddress,
    TransactionArgument,
    TransactionArgument__Bool,
    TransactionArgument__U8,
    TransactionArgument__U64,
    TransactionArgument__U128,
    TransactionArgument__Address,
    TransactionArgument__U8Vector,
)

@dataclass(frozen=True)
class ScriptCall__Meta42Accept(ScriptFunctionCall):
    pass

@dataclass(frozen=True)
class ScriptCall__Meta42Initialize(ScriptFunctionCall):
    pass

@dataclass(frozen=True)
class ScriptCall__Meta42MintToken(ScriptFunctionCall):
    hdfs_path: bytes

@dataclass(frozen=True)
class ScriptCall__Meta42ShareTokenById(ScriptFunctionCall):
    recevier: diem_types.AccountAddress
    token_id: bytes
    metadata: bytes

@dataclass(frozen=True)
class ScriptCall__Meta42ShareTokenByIndex(ScriptFunctionCall):
    recevier: diem_types.AccountAddress
    index: st.uint64
    metadata: bytes


@dataclass(frozen=True)
class ScriptFunctionCall__Meta42Accept(ScriptFunctionCall):
    pass

@dataclass(frozen=True)
class ScriptFunctionCall__Meta42Initialize(ScriptFunctionCall):
    pass

@dataclass(frozen=True)
class ScriptFunctionCall__Meta42MintToken(ScriptFunctionCall):
    hdfs_path: bytes

@dataclass(frozen=True)
class ScriptFunctionCall__Meta42ShareTokenById(ScriptFunctionCall):
    recevier: diem_types.AccountAddress
    token_id: bytes
    metadata: bytes

@dataclass(frozen=True)
class ScriptFunctionCall__Meta42ShareTokenByIndex(ScriptFunctionCall):
    recevier: diem_types.AccountAddress
    index: st.uint64
    metadata: bytes


def encode_meta42_accept_script():
    return Script(
        code=META42_ACCEPT,
        ty_args=[],
        args=[],
    )

def encode_meta42_accept_script_function():
    return TransactionPayload__ScriptFunction(
        value=ScriptFunction(
            module=ModuleId(
                address=AccountAddress.from_hex("00000000000000000000000000000002"),
                name=Identifier("Meta42"),
            ),
            function=Identifier("accept"),
            ty_args=[],
            args=[],
        )
    )

def encode_meta42_initialize_script():
    return Script(
        code=META42_INITIALIZE,
        ty_args=[],
        args=[],
    )

def encode_meta42_initialize_script_function():
    return TransactionPayload__ScriptFunction(
        value=ScriptFunction(
            module=ModuleId(
                address=AccountAddress.from_hex("00000000000000000000000000000002"),
                name=Identifier("Meta42"),
            ),
            function=Identifier("initialize"),
            ty_args=[],
            args=[],
        )
    )

def encode_meta42_mint_token_script(hdfs_path: bytes):
    return Script(
        code=META42_MINT_TOKEN,
        ty_args=[],
        args=[TransactionArgument__U8Vector(value=hdfs_path)],
    )

def encode_meta42_mint_token_script_function(hdfs_path: bytes):
    return TransactionPayload__ScriptFunction(
        value=ScriptFunction(
            module=ModuleId(
                address=AccountAddress.from_hex("00000000000000000000000000000002"),
                name=Identifier("Meta42"),
            ),
            function=Identifier("mint_token"),
            ty_args=[],
            args=[hdfs_path],
        )
    )

def encode_meta42_share_token_by_id_script(recevier: diem_types.AccountAddress, token_id: bytes, metadata: bytes):
    return Script(
        code=META42_SHARE_TOKEN_BY_ID,
        ty_args=[],
        args=[
            TransactionArgument__Address(value=recevier),
            TransactionArgument__U8Vector(value=token_id),
            TransactionArgument__U8Vector(value=metadata)
        ]
    )

def encode_meta42_share_token_by_id_script_function(recevier: diem_types.AccountAddress, token_id: bytes, metadata: bytes):
    return TransactionPayload__ScriptFunction(
        value=ScriptFunction(
            module=ModuleId(
                address=AccountAddress.from_hex("00000000000000000000000000000002"),
                name=Identifier("Meta42"),
            ),
            function=Identifier("share_token_by_id"),
            ty_args=[],
            args=[recevier, token_id, metadata]
        )
    )


def encode_meta42_share_token_by_index_script(recevier: diem_types.AccountAddress, index: st.uint64, metadata: bytes):
    return Script(
        code=META42_SHARE_TOKEN_BY_INDEX,
        ty_args=[],
        args=[
            TransactionArgument__Address(value=recevier),
            TransactionArgument__U64(value=index),
            TransactionArgument__U8Vector(value=metadata)
        ]
    )

def encode_meta42_share_token_by_index_script_function(recevier: diem_types.AccountAddress, token_id: bytes, metadata: bytes):
    return TransactionPayload__ScriptFunction(
        value=ScriptFunction(
            module=ModuleId(
                address=AccountAddress.from_hex("00000000000000000000000000000002"),
                name=Identifier("Meta42"),
            ),
            function=Identifier("share_token_by_index"),
            ty_args=[],
            args=[recevier, token_id, metadata]
        )
    )

def decode_meta42_accept_script(script: Script) -> ScriptCall:
    return ScriptCall__Meta42Accept()


def decode_meta42_accept_script_function(script: TransactionPayload) -> ScriptFunctionCall:
    if not isinstance(script, ScriptFunction):
        raise ValueError("Unexpected transaction payload")
    return ScriptFunctionCall__Meta42Accept(
    )

def decode_meta42_initialize_script(script: Script) -> ScriptCall:
    return ScriptCall__Meta42Initialize()

def decode_meta42_initialize_script_function(script: TransactionPayload) -> ScriptFunctionCall:
    if not isinstance(script, ScriptFunction):
        raise ValueError("Unexpected transaction payload")
    return ScriptFunctionCall__Meta42Initialize()

def decode_meta42_mint_token_script(script: Script) -> ScriptCall:
    return ScriptCall__Meta42MintToken(
        hdfs_path=decode_u8vector_argument(script.ty_args[0])
    )

def decode_meta42_mint_token_script_function(script: TransactionPayload) -> ScriptFunctionCall:
    if not isinstance(script, ScriptFunction):
        raise ValueError("Unexpected transaction payload")
    return ScriptFunctionCall__Meta42MintToken(
        hdfs_path=bcs_deserialize(script.ty_args[0], bytes)[0]
    )

def decode_meta42_share_token_by_id_script(script: Script) -> ScriptCall:
    return ScriptCall__Meta42ShareTokenById(
        recevier=decode_address_argument(script.args[0]),
        token_id=decode_u8vector_argument(script.ty_args[1]),
        metadata=decode_u8vector_argument(script.ty_args[2]),

    )

def decode_meta42_share_token_by_id_script_function(script: TransactionPayload) -> ScriptFunctionCall:
    if not isinstance(script, ScriptFunction):
        raise ValueError("Unexpected transaction payload")
    return ScriptFunctionCall__Meta42Accept(
        recevier=bcs_deserialize(script.ty_args[0], AccountAddress)[0],
        token_id=bcs_deserialize(script.ty_args[1],bytes)[0],
        metadata=bcs_deserialize(script.ty_args[2],bytes)[0],
    )

def decode_meta42_share_token_by_index_script(script: Script) -> ScriptCall:
    return ScriptCall__Meta42ShareTokenById(
        recevier=decode_address_argument(script.args[0]),
        index=decode_u64_argument(script.ty_args[1]),
        metadata=decode_u8vector_argument(script.ty_args[2]),

    )

def decode_meta42_share_token_by_index_script_function(script: TransactionPayload) -> ScriptFunctionCall:
    if not isinstance(script, ScriptFunction):
        raise ValueError("Unexpected transaction payload")
    return ScriptFunctionCall__Meta42Accept(
        recevier=bcs_deserialize(script.ty_args[0], AccountAddress)[0],
        index=bcs_deserialize(script.ty_args[1],st.uint64)[0],
        metadata=bcs_deserialize(script.ty_args[2],bytes)[0],
    )

META42_ACCEPT=b"\xa1\x1c\xeb\x0b\x03\x00\x00\x00\x05\x01\x00\x02\x03\x02\x05\x05\x07\x06\x07\x0d\x0e\x08\x1b\x10\x00\x00\x00\x01\x02\x01\x00\x01\x0c\x00\x01\x06\x0c\x06\x4d\x65\x74\x61\x34\x32\x06\x61\x63\x63\x65\x70\x74\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x01\x03\x0e\x00\x11\x00\x02"
META42_INITIALIZE=b"\xa1\x1c\xeb\x0b\x03\x00\x00\x00\x05\x01\x00\x02\x03\x02\x05\x05\x07\x06\x07\x0d\x12\x08\x1f\x10\x00\x00\x00\x01\x02\x01\x00\x01\x0c\x00\x01\x06\x0c\x06\x4d\x65\x74\x61\x34\x32\x0a\x69\x6e\x69\x74\x69\x61\x6c\x69\x7a\x65\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x01\x03\x0e\x00\x11\x00\x02"
META42_MINT_TOKEN=b"\xa1\x1c\xeb\x0b\x03\x00\x00\x00\x05\x01\x00\x02\x03\x02\x05\x05\x07\x0a\x07\x11\x12\x08\x23\x10\x00\x00\x00\x01\x02\x01\x00\x02\x0c\x0a\x02\x00\x02\x06\x0c\x0a\x02\x06\x4d\x65\x74\x61\x34\x32\x0a\x6d\x69\x6e\x74\x5f\x74\x6f\x6b\x65\x6e\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x01\x04\x0e\x00\x0b\x01\x11\x00\x02"
META42_SHARE_TOKEN_BY_ID=b"\xa1\x1c\xeb\x0b\x03\x00\x00\x00\x05\x01\x00\x02\x03\x02\x05\x05\x07\x10\x07\x17\x19\x08\x30\x10\x00\x00\x00\x01\x02\x01\x00\x04\x0c\x05\x0a\x02\x0a\x02\x00\x04\x06\x0c\x05\x0a\x02\x0a\x02\x06\x4d\x65\x74\x61\x34\x32\x11\x73\x68\x61\x72\x65\x5f\x74\x6f\x6b\x65\x6e\x5f\x62\x79\x5f\x69\x64\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x01\x06\x0e\x00\x0a\x01\x0b\x02\x0b\x03\x11\x00\x02"
META42_SHARE_TOKEN_BY_INDEX=b"\xa1\x1c\xeb\x0b\x03\x00\x00\x00\x05\x01\x00\x02\x03\x02\x05\x05\x07\x0e\x07\x15\x1c\x08\x31\x10\x00\x00\x00\x01\x02\x01\x00\x04\x0c\x05\x03\x0a\x02\x00\x04\x06\x0c\x05\x03\x0a\x02\x06\x4d\x65\x74\x61\x34\x32\x14\x73\x68\x61\x72\x65\x5f\x74\x6f\x6b\x65\x6e\x5f\x62\x79\x5f\x69\x6e\x64\x65\x78\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x01\x06\x0e\x00\x0a\x01\x0a\x02\x0b\x03\x11\x00\x02"


META42_SCRIPT_ENCODER_MAP: typing.Dict[typing.Type[ScriptCall], typing.Callable[[ScriptCall], Script]] = {
    ScriptCall__Meta42Accept: encode_meta42_accept_script,
    ScriptCall__Meta42Initialize: encode_meta42_initialize_script,
    ScriptCall__Meta42MintToken: encode_meta42_mint_token_script,
    ScriptCall__Meta42ShareTokenById: encode_meta42_share_token_by_id_script,
    ScriptCall__Meta42ShareTokenByIndex: encode_meta42_share_token_by_index_script
}

# pyre-ignore
META42_SCRIPT_FUNCTION_ENCODER_MAP: typing.Dict[
    typing.Type[ScriptFunctionCall], typing.Callable[[ScriptFunctionCall], TransactionPayload]
] = {
    ScriptFunctionCall__Meta42Accept: encode_meta42_accept_script_function,
    ScriptFunctionCall__Meta42Initialize: encode_meta42_initialize_script_function,
    ScriptFunctionCall__Meta42MintToken: encode_meta42_mint_token_script_function,
    ScriptFunctionCall__Meta42ShareTokenById: encode_meta42_share_token_by_id_script_function,
    ScriptFunctionCall__Meta42ShareTokenByIndex: encode_meta42_share_token_by_index_script_function,
}


META42_SCRIPT_DECODER_MAP: typing.Dict[bytes, typing.Callable[[Script], ScriptCall]] = {
    META42_ACCEPT: decode_meta42_accept_script,
    META42_INITIALIZE: decode_meta42_initialize_script,
    META42_MINT_TOKEN: decode_meta42_mint_token_script,
    META42_SHARE_TOKEN_BY_ID: decode_meta42_share_token_by_id_script,
    META42_SHARE_TOKEN_BY_INDEX: decode_meta42_share_token_by_index_script,
}

META42_SCRIPT_FUNCTION_DECODER_MAP: typing.Dict[str, typing.Callable[[TransactionPayload], ScriptFunctionCall]] = {
    "META42accept": decode_meta42_accept_script_function,
    "META42initialize": decode_meta42_initialize_script_function,
    "META42mint_token": decode_meta42_mint_token_script_function,
    "META42share_token_by_id": decode_meta42_share_token_by_id_script_function,
    "META42share_token_by_index": decode_meta42_share_token_by_index_script_function,
}