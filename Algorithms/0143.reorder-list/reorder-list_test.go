package Problem0143

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Golang/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head []int
	ans  []int
}{

	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
	},

	{
		[]int{1, 2, 3, 4},
		[]int{1, 4, 2, 3},
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]int{1, 8, 2, 7, 3, 6, 4, 5},
	},

	// 可以有多个 testcase
}

func Test_reorderList(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		head := kit.Ints2List(tc.head)
		reorderList(head)
		ans := kit.List2Ints(head)

		ast.Equal(tc.ans, ans, "输入:%v", tc)
	}
}

func Benchmark_reorderList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reorderList(kit.Ints2List(tc.head))
		}
	}
}