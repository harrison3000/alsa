package alsa

import "testing"

func TestMidiList(t *testing.T) {
	//not really a automatic test, just a way to set a breakpoint and let it rip

	c, e := NewSequencerClient("go-alsa test client")
	if e != nil {
		t.Fatal(e)
	}

	e = c.CreatePort("test port 0")
	if e != nil {
		t.Fatal(e)
	}

}
