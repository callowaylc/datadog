package datadog_go_cli

import(
  _ "os"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
  Address string `short:"a" long:"address" description:"Statsd agent host:port" default:"127.0.0.1:8125"`
}
var o Options

func InitOptions() {
  flags.Parse(&o)

  //if err != nil {
  //  panic(err)
  //  os.Exit(1)
  //}
}

func GetOptions() (Options, error) {
  return o, nil
}
