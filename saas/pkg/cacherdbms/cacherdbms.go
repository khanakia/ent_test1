package cacherdbms

import (
	"context"
	"encoding/json"
	"saas/gen/ent"
	"saas/gen/ent/kache"
	"time"
)

type Rdbms struct {
	client *ent.Client
}

func (pkg Rdbms) Version() string {
	return "0.01"
}

func New(client *ent.Client) *Rdbms {
	rdbms := &Rdbms{client: client}
	return rdbms
}

func (a *Rdbms) Put(key string, val interface{}, ttl int) (bool, error) {
	p, err := json.Marshal(val)
	if err != nil {
		return false, err
	}

	expire := int(time.Now().UnixNano()/int64(time.Second) + int64(ttl))

	entity, err := a.client.Kache.Query().Where(kache.Key(key)).Only(context.Background())
	if err != nil && ent.IsNotFound(err) {
		_, err := a.client.Kache.Create().SetKey(key).SetValue(string(p)).SetExpires(expire).Save(context.Background())
		if err != nil {
			return false, err
		}
	} else {
		entity.Update().SetValue(string(p)).SetExpires(expire).Save(context.Background())
	}

	return true, nil
}

func (a *Rdbms) Get(key string) interface{} {
	entity, err := a.client.Kache.Query().Where(kache.Key(key)).Only(context.Background())
	if err != nil {
		return nil
	}

	now := int(time.Now().UnixNano() / int64(time.Second))
	if now > entity.Expires {
		a.Del(key)
		return nil
	}

	var v interface{}
	err = json.Unmarshal([]byte(entity.Value), &v)
	if err != nil {
		return entity.Value
	}

	return v
}

func (a *Rdbms) Del(key string) {
	a.client.Kache.Delete().Where(kache.Key(key)).Exec(context.Background())
}

func (a *Rdbms) Flush() {
	a.client.ExecContext(context.Background(), "Truncate TABLE "+kache.Table)
}
