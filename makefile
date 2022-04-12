gen_proto:
	protoc --go_out=rhombic/ohs/local/pl/dto --go-grpc_out=rhombic/ohs/local/pl/dto rhombic/ohs/local/pl/dto/*.proto
	protoc --go_out=rhombic/acl/adapters/pl/dto --go-grpc_out=rhombic/acl/adapters/pl/dto rhombic/acl/adapters/pl/dto/*.proto


gen_proto_errors:
	protoc --go_out=rhombic/ohs/local/pl/errors --go-grpc_out=rhombic/ohs/local/pl/errors rhombic/ohs/local/pl/errors/*.proto

server:
	go run main.go

mock:
	go install github.com/golang/mock/mockgen@v1.6.0
	mockgen -source=rhombic/acl/ports/repositories/account.go -package=mock -destination=utils/mock/account.go # TODO windows环境不兼容编译
