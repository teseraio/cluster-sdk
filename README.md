
# Cluster-SDK

Cluster-sdk is an end-to-end framework to build high available applications based on Raft.

## Features

- Extensible GRPC server.
- Gossip discovery based on [Serf](github.com/hashicorp/serf).
- High availability with [Raft](github.com/hashicorp/raft).

## Example

```
$ go run ./example/main.go --node-name one --bootstrap 3 --grpc-port 5000

$ go run ./example/main.go --node-name two --bootstrap 3 --grpc-port 5001

$ go run ./example/main.go --node-name three --bootstrap 3 --grpc-port 5002
```
