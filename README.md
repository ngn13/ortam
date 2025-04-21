# ortam | simple config library for Go

![test workflow status](https://img.shields.io/github/actions/workflow/status/ngn13/ortam/test.yml?label=tests)

tiny and simple, environment variable based configuration library for Go with zero dependencies

### install
you can add ortam to your project using the `go get` command:
```bash
go get -u github.com/ngn13/ortam
```

### usage
first define a structure containing all the configuration options you want:
```go
type MyConfig struct {
    ListenAddr string
    Debug bool
}
```
your structure can only contain integer types (`uint...`, `int...` and `float...` types),
the `bool` type, `string` type, `*url.URL` type or `time.Duration` type, any other type
will be ignored, you can also use other structures to build up your configuration structure

here are example values for all the different types:

| type          | values                                                          | note                                                                 |
| ------------- | --------------------------------------------------------------- | -------------------------------------------------------------------- |
| int...        | 42, -12, 1098 etc.                                              | value should fit in given type                                       |
| uint...       | 42, 1098, 0 etc.                                                | value should fit in given type                                       |
| float...      | 3.3, 42.983, 99.99 etc.                                         | value should fit in given type                                       |
| bool          | true, 1, false, 0                                               |                                                                      |
| string        | "hello world", "127.0.0.1:8080" etc.                            |                                                                      |
| *url.URL      | "gemini://geminiprotocol.net", "ftp://ftp.gnu.org/gnu/gcc" etc. | value should be a full URL (scheme and host is required)             |
| time.Duration | "1h", "42m", "30s" etc.                                         | value requires a "h" (hour), a "m" (minute) or a "s" (second) suffix |

after defining your configuration variable using the new sturcture you created (optionally
with the default values), you can use `ortam.Load()` to load the configuration from the
environment variables
```go
var config MYConfig = MyConfig{
    // by default, config.Debug is false
    Debug: false,
}
ortam.Load(&config)
```
simple, right? now in this case `ListenAddr` will be loaded from the `LISTEN_ADDR` environment
variable, and `Debug` will be loaded from the `DEBUG` environment variable, so the environment
variable is named after the name of the member

if you want you can also define custom names for different members:
```go
type MyConfig struct {
    // use the environment var "ADDR"
    ListenAddr string `ortam:"ADDR"`
    Debug bool
}
```
you can also specify a prefix for all the environment variables:
```go
ortam.Load(&config, "APP")
```
now `ListenAddr` will be loaded from `APP_ADDR` and `Debug` will be loaded from `APP_DEBUG`

### license notice
this library is licensed under AGPL-3.0, meaning if you use this library (or a modified version of it)
in your program, **you have to** release your program as free (as in freedom), learn more about free
software [here](https://www.gnu.org/philosophy/free-sw.html)
