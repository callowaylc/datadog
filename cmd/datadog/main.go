package main

import(
	"os"
	_ "log"
	_ "fmt"
	_ "github.com/DataDog/datadog-go/statsd"
  a "github.com/callowaylc/datadog-go-cli"
)

func init() {
  a.InitOptions()
}

func main() {
	os.Exit(0)
}
