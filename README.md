

[![Build Status](https://travis-ci.org/splatpm/gcl.svg?branch=master)](https://travis-ci.org/splatpm/gcl)
# Gout (Go OUTput)

Gout is a library for handling ascii color and video attributes, output formatting,
program output and logging.

## Formatting functions

There are time and data size formatting functions available, HumanTimeParse() is
mostly only internally useful, but is useful enough that it was included in the
publicly exported API

* HumanSize(size int64) string
* HumanTimeColon(secs int64) string
* HumanTimeConcise(secs int64) string
* HumanTimeParse(secs int64) map[string]int64

#### Output functions

Output functions will optionally also push the messages to a logfile.
Look below for examples of how to use a logfile and how to set verbosity
or prompt options for the console.

* Info(string, args ...interface{})
* Debug(string, args ...interface{}election results     )
* Warn(string, args ...interface{})
* Error(string, args ...interface{})
* Status(string, args ...interface{})

*Example: Basic usage*

```go
package main

import (
  "git.thwap.org/splat/gout"
)

func main() {
  gout.Setup(true, false, true, "") // debug, quiet, verbose, logfile
  gout.Info("Test %s message", "info")
  gout.Debug("Test %s message", "debug")
  gout.Warn("Test %s %d", "warning", 1)
  gout.Error("error message")
}
```

*Example: Changing the output headers*
```go
package main

import (
  "git.thwap.org/splat/gout"
)

func main() {
  gout.Setup(true, false, true, "")
  gout.Info("Before")
  gout.Output.Prompts["info"] = String("###").Underline().Green()
  gout.Info("After")
}
```

*Example: Setting an output logfile*
```go
package main

import (
  "git.thwap.org/splat/gout"
)

func main() {
  gout.Setup(true, false, true, "/tmp/my-output.log")
  gout.Info("Log message")
  // prompts have nothing to do with logfile as only the type of promp
  // or the key (ie: info, debug, etc) is used.
}
```

#### String type methods

The String type is a alias for string, with the following methods.

* Black() String
* Red() String
* Green() String
* Yellow() String
* Blue() String
* Purple() String
* Cyan() String
* White() String
* Bold() String
* Underline() String
* Blink() String
* Reverse() String
* Conceal() String

*Example:*

```go
package main

import (
  "fmt"
  "git.thwap.org/splat/gout"
)

func main() {
  fmt.Println(gout.String("TEST").Bold().Red())
  fmt.Println(gout.String("TEST").Blink().Green())
}
```
