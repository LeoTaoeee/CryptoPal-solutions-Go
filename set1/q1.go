package set1

import (
	"encoding/hex"
	"encoding/base64"
)

//given a hexstring, turn it into base64
func HexToBase64(hexStr string) []byte {
	//decode hex, check for error
	hexBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}

	//make an empty array
	result := make([]byte, base64.StdEncoding.EncodedLen(len(hexBytes)))

	//write base64 encoding to said array
	base64.StdEncoding.Encode(result, hexBytes)

	//return
	return result
}
