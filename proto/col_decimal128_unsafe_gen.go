//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes Decimal128 rows from *Reader.
func (c *ColDecimal128) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([]Decimal128, rows)...)
	s := *(*slice)(unsafe.Pointer(c))
	s.Len *= 16
	s.Cap *= 16
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}

// EncodeColumn encodes Decimal128 rows to *Buffer.
func (c ColDecimal128) EncodeColumn(b *Buffer) {
	if len(c) == 0 {
		return
	}
	offset := len(b.Buf)
	const size = 128 / 8
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)

	s := *(*slice)(unsafe.Pointer(&c))
	s.Len *= 16
	s.Cap *= 16
	src := *(*[]byte)(unsafe.Pointer(&s))
	dst := b.Buf[offset:]
	copy(dst, src)
}
