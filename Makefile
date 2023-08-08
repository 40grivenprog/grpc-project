proto:
	rm -f pb/*.proto
	protoc --experimental_allow_proto3_optional --proto_path=proto \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		proto/*.proto

terminal_chat_server:
	go run terminal_chat_server/main.go

terminal_chat_client:
	go run terminal_chat_client/main.go


.PHONY: terminal_chat_client terminal_chat_server
