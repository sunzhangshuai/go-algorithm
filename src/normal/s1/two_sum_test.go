package s1

import (
	"reflect"
	"testing"
)

func TestHandle(t *testing.T) {
	ret := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	if !reflect.DeepEqual(ret, want) {
		t.Error("测试失败")
	}
}
