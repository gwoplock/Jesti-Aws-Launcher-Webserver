package webserv

import (
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

var wg *sync.WaitGroup

func Init(lwg *sync.WaitGroup) *http.Server {
	logrus.Info("Starting web server")

	wg = lwg

	srv := &http.Server{Addr: ":8080"}
	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/ws", wsHandler)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed { //TODO change to 80 for prod
			logrus.Fatal("Web server error: ", err)
		}
	}()

	logrus.Info("web server started")

	return srv
}
