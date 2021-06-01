import os

initCmd = 'cd $HOME/quorum; rm -rf data; mkdir data; cp nodekeys/{}/nodekey data; cp static-nodes/{}/static-nodes.json data; geth --datadir data init raft.json'

startCmd = 'cd $HOME/quorum; export PRIVATE_CONFIG=ignore; nohup geth --datadir data --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21000 {} > log.txt 2>&1 &' 

stopCmd = 'killall -KILL geth'

rebootCmd = 'sudo reboot'

installGethCmd = 'sudo cp geth_quorum /usr/bin/geth'

sendDirCmd = 'scp -i pikey -r ../quorum pi@{}:'

def get_nodes():
  f = open("hosts", "r")
  lines = f.readlines()
  nodes = []
  for l in lines:
    nodes.append(l.strip())
  return nodes

def init(m=2):
  nodes = get_nodes()

  for i in range(int(m)):
    rExecute(nodes[i], initCmd.format(i+1, m))

def start(m=2):
  nodes = get_nodes()

  for i in range(int(m)):
    if i == 0:
      rExecute(nodes[i], startCmd.format(''))
    else:
      rExecute(nodes[i], startCmd.format('--raftjoinexisting {}'.format(i+1)))

def stop(m=2):
  nodes = get_nodes()

  for i in range(int(m)):
    rExecute(nodes[i], stopCmd)

def reboot(m=2):
  nodes = get_nodes()

  for i in range(int(m)):
    rExecute(nodes[i], rebootCmd)

def install():
  nodes = get_nodes()

  for n in nodes:
    rExecute(n, installGethCmd)

def sendDir():
  nodes = get_nodes()

  for n in nodes:
    execute(sendDirCmd.format(n))

def rExecute(host, cmd):
  rCmd = "ssh -i pikey pi@{} '{}'".format(host, cmd)
  execute(rCmd)

def execute(cmd):
  print(cmd)
  os.system(cmd)
