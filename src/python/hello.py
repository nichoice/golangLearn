
import ctypes
import os
from ctypes import cdll


base_path = os.path.dirname(os.path.realpath(__file__))

lib = cdll.LoadLibrary(os.path.join(base_path, 'hello.so'))

res = lib.Hello
res.restype = ctypes.c_char_p
print(res())


cstr = lib.Cstr
cstr.argtype = ctypes.c_char_p
cstr.restype = ctypes.c_char_p

str_test = "hello".encode('utf-8')
print(str_test.decode('utf-8', 'ignore'))
