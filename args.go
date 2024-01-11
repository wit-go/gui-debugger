package debugger

// initializes logging and command line options

import (
	"go.wit.com/arg"
	"go.wit.com/log"
)

var INFO *log.LogFlag
var POLL *log.LogFlag
var CHAN *log.LogFlag
var WARN *log.LogFlag

var argDebugger ArgsDebugger

// This struct can be used with the go-arg package
type ArgsDebugger struct {
	Debugger bool `arg:"--debugger" help:"open the debugger window"`
}

// returns true if --gui-debug was passed from the command line
func ArgDebug() bool {
	return argDebugger.Debugger
}

func init() {
	arg.Register(&argDebugger)

	full := "go.wit.com/gui/debugger"
	short := "bugger"

	INFO = log.NewFlag("INFO", false, full, short, "simple debugging Info()")
	POLL = log.NewFlag("POLL", false, full, short, "watch the debugger poll things")
	CHAN = log.NewFlag("CHAN", true,  full, short, "chan() test code output")
	WARN = log.NewFlag("WARN", true,  full, short, "should warn the user")
}
