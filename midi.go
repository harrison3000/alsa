package alsa

import (
	"os"

	"github.com/yobert/alsa/alsatype"
)

func NewSequencerClient(name string) (*SeqPort, error) {
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

	err = port.fh.ioctlRead(cmdSeqClientId, &port.clientId)
	if err != nil {
		return nil, err
	}

	clInfo := alsatype.SeqClientInfo{
		ClientId: port.clientId,
	}

	err = port.fh.ioctlRW(cmdSeqGetClientInfo, &clInfo)
	if err != nil {
		return nil, err
	}

	if len(name) > 63 {
		name = name[:63]
	}
	namebs := clInfo.Name[:0]
	namebs = append(namebs, name...)
	namebs = append(namebs, 0)

	err = port.fh.ioctlWrite(cmdSeqSetClientInfo, &clInfo)
	if err != nil {
		return nil, err
	}

	return &port, nil
}

type SeqPort struct {
	fh file_handle

	clientId int32
	pversion alsatype.PVersion
}
