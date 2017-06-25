package log

import "fmt"
import "time"
import "github.com/fatih/color"

type LogConfig struct {
	MinLevel, MaxLevel int
}

var Config = LogConfig{
	CatTrace,
	CatFatal,
}

const (
	CatTrace = iota
	CatDebug
	CatInfo
	CatWarn
	CatError
	CatFatal
)

func nameForCategory(cat int) string {
	switch cat {
	case CatTrace: return "TRACE"
	case CatDebug: return "DEBUG"
	case CatInfo: return "INFO"
	case CatWarn: return "WARN"
	case CatError: return "ERROR"
	case CatFatal: return "FATAL"
	}
	return "???"
}

func colorForCategory(cat int) color.Attribute {
	switch cat {
	case CatTrace: return color.FgWhite
	case CatDebug: return color.FgHiCyan
	case CatInfo: return color.FgHiWhite
	case CatWarn: return color.FgHiMagenta
	case CatError: return color.FgHiYellow
	case CatFatal: return color.FgHiRed
	}
	return color.FgWhite
}

func writeLine(cat int, s string, args ...interface{}) {
	if !passesFilter(cat) {
		return
	}

	tmNow := time.Now()
	tmFormatted := fmt.Sprintf("%02d:%02d:%02d.%03d", tmNow.Hour(), tmNow.Minute(), tmNow.Second(), tmNow.Nanosecond() / 1000000)
	line := fmt.Sprintf(s, args...)

	catName := nameForCategory(cat)
	catColor := colorForCategory(cat)

	color.Set(catColor)
	fmt.Printf("[%s] [% 5s] %s\n", tmFormatted, catName, line)
	color.Unset()
}

func passesFilter(cat int) bool {
	return cat >= Config.MinLevel && cat <= Config.MaxLevel
}

func Open(minLevel, maxLevel int) {
	Config.MinLevel = minLevel
	Config.MaxLevel = maxLevel
}

func Trace(s string, args ...interface{}) {
	writeLine(CatTrace, s, args...)
}

func Debug(s string, args ...interface{}) {
	writeLine(CatDebug, s, args...)
}

func Info(s string, args ...interface{}) {
	writeLine(CatInfo, s, args...)
}

func Warn(s string, args ...interface{}) {
	writeLine(CatWarn, s, args...)
}

func Error(s string, args ...interface{}) {
	writeLine(CatError, s, args...)
}

func Fatal(s string, args ...interface{}) {
	writeLine(CatFatal, s, args...)
}
