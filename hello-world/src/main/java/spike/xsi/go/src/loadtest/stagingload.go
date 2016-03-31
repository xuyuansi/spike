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
    URL:    "https://st1web501.atl.xoom.com:5567/services/CA",
  }, 
  vegeta.Target{
    Method: "GET",
    URL:    "https://st1web501.atl.xoom.com:5567/services/PH",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://st1web501.atl.xoom.com:5567/services/MX",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://st1web501.atl.xoom.com:5567/services/IN",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://st1web501.atl.xoom.com:5567/services/HK",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://st1web501.atl.xoom.com:5567/services/AR",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://st1web501.atl.xoom.com:5567/services/GT",
  },
  vegeta.Target{
    Method: "GET",
    URL:    "https://st1web501.atl.xoom.com:5567/services/FR",
  })
  attacker := vegeta.NewAttacker()

  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration) {
    metrics.Add(res)
  }
  metrics.Close()

  fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
