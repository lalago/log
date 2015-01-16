package log

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	// external writer
	var f, err = os.Create("./writer.out")
	if err != nil {
		Log.Error("something %s", "wrong")
	}
	var l = New(f, DEBUG, "log test")
	l.Debug("I love debug log!")
	l.Debug("I love %s log!", "debug")

	l.Info("I love %s log!", "info")

	l.Warning("I love %s log!", "warning")

	l.Error("I love %s log!", "error")

	// Log (using os.Stderr)
	var stdlog = New(os.Stderr, DEBUG, "")
	stdlog.Debug("I love debug log!")
	stdlog.Debug("I love %s log!", "debug")
	stdlog.Info("I love %s log!", "info")
	stdlog.Warning("I love %s log!", "warning")
	stdlog.Error("I love %s log!", "error")
}
