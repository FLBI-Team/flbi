import os

initCmd = 'geth init $HOME/ethereum/genesis/m{}.json > /dev/null 2>&1'

startCmd = 'nohup geth --bootnodes "{}" --rpc --rpcapi eth,web3,net --rpcaddr "0.0.0.0" --rpccorsdomain "*" --allow-insecure-unlock --gcmode archive --networkid 15416 --keystore $HOME/ethereum/account --unlock $(cat $HOME/ethereum/account.txt) --password $HOME/ethereum/password.txt --mine > $HOME/ethereum/eth.log 2>&1 &'

stopCmd = 'killall -KILL geth'

rmCmd = 'rm -rf $HOME/.ethereum/geth/chaindata $HOME/.ethereum/geth/transactions.rlp && rm $HOME/ethereum/eth.log'

rebootCmd = 'sudo reboot'

sendGethCmd = 'scp -i pikey ../geth pi@{}:'

installGethCmd = 'sudo mv geth /usr/bin'

sendDirCmd = 'scp -i pikey -r ../ethereum pi@{}:'

def get_nodes():
  f = open("hosts", "r")
  lines = f.readlines()
  nodes = []
  for l in lines:
    nodes.append(l.strip())
  return nodes

def get_bootnodes():
  f = open("bootnodes", "r")
  lines = f.readlines()
  bootnodes = []
  for l in lines:
    bootnodes.append(l.strip())
  return bootnodes

def init(m=1):
  nodes = get_nodes()

  cmd = initCmd.format(m)

  for i in range(int(m)):
    rExecute(nodes[i], cmd)

def start(m=1):
  nodes = get_nodes()

  bootnodes = get_bootnodes()

  if m < len(bootnodes):
    bootnodes = bootnodes[:m]

  bootnodestr = ','.join(bootnodes)

  cmd = startCmd.format(bootnodestr)
  for i in range(int(m)):
    rExecute(nodes[i], cmd)

def stop(m=1):
  nodes = get_nodes()

  for i in range(int(m)):
    rExecute(nodes[i], stopCmd)

def rm(m=1):
  nodes = get_nodes()

  for i in range(int(m)):
    rExecute(nodes[i], rmCmd)

def reboot(m=1):
  nodes = get_nodes()

  for i in range(int(m)):
    rExecute(nodes[i], rebootCmd)

def install():
  nodes = get_nodes()

  for n in nodes:
    execute(sendGethCmd.format(n))
    rExecute(n, installGethCmd)

def setupDir():
  nodes = get_nodes()

  for n in nodes:
    execute(sendDirCmd.format(n))

def rExecute(host, cmd):
  rCmd = "ssh -i pikey pi@{} '{}'".format(host, cmd)
  execute(rCmd)

def execute(cmd):
  print(cmd)
  os.system(cmd)
