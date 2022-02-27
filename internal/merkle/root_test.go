package merkle_test

import (
	"reflect"
	"testing"

	"github.com/erdongli/pbchain/internal/merkle"
)

func TestComputeRoot(t *testing.T) {
	tests := []struct {
		in   [][]byte
		want []byte
	}{
		{
			in: [][]byte{},
			want: []byte{
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
		},
		{
			in: [][]byte{
				[]byte("hello world!"),
			},
			want: []byte{
				0x13, 0x32, 0xf5, 0x13,
				0x2b, 0x45, 0xde, 0x03,
				0xe7, 0x29, 0xe7, 0x64,
				0xb3, 0x8f, 0x9a, 0x2e,
				0x40, 0x57, 0xd6, 0x76,
				0x0a, 0x88, 0x07, 0x66,
				0x23, 0x31, 0xb3, 0x7d,
				0xde, 0xa9, 0x36, 0xe4,
			},
		},
		{
			in: [][]byte{
				[]byte("hello world!"),
				[]byte("hello world!"),
			},
			want: []byte{
				0x6b, 0x27, 0x36, 0x8d,
				0x5f, 0x46, 0xf3, 0x85,
				0x61, 0xcf, 0x76, 0xfe,
				0x5a, 0xd0, 0xad, 0x15,
				0x96, 0x8e, 0x41, 0xaa,
				0x01, 0xbc, 0x90, 0xc4,
				0xb8, 0x26, 0xa2, 0x8f,
				0xf6, 0x00, 0xb7, 0x48,
			},
		},
		{
			in: [][]byte{
				[]byte("hello world!"),
				[]byte("hello world!"),
				[]byte("hello world!"),
			},
			want: []byte{
				0x27, 0xc5, 0x24, 0xcd,
				0x10, 0x0c, 0xde, 0x53,
				0xd8, 0x42, 0x56, 0x0e,
				0x53, 0xf4, 0x61, 0x9f,
				0xc4, 0x6b, 0xcd, 0x20,
				0xf4, 0xd8, 0x5e, 0x38,
				0xe4, 0xdc, 0x30, 0x0d,
				0xed, 0x5f, 0x83, 0xbf,
			},
		},
		{
			in: [][]byte{
				[]byte("hello world!"),
				[]byte("hello world!"),
				[]byte("hello world!"),
				[]byte("hello world!"),
			},
			want: []byte{
				0x27, 0xc5, 0x24, 0xcd,
				0x10, 0x0c, 0xde, 0x53,
				0xd8, 0x42, 0x56, 0x0e,
				0x53, 0xf4, 0x61, 0x9f,
				0xc4, 0x6b, 0xcd, 0x20,
				0xf4, 0xd8, 0x5e, 0x38,
				0xe4, 0xdc, 0x30, 0x0d,
				0xed, 0x5f, 0x83, 0xbf,
			},
		},
	}

	for _, tc := range tests {
		got := merkle.ComputeRoot(tc.in)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
