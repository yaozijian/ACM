package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	point struct {
		row, col int      // 坐标
		val      int      // 值
		path     []*point // 以当前点为终点的最长路径
		pathval  int      // 最长路径长度
	}
	triangle [][]*point
)

func main() {
	rand.Seed(time.Now().Unix())
	this := genRandomTriangle()
	for this.findBestPath() {
	}
	this.outputBestPath()
}

func (this triangle) findBestPath() (changed bool) {

	merge := func(cur, cmp *point) bool {
		// 从上一行某个点的路径到当前点的路径更长
		if cmp.pathval+cur.val > cur.pathval {
			cur.pathval = cmp.pathval + cur.val
			cur.path = make([]*point, len(cmp.path)+1)
			copy(cur.path, cmp.path)
			cur.path[len(cmp.path)] = cur
			return true
		} else {
			return false
		}
	}

	for y, row := range this {
		if y == 0 {
			continue
		}
		for x, cur := range row {
			// 上一行的列数
			upcols := len(this[y-1])
			// 左上方的点
			ltx, lty := x-1, y-1
			if ltx >= 0 && ltx < upcols && lty >= 0 {
				if merge(cur, this[lty][ltx]) {
					changed = true
				}
			}
			// 右上方的点
			ltx, lty = x, y-1
			if ltx >= 0 && ltx < upcols && lty >= 0 {
				if merge(cur, this[lty][ltx]) {
					changed = true
				}
			}
		}
	}

	return
}

func (this triangle) outputBestPath() {

	rows := len(this)
	lastRow := this[rows-1]

	maxLen := 0
	maxCol := 0

	// 找最后一行中路径最长的列
	for col, pt := range lastRow {
		if pt.pathval > maxLen {
			maxCol = col
			maxLen = pt.pathval
		}
	}

	// 输出
	fmt.Println("\n最长路径: ")
	pt := lastRow[maxCol]
	for _, pt := range pt.path {
		fmt.Printf("(%d,%d,%d) -> ", pt.row, pt.col, pt.val)
	}
	fmt.Printf(" %d\n\n", pt.pathval)
}

func genRandomTriangle() (obj triangle) {

	// 随机确定行数
	rows := rand.Intn(100) + 1
	obj = make([][]*point, rows)

	w := 8
	maxw := (rows-1)*w + 2
	space := maxw / 2
	delta := w / 2

	cols := 0
	for row := 0; row < rows; row++ {
		// 列数从1依次递增
		cols++
		obj[row] = make([]*point, cols)
		fmt.Printf("%*s", space, " ")

		// 依次生成各列
		for col := range obj[row] {
			pt := &point{
				row: row,
				col: col,
				val: rand.Intn(100),
			}
			pt.pathval = pt.val
			pt.path = []*point{pt}
			obj[row][col] = pt

			fmt.Printf("%-*d", w, pt.val)
		}
		fmt.Println()
		space -= delta
	}
	return
}
