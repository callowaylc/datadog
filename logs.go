package cli

import (
  "os"
  "time"
  "fmt"
  log "github.com/Sirupsen/logrus"
)

type LogFields log.Fields

func InitLog() {
  log.SetFormatter(&log.JSONFormatter{})
  log.SetOutput(os.Stderr)
  log.SetLevel(log.InfoLevel)
}

func Logs(message string, fields LogFields) {
  if fields == nil {
    fields = LogFields{}
  }
  fields["time"] = time.Now()  
  log.WithFields(map[string]interface{}(fields)).Info(fmt.Sprintf("%v", message)) 
}
