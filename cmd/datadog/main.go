package main

import(
	"os"
	"log"
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
)

func init() {

}

func main() {
  c, err := statsd.New("127.0.0.1:8125")
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("%+v\n", c)
	os.Exit(0)
}
