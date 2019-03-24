# bootstrap geth in testnet 

# 8545 is the default HTTP-RPC server listening port
# 30303 is the default network listening port
# --rpc enables the HTTP-RPC API
# --rpcaddr allows accessing RPC from other containers and/or hosts
# --ipcpath redirect the file to address the problem that we  can't put your geth.ipc file (unix socket) on a non unix partition as explained at https://github.com/ethereum/go-ethereum/issues/16215
docker run --rm --name ethereum -p 30303:30303 -p 8545:8545 -v /e/ethereum:/root/.ethereum ethereum/client-go:alltools-v1.8.23 geth --testnet --syncmode "fast" --rpc --rpcaddr 0.0.0.0 --ipcpath /root/ethereum/testnet/geth.ipc