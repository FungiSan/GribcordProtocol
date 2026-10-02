app:
	rm -rf gen/app
	protoc -I proto proto/gribcord.proto --go_out=./gen/app/ --go-grpc_out=./gen/app

identity:
	rm -rf gen/identity
	protoc -I proto proto/identity.proto --go_out=./gen/identity/ --go-grpc_out=./gen/identity/

