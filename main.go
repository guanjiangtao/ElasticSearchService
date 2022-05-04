package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"ElasticSearchService/tars-protocol/ElasticSearchServiceApp"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(WebApiImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("WebApiImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(ElasticSearchServiceApp.WebApi)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".WebApiObj")

	// Run application
	tars.Run()
}
