module github.com/hikata101/health_data_service

go 1.24.10

replace github.com/hikata101/health_data_service/gen => ./gen

require (
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.76.0
	google.golang.org/protobuf v1.36.10
)

require (
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
)
