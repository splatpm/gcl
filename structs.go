package gout

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

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
