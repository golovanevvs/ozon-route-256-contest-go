package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type item struct {
	x, y int
	st   int
	cost int
}

type queue []item

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
		var x1, y1, x2, y2 int
		fmt.Fscanln(in, &x1, &y1)
		fmt.Fscanln(in, &x2, &y2)

		fmt.Fprintln(out, tTaskSolving(strs, n, m, x1, y1, x2, y2))
	}
}

func tTaskSolving(strs []string, n, m, x1, y1, x2, y2 int) int {
	dxEven := [6]int{-1, -1, 0, 0, 1, 1}
	dyEven := [6]int{0, 1, -1, 1, 0, 1}
	dxOdd := [6]int{-1, -1, 0, 0, 1, 1}
	dyOdd := [6]int{-1, 0, -1, 1, -1, 0}

	x1--
	y1--
	x2--
	y2--

	big := math.MaxInt64
	dist := make([][][]int, n)
	for i := range n {
		dist[i] = make([][]int, m)
		for j := range m {
			dist[i][j] = []int{big, big}
		}
	}

	init := 0
	if !isLand(strs, n, m, x1, y1) {
		init = 1
	}
	dist[x1][y1][init] = 0

	q := &queue{}

	heap.Init(q)
	heap.Push(q, item{x1, y1, init, 0})

	for q.Len() > 0 {
		u := heap.Pop(q).(item)

		if u.cost != dist[u.x][u.y][u.st] {
			continue
		}

		evenODD := u.x & 1

		var dx, dy [6]int

		if evenODD == 0 {
			dx = dxEven
			dy = dyEven
		} else {
			dx = dxOdd
			dy = dyOdd
		}

		for k := range 6 {
			nx := u.x + dx[k]
			ny := u.y + dy[k]

			current := isLand(strs, n, m, u.x, u.y)
			next := isLand(strs, n, m, nx, ny)

			nextSt := 0

			if !next {
				nextSt = 1
			}

			cost := u.cost

			if current != next {
				cost++
			}

			if nx >= 0 && nx < n && ny >= 0 && ny < m {
				if cost < dist[nx][ny][nextSt] {
					dist[nx][ny][nextSt] = cost
					heap.Push(q, item{nx, ny, nextSt, cost})
				}
			}
		}
	}
	return min(dist[x2][y2][0], dist[x2][y2][1])
}

func isLand(data []string, n, m, x, y int) bool {
	if x < 0 || x >= n || y < 0 || y >= len(data[x]) {
		return false
	}
	return data[x][y] == ' '
}

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].cost < q[j].cost
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x any) {
	*q = append(*q, x.(item))
}

func (q *queue) Pop() any {
	temp := *q
	n := len(temp)
	x := temp[n-1]
	*q = temp[:n-1]

	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
