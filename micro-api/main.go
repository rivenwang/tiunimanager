package main

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/pingcap-inc/tiem/docs"
	"github.com/pingcap-inc/tiem/library/firstparty/client"
	"github.com/pingcap-inc/tiem/library/framework"
	"github.com/pingcap-inc/tiem/micro-api/route"
	clusterPb "github.com/pingcap-inc/tiem/micro-cluster/proto"
)

// @title TiEM UI API
// @version 1.0
// @description TiEM UI API

// @contact.name zhangpeijin
// @contact.email zhangpeijin@pingcap.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1/
func main() {
	f := framework.InitBaseFrameworkFromArgs(framework.MetaDBService,
		initGinEngine,
	)

	f.PrepareClientClient(map[framework.ServiceNameEnum]framework.ClientHandler{
		framework.ClusterService: func(service micro.Service) error {
			client.ClusterClient = clusterPb.NewClusterService(string(framework.ClusterService), service.Client())
			return nil
		},
	})

	f.StartService()
}

func initGinEngine(d *framework.BaseFramework) error {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	route.Route(g)

	port := d.GetServiceMeta().ServicePort

	addr := fmt.Sprintf(":%d", port)

	if err := g.Run(addr); err != nil {
		d.GetLogger().Fatal(err)
	}

	return nil
}
