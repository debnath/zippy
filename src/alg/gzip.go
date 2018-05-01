package alg

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
)

type GzipCompression struct{}

func (c GzipCompression) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	//need to pass in the bytes.Buffer in this way...the io.Writer interface screws it up if we try to instantiate it inline
	err := gzipWrite(&buf, data)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c GzipCompression) Decompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	err := gunzipWrite(&buf, data)

	if err != nil {
		return nil, err
	}
	//buf.Write(data)

	return buf.Bytes(), nil
}

// Write gzipped data to a Writer
func gzipWrite(w io.Writer, data []byte) error {
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer gw.Close()
	gw.Write(data)
	return err
}

// Write gunzipped data to a Writer
func gunzipWrite(w io.Writer, zip []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(zip))
	defer gr.Close()
	zip, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(zip)
	return nil
}
