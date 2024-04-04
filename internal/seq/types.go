package seq

type SeqClientType int32

type SeqClientInfo struct {
	ClientId       int32
	Type           SeqClientType
	Name           [64]byte
	FilterFlags    uint32
	MulticasFilter [8]byte
	EventFilter    [32]byte
	NumPorts       int32 //RO
	EventsLost     int32
	Card           int32 //RO
	Pid            int32 //RO
	MIDIVersion    int32
	GroupFilter    int32

	reserved [48]byte
}

type Addr struct {
	Client, Port byte
}

type PortInfo struct {
	Address      Addr
	Name         [64]byte
	Capability   int32
	Type         int32
	MidiChannels int32
	MidiVoices   int32
	SynthVoices  int32
	ReadUse      int32
	WriteUse     int32
	kernel       uintptr
	Flags        int32 //TODO specific type?
	TimeQueue    byte
	reserved     [59]byte
}
