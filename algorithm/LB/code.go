package main

import (
	"fmt"
	"math"
	"math/rand"
)

// RoundRobin 轮训算法
type RoundRobin struct {
	server []string
	index  int
}

func (rr *RoundRobin) GetServer() string {
	server := rr.server[rr.index]
	rr.index = rr.index + 1%len(rr.server)
	return server
}

// WeightedRoundRobin 加权轮询
type WeightedRoundRobin struct {
	servers []string
	weights []int
	index   int
}

func (wrr *WeightedRoundRobin) GetWeightRoundRobin() string {
	totalWeight := 0
	for _, weight := range wrr.weights {
		totalWeight += weight
	}
	valWeight := rand.Intn(totalWeight)
	for i, weight := range wrr.weights {
		valWeight -= weight
		if valWeight < 0 {
			return wrr.servers[i]
		}
	}
	return wrr.servers[0]
}

type LeastConnections struct {
	servers []string
	conns   map[string]int
}

func (lc *LeastConnections) addConnCount(server string) {
	lc.conns[server]++
}

func (lc *LeastConnections) GetConnServer() string {
	maxCount := math.MaxInt
	serverS := ""
	for server, count := range lc.conns {
		if count < maxCount {
			maxCount = count
			serverS = server
		}
	}

	return serverS

}

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	var backtrack func(*[][]int, []int)
	backtrack = func(res *[][]int, path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			*res = append(*res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(res, path)
			path = path[:len(path)-1]
			used[i] = false
		}

	}

	res := [][]int{}
	path := []int{}
	backtrack(&res, path)
	return res

}

func main() {
	nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Println(res)
	return
	lc := &LeastConnections{servers: []string{"1", "2"}, conns: make(map[string]int)}
	lc.addConnCount("1")
	lc.addConnCount("1")
	lc.addConnCount("2")
	println(lc.conns)
}
