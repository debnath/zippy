package alg

import "fmt"

//Compression interface for all the various compression methods.
type Compression interface {
	Zip(u []byte) []byte
	Unzip(z []byte) ([]byte, error)
}

func nilContentError() error {
	return fmt.Errorf("Cannot zip or unzip nil content")
}