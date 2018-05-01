package zippy_test

import (
	"testing"

	"reflect"

	"fmt"

	"github.com/debnath/zippy"
)

//Nothing fancy....  probably should switch to ginkgo/gomega at some point.
func assertEqual(t *testing.T, exp interface{}, act interface{}, msg string) {
	if !reflect.DeepEqual(exp, act) {
		t.Fatalf("%s: EXPECTED: %s != ACTUAL: %s", msg, exp, act)
	}
}

func TestSnappy(t *testing.T) {
	b := []byte("testing snappy")

	zpy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_SNAPPY,
	})

	z, _ := zpy.Compress(b)
	unzp, err := zpy.Decompress(z)

	expected := b
	actual := unzp

	assertEqual(t, string(expected[:]), string(actual[:]), "SNAPPY zip and unzip")
	assertEqual(t, err, nil, "SNAPPY error is nil")
}

//Snappy errors out if you try to decompress nil values... needs some investigation
func TestSnappyNil(t *testing.T) {
	zpy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_SNAPPY,
	})

	cmpr, err := zpy.Compress(nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	dcmp, errd := zpy.Decompress(cmpr)

	if errd != nil {
		fmt.Println(errd.Error())
	}
	nilBytes := []byte(nil)

	assertEqual(t, nilBytes, cmpr, "SNAPPY compressed with nil input")
	assertEqual(t, nilBytes, dcmp, "SNAPPY decompressed with nil input")
	assertEqual(t, nil, err, "SNAPPY compress err is nil with nil input")
	assertEqual(t, nil, errd, "SNAPPY decompress err is nil with nil input")

}

func TestGzip(t *testing.T) {
	b := []byte("testing gzip")

	zpy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_GZIP,
	})

	cmpr, err := zpy.Compress(b)
	if err != nil {
		fmt.Println("error on compression", err.Error())
	}
	dcmp, errd := zpy.Decompress(cmpr)

	if errd != nil {
		fmt.Println("error on decompression", errd.Error())
	}

	assertEqual(t, err, nil, "GZIP no error when compressing")
	assertEqual(t, errd, nil, "GZIP no error when decompressing")
	assertEqual(t, string(b[:]), string(dcmp[:]), "GZIP can compress and decompress byte string")


}

//gzip does not error out when compressing or decompression nil values... so there is no nil handling
func TestGzipNil(t *testing.T) {
	zpy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_GZIP,
	})

	cmpr, err := zpy.Compress(nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	unzp, errd := zpy.Decompress(cmpr)
	if errd != nil {
		fmt.Println(errd.Error())
	}

	nilBytes := []byte(nil)

	assertEqual(t, nilBytes, unzp, "GZIP decompressed with nil input")
	assertEqual(t, nil, err, "GZIP err is nil with COMPRESSING nil input")
	assertEqual(t, nil, errd, "GZIP err is nil with DECOMPRESSING nil input")
}

func TestNone(t *testing.T) {
	b := []byte("testing no compression")

	zpy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_NONE,
	})

	z, _ := zpy.Compress(b)
	unzp, err := zpy.Decompress(z)

	expected := b
	actual := unzp

	assertEqual(t, string(expected[:]), string(actual[:]), "NO COMPRESSION zip and unzip")
	assertEqual(t, err, nil, "NO COMPRESSION error is nil")
}

func TestNoneNil(t *testing.T) {
	zpy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_NONE,
	})

	zp, _ := zpy.Compress(nil)
	unzp, err := zpy.Decompress(zp)
	nilBytes := []byte(nil)

	assertEqual(t, nilBytes, zp, "NO COMPRESSION compressed with nil input")
	assertEqual(t, nilBytes, unzp, "NO COMPRESSION decompressed with nil input")
	assertEqual(t, nil, err, "NO COMPRESSION err is nil with nil input")
}
