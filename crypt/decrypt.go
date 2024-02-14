package crypt

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"

	"github.com/ProtonMail/go-crypto/eax"
	"golang.org/x/crypto/twofish"
)

func decompress(src []byte) ([]byte, error) {
	reader := bytes.NewReader(src)
	z, err := zlib.NewReader(reader)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	io.Copy(w, z)

	return b.Bytes(), nil
}

func Decrypt(src []byte) ([]byte, error) {
	key := [...]byte{137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137, 137}
	iv := [...]byte{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16}

	length := len(src)
	processed := make([]byte, length)

	// deobfuscation
	for i := 0; i < length; i++ {
		processed[i] = src[length+^i] ^ byte(length-i*length)
	}

	// decryption
	cipher, _ := twofish.NewCipher(key[:])
	decryptor, _ := eax.NewEAX(cipher)
	output, err := decryptor.Open(nil, iv[:], processed, nil)
	if err != nil {
		return nil, err
	}

	// deobfuscation
	for i := 0; i < len(output); i++ {
		output[i] = output[i] ^ byte(len(output)-i)
	}

	// decompression
	b, err := decompress(output[4:])
	if err != nil {
		return nil, err
	}

	return b, nil
}
