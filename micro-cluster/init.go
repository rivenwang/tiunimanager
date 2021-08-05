package main

import (
	"crypto/tls"

	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/transport"
	"github.com/pingcap/tiem/library/firstparty/config"
	"github.com/pingcap/tiem/library/secondparty/libtiup"
	"github.com/pingcap/tiem/library/thirdparty/logger"
	"github.com/pingcap/tiem/library/thirdparty/tracer"
	cluster "github.com/pingcap/tiem/micro-cluster/proto"
	"github.com/pingcap/tiem/micro-cluster/service"
	"github.com/pingcap/tiem/micro-cluster/service/tenant/adapt"
	dbclient "github.com/pingcap/tiem/micro-metadb/client"
)

// Global LogRecord object
var log *logger.LogRecord

func initConfig() {
	config.InitForMonolith()
}

func initLogger() {
	log = logger.GetLogger()
	service.InitClusterLogger()

	log.Debug("init logger completed!")
}

func initClusterOperator() {
	libtiup.MicroInit("./tiupmgr/tiupmgr", "tiup", "")
}

func initService() {
	cert, err := tls.LoadX509KeyPair(config.GetCertificateCrtFilePath(), config.GetCertificateKeyFilePath())
	if err != nil {
		log.Fatal(err)
		return
	}
	tlsConfigPtr := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	srv1 := micro.NewService(
		micro.Name(service.TiEMClusterServiceName),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapClient(opentracing.NewClientWrapper(tracer.GlobalTracer)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer.GlobalTracer)),
		micro.Transport(transport.NewHTTPTransport(transport.Secure(true), transport.TLSConfig(tlsConfigPtr))),
		micro.Registry(etcd.NewRegistry(registry.Addrs(config.GetRegistryAddress()...))),
	)

	srv1.Init()

	cluster.RegisterClusterServiceHandler(srv1.Server(), new(service.ClusterServiceHandler))

	if err := srv1.Run(); err != nil {
		log.Fatal(err)
	}

	srv2 := micro.NewService(
		micro.Name(service.TiEMManagerServiceName),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapClient(opentracing.NewClientWrapper(tracer.GlobalTracer)),
		micro.WrapHandler(opentracing.NewHandlerWrapper(tracer.GlobalTracer)),
		micro.Transport(transport.NewHTTPTransport(transport.Secure(true), transport.TLSConfig(tlsConfigPtr))),
		micro.Registry(etcd.NewRegistry(registry.Addrs(config.GetRegistryAddress()...))),
	)
	srv2.Init()

	cluster.RegisterTiEMManagerServiceHandler(srv2.Server(), new(service.ManagerServiceHandler))

	if err := srv2.Run(); err != nil {
		log.Fatal(err)
	}
}

func initClient() {
	dbclient.InitDBClient()
}

func initPort() {
	adapt.InjectionMetaDbRepo()
}
