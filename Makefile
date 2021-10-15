all: build
build: generate
	go build main.go
generate:
	protoc -I ./api/v1 \
	--go_out=./api/v1 --go_opt=paths=source_relative \
    --go-grpc_out=./api/v1 --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=./api/v1 --grpc-gateway_opt=logtostderr=true,paths=source_relative,generate_unbound_methods=true \
	--gotag_out=xxx="bson+\"-\"",paths=source_relative,outdir=./api/v1:. \
	api/v1/jukebox.proto