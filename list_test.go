package demo

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	pool, err := NewResourcePool()
	if err != nil {
		t.Fatal(err)
		return
	}
	if pool.Len() == 0 {
		t.Fatal("length of pool should not be 0")
	}
	fmt.Println("resource pool len", pool.Len())
	for _, re := range pool.List() {
		fmt.Println(re.Name)
		fmt.Println(re.Text)
		fmt.Println(len(re.Audio))
	}
}
