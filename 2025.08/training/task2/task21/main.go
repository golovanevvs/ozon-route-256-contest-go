package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	name string
	let  string
	n    int
	m    int
}

func main() {
	file, err := os.Open("../tests/1")
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
		var n, m int
		fmt.Fscanln(in, &n, &m)

		strs := make([]string, 0, n)

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

		for _, k := range getPath(strs) {
			fmt.Fprintln(out, strings.Join(k, ""))
		}
	}
}

func getPath(source []string) [][]string {
	pointA := point{
		name: "A",
		let:  "a",
	}
	pointB := point{
		name: "B",
		let:  "b",
	}

	data := getMainSlice(source, &pointA, &pointB)

	var result [][]string
	switch {
	case pointA.n < pointB.n:
		result = getPathOnePoint(getPathOnePoint(data, pointA, "left"), pointB, "right")
	default:
		switch {
		case pointA.m < pointB.m:
			result = getPathOnePoint(getPathOnePoint(data, pointA, "left"), pointB, "right")
		default:
			result = getPathOnePoint(getPathOnePoint(data, pointB, "left"), pointA, "right")
		}
	}
	return result
}

func getMainSlice(source []string, pointA, pointB *point) [][]string {
	// data[строка][столбец]
	data := make([][]string, len(source))
	for i := range data {
		data[i] = make([]string, len(source[i]))
	}
	for i := 0; i < len(source); i++ {
		for j, v := range source[i] {
			switch v {
			case 'A':
				pointA.n = i
				pointA.m = j
			case 'B':
				pointB.n = i
				pointB.m = j
			}
			data[i][j] = string(v)
		}
	}
	return data
}

func getPathOnePoint(data [][]string, point point, dir string) [][]string {
	n := len(data) - 1
	m := len(data[0]) - 1
	switch dir {
	case "left":
		switch {
		case (point.n+1)%2 == 0:
			for point.n > 0 {
				point.n--
				data[point.n][point.m] = point.let
			}
			for point.m > 0 {
				point.m--
				data[point.n][point.m] = point.let
			}

		default:
			for point.m > 0 {
				point.m--
				data[point.n][point.m] = point.let
			}
			for point.n > 0 {
				point.n--
				data[point.n][point.m] = point.let
			}
		}
	case "right":
		switch {
		case (point.n+1)%2 == 0:
			for point.n < n {
				point.n++
				data[point.n][point.m] = point.let
			}
			for point.m < m {
				point.m++
				data[point.n][point.m] = point.let
			}

		default:
			for point.m < m {
				point.m++
				data[point.n][point.m] = point.let
			}
			for point.n < n {
				point.n++
				data[point.n][point.m] = point.let
			}
		}
	}
	return data
}
