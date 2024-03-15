package configmgr

import (
	"configmgr/gen/appconfig"
	"configmgr/gen/cacheconfig"
	"context"
)

func All() {

}
func GetAppConfig() *appconfig.AppConfig {
	cfg, err := appconfig.LoadFromPath(context.Background(), "../config/pkl/local/app.pkl")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("I'm running on host %s\n", cfg.Hostname)
	return cfg
}

func GetCacheConfig() *cacheconfig.CacheConfig {
	cfg, err := cacheconfig.LoadFromPath(context.Background(), "../config/pkl/local/cache.pkl")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("I'm running on host %s\n", cfg.Hostname)
	return cfg
}
