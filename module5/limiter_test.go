package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	l := NewSlidingWindow(100*time.Millisecond, time.Second, 10)
	for i := 0; i < 5; i++ {
		fmt.Println(l.AllowRequest())
	}
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 5; i++ {
		fmt.Println(l.AllowRequest())
	}
	r := l.AllowRequest()
	fmt.Println(r)
	if r {
		t.Fatal()
	}
	time.Sleep(501 * time.Millisecond)
	for i := 0; i < 5; i++ {
		r := l.AllowRequest()
		fmt.Println(r)
		if !r {
			t.Failed()
		}
	}
}
