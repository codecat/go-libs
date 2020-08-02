# Codecat Go libs

[![Build Status](https://travis-ci.org/codecat/go-libs.svg?branch=master)](https://travis-ci.org/codecat/go-libs)

These are the Go libraries I've written that I use in my Go applications.

To use these, install them using `go get`, for example:

```
$ go get github.com/codecat/go-libs/log
$ go get github.com/codecat/go-libs/pacman
$ go get github.com/codecat/go-libs/settings
```

## Libraries

### `log`
Call any of these to actually write to the log:
* `log.Trace(format string, args ...interface{})`
* `log.Debug(format string, args ...interface{})`
* `log.Info(format string, args ...interface{})`
* `log.Warn(format string, args ...interface{})`
* `log.Error(format string, args ...interface{})`
* `log.Fatal(format string, args ...interface{})`

### `pacman`
Binary packing streams.

* To pack data directly to a file:
  ```go
  file, _ := os.Create("test.bin")
	packer, _ := pacman.NewPacker(file)
  ```
* To create a packer for a memory stream:
  ```go
  buffer := new(bytes.Buffer)
  packer, _ := pacman.NewPacker(buffer)
  packer.WriteUInt32(0)
  ```
* To unpack data from a file:
  ```go
  file, _ := os.Open("test.bin")
	unpacker, _ := pacman.NewUnpacker(file)
  ```
* To unpack data from a byte slice:
  ```go
  buffer := bytes.NewBuffer(slice)
  unpacker, _ := pacman.NewUnpacker(buffer)
  ```

Also has support for "blocks", which is useful if you want to be able to skip blocks for whatever reason.

```go
packer, _ := pacman.NewPacker(buffer)
b := packer.BeginBlock()
b.WriteInt32(1)
b.WriteInt32(2)
b.WriteInt32(3)
packer.EndBlock()
packer.WriteInt32(4)
```

```go
unpacker, _ := pacman.NewUnpacker(buffer)
b := unpacker.ReadBlock()
x := b.ReadInt32() // 1
y := b.ReadInt32() // 2
// We can ignore (or read later) the 3rd integer in the block
unpacker.ReadInt32() // 4
z := b.ReadInt32() // 3
```

### `settings`
Load and save a settings yaml file.

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
