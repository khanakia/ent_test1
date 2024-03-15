// Code generated from Pkl module `MainConfig`. DO NOT EDIT.
package mainconfig

import (
	"context"

	"configmgr/gen/appconfig"
	"configmgr/gen/cacheconfig"
	"github.com/apple/pkl-go/pkl"
)

type MainConfig struct {
	// The hostname that this server responds to.
	App *appconfig.AppConfig `pkl:"app"`

	Cache *cacheconfig.CacheConfig `pkl:"cache"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a MainConfig
func LoadFromPath(ctx context.Context, path string) (ret *MainConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a MainConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*MainConfig, error) {
	var ret MainConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
