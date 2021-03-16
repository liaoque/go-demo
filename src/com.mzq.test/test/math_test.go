package test

import (
	"com.mzq.test/user"
	"testing"
)

func TestAdd(t *testing.T) {
	if  1 != user.Add(1, 0){
		t.Error("1+0!=3")
	}
}


func TestAdd2(t *testing.T) {
	if  3 != user.Add2(1, 0){
		t.Error("1+0!=3")
	}
}


func BenchmarkAdd(b *testing.B) {
	for i :=0; i< 10000000; i++ {
		user.Add(1, 2)
	}
}