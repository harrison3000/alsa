package alsa

import (
	"os"

	"github.com/yobert/alsa/alsatype"
)

func NewSequencerPort(name string) (*SeqPort, error) {
	var port SeqPort
	var err error

	port.fh.File, err = os.Open("/dev/snd/seq")
	if err != nil {
		return nil, err
	}

	err = port.fh.ioctlRead(cmdSeqPVersion, &port.pversion)
	if err != nil {
		return nil, err
	}

	return &port, nil
}

type SeqPort struct {
	fh file_handle

	pversion alsatype.PVersion
}
