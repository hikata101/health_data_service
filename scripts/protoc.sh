echo "Generating health_data_service gRPC codes..."
protoc -I=. --go-grpc_out=./gen --go_out=./gen ./proto/*.proto