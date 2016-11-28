package cli

import(
  "os"
  "fmt"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
  Address string `default:"127.0.0.1:8125" short:"a" long:"address" description:"statsd agent host:port"`
  Name string `required:"true" short:"n" long:"name" description:"metric name"`
  Type string `required:"true" short:"t" long:"type" description:"metric type ex. counter, histogram, gauge"`
  ValueFloat float64 `long:"value-float" description:"float based metric value"`
  ValueInt int64 `long:"value-int" description:"int based metric value"`
  ValueString string `long:"value-string" description:"string based metric value"`
  Rate float64 `default:"1" short:"r" long:"rate" description:"sample rate"`
  Tags []string `long:"tag" description:"metric tags"`
  Logs bool `long:"logs" description:"enables logging to stderr"`
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
    default:
      panic(err)
      os.Exit(2)
    }
  
  } else if err != nil {
    panic(err)
  }

  fmt.Println("asdf")
}

func GetOptions() Options {
  return o
}
