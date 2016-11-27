package cli

import(
  "os"
  _ "fmt"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
  Address string `default:"127.0.0.1:8125" short:"a" long:"address" description:"statsd agent host:port" default:"127.0.0.1:8125"`
  Name string `required:"true" short:"n" long:"name" description:"metric name"`
  Type string `required:"true" short:"t" long:"type" description:"metric type ex. counter, histogram, gauge"`
  Value interface{} `required:"true" short:"v" long:"value" description:"metric value for our given type"`
  Tags []string `long:"tag" description:"metric tags"`

}
var o Options

func InitOptions() {
  _, err := flags.Parse(&o)
  
  if flagsErr, ok := err.(*flags.Error); ok {
    switch flagsErr.Type {
    case flags.ErrHelp:
      os.Exit(0)  
    case flags.ErrTag:
      panic(err)
      os.Exit(1)
    default:
      os.Exit(2)
    }
  
  } else if err != nil {
    panic(err)
    os.Exit(1)
  }
}

func GetOptions() Options {
  return o
}

