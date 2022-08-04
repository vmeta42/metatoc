from ..canoser import Uint64, Struct

GasCarrier = Uint64

ZERO_GAS_UNITS = 0
MAX_ABSTRACT_MEMORY_SIZE = Uint64.max_value
WORD_SIZE = 8
CONST_SIZE = 1
REFERENCE_SIZE = 8
STRUCT_SIZE = 2
DEFAULT_ACCOUNT_SIZE = 32
LARGE_TRANSACTION_CUTOFF = 600

class GasConstants(Struct):
    _fields = [
        ("global_memory_per_byte_cost", Uint64),
        ("global_memory_per_byte_write_cost", Uint64),
        ("min_transaction_gas_units", Uint64),
        ("large_transaction_cutoff", Uint64),
        ("instrinsic_gas_per_byte", Uint64),
        ("maximum_number_of_gas_units", Uint64),
        ("min_price_per_gas_unit", Uint64),
        ("max_price_per_gas_unit", Uint64),
        ("max_transaction_size_in_bytes", Uint64)
    ]

    @classmethod
    def default(cls):
        return cls(8, 8, 600, 600, 8, 2_000_000, 0, 10_000, 4096)


class GasCost(Struct):
    _fields = [
        ("instruction_gas", Uint64),
        ("memory_gas", Uint64),
    ]

class CostTable(Struct):
    _fields = [
        ("instruction_table", [GasCost]),
        ("native_table", [GasCost]),
        ("gas_constants", GasConstants),
    ]
