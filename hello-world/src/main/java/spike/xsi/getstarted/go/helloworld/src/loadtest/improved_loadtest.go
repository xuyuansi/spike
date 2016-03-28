package main

import (
  "fmt"
  "time"

  vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
  rate := uint64(30) // per second
  duration := 4 * time.Second
  targeter := vegeta.NewStaticTargeter(vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/CA",
  }, 
  vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/PH",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/MX",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/IN",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/HK",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/AR",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/GT",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "http://catalog.1.xoomapi.:1111/services/improved/FR",
  })
  attacker := vegeta.NewAttacker()

  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration) {
    metrics.Add(res)
  }
  metrics.Close()

  fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
