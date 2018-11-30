package wrapper

import (
	"context"
	"time"

	"github.com/whatvn/dqueue/protobuf"

	log "github.com/golang/glog"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
	op "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

const (
	Domain      = "DelayQueue"
	Consul      = "127.0.0.1:8500"
	OpenTracing = "127.0.0.1:8200"
)

type logWrapper struct {
	client.Client
}

func (c *logWrapper) Call(ctx context.Context, req client.Request,
	rsp interface{}, opts ...client.CallOption) error {
	log.Infof("[client][%v] request: %v; options: %v", time.Now().Format(time.StampMilli), req.Method(), opts)
	return c.Client.Call(ctx, req, rsp, opts...)
}

func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

type clientWrapper struct {
	tracer  opentracing.Tracer
	service micro.Service
}

// NewClient return client wrapper for go-micro ecosystem
func NewClient(domain string, consulAddr string, opts ...client.Option) *clientWrapper {
	log.Info("[Start InitClient]", domain, consulAddr)
	defer log.Info("[End InitClient]")

	// add options to client
	opts = append(opts, client.Retries(3))
	opts = append(opts, client.PoolSize(20))

	c := clientWrapper{}
	c.tracer = opentracing.GlobalTracer()
	c.initService(domain, consulAddr, opts...)
	log.Info("Init %s client done!", domain)
	return &c
}

func (c *clientWrapper) GetClient() client.Client { return c.service.Client() }

func (c *clientWrapper) GetService() micro.Service { return c.service }

func (c *clientWrapper) initService(domain string, consulAddr string, opts ...client.Option) {

	r := registry.NewRegistry(registry.Addrs(consulAddr))
	c.service = grpc.NewService(
		micro.Name(domain),
		micro.Registry(r),
		micro.Selector(selector.NewSelector(selector.Registry(r))),
		micro.WrapClient(
			logWrap,
			op.NewClientWrapper(opentracing.GlobalTracer()),
		),
	)
	c.GetClient().Init(opts...)
}

func NewDelayQueueClient() delayQueue.DelayQueueService {
	service := NewClient(Domain, Consul).GetService()
	return delayQueue.NewDelayQueueService(Domain, service.Client())
}
