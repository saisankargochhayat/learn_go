package main

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Entry {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Out = os.Stdout

	return l.WithFields(log.Fields{
		"serviceBuildFieldName": "sailogger",
		"Tag":                   "2",
	})
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

}

func main() {
	var loggerTest *logrus.Entry = GetLogger()
	loggerTest.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	loggerTest.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	loggerTest.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
