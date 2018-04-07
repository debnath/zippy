package alg

//Compression is an interface for all the various compression and decompression methods.
type Compression interface {
	Compress(u []byte) ([]byte, error)
	Decompress(z []byte) ([]byte, error)
}
