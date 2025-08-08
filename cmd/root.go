package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func root(cmd *cobra.Command, args []string) {
	zl, err := initZapLogger()
	if err != nil {
		log.Fatal("unable to start zap logger")
	}

	var clPool []io.Closer

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func(logger *zap.Logger) {
		<-c
		logger.Info("cancelling")

		cancel()
	}(zl)

	// TODO: start the daemon

	err = notifySystemdReady(zl)
	if err != nil {
		zl.Fatal("failed to notify systemd ready", zap.Error(err))
	}

	zl.Info("daemon started")

	<-ctx.Done()

	err = gracefulClose(clPool)
	if err != nil {
		zl.Error("failed to close the daemon", zap.Error(err))
	}

	err = notifySystemdStopping(zl)
	if err != nil {
		zl.Error("failed to notify systemd", zap.Error(err))
	}
}
