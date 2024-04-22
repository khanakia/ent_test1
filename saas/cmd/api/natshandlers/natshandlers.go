package natshandlers

import (
	"fmt"
	"lace/natso"
	"saas/pkg/app"
	"time"

	"github.com/invopop/jsonschema"
	"github.com/nats-io/nats.go/micro"
)

type Config struct {
	Natso natso.Natso
	// EntClient *ent.Client
	// Nlog      nlog.Logger
}

func SchemaFor(t any) string {
	schema := jsonschema.Reflect(t)
	data, _ := schema.MarshalJSON()
	return string(data)
}

type runFunc func(req micro.Request)

func New(config Config) {

	plugin := app.GetPlugins()

	nc := config.Natso.GetConn()
	// ec := config.Natso.GetEncodedConn()
	// va := validator.New()

	srv, err := micro.AddService(nc, micro.Config{
		Name:    "anago", // just a name
		Version: "0.0.1",
	})

	if err != nil {
		plugin.Nlog.Sugar.Fatalf("FATAL ADD NATS SERVICE - %s", err.Error())
	}

	// fmt.Println(srv)

	m := srv.AddGroup("anago")

	m.AddEndpoint(
		"ping",
		micro.HandlerFunc(func(req micro.Request) {
			time.Sleep(time.Second * 10)
			fmt.Println("test")
			req.Respond(req.Data())
		}),
		micro.WithEndpointMetadata(map[string]string{
			"format": "application/json",
		}),
	)
}

func runRequest(fn runFunc) func(req micro.Request) {
	return func(req micro.Request) {
		go func() {
			fn(req)
		}()
	}
}
