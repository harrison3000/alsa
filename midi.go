package alsa

import (
	"os"

	"github.com/yobert/alsa/alsatype"
)

const (
	cmdSeqPVersion = 0x8000_5300
)

func NewSequencerClient(name string) (*SeqClient, error) {
	var port SeqClient

	fh, err := os.Open("/dev/snd/seq")
	if err != nil {
		return nil, err
	}
	port.fh = fh

	err = ioctl2(fh, cmdSeqPVersion, &port.pversion)
	if err != nil {
		return nil, err
	}

	return &port, nil
}

type SeqClient struct {
	fh *os.File

	pversion alsatype.PVersion
}
