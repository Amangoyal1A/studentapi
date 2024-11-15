package main

import (
	"github.com/Amangoyal1A/studentapi/app"
	"github.com/Amangoyal1A/studentapi/configs"
)

func main() {
	configs.LoadConfig()
	app.Start()
}
