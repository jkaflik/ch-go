// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ClickHouse/ch-go/internal/gold"
)

func TestColDateTime64_DecodeColumn(t *testing.T) {
	t.Parallel()
	const rows = 50
	var data ColDateTime64
	for i := 0; i < rows; i++ {
		data.Data = append(data.Data, DateTime64(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)
	t.Run("Golden", func(t *testing.T) {
		t.Parallel()
		gold.Bytes(t, buf.Buf, "col_datetime64")
	})
	t.Run("Ok", func(t *testing.T) {
		br := bytes.NewReader(buf.Buf)
		r := NewReader(br)

		var dec ColDateTime64
		require.NoError(t, dec.DecodeColumn(r, rows))
		require.Equal(t, data, dec)
		require.Equal(t, rows, dec.Rows())
		dec.Reset()
		require.Equal(t, 0, dec.Rows())

	})
	t.Run("ZeroRows", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColDateTime64
		require.NoError(t, dec.DecodeColumn(r, 0))
	})
	t.Run("ErrUnexpectedEOF", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColDateTime64
		require.ErrorIs(t, dec.DecodeColumn(r, rows), io.ErrUnexpectedEOF)
	})
	t.Run("NoShortRead", func(t *testing.T) {
		var dec ColDateTime64
		requireNoShortRead(t, buf.Buf, colAware(&dec, rows))
	})
	t.Run("ZeroRowsEncode", func(t *testing.T) {
		var v ColDateTime64
		v.EncodeColumn(nil) // should be no-op
	})
}

func BenchmarkColDateTime64_DecodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColDateTime64
	for i := 0; i < rows; i++ {
		data.Data = append(data.Data, DateTime64(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	br := bytes.NewReader(buf.Buf)
	r := NewReader(br)

	var dec ColDateTime64
	if err := dec.DecodeColumn(r, rows); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		br.Reset(buf.Buf)
		r.raw.Reset(br)
		dec.Reset()

		if err := dec.DecodeColumn(r, rows); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkColDateTime64_EncodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColDateTime64
	for i := 0; i < rows; i++ {
		data.Data = append(data.Data, DateTime64(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		data.EncodeColumn(&buf)
	}
}
