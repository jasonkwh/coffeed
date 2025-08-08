package handler

import (
	"context"

	"github.com/jasonkwh/coffeed/internal/adapter"
	"go.uber.org/zap"
)

var _ = adapter.DaemonHandler(&daemonHandler{})

type daemonHandler struct {
	zl *zap.Logger
}

func NewDaemonHandler(
	opts ...handlerOption,
) adapter.DaemonHandler {
	hdl := &daemonHandler{}

	for _, opt := range opts {
		opt(hdl)
	}

	if hdl.zl == nil {
		hdl.zl = zap.NewNop()
	}

	return hdl
}

func (hdl *daemonHandler) SetBusy(
	ctx context.Context,
	req *proto.SetBusyRequest,
) (*proto.SetBusyResponse, error) {
	// TODO: Implement
)
