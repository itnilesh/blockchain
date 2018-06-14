# blockchain
Sample in memory blockchain based on grpc client server in golang

Make sure that golang is installed and GOROOT and GOPATH are exported.
#### GRPC Server 

GRPC server created blockchain in memory and gives grpc end points to add and list blockchain.

#### GRPC Client

GRPC Go client adds blocks to the blockchain and then request to list it.

#### How to run ?
- Run GRPC  server from command line from root 'blockchain' of this project  
`bash-3.2$ go run ./salpe/com/server/main.go`    
`Blockchain server started....now you can do  grpc requests`

- Run GRPC command line to do grpc client requests from root `blockchain` .   
`go run ./salpe/com/client/main.go --Add=true` .   
`New block added` .   
`go run main.go  --List=true` .  
`==> genesis` .   
`==> nilesh` .   








