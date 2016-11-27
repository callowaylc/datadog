package cli

import (
  "os"
  "time"
  "fmt"
  log "github.com/Sirupsen/logrus"
)

type LogFields log.Fields

func InitLogs() {
  log.SetFormatter(&log.JSONFormatter{})
  log.SetOutput(os.Stderr)
  log.SetLevel(log.InfoLevel)
}

func Logs(message string, fields LogFields) {
  options := GetOptions()
  if options.Logs {
    if fields == nil {
      fields = LogFields{}
    }
    fields["time"] = time.Now()  
    log.WithFields(map[string]interface{}(fields)).Info(fmt.Sprintf("%v", message))
  }
}
