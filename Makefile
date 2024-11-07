.PHONY: proto
proto:
	protoc \
		--proto_path=api/grpc/proto \
		--go_out=api/grpc/pb --go_opt=paths=source_relative \
		--go-grpc_out=api/grpc/pb --go-grpc_opt=paths=source_relative \
		api/grpc/proto/pedido.proto
