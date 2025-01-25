package kannri

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	apiv1 "github.com/aita/kannri/kannripb/kannri/v1"
)

type Manager struct {
	services []*service

	mu sync.Mutex
}

func New() *Manager {
	return &Manager{}
}

func (m *Manager) CreateService(ctx context.Context, req *apiv1.CreateServiceRequest) (*apiv1.Service, error) {
	srv := &service{
		name:        req.Name,
		command:     req.Command,
		args:        req.Args,
		environment: req.Environment,
		status:      apiv1.ServiceStatus_SERVICE_STATUS_STOPPED,
	}

	m.mu.Lock()
	m.services = append(m.services, srv)
	m.mu.Unlock()

	return srv.toProto(), nil
}

func (m *Manager) Run(ctx context.Context) error {
	if err := m.start(ctx); err != nil {
		return err
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT)

	var err error
	select {
	case <-ctx.Done():
		err = ctx.Err()
		slog.ErrorContext(ctx, "Context done, stopping services", "error", err)
		break

	case sig := <-sigch:
		slog.InfoContext(ctx, "Received signal, stopping services", "signal", sig)
		return nil
	}

	if terminateErr := m.terminate(ctx); terminateErr != nil {
		err = errors.Join(err, terminateErr)
	}

	return err
}

func (m *Manager) start(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var err error
	for _, srv := range m.services {
		if startErr := srv.start(ctx); startErr != nil {
			err = errors.Join(err, startErr)
		}
	}
	return err
}

func (m *Manager) terminate(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var err error
	for _, srv := range m.services {
		if stopErr := srv.stop(ctx); stopErr != nil {
			err = errors.Join(err, stopErr)
		}
	}
	return err
}
