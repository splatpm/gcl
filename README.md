

[![Build Status](https://travis-ci.org/splatpm/gcl.svg?branch=master)](https://travis-ci.org/splatpm/gcl)
## Gout (Go OUTput is a library for handling ascii color and video attributes.)

#### Formatting functions

* HumanSize(size int64) string
* HumanTimeColon(secs int64) string
* HumanTimeConcise(secs int64) string
* HumanTimeParse(secs int64) map[string]int64

#### Output functions

Output functions will optionally also push the messages to a logfile.
Look below for examples of how to use a logfile and how to set verbosity
or prompt options for the console.

* Info(string, args ...interface{})
* Debug(string, args ...interface{})
* Warn(string, args ...interface{})
* Error(string, args ...interface{})
* Status(string, args ...interface{})

#### String type methods

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
