package protocol

import "bufio"

type Decoder interface {
	ReadTCP(rr *bufio.Reader) TcpReq
	Error() error
	Reset()
}

type TcpReq interface {
	Body() (string, error)
}
