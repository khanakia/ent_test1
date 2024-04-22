package natsresp

import (
	"fmt"
	"lace/natsutil"
	"lace/nlog"

	"github.com/nats-io/nats.go/micro"
)

type NatsMicroResp struct {
	req  micro.Request
	nlog nlog.Logger
}

// NatsResp pkg is used to create response for NatsMicro request
//
//	natsResp := natsresp.NewNatsMicroResp(req, config.Nlog)
//
//	natsResp.SendError(natsresp.WithError(fmt.Errorf("some error")))
//
//	natsResp.Send(PipelineRunReq{
//		ID: "2343",
//	})
func NewNatsMicroResp(req micro.Request, nlog nlog.Logger) *NatsMicroResp {
	return &NatsMicroResp{
		req:  req,
		nlog: nlog,
	}
}

func (mr NatsMicroResp) Send(data interface{}) {
	if err := mr.req.RespondJSON(natsutil.CreateRespWithData(data)); err != nil {
		mr.nlog.Sugar.Errorw("FATAL NATS ERROR: respondWithData - " + err.Error())
	}
}

func (mr NatsMicroResp) SendError(opts ...ErrOption) {
	cfg := errConfig{
		statusCode: 400,
		errCode:    "error",
		err:        fmt.Errorf("server error"),
	}

	cfg.errOptions(opts...)

	publishErr := mr.req.RespondJSON(natsutil.CreateRespWithErr(cfg.err.Error(), cfg.errCode, cfg.statusCode))

	if publishErr != nil {
		mr.nlog.Sugar.Errorw("NATS ERROR: respondWithError - " + publishErr.Error())
	}
}

type errConfig struct {
	statusCode uint
	errCode    string
	err        error
}

func (c *errConfig) errOptions(opts ...ErrOption) {
	for _, opt := range opts {
		opt(c)
	}
}

type ErrOption func(*errConfig)

func WithStatusCode(statusCode uint) ErrOption {
	return func(c *errConfig) {
		c.statusCode = statusCode
	}
}

func WithErrorCode(code string) ErrOption {
	return func(c *errConfig) {
		c.errCode = code
	}
}

func WithError(err error) ErrOption {
	return func(c *errConfig) {
		c.err = err
	}
}
