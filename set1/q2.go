package set1

import (
	"encoding/hex"
	"errors"
)

func Fixed_Xor(hexStr1 string, hexStr2 string) ([]byte, error){
	//check input length
	if len(hexStr1) != len(hexStr2) {
		return []byte{}, errors.New("input strings lengths not equal")
	}
	//decode hex, check for error
	hexBytes1, err := hex.DecodeString(hexStr1)
	if err != nil {
		panic(err)
	}

	//repeat
	hexBytes2, err := hex.DecodeString(hexStr2)
	if err != nil {
		panic(err)
	}

	//bit wise Xor
	for i := range hexBytes1 {
		hexBytes2[i] ^= hexBytes1[i]
	}

	//make empty array
	var result = make([]byte, len(hexStr1))

	//encode into result
	hex.Encode(result, hexBytes2)

	//return
	return result, nil
}