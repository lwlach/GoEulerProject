package main

import "testing"

func TestFindMex(t *testing.T) {
	t.Run("find mex", func(t *testing.T) {
		arr := []int{0, 3, 2}
		mex := findMex(arr)
		if mex != 1 {
			t.Fail()
		}
	})
	t.Run("find mex2", func(t *testing.T) {
		arr := []int{0, 3, 2}
		mex := findMex2(arr)
		if mex != 1 {
			t.Fail()
		}
	})
}
