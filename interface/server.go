package handler

import (
	"context"

	"github.com/hikata101/health_data_service/domain"
	pb "github.com/hikata101/health_data_service/gen/github.com/hikata101/health_data_service/v1"
	"github.com/hikata101/health_data_service/logger"
	"google.golang.org/grpc"
)

type healthDataServer struct {
	pb.UnimplementedDatasetServer
}

func NewhealthDataServer() healthDataServer {
	return healthDataServer{}
}

func (s healthDataServer) Download(req *pb.DownloadRequest, stream grpc.ServerStreamingServer[pb.DownloadReply]) error {
	logger.Logger.Info("Download request received")
	err := domain.DownloadDataset(req, stream)
	if err != nil {
		logger.Logger.Error("Error processing download request: " + err.Error())
		return err
	}
	logger.Logger.Info("Download request processed successfully")
	return nil
}

func (s healthDataServer) Upload(stream grpc.ClientStreamingServer[pb.UploadChunk, pb.UploadReply]) error {
	logger.Logger.Debug("Upload request received")
	logger.Logger.Debug("Upload functionality not implemented yet")
	return nil
}

func (s healthDataServer) List(ctx context.Context, req *pb.ListDatasetsRequest) (*pb.ListDatasetsReply, error) {
	logger.Logger.Debug("List request received")
	logger.Logger.Debug("List functionality not implemented yet")
	return nil, nil
}
