package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortConfig struct {
	column     int
	numeric    bool
	reverse    bool
	unique     bool
	check      bool
	month      bool
	human      bool
	trimSpaces bool
	filename   string
}

func parseArgs() SortConfig {
	cfg := SortConfig{}

	flag.IntVar(&cfg.column, "k", 1, "sort by column number (default 1)")
	flag.BoolVar(&cfg.numeric, "n", false, "numeruc sort")
	flag.BoolVar(&cfg.reverse, "r", false, "reverse order")
	flag.BoolVar(&cfg.unique, "u", false, "unique lines")
	flag.BoolVar(&cfg.check, "c", false, "check if sorted")
	flag.BoolVar(&cfg.month, "M", false, "sort by month name (Jan..Dec)")
	flag.BoolVar(&cfg.trimSpaces, "b", false, "ignore trailing blanks")
	flag.BoolVar(&cfg.human, "h", false, "human-readable numeric sort (1K, 2M)")
	flag.Parse()

	for _, arg := range os.Args[1:] {
		if len(arg) > 2 && strings.HasPrefix(arg, "-") {
			for _, c := range arg[1:] {
				switch c {
				case 'n':
					cfg.numeric = true
				case 'r':
					cfg.reverse = true
				case 'u':
					cfg.unique = true
				case 'c':
					cfg.check = true
				case 'M':
					cfg.month = true
				case 'b':
					cfg.trimSpaces = true
				case 'h':
					cfg.human = true
				}
			}
		}
	}

	args := flag.Args()
	if len(args) > 0 {
		cfg.filename = args[len(args)-1]
	}
	return cfg
}

func readlines(filename string) ([]string, error) {
	var scanner *bufio.Scanner
	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func getColumn(line string, col int) string {
	fields := strings.Split(line, "\t")
	if col <= len(fields) {
		return fields[col-1]
	}
	return line
}

func parsHumanNumber(s string) float64 {
	if s == "" {
		return 0
	}
	s = strings.TrimSpace(s)
	mult := 1.0
	last := s[len(s)-1]
	switch last {
	case 'K', 'k':
		mult = 1e3
		s = s[:len(s)-1]
	case 'M', 'm':
		mult = 1e6
		s = s[:len(s)-1]
	case 'G', 'g':
		mult = 1e9
		s = s[:len(s)-1]
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v * mult
}

var monthMap = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
	"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
	"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

func lessFunc(a, b string, cfg SortConfig) bool {
	ka := getColumn(a, cfg.column)
	kb := getColumn(b, cfg.column)

	if cfg.trimSpaces {
		ka = strings.TrimSpace(ka)
		kb = strings.TrimSpace(kb)
	}

	var less bool
	switch {
	case cfg.numeric:
		af, err1 := strconv.ParseFloat(ka, 64)
		bf, err2 := strconv.ParseFloat(kb, 64)
		if err1 == nil && err2 == nil {
			less = af < bf
		} else {
			less = ka < kb
		}

	case cfg.month:
		less = monthMap[ka] < monthMap[kb]

	case cfg.human:
		less = parsHumanNumber(ka) < parsHumanNumber(kb)

	default:
		less = ka < kb
	}

	if cfg.reverse {
		return !less
	}
	return less
}

func checkSorted(lines []string, cfg SortConfig) bool {
	for i := 1; i < len(lines); i++ {
		if lessFunc(lines[i], lines[i-1], cfg) {
			return false
		}
	}
	return true
}

func uniqueLines(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	result := []string{lines[0]}
	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			result = append(result, lines[i])
		}
	}
	return result
}

func main() {
	cfg := parseArgs()

	lines, err := readlines(cfg.filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if cfg.check {
		if checkSorted(lines, cfg) {
			fmt.Println("Data is sorted")
		} else {
			fmt.Println("Data is not sorted")
			os.Exit(1)
		}
		return
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return lessFunc(lines[i], lines[j], cfg)
	})

	if cfg.unique {
		lines = uniqueLines(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
