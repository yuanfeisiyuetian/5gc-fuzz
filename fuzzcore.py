from argparse import Namespace
from pyjfuzz.lib import *

config = PJFConfiguration(Namespace(json={"test": ["1", 2, True]}, nologo=True, level=6))
fuzzer = PJFFactory(config)
while True:
    print (fuzzer.fuzzed)
