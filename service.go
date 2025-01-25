package kannri

import (
	"context"
	"log/slog"
	"os"
	"os/exec"

	apiv1 "github.com/aita/kannri/kannripb/kannri/v1"
)

type service struct {
	name        string
	description string
	command     string
	args        []string
	environment []string
	status      apiv1.ServiceStatus

	cmd *exec.Cmd
}

func (srv *service) toProto() *apiv1.Service {
	return &apiv1.Service{
		Name:        srv.name,
		Status:      srv.status,
		Description: srv.description,
		Command:     srv.command,
		Args:        srv.args,
	}
}

func (srv *service) start(ctx context.Context) error {
	srv.status = apiv1.ServiceStatus_SERVICE_STATUS_STARTING
	slog.InfoContext(ctx, "Starting service", srv.logGroup())

	srv.cmd = exec.CommandContext(ctx, srv.command, srv.args...)
	srv.cmd.Env = srv.environment
	srv.cmd.Stdout = os.Stdout
	srv.cmd.Stderr = os.Stderr

	if err := srv.cmd.Start(); err != nil {
		srv.status = apiv1.ServiceStatus_SERVICE_STATUS_FAILED
		slog.ErrorContext(ctx, "Failed to start service", srv.logGroup())
		return err
	}

	return nil
}

func (srv *service) stop(ctx context.Context) error {
	srv.status = apiv1.ServiceStatus_SERVICE_STATUS_STOPPING
	slog.InfoContext(ctx, "Stopping service", srv.logGroup())

	if err := srv.cmd.Process.Kill(); err != nil {
		return err
	}
	return nil
}

func (srv *service) logGroup() slog.Attr {
	return slog.Group(
		"service",
		"name", srv.name,
		"status", srv.status.String(),
		"command", srv.command,
		"args", srv.args,
		"environment", srv.environment,
	)
}
