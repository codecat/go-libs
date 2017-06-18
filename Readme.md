# Codecat Go libs

These are the Go libraries I've written that I use in my Go applications.

To use these, install them using `go get`, for example:

```
$ go get github.com/codecat/go-libs/log
```

## Libraries

### `log`

* First call `log.Open(log.CatTrace, log.CatFatal)` where the first parameter is the minimum log level and the second parameter is the maximum log level.
* You can then call any of these:
  * `log.Trace(format string, args ...interface{})`
  * `log.Debug(format string, args ...interface{})`
  * `log.Info(format string, args ...interface{})`
  * `log.Warn(format string, args ...interface{})`
  * `log.Error(format string, args ...interface{})`
  * `log.Fatal(format string, args ...interface{})`
