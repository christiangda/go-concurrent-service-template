package service3

import "log/slog"

type Server3 struct {
	stopCh chan struct{}
}

func NewServer3() *Server3 {
	return &Server3{
		stopCh: make(chan struct{}),
	}
}

func (s *Server3) Start() {
	// do something here for long running tasks
	// like a gRPC server
	slog.Info("Starting service 3...")

	<-s.stopCh
}

func (s *Server3) Stop() {
	// do something here to stop the service
	slog.Info("Stopping service 3...")

	close(s.stopCh)
}
