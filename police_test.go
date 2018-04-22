package police

import "testing"

func TestLimit(t *testing.T) {
	t.Run("create limited channel", func(t *testing.T) {
		a := &Arrival{}
		want := 3
		a.Limit(3)
		got := cap(a.Ch)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestBlock(t *testing.T) {
	t.Run("add struct{} to channel", func(t *testing.T) {
		a := &Arrival{}
		want := 1
		a.Limit(1)
		a.Block()
		got := len(a.Ch)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestRelease(t *testing.T) {
	t.Run("remove struct{} from channel", func(t *testing.T) {
		a := &Arrival{}
		want := 0
		a.Limit(1)
		a.Ch <- struct{}{}
		a.Release()
		got := len(a.Ch)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
