package zippy

import (
	"strings"

	"github.com/debnath/zippy/src/alg"
)

const (
	COMPRESSION_SNAPPY = "snappy"
	COMPRESSION_GZIP   = "gzip"
	COMPRESSION_NONE   = "none"
)

type Config struct {
	CompressionFormat string
}

type Zippy interface {
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}

func New(c Config) Zippy {
	var cmp alg.Compression
	switch strings.ToLower(c.CompressionFormat) {
	case COMPRESSION_SNAPPY:
		cmp = alg.SnappyCompression{}
	case COMPRESSION_GZIP:
		cmp = alg.GzipCompression{}
	case COMPRESSION_NONE:
		cmp = alg.NoCompression{}
	default:
		cmp = alg.SnappyCompression{}
	}

	return cmp
}
