package main

import (
  "fmt"
  "time"

  vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
  rate := uint64(100) // per second
  duration := 4 * time.Second
  targeter := vegeta.NewStaticTargeter(vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/CA",
  }, 
  vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/PH",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/MX",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/IN",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/HK",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/AR",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/GT",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://catalog.1.xoomapi.:1111/services/FR",
  })
  attacker := vegeta.NewAttacker()

  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration) {
    metrics.Add(res)
  }
  metrics.Close()

  fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
