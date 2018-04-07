package alg

type NoCompression struct{}

func (c NoCompression) Compress(data []byte) ([]byte, error) {
	return data, nil
}

func (c NoCompression) Decompress(data []byte) ([]byte, error) {
	return data, nil
}
