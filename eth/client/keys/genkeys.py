import json, os
from web3 import Web3

for i in range(200):
  os.system("geth --keystore . account new --password password.txt")
