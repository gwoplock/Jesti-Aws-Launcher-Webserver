package main

import (
	"fmt"

	"github.com/gwoplock/Jesti-Aws-Launcher-Webserver/internal/webserv"

	"github.com/gwoplock/Jesti-Aws-Launcher-Webserver/internal/logging"
)

func main() {
	fmt.Println("hello world")
	logging.Init()
	webserv.Init()
}
