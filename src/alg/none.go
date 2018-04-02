package alg

type NoCompression struct{}

func (c NoCompression) Zip(unzip []byte) []byte {
	return unzip
}

func (c NoCompression) Unzip(zip []byte) ([]byte, error) {
	return zip, nil
}
