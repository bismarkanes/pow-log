# pow-log

pow-log is go package for logger custom engine with another additional logger which can be optionally enable. The default engine is using Go package ```log```. The other engine can be defined for example logger to file or to another storage.

## Installation

Use go get.

```bash
get get https://github.com/bismarkanes/pow-log
```

## Usage

```go
log := NewLogConsole()
log.Info("This is info %s", "Hello World")
log.Label("This is label %s", "MYLABEL", "Hello World")

useConsole := true
useFile := true
logf := NewLogFile("/tmp/mylog.log", useConsole, useFile)
log.Info("This is info %s", "Hello World")
log.Label("This is label %s", "MYLABEL", "Hello World")
```

## License

© Bismark
[MIT](https://choosealicense.com/licenses/mit/)
