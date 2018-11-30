package helper

import (
	"time"

	"context"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	micro_server "github.com/micro/go-micro/server"
	opWrapper "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	opConfig "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics/expvar"
	log "github.com/golang/glog"
)

type server struct {
	tracer  opentracing.Tracer
	service micro.Service
}

//NewServer return server wrapper for go-micro ecosystem
func NewServer(domain string, consulAddr string, opAddress string) *server {
	s := server{}
	s.initTracer(domain, opAddress)
	s.initService(domain, consulAddr)
	return &s
}

func (s *server) GetService() micro.Service { return s.service }

func (s *server) initService(domain string, consulAddr string) {
	s.service = grpc.NewService(
		micro.WrapHandler(LogHandlerWrapper()),
		micro.WrapHandler(opWrapper.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Name(domain), micro.Version("latest"),
		micro.Registry(registry.NewRegistry(registry.Addrs(consulAddr))),
		//micro.WrapHandler(ValidateHandlerWrapper()),
	)
	s.service.Server().Init(micro_server.Wait(true))
	s.service.Init()
}

func (s *server) initTracer(domain string, opAddress string) {
	s.tracer, _, _ = opConfig.Configuration{
		ServiceName: domain,
		Sampler: &opConfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &opConfig.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  opAddress,
		}}.
		NewTracer(
		opConfig.Observer(rpcmetrics.NewObserver(
			expvar.NewFactory(10).Namespace(domain, nil),
			rpcmetrics.DefaultNameNormalizer)))
	log.Info("init tracer done")
	opentracing.SetGlobalTracer(s.tracer)
}

//LogHandlerWrapper log middleware wrapper for server
func LogHandlerWrapper() micro_server.HandlerWrapper {
	return func(h micro_server.HandlerFunc) micro_server.HandlerFunc {
		return func(ctx context.Context, req micro_server.Request, rsp interface{}) error {
			log.Infof("request: %v\n", req.Method())
			return h(ctx, req, rsp)
		}
	}
}
