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
	github.com/c2h5oh/datasize v0.0.0-20231215233829-aa82cc1e6500 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/influxdata/tdigest v0.0.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/smarty/assertions v1.15.0 // indirect
	github.com/smartystreets/goconvey v1.8.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tsenart/go-tsz v0.0.0-20180814235614-0bd30b3df1c3 // indirect
	github.com/tsenart/vegeta v12.7.0+incompatible // indirect
	go.elastic.co/ecszap v1.0.2 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240429193739-8cf5692501f6 // indirect
)
