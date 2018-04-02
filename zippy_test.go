package zippy_test

import (
	"testing"

	"reflect"

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

	zippy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_SNAPPY,
	})

	z := zippy.Zip(b)
	unzp, err := zippy.Unzip(z)

	expected := b
	actual := unzp

	assertEqual(t, string(expected[:]), string(actual[:]), "SNAPPY zip and unzip")
	assertEqual(t, err, nil, "SNAPPY error is nil")
}

func TestSnappyNil(t *testing.T) {
	zippy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_SNAPPY,
	})

	zp := zippy.Zip(nil)
	unzp, err := zippy.Unzip(zp)
	nilBytes := []byte(nil)

	assertEqual(t, nilBytes, zp, "SNAPPY zipped with nil input")
	assertEqual(t, nilBytes, unzp, "SNAPPY unzipped with nil input")
	assertEqual(t, nil, err, "SNAPPY err is nil with nil input")
}

func TestNone(t *testing.T) {
	b := []byte("testing no compression")

	zippy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_NONE,
	})

	z := zippy.Zip(b)
	unzp, err := zippy.Unzip(z)

	expected := b
	actual := unzp

	assertEqual(t, string(expected[:]), string(actual[:]), "NO COMPRESSION zip and unzip")
	assertEqual(t, err, nil, "NO COMPRESSION error is nil")
}

func TestNoneNil(t *testing.T) {
	zippy := zippy.New(zippy.Config{
		CompressionFormat: zippy.COMPRESSION_NONE,
	})

	zp := zippy.Zip(nil)
	unzp, err := zippy.Unzip(zp)
	nilBytes := []byte(nil)

	assertEqual(t, nilBytes, zp, "NO COMPRESSION zipped with nil input")
	assertEqual(t, nilBytes, unzp, "NO COMPRESSION unzipped with nil input")
	assertEqual(t, nil, err, "NO COMPRESSION err is nil with nil input")
}
