package util

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmlogrus"
)

var log = &logrus.Logger{
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "function.name",
		},
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	},
	ReportCaller: true,
}

func init() {
	if os.Getenv("FILEBEAT_ADDRESS") == "" {
		log.SetOutput(os.Stdout)
		return
	}

	conn, err := net.Dial("udp", os.Getenv("FILEBEAT_ADDRESS"))
	if err != nil {
		fmt.Println(err)
	}

	mw := io.MultiWriter(os.Stdout, conn)
	log.SetOutput(mw)
}

// LogEntry is Log caller to initiate logrus with APM
func LogEntry(ctx context.Context) *logrus.Entry {
	return log.WithFields(apmlogrus.TraceContext(ctx))
}
