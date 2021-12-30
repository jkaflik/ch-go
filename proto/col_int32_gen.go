// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
)

var _ = binary.LittleEndian // clickHouse uses LittleEndian

// ColInt32 represents Int32 column.
type ColInt32 []int32

// Compile-time assertions for ColInt32.
var (
	_ ColInput  = ColInt32{}
	_ ColResult = (*ColInt32)(nil)
	_ Column    = (*ColInt32)(nil)
)

// Type returns ColumnType of Int32.
func (ColInt32) Type() ColumnType {
	return ColumnTypeInt32
}

// Rows returns count of rows in column.
func (c ColInt32) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColInt32) Row(i int) int32 {
	return c[i]
}

// Append int32 to column.
func (c *ColInt32) Append(v int32) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt32) Reset() {
	*c = (*c)[:0]
}

// NewArrInt32 returns new Array(Int32).
func NewArrInt32() *ColArr {
	return &ColArr{
		Data: new(ColInt32),
	}
}

// AppendInt32 appends slice of int32 to Array(Int32).
func (c *ColArr) AppendInt32(data []int32) {
	d := c.Data.(*ColInt32)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Int32 rows to *Buffer.
func (c ColInt32) EncodeColumn(b *Buffer) {
	const size = 32 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint32(
			b.Buf[offset:offset+size],
			uint32(v),
		)
		offset += size
	}
}
