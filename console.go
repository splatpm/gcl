package gout

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	Output  output
	Winsize winsize
)

// data structures

type output struct {
	Prompts map[string]string
	Debug   bool
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// console info and utilities

func consInfo() winsize {
	ws := winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))
	if int(retCode) == -1 {
		panic(errno)
	}
	return ws
}

func padding(spaces int) string {
	pad := " "
	retv := ""
	for i := 0; i < spaces; i++ {
		retv = fmt.Sprintf("%s%s", retv, pad)
	}
	return retv
}

// Output functions
func consoleOutput(t string, e string, f string, args ...interface{}) {
	fmt.Printf("%s %s %s%s",
		Output.Prompts[t],
		fmt.Sprintf(f, args...),
		padding(int(Winsize.Col)-(7+len(fmt.Sprintf(f, args...)))),
		e)
}

func Info(f string, args ...interface{}) {
	consoleOutput("info", "\n", f, args)
}

func Debug(f string, args ...interface{}) {
	if Output.Debug {
		consoleOutput("debug", "\n", f, args)
	}
}

func Warn(f string, args ...interface{}) {
	consoleOutput("warn", "\n", f, args)
}

func Error(f string, args ...interface{}) {
	consoleOutput("error", "\n", f, args)
}

func Status(f string, args ...interface{}) {
	consoleOutput("status", "\r", f, args)
}

// Setup
func init() {
	Winsize = consInfo()
	Output = output{Prompts: make(map[string]string), Debug: true}
	Output.Prompts["info"] = fmt.Sprintf("%s%s%s",
		String(".").Cyan(),
		String(".").Bold().Cyan(),
		String(".").Bold().White())
	Output.Prompts["warn"] = fmt.Sprintf("%s%s%s",
		String(".").Yellow(),
		String(".").Bold().Yellow(),
		String(".").Bold().White())
	Output.Prompts["debug"] = fmt.Sprintf("%s%s%s",
		String(".").Purple(),
		String(".").Bold().Purple(),
		String(".").Bold().White())
	Output.Prompts["error"] = fmt.Sprintf("%s%s%s",
		String(".").Red(),
		String(".").Bold().Red(),
		String(".").Bold().White())
	Output.Prompts["status"] = fmt.Sprintf("%s%s%s",
		String("-").Cyan(),
		String("-").Bold().Cyan(),
		String("-").Bold().White())
}
