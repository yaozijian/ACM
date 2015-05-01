package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type (
	rect struct {
		width, height int
		order         []int //以本矩形为最后一个矩形的序列的嵌套次序(从小到大)
	}
)

func main() {

	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)

	cnt := nextInt(scan)
	for idx := 1; idx <= cnt; idx++ {
		// 读取一组数据
		rects := nextArray(scan, idx)

		// 找最佳次序
		for findBestPath(rects) {
		}

		maxLen := 0
		maxIdx := 0
		for x, rect := range rects {
			if len(rect.order) > maxLen {
				maxLen = len(rect.order)
				maxIdx = x
			}
		}

		rect := rects[maxIdx]
		fmt.Println("\n\n最佳次序: ")
		for _, x := range rect.order {
			this := rects[x]
			fmt.Printf("(%4d,%4d) -> ", this.width, this.height)
		}
		fmt.Println("\n\n")
	}
}

func findBestPath(rects []*rect) (haschange bool) {
	for x, rect := range rects {
		for y, other := range rects {
			if y != x {
				cnt := len(other.order)
				if rect.canInclude(other) && cnt+1 > len(rect.order) {
					rect.order = make([]int, cnt+1)
					copy(rect.order, other.order)
					rect.order[cnt] = x
					haschange = true
				}
			}
		}
	}
	return
}

// 可以包含其他矩形吗?
func (this *rect) canInclude(other *rect) bool {
	if this.width > other.width && this.height > other.height {
		return true
	} else if this.width > other.height && this.height > other.width {
		return true
	} else {
		return false
	}
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

func nextArray(scan *bufio.Scanner, index int) (rects []*rect) {

	cnt := nextInt(scan)

	fmt.Printf("\n----- 测试数据%d cnt=%d -----\n", index, cnt)
	rects = make([]*rect, cnt)
	for idx := range rects {
		w := nextInt(scan)
		h := nextInt(scan)
		rects[idx] = &rect{width: w, height: h, order: []int{idx}}
		fmt.Printf("矩形%2d: w=%-4d h=%-4d\n", idx+1, w, h)
	}

	return
}
