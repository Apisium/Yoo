package yoo

import (
  "fmt"
  "time"
)

func copyMap(obj *Variables) *Variables {
  oldMap := *obj
  newMap := make(Variables, len(oldMap))
  for key, value := range oldMap {
    newMap[key] = value
  }
  return &newMap
}

func trace() func() {
  start := time.Now()
  return func() {
    fmt.Printf("Cost: %s\n", time.Since(start))
  }
}
