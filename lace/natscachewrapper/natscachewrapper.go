package cache_natsapi

import (
	"fmt"
	"lace/cache"
	"lace/natso"
	"lace/natsutil"

	"github.com/nats-io/nats.go"
)

// This packaeg is receive request so this will behave more like an api and will be included in cache microservices

const (
	NATS_CACHE_PUT   = "cache.put"
	NATS_CACHE_GET   = "cache.get"
	NATS_CACHE_DEL   = "cache.del"
	NATS_CACHE_FLUSH = "cache.flush"
)

type config struct {
	Natso natso.Natso
	Cache cache.Store
}

type NatsCacheWrapper struct {
	config
}

type res struct {
	Code string
}

func New(config config) *NatsCacheWrapper {
	ec := config.Natso.GetEncodedConn()
	PutSubs(ec, config.Cache)
	GetSubs(ec, config.Cache)
	DelSubs(ec, config.Cache)
	FlushSubs(ec, config.Cache)

	return &NatsCacheWrapper{
		config: config,
	}
}

type CachePutReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Ttl   int    `json:"ttl"`
}

func PutSubs(ec *nats.EncodedConn, cache cache.Store) {
	ec.Subscribe(NATS_CACHE_PUT, func(subj, reply string, msg CachePutReq) {
		fmt.Println(msg)
		cache.Put(msg.Key, msg.Value, msg.Ttl)
	})
}

type CacheGetReq struct {
	Key string `json:"key"`
}

func GetSubs(ec *nats.EncodedConn, cache cache.Store) {
	ec.Subscribe(NATS_CACHE_GET, func(subj, reply string, msg CacheGetReq) {
		ec.Publish(reply, natsutil.CreateRespWithData(cache.Get(msg.Key)))
	})
}

type CacheDelReq struct {
	Key string `json:"key"`
}

func DelSubs(ec *nats.EncodedConn, cache cache.Store) {
	ec.Subscribe(NATS_CACHE_DEL, func(subj, reply string, msg CacheDelReq) {
		cache.Del(msg.Key)
	})
}

func FlushSubs(ec *nats.EncodedConn, cache cache.Store) {
	ec.Subscribe(NATS_CACHE_FLUSH, func(subj, reply string, msg interface{}) {
		cache.Flush()
	})
}
