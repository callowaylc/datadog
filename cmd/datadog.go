package main

import(
	"os"
  _ "time"
	_ "log"
	_ "fmt"
	"github.com/DataDog/datadog-go/statsd"
  a "github.com/callowaylc/datadog"
)

func init() {
  a.InitLogs()
  a.InitOptions()

  a.Logs("passed options", a.LogFields{ 
    "options": a.GetOptions(), 
  })
}

func main() {
  options := a.GetOptions()

  // Create the client
  c, err := statsd.New(options.Address)
  if err != nil {
    a.Logs("failed to open connection to datadog aggent", a.LogFields{
      "error": err,
      "host": options.Address,
    })
    panic(err)
  }

  switch options.Type {
  case "Gauge":
    err = c.Gauge(
      options.Name, options.ValueFloat, options.Tags, options.Rate,
    )
  case "Histogram":
    err = c.Histogram(
      options.Name, options.ValueFloat, options.Tags, options.Rate,
    )
  case "Timing":
    panic("not implemented!")
    //err = c.Timing(
    //  options.Name, options.ValueString.(time.Duration), options.Tags, options.Rate,
    //)
  case "TimeInMilliseconds ":
    err = c.TimeInMilliseconds(
      options.Name, options.ValueFloat, options.Tags, options.Rate,
    )        
  case "ServiceCheck":
    panic("not implemented!")
    //err := c.Histogram(
    //  options.Name, options.Value.(float64), options.Tags, options.Rate
    //)    
  case "Set":
    err = c.Set(
      options.Name, options.ValueString, options.Tags, options.Rate,
    )    
  case "Incr":
    err = c.Incr(
      options.Name, options.Tags, options.Rate,
    )    
  case "Decr":
    err = c.Decr(
      options.Name, options.Tags, options.Rate,
    )     
  case "Count":
    err = c.Count(
      options.Name, options.ValueInt, options.Tags, options.Rate,
    )    
  }

  if err != nil {
    a.Logs("failed to send stat", a.LogFields{
      "error": err,
    })
  }

	os.Exit(0)
}
