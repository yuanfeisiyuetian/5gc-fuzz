from argparse import Namespace
from pyjfuzz.lib import *

config = PJFConfiguration(Namespace(json={"test": ["1", 2, True]}, nologo=True, level=6))
fuzzer = PJFFactory(config)
path = "test.json"
with open(path,"w") as f:
    while True:
        print (fuzzer.fuzzed)
        f.writelines(fuzzer.fuzzed)
