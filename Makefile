.PHONY: gen
gen:
	protoc -I internal/service/proto/ \
		--go_out=pkg/competition/api --go_opt=paths=source_relative \
		--go-grpc_out=pkg/competition/api --go-grpc_opt=paths=source_relative \
		internal/service/proto/*.proto
