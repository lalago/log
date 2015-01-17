# log
logging lib for golang

##Usage:
**Directly**
```go
package main

import (
  "github.com/lalago/log"
  "os"
)


func main() {
	// new logger
	f, err := os.Create("log.out")
	if err != nil {
		// using default Stderr log
		log.Error("Creating log file fails")
	}

	var l = log.New(f, DEBUG, "lalago")
	l.Info("Hello, %s!", "lalago")
}


```
