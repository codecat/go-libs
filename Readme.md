# Codecat Go libs

[![Build Status](https://travis-ci.org/codecat/go-libs.svg?branch=master)](https://travis-ci.org/codecat/go-libs)

These are the Go libraries I've written that I use in my Go applications.

To use these, install them using `go get`, for example:

```
$ go get github.com/codecat/go-libs/log
$ go get github.com/codecat/go-libs/settings
```

## Libraries

### `log`

* Optionally, first call `log.Open(log.CatTrace, log.CatFatal)` where the first parameter is the minimum log level and the second parameter is the maximum log level.
* Then call any of these to actually write to the log:
  * `log.Trace(format string, args ...interface{})`
  * `log.Debug(format string, args ...interface{})`
  * `log.Info(format string, args ...interface{})`
  * `log.Warn(format string, args ...interface{})`
  * `log.Error(format string, args ...interface{})`
  * `log.Fatal(format string, args ...interface{})`

### `settings`

* Instantiate a structure:
  ```go
  var config struct {
    Foo string
    Bar string
  }
  ```
* Call `settings.Load(filename string, out interface{})` on the structure:
  ```go
  func main() {
    err := settings.Load("config.yaml", &config)
    if err != nil {
      fmt.Println(err.Error())
    }
  }
  ```
* After modifying programmatically, save using `settings.Save(filename string, in interface{})`:
  ```go
  err := settings.Save("config.yaml", &config)
  if err != nil {
    fmt.Println(err.Error())
  }
  ```

## License

These libraries are licensed under the MIT license.

Copyright © 2017 github.com/codecat

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the “Software”), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.
