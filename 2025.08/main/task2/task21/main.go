package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("../tests/2")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v", err)
		return
	}
	defer file.Close()

	in := bufio.NewReader(file)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)
}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var t int
	fmt.Fscanln(in, &t)

	for range t {
		var n int
		fmt.Fscanln(in, &n)

		strs := make([]string, 0)

		for range n {

			var str string
			strWithSuffix, _ := in.ReadString('\n')
			if strings.HasSuffix(strWithSuffix, "\r\n") {
				str = strings.TrimRight(strWithSuffix, "\r\n")
			} else if strings.HasSuffix(strWithSuffix, "\n") {
				str = strings.TrimRight(strWithSuffix, "\n")
			} else {
				str = strWithSuffix
			}

			strs = append(strs, str)
		}

		fmt.Fprintln(out, tTaskSolving(strs))
	}
}

func tTaskSolving(strs []string) string {
	part1 := strings.Fields(strs[0])
	do := strings.TrimRight(part1[len(part1)-1], "!")

	mapWhoPoints := make(map[string]int)

	for _, str := range strs {
		parts := strings.Fields(str)

		whoSpk := strings.TrimRight(parts[0], ":")
		points := mapWhoPoints[whoSpk]
		mapWhoPoints[whoSpk] = points

		if parts[2] == "am" {
			points = mapWhoPoints[whoSpk]
			if len(parts) == 4 {
				mapWhoPoints[whoSpk] = points + 2
			} else {
				mapWhoPoints[whoSpk] = points - 1
			}
		} else {
			points = mapWhoPoints[parts[1]]
			if len(parts) == 4 {
				mapWhoPoints[parts[1]] = points + 1
			} else {
				mapWhoPoints[parts[1]] = points - 1
			}
		}

	}

	resSlice := make([]string, 0)
	max := math.MinInt
	for k, v := range mapWhoPoints {
		if v > max {
			max = v
			resSlice = make([]string, 0)
			resSlice = append(resSlice, k)
		} else if v == max {
			resSlice = append(resSlice, k)
		}
	}

	sort.Strings(resSlice)

	res := make([]string, len(resSlice))
	for i, v := range resSlice {
		var builder strings.Builder
		builder.WriteString(v)
		builder.WriteString(" is ")
		builder.WriteString(do)
		builder.WriteString(".")
		res[i] = builder.String()
	}

	resString := strings.Join(res, "\n")

	return resString
}
