.PHONY: pb
pb:
	@protoc -I protos/ protos/*.proto --go_out=plugins=grpc:pb