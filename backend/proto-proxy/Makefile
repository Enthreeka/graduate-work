.PHONY: proto
proto:
	docker build -t proto:latest .
	docker run --rm --name=proto -v `pwd`:/build proto:latest

mock-gen:
	#go install go.uber.org/mock/mockgen@latest
	mockgen -source=go/rpc_grpc.pb.go -destination=mock/rpc_mock.go -package=mock