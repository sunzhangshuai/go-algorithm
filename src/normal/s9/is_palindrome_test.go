package s9

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(121) {
		t.Error("测试失败")
	}
}
