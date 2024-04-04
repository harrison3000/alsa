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

	cmdSeqCreatePort ioctl_e = 0xc000_5320
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

	//TODO user pversion

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

func (cli *SeqClient) CreatePort(name string) error {
	var p seq.PortInfo

	str2cstr(name, p.Name[:])
	p.Address.Client = byte(cli.clientId)

	//TODO a bunch of settings

	err := ioctl2(cli.fh, cmdSeqCreatePort, &p)
	if err != nil {
		return err
	}

	return nil
}

type SeqClient struct {
	fh *os.File

	clientId int32
	pversion alsatype.PVersion
}

type SeqPort struct {
	name string
}
