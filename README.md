This Repo have implementation of:

1) Unary call
2) Server Streaming
3) Client streaming
4) Bi-Directional Streaming


To run the project:
```
go mod tidy
open 2 terminal
  Terminal 1: brew install protobuf
  Terminal 1: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  Terminal 1: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  Terminal 1: export PATH="$PATH:$(go env GOPATH)/bin"
  Terminal 1: protoc --go_out=. --go-grpc_out=. proto/greet.proto
  Termianl 1: /project/server: go run *.go
  Termianl 2: /project/client: go run *.go
```
