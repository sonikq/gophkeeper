proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --experimental_allow_proto3_optional  --go-grpc_opt=paths=source_relative internal/delivery/grpc/v1/gophkeeper.proto
run-server:
	go run cmd/server/main.go
run-client:
	go run cmd/client/main.go