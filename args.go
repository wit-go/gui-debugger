package debugger

// initializes logging and command line options

import (
	arg "github.com/alexflint/go-arg"
	"go.wit.com/log"
)

var INFO log.LogFlag
var POLL log.LogFlag
var BUG log.LogFlag
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

	INFO.B = false
	INFO.Name = "INFO"
	INFO.Subsystem = "bugger"
	INFO.Desc = "simple debugging Info()"
	INFO.Register()

	POLL.B = false
	POLL.Name = "POLL"
	POLL.Subsystem = "bugger"
	POLL.Desc = "watch the debugger poll things"
	POLL.Register()
}
