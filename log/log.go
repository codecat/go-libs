package log

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Config contains instructions on how to write log to stdout
type Config struct {
	MinLevel        int
	MaxLevel        int
	Timestamp       bool
	TimestampFormat string
	Category        bool
	CategoryFormat  string
}

// CurrentConfig is the configuration currently used for logging
var CurrentConfig = Config{
	MinLevel:        CatTrace,
	MaxLevel:        CatFatal,
	Timestamp:       true,
	TimestampFormat: "%02d:%02d:%02d.%03d | ",
	Category:        true,
	CategoryFormat:  "% 5s | ",
}

var logLock sync.Mutex

const (
	// CatTrace is the level for trace messages
	CatTrace = iota
	// CatDebug is the level for debug messages
	CatDebug
	// CatInfo is the level for info messages
	CatInfo
	// CatWarn is the level for warn messages
	CatWarn
	// CatError is the level for error messages
	CatError
	// CatFatal is the level for fatal messages
	CatFatal
)

func nameForCategory(cat int) string {
	switch cat {
	case CatTrace:
		return "TRACE"
	case CatDebug:
		return "DEBUG"
	case CatInfo:
		return "INFO"
	case CatWarn:
		return "WARN"
	case CatError:
		return "ERROR"
	case CatFatal:
		return "FATAL"
	}
	return "???"
}

func colorForCategory(cat int) color.Attribute {
	switch cat {
	case CatTrace:
		return color.FgWhite
	case CatDebug:
		return color.FgHiCyan
	case CatInfo:
		return color.FgHiWhite
	case CatWarn:
		return color.FgHiMagenta
	case CatError:
		return color.FgHiYellow
	case CatFatal:
		return color.FgHiRed
	}
	return color.FgWhite
}

func writeLine(cat int, s string, args ...interface{}) {
	if !passesFilter(cat) {
		return
	}

	logLock.Lock()
	defer logLock.Unlock()

	fullOutput := fmt.Sprintf(s, args...)
	lines := strings.Split(fullOutput, "\n")
	for _, line := range lines {
		catColor := colorForCategory(cat)
		color.Set(catColor)

		if CurrentConfig.Timestamp {
			tmNow := time.Now()
			fmt.Printf(CurrentConfig.TimestampFormat, tmNow.Hour(), tmNow.Minute(), tmNow.Second(), tmNow.Nanosecond()/1000000)
		}

		if CurrentConfig.Category {
			catName := nameForCategory(cat)
			fmt.Printf(CurrentConfig.CategoryFormat, catName)
		}

		color.Unset()

		fmt.Printf("%s\n", line)
	}
}

func passesFilter(cat int) bool {
	return cat >= CurrentConfig.MinLevel && cat <= CurrentConfig.MaxLevel
}

// Trace log message
func Trace(s string, args ...interface{}) {
	writeLine(CatTrace, s, args...)
}

// Debug log message
func Debug(s string, args ...interface{}) {
	writeLine(CatDebug, s, args...)
}

// Info log message
func Info(s string, args ...interface{}) {
	writeLine(CatInfo, s, args...)
}

// Warn log message
func Warn(s string, args ...interface{}) {
	writeLine(CatWarn, s, args...)
}

// Error log message
func Error(s string, args ...interface{}) {
	writeLine(CatError, s, args...)
}

// Fatal log message
func Fatal(s string, args ...interface{}) {
	writeLine(CatFatal, s, args...)
}
