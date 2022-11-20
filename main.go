package main

import (
	"github.com/adridevelopsthings/openapi-change-notification/api"
)

func main() {
	api.BuildNewContext()
	go api.StartDeprecationWorker(api.BuildContext())
	api.StartServer()
}
