package serverimplement

import (
	"context"

	contracts "github.com/atrop1ne/logger-service-api/gen/go/contracts/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	contracts.UnimplementedLogsServer
}

func (s *GRPCServer) GetLogsLevels(ctx context.Context, req *emptypb.Empty) (resp *contracts.GetLogsLevelsResponse, err error) {
	resp = &contracts.GetLogsLevelsResponse{
		LogsLevels: []*contracts.LogLevel{
			{Id: 1, Name: "INFO"},
			{Id: 1, Name: "ERROR"},
		},
	}
	return
}

func (s *GRPCServer) GetLogs(ctx context.Context, req *contracts.GetLogsRequest) (resp *contracts.GetLogsResponse, err error) {
	resp = &contracts.GetLogsResponse{
		Logs: []*contracts.Log{
			{Id: 1, LevelId: 1, Message: "Test"},
			{Id: 1, LevelId: 1, Message: "Test2"},
		},
	}
	return
}
