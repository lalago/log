# log
logging lib for golang

###Usage:
**Features**
- colorizing level letter
 - Blue for Level Debug **D**
 - Green for Level Info **I**
 - Magenta for Level Warning **W**
 - Red for Level Error **E**
 

**Examples**
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
		log.Fatal("Creating log file fails")
	}

	var l = log.New(f, DEBUG, "lalago")
	l.Info("Hello, %s!", "lalago")
}


```
