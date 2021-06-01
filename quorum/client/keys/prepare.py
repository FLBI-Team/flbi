import json, os
from web3 import Web3


def keyFunc(f):
  if f.startswith("UTC"):
    return "z"+f
  return f

filenames = os.listdir()
filenames.sort(key=keyFunc)

addrList = []
i = 0

for filename in filenames:

  if filename.endswith(".py") or filename.endswith(".txt"):
    continue

  with open(filename) as file:
    data = json.load(file)
    addrList.append( data["address"])
  
  os.rename(filename, "clientkey{}".format(str(i).zfill(3)))

  i += 1


alloc = {}

for addr in addrList:
  alloc[addr] = {
    "balance": "0x200000000000000000000000000000000000000000000000000000000000000"
  }

allocToPrint = {
  "alloc": alloc
}
print(json.dumps(allocToPrint, indent=2))

print("\n\n")

print(
  "address[] memory clientkeys = new address[]({});\n".format(
  len(addrList))
)

for i, addr in enumerate(addrList):
  print("clientkeys[{}] = {};".format(i, Web3.toChecksumAddress(addr)))
