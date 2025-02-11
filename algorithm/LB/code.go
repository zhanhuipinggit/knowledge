package main

import (
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

func main() {
	lc := &LeastConnections{servers: []string{"1", "2"}, conns: make(map[string]int)}
	lc.addConnCount("1")
	lc.addConnCount("1")
	lc.addConnCount("2")
	println(lc.conns)
}
