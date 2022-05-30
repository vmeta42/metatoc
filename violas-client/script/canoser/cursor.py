
class Cursor:
    def __init__(self, buffer, offset=0):
        self.buffer = buffer
        if isinstance(buffer, list):
            self.buffer = bytes(buffer)
        if isinstance(buffer, bytearray):
            self.buffer = bytes(buffer)
        self.offset = offset
        self.buffer_len = len(self.buffer)

    def read_bytes(self, size):
        end = self.offset + size
        if end > self.buffer_len:
            raise IOError("{} exceed buffer size: {}".format(end, self.buffer_len))
        ret = self.buffer[self.offset:end]
        self.offset = end
        return ret

    def read_to_end(self):
        ret = self.buffer[self.offset:]
        self.offset = self.buffer_len
        return ret

    def peek_bytes(self, size):
        end = self.offset + size
        if end > self.buffer_len:
            raise IOError("{} exceed buffer size: {}".format(end, total))
        return self.buffer[self.offset:end]

    def is_finished(self):
        return self.offset == self.buffer_len

    def position(self):
        return self.offset

    def read_u8(self):
        arr = self.read_bytes(1)
        return int(arr[0])