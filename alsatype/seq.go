package alsatype

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
