package main

import (
	"math"
	"sync"
)

func main() {
	//同时并发的进程
	cLen := 50
	//当前并发执行进程
	process := 0

	list := []interface{}{}
	for i := 0; i <= int(math.Ceil(float64(len(list))/float64(cLen))); i++ {
		wg := sync.WaitGroup{} //分段并发获取结果
		for j := 0; j < cLen; j++ {
			if process >= len(list) {
				break
			}

			wg.Add(1)
			go func(data interface{}) {
				defer wg.Done()
				//todo something
			}(list[process])
			process++
		}
		wg.Wait()
	}

}
