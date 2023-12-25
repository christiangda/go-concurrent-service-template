package service2

import "log/slog"

type Server2 struct {
	stopCh chan struct{}
}

func NewServer2() *Server2 {
	return &Server2{
		stopCh: make(chan struct{}),
	}
}

func (s *Server2) Start() {
	// do something here for long running tasks
	// like a gRPC server
	slog.Info("Starting service 2...")

	<-s.stopCh
}

func (s *Server2) Stop() {
	// do something here to stop the service
	slog.Info("Stopping service 2...")

	close(s.stopCh)
}
