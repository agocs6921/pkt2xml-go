package crypt

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"

	"github.com/ProtonMail/go-crypto/eax"
	"golang.org/x/crypto/twofish"
)

func compress(src []byte) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	writer := zlib.NewWriter(buffer)
	defer writer.Close()
	_, err := writer.Write(src)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func Encrypt(src []byte) ([]byte, error) {
	key := [...]byte{137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137}
	iv := [...]byte{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16}

	// compression
	compressed, err := compress(src)
	if err != nil {
		return nil, err
	}

	// adding size bytes
	size := uint32(len(src))
	size_slice := make([]byte, 4)
	binary.BigEndian.PutUint32(size_slice, size)
	compressed = append(size_slice, compressed...)

	// obfuscation
	for i := 0; i < len(compressed); i++ {
		compressed[i] ^= byte(len(compressed) - i)
	}

	// encryption
	cipher, _ := twofish.NewCipher(key[:])
	encryptor, _ := eax.NewEAX(cipher)
	encrypted := encryptor.Seal(nil, iv[:], compressed, nil)

	// obfuscation
	length := len(encrypted)
	output := make([]byte, length)
	for i := 0; i < length; i++ {
		output[length+^i] = encrypted[i] ^ byte(length-i*length)
	}

	return output, nil
}
