package main

import (
	"github.com/chyuhung/my-dashboard/config"
	"github.com/chyuhung/my-dashboard/routes"
)

func main() {
	config.Init()
	routes.Setup()
}
