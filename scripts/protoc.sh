echo "Generating health_data_service gRPC codes..."
protoc -I=./proto --go-grpc_out=./gen --go_out=./gen ./proto/*.proto

cd ./gen/github.com/hikata101/health_data
go mod init github.com/hikata101/health_data
go mod tidy

cd ../../../..
go mod tidy
echo "gRPC code generation completed."