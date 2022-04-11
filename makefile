gen_proto:
	protoc --go_out=rhombic/ohs/local/pl/dto --go-grpc_out=rhombic/ohs/local/pl/dto rhombic/ohs/local/pl/dto/*.proto

gen_proto_errors:
	protoc --go_out=rhombic/ohs/local/pl/errors --go-grpc_out=rhombic/ohs/local/pl/errors rhombic/ohs/local/pl/errors/*.proto

server:
	go run main.go