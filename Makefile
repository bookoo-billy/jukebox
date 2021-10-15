all: build
build: generate
	go build main.go
generate:
	mkdir -p gen/api/v1
	protoc -I ./api/v1 \
		--go_out=./gen/api/v1 --go_opt=paths=source_relative \
		--go-grpc_out=./gen/api/v1 --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./gen/api/v1 --grpc-gateway_opt=logtostderr=true,paths=source_relative,generate_unbound_methods=true \
		api/v1/jukebox.proto

	protoc -I ./api/v1 \
		--gotag_out=xxx="bson+\"-\"",paths=source_relative,outdir=./gen/api/v1:. \
		api/v1/jukebox.proto