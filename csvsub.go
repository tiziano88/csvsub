package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	r := csv.NewReader(os.Stdin)

	flag.Parse()
	input := flag.Arg(0)

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
