package tracker_common

import (
	"context"
	"github.com/automation-and-vision/tracker-common/grpc/issue_grpc"
)

type IssueServer struct {
	issue_grpc.UnimplementedIssueServiceServer
}

func (s IssueServer) OnProjectCreated(ctx context.Context, req *issue_grpc.CreateProjectRequest) (*issue_grpc.CreateProjectResponse, error) {
	return &issue_grpc.CreateProjectResponse{
		Response: "response",
	}, nil
}
