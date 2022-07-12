import sys
import os

file_name = sys.argv[1]
file = open(file_name, "a+")
for path in sys.argv[2:]:
    with open(path, "rb") as input:
        code = input.read()
        s = "b\""
        for i in code:
            s += "\\x"
            s += "{:0>2x}".format(i)
        s += "\""
        file.write(os.path.basename(path)[:-3].upper())
        file.write("=")
        file.write(s)
        file.write("\n")
