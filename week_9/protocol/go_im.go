package protocol

// 实现一个从 socket connection 中解码出 goim 协议的解码器。
import (
	"bufio"

	"github.com/pkg/errors"
)

var (
	ErrIllegalHeader = errors.New("header length error")
)

const (
	MaxBodySize = int32(1 << 12)
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

type GoImReq struct {
	ProtocolVersion int32 //协议版本
	Operation       int32 //操作码
	Seq             int32 //请求序号 ID
	Content         string
}

func (p *GoImReq) Body() (string, error) {
	return p.Content, nil
}

type GoImProto struct {
	buf []byte
	err error
	dst *GoImReq
}

func (p *GoImProto) Reset() {
	p.err = nil
	p.dst = nil
	p.buf = nil
}

func (p *GoImProto) ReadTCP(rr *bufio.Reader) TcpReq {
	p.buf, p.err = rr.Peek(_rawHeaderSize)
	if p.err != nil {
		return nil
	}
	pkgLen := p.readInt32(_packOffset, _packSize)
	headerLen := p.readInt32(_headerOffset, _headerSize)

	if headerLen != _rawHeaderSize {
		p.err = ErrIllegalHeader
		return nil
	}
	p.dst.ProtocolVersion = p.readInt32(_verOffset, _verSize)
	p.dst.Operation = p.readInt32(_opOffset, _opSize)
	p.dst.Seq = p.readInt32(_seqOffset, _seqSize)

	if bodyLen := pkgLen - headerLen; bodyLen > 0 {
		p.buf, p.err = rr.Peek(int(pkgLen))
		if p.err != nil {
			return nil
		}
		p.dst.Content = p.readString(_rawHeaderSize, int(bodyLen))
	}
	return p.dst
}

func (p *GoImProto) readInt32(i int, step int) int32 {
	if step <= 0 {
		step = 1
	}
	bytes := p.buf[i : i+step]
	var tmp = int32(bytes[step-1])
	if step == 1 {
		return tmp
	}
	for i := 1; i <= step-1; i++ {
		tmp = tmp | int32(bytes[step-1-i])<<i*8
	}
	return tmp
}

func (p *GoImProto) readString(start int, step int) string {
	return string(p.buf[start : start+step])
}

func (p *GoImProto) Error() error {
	return p.err
}

func NewGoImProto() *GoImProto {
	return &GoImProto{
		dst: &GoImReq{},
	}
}
