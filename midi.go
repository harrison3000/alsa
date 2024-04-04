package alsa

import (
	"os"

	"github.com/yobert/alsa/alsatype"
	"github.com/yobert/alsa/internal/seq"
)

const (
	cmdSeqPVersion      ioctl_e = 0x8000_5300
	cmdSeqClientId      ioctl_e = 0x8000_5301
	cmdSeqGetClientInfo ioctl_e = 0xc000_5310
	cmdSeqSetClientInfo ioctl_e = 0x4000_5311
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

	err = ioctl2(fh, cmdSeqClientId, &port.clientId)
	if err != nil {
		return nil, err
	}

	clInfo := seq.SeqClientInfo{
		ClientId: port.clientId,
	}

	err = ioctl2(fh, cmdSeqGetClientInfo, &clInfo)
	if err != nil {
		return nil, err
	}

	str2cstr(name, clInfo.Name[:])

	err = ioctl2(fh, cmdSeqSetClientInfo, &clInfo)
	if err != nil {
		return nil, err
	}

	return &port, nil
}

type SeqClient struct {
	fh *os.File

	clientId int32
	pversion alsatype.PVersion
}
