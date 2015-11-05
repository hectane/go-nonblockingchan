package nbc

import (
	"github.com/hectane/go-attest"

	"testing"
	"time"
)

// Because there is no synchronization between the testing goroutine and the
// channel's goroutine, a small delay must be introduced between operations.
func pause() {
	<-time.After(50 * time.Millisecond)
}

func TestNonBlockingChan(t *testing.T) {
	n := New()
	pause()
	if err := attest.ChanSend(n.Send, true); err != nil {
		t.Fatal(err)
	}
	pause()
	l := n.Len()
	if l != 1 {
		t.Fatalf("%d != 1", 1)
	}
	i, err := attest.ChanRecv(n.Recv)
	if err != nil {
		t.Fatal(err)
	}
	if i != true {
		t.Fatalf("%v != true", i)
	}
	pause()
	l = n.Len()
	if l != 0 {
		t.Fatalf("%d != 0", 0)
	}
	close(n.Send)
	pause()
	if err := attest.ChanClosed(n.Recv); err != nil {
		t.Fatal(err)
	}
}
