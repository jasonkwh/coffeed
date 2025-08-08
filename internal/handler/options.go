package handler

import "go.uber.org/zap"

type handlerOption func(*daemonHandler)

func WithLogger(zl *zap.Logger) handlerOption {
	return func(h *daemonHandler) {
		h.zl = zl
	}
}
