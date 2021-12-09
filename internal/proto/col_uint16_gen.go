// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import "github.com/go-faster/errors"

// ColumnUInt16 represents UInt16 column.
type ColumnUInt16 []uint16

// Compile-time assertions for ColumnUInt16.
var (
	_ Input  = ColumnUInt16{}
	_ Result = (*ColumnUInt16)(nil)
)

// Type returns ColumnType of UInt16.
func (ColumnUInt16) Type() ColumnType {
	return ColumnTypeUInt16
}

// Rows returns count of rows in column.
func (c ColumnUInt16) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColumnUInt16) Reset() {
	*c = (*c)[:0]
}

// EncodeColumn encodes UInt16 rows to *Buffer.
func (c ColumnUInt16) EncodeColumn(b *Buffer) {
	for _, v := range c {
		b.PutUInt16(v)
	}
}

// DecodeColumn decodes UInt16 rows from *Reader.
func (c *ColumnUInt16) DecodeColumn(r *Reader, rows int) error {
	const size = 16 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			bin.Uint16(data[i:i+size]),
		)
	}
	*c = v
	return nil
}
