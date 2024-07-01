package main

import (
	"fmt"
)

type UnionFind struct {
	parent map[int]int
	rank   map[int]int
	size   map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[int]int),
		rank:   make(map[int]int),
		size:   make(map[int]int),
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
			uf.size[rootX] += uf.size[rootY]
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
			uf.size[rootY] += uf.size[rootX]
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
			uf.size[rootX] += uf.size[rootY]
		}
	}
}

func (uf *UnionFind) Add(x int) {
	if _, found := uf.parent[x]; !found {
		uf.parent[x] = x
		uf.rank[x] = 0
		uf.size[x] = 1
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uf := NewUnionFind()
	for _, num := range nums {
		uf.Add(num)
		if _, found := uf.parent[num-1]; found {
			uf.Union(num, num-1)
		}
		if _, found := uf.parent[num+1]; found {
			uf.Union(num, num+1)
		}
	}

	longest := 0
	for _, size := range uf.size {
		if size > longest {
			longest = size
		}
	}

	return longest
}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println("Longest consecutive sequence length (example 1):", longestConsecutive(nums1))
	fmt.Println("Longest consecutive sequence length (example 2):", longestConsecutive(nums2))
}
