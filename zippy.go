package zippy

import (
	"github.com/debnath/zippy/src/alg"
	"strings"
)

const (
	COMPRESSION_SNAPPY = "snappy"
	COMPRESSION_GZIP = "gzip"
	COMPRESSION_NONE = "none"
)

type Config struct {
	CompressionFormat string
}

type Zippy struct {
	Zip func([]byte) []byte
	Unzip func([]byte) ([]byte, error)
}

func New(c Config) Zippy{
	var cmp alg.Compression
	switch strings.ToLower(c.CompressionFormat) {
	case COMPRESSION_SNAPPY:
		cmp = alg.SnappyCompression{}
	case COMPRESSION_NONE:
		cmp = alg.NoCompression{}
	default:
		cmp = alg.SnappyCompression{}
	}

	z := Zippy{}
	z.Zip = cmp.Zip
	z.Unzip = cmp.Unzip

	return z
}

/*
func main() {
	content := []byte("test string for compression")

	zippy := New(Config{
		CompressionFormat: "none",
	})

	//Zipping
	zp := zippy.Zip(content)

	//Unzipping
	uz, _ := zippy.Unzip(zp)

	fmt.Println("unzipped string:", string(uz[:]))
}
*/