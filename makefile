
run:
	go run main.go

ent:
	go generate ./repository/ent



protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --validate_out=paths=source_relative,lang=go:. server/schema/helloworld.proto