package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type (
	point  struct{ row, col int }
	record struct {
		steps []*point
	}
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	cnt := nextInt(scan)
	for idx := 1; idx <= cnt; idx++ {
		// 读取一组数据
		heights := nextArray(scan, idx)
		rows := len(heights)
		cols := len(heights[0])

		// 初始化路径记录
		pathrec := make([][]*record, rows)
		for row := range pathrec {
			pathrec[row] = make([]*record, cols)
			for col := range pathrec[row] {
				pathrec[row][col] = &record{steps: []*point{&point{row, col}}}
			}
		}

		// 查找最佳路径
		for findBestPath(heights, pathrec) {
		}

		// 输出结果
		maxLen := 0
		maxPt := &point{}
		for x, row := range pathrec {
			for y, rec := range row {
				if len(rec.steps) > maxLen {
					maxLen = len(rec.steps)
					maxPt = &point{x, y}
				}
			}
		}

		rec := pathrec[maxPt.row][maxPt.col]
		fmt.Println("\n结果: 长度=", len(rec.steps))
		for _, pt := range rec.steps {
			fmt.Printf("(%d,%d,%d) -> ", pt.row, pt.col, heights[pt.row][pt.col])
		}
		fmt.Println("\n\n")
	}
}

func findBestPath(heights [][]int, pathrec [][]*record) bool {

	rows := len(heights)
	cols := len(heights[0])
	haschange := false

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			curHeight := heights[row][col]
			curRecord := pathrec[row][col]
			maxPrevSteps := len(curRecord.steps) - 1

			var bestPrevPoint *point
			var prevpts []*point

			// 可能的前一个点
			prevpts = append(prevpts, &point{row - 1, col})
			prevpts = append(prevpts, &point{row + 1, col})
			prevpts = append(prevpts, &point{row, col - 1})
			prevpts = append(prevpts, &point{row, col + 1})

			for _, pt := range prevpts {
				if pt.row < 0 || pt.row >= rows {
					continue
				} else if pt.col < 0 || pt.col >= cols {
					continue
				} else {
					prevHeight := heights[pt.row][pt.col]
					prevRecord := pathrec[pt.row][pt.col]
					// 当前点高度小于前一个点(可以从前一个点滑向当前点)
					// 而且前一个点已经经过的路径更长
					if prevRecord != nil && curHeight < prevHeight && len(prevRecord.steps) > maxPrevSteps {
						maxPrevSteps = len(prevRecord.steps)
						bestPrevPoint = pt
					}
				}
			}

			// 找到更好的、以当前点(row,col)为终点的路径
			if bestPrevPoint != nil {
				prevrec := pathrec[bestPrevPoint.row][bestPrevPoint.col]
				curRecord.steps = append(prevrec.steps, &point{row, col})
				haschange = true
			}
		}
	}

	return haschange
}

//--------------------------------------------------------------------

func nextInt(scan *bufio.Scanner) int {
	scan.Scan()
	val, err := strconv.Atoi(scan.Text())
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func nextArray(scan *bufio.Scanner, index int) (heights [][]int) {

	rows := nextInt(scan)
	cols := nextInt(scan)

	fmt.Printf("\n----- 测试数据%d rows=%d cols=%d -----\n", index, rows, cols)
	heights = make([][]int, rows)
	for row := range heights {
		heights[row] = make([]int, cols)
		for col := range heights[row] {
			heights[row][col] = nextInt(scan)
			fmt.Printf("%4d", heights[row][col])
		}
		fmt.Println()
	}

	return
}
