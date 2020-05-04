package myprotocol

import (
	"crypto/cipher"
)

type PacketPacker struct {
	transform.NopResetter
}

func (p *PacketPacker) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src)
	binary.LittleEndian.PutUint16(dst, uint16(nSrc))
	nDst = copy(dst[2:], src)
	if nDst < nSrc {
		err = transform.ErrShortSrc
	}

	nDst += 2
	return
}

type EncryptedPacketPacker struct {
	transform.NopResetter
	cip cipher.Block
}

func (p *EncryptedPacketPacker) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nSrc = len(src)
	psrc := padPKC7(src)
	p.cip.Encrypt(dst, psrc)
	nDst = len(psrc)
	return
}
