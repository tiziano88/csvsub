package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "%s <format-string>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "<format-string> contains substitution placeholders in the format \"{header}\", where header corresponds to one of the input CSV file header fields.\n")
	}
}

func main() {
	r := csv.NewReader(os.Stdin)

	flag.Parse()

	input := flag.Arg(0)
	if input == "" {
		flag.Usage()
		os.Exit(-1)
	}

	rs, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	fieldMap := make(map[string]int)

	re, err := regexp.Compile("{\\w*}")
	if err != nil {
		panic(err)
	}

	// Go through each field and reformat it.
	for i, row := range rs {
		if i == 0 {
			// Header row.
			for j, fName := range row {
				fieldMap[fName] = j
			}
		} else {
			fmt.Println(re.ReplaceAllStringFunc(input, func(m string) string {
				fName := strings.TrimSuffix(strings.TrimPrefix(m, "{"), "}")
				fNum, ok := fieldMap[fName]
				if !ok {
					return m
				}
				return row[fNum]
			}))
		}
	}
}
