package alg

import (
	"snappy"
)

/*
	Snappy was designed with CPU performance in mind, not compression %.
	On a single core of a Core i7 processor in 64-bit mode, it compresses at about
	250 MB/sec or more and decompresses at about 500 MB/sec or more
*/

type SnappyCompression struct{}

func (c SnappyCompression) Compress(data []byte) ([]byte, error) {
	if data == nil { //@todo research test handling of nil values with snappy
		return nil, nil
	}

	var zip []byte
	return snappy.Encode(zip, data), nil
}

func (c SnappyCompression) Decompress(data []byte) ([]byte, error) {
	if data == nil {
		return nil, nil
	}

	var dst []byte
	dcmp, err := snappy.Decode(dst, data)
	if err != nil {
		return nil, err
	}

	return dcmp, nil
}
