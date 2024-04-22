// Code generated from Pkl module `AppConfig`. DO NOT EDIT.
package appconfig

import (
	"context"

	"configmgr/gen/appconfig/environment"
	"github.com/apple/pkl-go/pkl"
)

type AppConfig struct {
	AppName string `pkl:"appName"`

	AppKey string `pkl:"appKey"`

	// The hostname that this server responds to.
	Hostname string `pkl:"hostname"`

	// The port to listen on.
	Port string `pkl:"port"`

	// The environment to deploy to.
	Environment environment.Environment `pkl:"environment"`

	LogFile string `pkl:"logFile"`

	// The database connection for this application
	Database *Database `pkl:"database"`

	Nats *Nats `pkl:"nats"`

	Graphql *Graphql `pkl:"graphql"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AppConfig
func LoadFromPath(ctx context.Context, path string) (ret *AppConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AppConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*AppConfig, error) {
	var ret AppConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
