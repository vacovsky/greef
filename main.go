package main

import (
	"sync"

	"github.com/vacovsky/greef/data"
	"github.com/vacovsky/greef/monitor"
	"github.com/vacovsky/greef/serve"
)

// WG is the waitgroup for the goroutines
var WG sync.WaitGroup

func init() {
	WG = sync.WaitGroup{}
}

func main() {
	data.MigrateDataSchema()
	go serve.Serve()
	go monitor.Start()
	WG.Add(1)
	WG.Wait()
}
