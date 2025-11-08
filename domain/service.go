package domain

import (
	"fmt"

	"github.com/hikata101/health_data_service/infrastructure"
	"github.com/hikata101/health_data_service/logger"

	pb "github.com/hikata101/health_data_service/gen/github.com/hikata101/health_data_service/v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func DownloadDataset(req *pb.DownloadRequest, stream grpc.ServerStreamingServer[pb.DownloadReply]) error {
	// Implement the logic to download the dataset here.
	// For now, just print a message.
	switch v := req.Request.(type) {
	case *pb.DownloadRequest_WhoEurope:
		// Handle WHO Europe download requests
		IndicatorCode, err := GetIndicatorCode(v.WhoEurope.Indicator)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("Error getting indicator code: %v", err))
			stream.Send(&pb.DownloadReply{
				Status: int32(codes.InvalidArgument),
			})
			return err
		}
		countryCode, err := GetCountryCode(v.WhoEurope.Country)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("Error getting country code: %v", err))
			stream.Send(&pb.DownloadReply{
				Status: int32(codes.InvalidArgument),
			})
			return err
		}
		query_params := "filter=COUNTRY:" + countryCode
		logger.Logger.Debug(fmt.Sprintf("Downloading WHO Europe dataset with parameters: %+v\n", query_params))
		// Call the external API to get the data
		resp, err := infrastructure.Execute(stream.Context(), IndicatorCode, query_params)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("Error executing WHO Europe request: %v", err))
			stream.Send(&pb.DownloadReply{
				Status: int32(codes.Unknown),
			})
			return err
		}
		logger.Logger.Debug(fmt.Sprintf("Received WHO Europe response: %s", resp))
		parsed, err := ParseWHOEuropeCSVToReply(resp)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("Error parsing WHO Europe response: %v", err))
			stream.Send(&pb.DownloadReply{
				Status: int32(codes.Internal),
			})
			return err
		}
		print(parsed)
		// Send the parsed protobuf message back to the client
		logger.Logger.Debug("Successfully parsed WHO Europe response into protobuf")
		stream.Send(&pb.DownloadReply{
			Status: int32(codes.OK),
			Reply: &pb.DownloadReply_WhoEuropeReply{
				WhoEuropeReply: parsed,
			}})
	default:
		logger.Logger.Error("Unknown download request type")
		return errors.New("unknown download request type")
	}
	return nil
}
