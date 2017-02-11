package gout

import (
	"fmt"
	"log"

	eval "github.com/Knetic/govaluate"
)

// data structures

type output struct {
	Prompts   map[string]string
	Debug     bool
	Quiet     bool
	Verbose   bool
	ToFile    bool
	Throbber  []string
	lastThrob int
}

// Output functions
func consoleOutput(t string, e string, f string, args ...interface{}) {
	Winsize = ConsInfo()
	if Output.ToFile {
		log.Printf("%s %s\n", t, fmt.Sprintf(f, args...))
	}
	fmt.Printf("%s %s %s%s",
		Output.Prompts[t],
		fmt.Sprintf(f, args...),
		padding(int(Winsize.Col)-(7+len(fmt.Sprintf(f, args...)))),
		e)
}

func Info(f string, args ...interface{}) {
	if !Output.Quiet {
		consoleOutput("info", "\n", f, args...)
	}
}

func Debug(f string, args ...interface{}) {
	if Output.Debug {
		consoleOutput("debug", "\n", f, args...)
	}
}

func Warn(f string, args ...interface{}) {
	if Output.Verbose {
		consoleOutput("warn", "\n", f, args...)
	}
}

func Error(f string, args ...interface{}) {
	consoleOutput("error", "\n", f, args...)
}

func Status(f string, args ...interface{}) {
	consoleOutput("status", "\r", f, args...)
}

func Progress(l int, p int) string {
	rl := l - 2
	var rslt interface{}
	var sp string
	if p < 100 {
		expr, _ := eval.NewEvaluableExpression(fmt.Sprintf("%d * 0.%02d", rl, p))
		rslt, _ = expr.Evaluate(nil)
		sp = repeat("#", int(rslt.(float64)))
	} else {
		rslt = rl
		sp = repeat("#", rslt.(int))
	}
	pd := padding(rl - len(sp))
	return fmt.Sprintf("[%s%s]", sp, pd)
}

func Throbber() string {
	if Output.lastThrob == len(Output.Throbber)+1 {
		Output.lastThrob = 1
	} else {
		Output.lastThrob++
	}
	return Output.Throbber[Output.lastThrob-1]
}
