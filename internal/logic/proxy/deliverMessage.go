package proxy

import (
	"context"
	"gim/pkg/protocol/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IDeliveMessageProxy interface {
	DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*emptypb.Empty, error)
}

var DeliveMessageProxy IDeliveMessageProxy
