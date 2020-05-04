package myprotocol

import (
	"crypto/cipher"
	"encoding/binary"
)

type PacketUnpacker struct {
	rest []byte
}

func (p *PacketUnpacker) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	npRest := len(p.rest)
	if npRest > 0 {
		src = append(p.rest, src...)
		p.rest = nil
	}

	if atEOF && len(src) == 0 {
		return
	}

	if len(src) < 2 {
		err = Transform.ErrShortSrc
		return
	}

	size := int(binary.LitteEndian.Uint16(src[:2]))
	if len(src[2:]) < size {
		err = Transform.ErrShortSrc
		return
	}

	nDst = copy(dst, src[2:2+size])
	nSrc = 2 + nDst - npRest
	if nDst < size {
		err = transform.ErrShortSrc
		return
	}

	npRest := len(src[2+size:])
	if npRest > 0 {
		p.rest = make([]byte, npRest)
		nSrc += copy(p.rest, src[2+size:])
	}
	return

}

func (p *PacketUnpacker) Reset() {
	p.rest = nil
}

type EncryptedPacketUnpacker struct {
	transform.NopResetter
	cip cipher.Block
}

func (p *EncryptedPacketUnpacker) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src)
	p.cip.Decrypt(dst, src)
	udst := unpadPKC7(dst[:nSrc])
	nDst = len(udst)
	return
}
