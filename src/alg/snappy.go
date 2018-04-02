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

func (c SnappyCompression) Zip(unzip []byte) []byte {
	if unzip == nil {
		return nil
	}

	var zip []byte
	return snappy.Encode(zip, unzip)
}

func (c SnappyCompression) Unzip(zip []byte) ([]byte, error) {
	if zip == nil {
		return nil, nil
	}

	var dst []byte
	unzip, err := snappy.Decode(dst, zip)
	if err != nil {
		return nil, err
	}

	return unzip, nil
}
