package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)


type RingBuffer struct {
	size  int
	items []lineEntry
	pos   int
	full  bool
}

type lineEntry struct {
	num  int
	text string
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		size:  size,
		items: make([]lineEntry, size),
	}
}

func (rb *RingBuffer) Add(num int, text string) {
	if rb.size == 0 {
		return
	}
	rb.items[rb.pos] = lineEntry{num: num, text: text}
	rb.pos = (rb.pos + 1) % rb.size
	if rb.pos == 0 {
		rb.full = true
	}
}


func (rb *RingBuffer) Items() []lineEntry {
	if !rb.full {
		return rb.items[:rb.pos]
	}
	return append(rb.items[rb.pos:], rb.items[:rb.pos]...)
}

func (rb *RingBuffer) Clear() {
	rb.pos = 0
	rb.full = false
}


type Grep struct {
	flagA int
	flagB int
	flagC int

	flagCount bool // -c
	flagIgnore bool // -i
	flagInvert bool // -v
	flagFixed  bool // -F
	flagNumber bool // -n

	pattern   string
	regex     *regexp.Regexp
	reader    io.Reader
}

func NewGrep() *Grep {
	return &Grep{}
}

func (g *Grep) Compile() error {
	if g.flagC > 0 {
		g.flagA = g.flagC
		g.flagB = g.flagC
	}

	if g.flagFixed {
		if g.flagIgnore {
			g.pattern = strings.ToLower(g.pattern)
		}
		return nil
	}

	p := g.pattern
	if g.flagIgnore {
		p = "(?i)" + p
	}

	re, err := regexp.Compile(p)
	if err != nil {
		return err
	}
	g.regex = re
	return nil
}

func (g *Grep) match(line string) bool {
	orig := line

	if g.flagIgnore && g.flagFixed {
		return strings.Contains(strings.ToLower(orig), g.pattern)
	}

	if g.flagFixed {
		return strings.Contains(orig, g.pattern)
	}

	return g.regex.MatchString(orig)
}

func (g *Grep) printLine(num int, text string) {
	if g.flagNumber {
		fmt.Printf("%d:%s\n", num, text)
	} else {
		fmt.Println(text)
	}
}

func (g *Grep) Run() error {
	scanner := bufio.NewScanner(g.reader)

	var (
		ringBuf = NewRingBuffer(g.flagB)
		after   = 0
		count   = 0

		lastPrinted = -1
	)

	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		matched := g.match(line)
		if g.flagInvert {
			matched = !matched
		}

		if matched {
			if g.flagCount {
				count++
				continue
			}

			for _, prev := range ringBuf.Items() {
				if prev.num > lastPrinted {
					g.printLine(prev.num, prev.text)
					lastPrinted = prev.num
				}
			}
			ringBuf.Clear()

			if lineNum > lastPrinted {
				g.printLine(lineNum, line)
				lastPrinted = lineNum
			}

			after = g.flagA
		} else {
			if after > 0 {
				if lineNum > lastPrinted {
					g.printLine(lineNum, line)
					lastPrinted = lineNum
				}
				after--
			} else {
				ringBuf.Add(lineNum, line)
			}
		}
	}

	if g.flagCount {
		fmt.Println(count)
	}

	return scanner.Err()
}


func main() {
	grep := NewGrep()

	flag.IntVar(&grep.flagA, "A", 0, "print N lines After match")
	flag.IntVar(&grep.flagB, "B", 0, "print N lines Before match")
	flag.IntVar(&grep.flagC, "C", 0, "print N lines of Context around match")
	flag.BoolVar(&grep.flagCount, "c", false, "print only count of matching lines")
	flag.BoolVar(&grep.flagIgnore, "i", false, "ignore case")
	flag.BoolVar(&grep.flagInvert, "v", false, "invert match")
	flag.BoolVar(&grep.flagFixed, "F", false, "match fixed string")
	flag.BoolVar(&grep.flagNumber, "n", false, "print line number")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: grep [flags] pattern [file]")
		os.Exit(1)
	}

	grep.pattern = flag.Arg(0)

	if flag.NArg() >= 2 {
		fname := flag.Arg(1)
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		grep.reader = file
	} else {
		grep.reader = os.Stdin
	}

	if err := grep.Compile(); err != nil {
		fmt.Fprintln(os.Stderr, "invalid pattern:", err)
		os.Exit(1)
	}

	if err := grep.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "processing error:", err)
		os.Exit(1)
	}
}