# @shumengineer

.PHONY: protos

protos:
	protoc -I=protos --go_out=. --go-grpc_out=. protos/*.proto