# log
logging lib for golang

##Usage:
**Directly**
```go
import (
  . "github.com/lalago/log"
)

func main() {
  Log.Debug("Cas you are amazing, just the way you are.")
}
```
**Make a new logger**
```go

var l = New(os.Stderr, DEBUG, "lalago")
l.Info("Hello, %s", "lalago")

```
