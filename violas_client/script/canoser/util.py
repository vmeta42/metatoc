import struct

def int_list_to_hex(ints):
    return bytes(ints).hex()

def int_list_to_bytes(ints):
    return bytes(ints)
    # return struct.pack("<{}B".format(len(ints)), *ints).hex()

def bytes_to_int_list(bytes_str):
    tp = struct.unpack("<{}B".format(len(bytes_str)), bytes_str)
    return list(tp)

def bytes_to_hex(bytes_str):
    return bytes_str.hex()

def hex_to_bytes(hex_str):
    return bytes.fromhex(hex_str)

def hex_to_int_list(hex_str):
    return bytes_to_int_list(bytes.fromhex(hex_str))