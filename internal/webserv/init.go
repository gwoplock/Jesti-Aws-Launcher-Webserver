package webserv

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.Info("Starting web server")
	http.Handle("/", http.FileServer(http.Dir("./www")))

	logrus.Info("web server started")

	err := http.ListenAndServe(":8080", nil) //TODO change to 80 for prod
	logrus.Fatal(err)
}
