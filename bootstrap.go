package kannri

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"buf.build/go/protoyaml"

	apiv1 "github.com/aita/kannri/kannripb/kannri/v1"
)

const currentVersion = "1"

func LoadBootstraps(filenames ...string) (*apiv1.Bootstrap, error) {
	bootstraps := make([]*apiv1.Bootstrap, 0, len(filenames))
	for _, filename := range filenames {
		bootstrap, err := ParseBootstrapFromFile(filename)
		if err != nil {
			return nil, err
		}
		bootstraps = append(bootstraps, bootstrap)
	}
	return MergeBootstraps(bootstraps...), nil
}

func ParseBootstrap(data []byte) (*apiv1.Bootstrap, error) {
	var bootstrap apiv1.Bootstrap
	err := protoyaml.Unmarshal(data, &bootstrap)
	if err != nil {
		return nil, err
	}

	if bootstrap.Version != "" && bootstrap.Version != currentVersion {
		return nil, fmt.Errorf("unsupported version: %s", bootstrap.Version)
	}
	if bootstrap.Version == "" {
		bootstrap.Version = currentVersion
	}

	return &bootstrap, nil
}

func ParseBootstrapFromFile(filename string) (bootstrap *apiv1.Bootstrap, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return ParseBootstrap(data)
}

func MergeBootstraps(bootstraps ...*apiv1.Bootstrap) *apiv1.Bootstrap {
	merged := &apiv1.Bootstrap{}
	for _, bootstrap := range bootstraps {
		if bootstrap == nil {
			continue
		}
		if merged.Version == "" {
			merged.Version = bootstrap.Version
		}
		merged.Services = append(merged.Services, bootstrap.Services...)
	}
	return merged
}

func ApplyBootstrap(ctx context.Context, manager *Manager, bootstrap *apiv1.Bootstrap) error {
	for _, service := range bootstrap.Services {
		if _, err := manager.CreateService(ctx, service); err != nil {
			return err
		}
	}
	return nil
}
