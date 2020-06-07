package main

import (
	"fmt"
	"sync"

	"github.com/gwoplock/Jesti-Aws-Launcher-Webserver/internal/webserv"

	"github.com/gwoplock/Jesti-Aws-Launcher-Webserver/internal/logging"
)

func main() {
	fmt.Println("hello world")
	logging.Init()

	wg := sync.WaitGroup{}

	webserv.Init(&wg)

	wg.Wait()
}
