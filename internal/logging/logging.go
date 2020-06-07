package logging

import (
	"fmt"

	//"github.com/Freman/eventloghook"
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	logrus.SetLevel(logrus.TraceLevel)

	hook, err := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename:   "./system.log",
			MaxSize:    100,
			MaxBackups: 1,
			Compress:   false,
		},
		logrus.InfoLevel,
		&logrus.JSONFormatter{},
		&lumberjackrus.LogFileOpts{},
	)
	if err != nil {
		panic(fmt.Errorf("Cant create file hook: %s", err))
	}

	logrus.AddHook(hook)

	//TODO: add windows event log logging
	// elog, err = eventlog.Open("jetsi launch")
	// if err != nil {
	// 	panic(fmt.Errorf("Cant create windows event log hook: %s", err))
	// }
	// defer elog.Close()
	// log.Hooks.Add(eventloghook.NewHook(elog))

	logrus.Info("Logging started")
}
