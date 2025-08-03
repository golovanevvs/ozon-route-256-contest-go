package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Структура элемента очереди
type item struct {
	x, y  int
	isSea bool
	cost  int
}

// Очередь с поддержкой приоритета
type priorityQueue []*item

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*item)) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Функция проверки принадлежности ячейки к суше
func isLand(data [][]byte, x, y int) bool {
	return data[x][y] != ' '
}

// Основная функция обработки задачи
func taskSolve(data [][]byte, startX, startY, endX, endY int) int {
	rows, cols := len(data), len(data[0])

	// Массив расстояний: dist[row][col][seaState]
	distance := make([][][]int, rows)
	for r := range distance {
		distance[r] = make([][]int, cols)
		for c := range distance[r] {
			distance[r][c] = []int{-1, -1} // [-1,-1]: неизведанная точка
		}
	}

	startIsSea := !isLand(data, startX, startY)

	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	// Начальная позиция
	heap.Push(&pq, &item{
		x:     startX,
		y:     startY,
		isSea: startIsSea,
		cost:  0,
	})

	distance[startX][startY][boolToIndex(startIsSea)] = 0

	// Определение возможных направлений
	dirsEven := []struct{ dx, dy int }{{-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}} // Чётные ряды
	dirsOdd := []struct{ dx, dy int }{{-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 1}}  // Нечётные ряды

	for pq.Len() > 0 {
		currItem := heap.Pop(&pq).(*item)
		x, y, seaState, currCost := currItem.x, currItem.y, currItem.isSea, currItem.cost

		// Если достигли цели
		if x == endX && y == endY {
			return currCost
		}

		// Выбираем правильные направления в зависимости от чётности строки
		var directions []struct{ dx, dy int }
		if x%2 == 0 {
			directions = dirsEven
		} else {
			directions = dirsOdd
		}

		// Рассматриваем всех соседей
		for _, dir := range directions {
			newX, newY := x+dir.dx, y+dir.dy

			// Игнорируем выходящие за пределы карты координаты
			if newX < 0 || newX >= rows || newY < 0 || newY >= cols {
				continue
			}

			// Определяем состояние следующей клетки
			nextIsSea := !isLand(data, newX, newY)

			// Стоимость перехода
			nextCost := currCost
			if seaState != nextIsSea { // Меняется среда?
				nextCost++ // Наказываем за переход
			}

			// Индекс нового состояния
			stateIndex := boolToIndex(nextIsSea)

			// Обновляем расстояние, если нашли лучший путь
			if distance[newX][newY][stateIndex] == -1 ||
				nextCost < distance[newX][newY][stateIndex] {
				distance[newX][newY][stateIndex] = nextCost
				heap.Push(&pq, &item{
					x:     newX,
					y:     newY,
					isSea: nextIsSea,
					cost:  nextCost,
				})
			}
		}
	}

	return -1 // Невозможно достичь цели
}

// Преобразование булевого значения в индекс массива
func boolToIndex(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Основной цикл программы
func main() {
	file, err := os.Open("../tests/1")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var numTests int
	fmt.Scanf("%d\n", &numTests)

	for testCase := 0; testCase < numTests; testCase++ {
		var rows, cols int
		fmt.Scanf("%d %d\n", &rows, &cols)

		data := make([][]byte, rows)
		for row := 0; row < rows; row++ {
			line := ""
			scanner.Scan()
			line = scanner.Text()
			data[row] = []byte(line)
		}

		var startX, startY, endX, endY int
		fmt.Scanf("%d %d\n%d %d\n", &startX, &startY, &endX, &endY)

		result := taskSolve(data, startX-1, startY-1, endX-1, endY-1)
		fmt.Fprintf(writer, "%d\n", result)
	}
}
