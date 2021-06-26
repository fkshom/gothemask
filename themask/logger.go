package themask

import (
	"log"
	"os"

	"github.com/k0kubun/pp"
)

type Logger interface {
	Warnf(string, ...interface{})
	Debugf(string, ...interface{})
	Debugp(string, ...interface{})
}

var logger Logger = &BasicLogger{
	Logger: log.New(os.Stderr, "", log.LstdFlags),
}


const (
	DEBUG int = iota
	WARN
	SILENT
)
var level int = WARN

func SetLogger(l Logger) {
	logger = l
}

func SetLevel(_level int) {
	level = _level
}

type BasicLogger struct {
	Logger *log.Logger
}

func (bl *BasicLogger) Warnf(format string, v ...interface{}) {
	if level <= WARN {
		bl.Logger.Printf(format, v...)
	}
}

func (bl *BasicLogger) Debugp(text string, v ...interface{}) {
	bl.Debugf(text)
	pptext := pp.Sprint(v...)
	bl.Debugf(pptext)
}
func (bl *BasicLogger) Debugf(format string, v ...interface{}) {
	if level <= DEBUG {
		bl.Logger.Printf(format, v...)
	}
}
