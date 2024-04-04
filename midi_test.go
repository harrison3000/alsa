package alsa

import "testing"

func TestMidiList(t *testing.T) {
	//not really a automatic test, just a way to set a breakpoint and let it rip

	p, e := NewSequencerClient("go-alsa test client")
	if e != nil {
		t.Fatal(e)
	}

	_ = p
}
