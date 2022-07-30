auth_proto:
	 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/auth/proto/auth.proto

filebox_proto:
	 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/filebox/proto/filebox.proto

server:
	go run cmd/main.go