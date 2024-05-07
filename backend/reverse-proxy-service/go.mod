module github.com/Enthreeka/reverse-proxy-service

//replace github.com/Enthreeka/proto-proxy => ../proto-proxy

go 1.22.0

require (
	//github.com/Entreeka/proto-proxy v1.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/joho/godotenv v1.5.1
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.63.2
)

require (
	github.com/redis/go-redis/v9 v9.5.1
	google.golang.org/protobuf v1.34.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240429193739-8cf5692501f6 // indirect
)
