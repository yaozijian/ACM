package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	object struct {
		price          int       // 价格
		weight         int       // 权重
		weightprice    int       // 价格*权重
		objects        []*object // 物品列表
		maxweightprice int
		maxprice       int // 包含此物品的最大价格
	}

	problem struct {
		money   int
		objects []*object
	}
)

func main() {

	rand.Seed(time.Now().Unix())

	this := genProblem()
	for this.findBestSolution() {
	}

	maxprice := 0
	maxidx := 0
	for idx, obj := range this.objects {
		if obj.maxweightprice > maxprice {
			maxidx = idx
			maxprice = obj.maxweightprice
		}
	}

	obj := this.objects[maxidx]
	fmt.Printf("\n最大权重价格: %d 总金额: %d\n", obj.maxweightprice, obj.maxprice)
	for idx, cur := range obj.objects {
		fmt.Printf("  选择物品%2d: 价格=%-3d 权重=%d 权重价格=%d\n", idx+1, cur.price, cur.weight, cur.weightprice)
	}
}

func (this *problem) findBestSolution() (moved bool) {

	for idx, cur := range this.objects {
		for idy, cmp := range this.objects {

			if idy == idx {
				continue
			}

			curUsed := false
			for _, obj := range cmp.objects {
				if obj == cur {
					curUsed = true
					break
				}
			}

			// 没有选择过这个物品,且满足总金额上限条件
			if !curUsed && cmp.maxprice+cur.price <= this.money {
				// 并且权重价格更大
				if cmp.maxweightprice+cur.weightprice > cur.maxweightprice {
					cur.maxweightprice = cmp.maxweightprice + cur.weightprice
					cur.maxprice = cmp.maxprice + cur.price
					cur.objects = make([]*object, len(cmp.objects)+1)
					copy(cur.objects, cmp.objects)
					cur.objects[len(cmp.objects)] = cur
					moved = true
				}
			}
		}
	}
	return
}

func genProblem() (this *problem) {

	this = &problem{money: rand.Intn(5e3) + 2e3}
	cnt := rand.Intn(10) + 5

	fmt.Printf("总金额上限: %d 物品数: %d\n", this.money, cnt)

	this.objects = make([]*object, cnt)
	for idx := range this.objects {
		obj := &object{
			price:  rand.Intn(500) + 50,
			weight: rand.Intn(5) + 1,
		}
		obj.weightprice = obj.weight * obj.price
		obj.maxweightprice = obj.weightprice
		obj.maxprice = obj.price
		obj.objects = []*object{obj}

		this.objects[idx] = obj
		fmt.Printf("  物品%2d: 价格=%-3d 权重=%d 权重价格=%d\n", idx+1, obj.price, obj.weight, obj.weightprice)
	}

	return
}
