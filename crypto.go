package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func newEncryptionKey() []byte{
	keyBuf := make([]byte,32)
	io.ReadFull(rand.Reader, keyBuf)
	return keyBuf
}

func copyStream(stream cipher.Stream,blockSize int ,src io.Reader, dst io.Writer)(int, error){
	var (
		buf = make([]byte, 32*1024) // max amount we are gonna copy to memory
		nw = blockSize
	)
	for{
		n, err := src.Read(buf)
		
		if n > 0{
			stream.XORKeyStream(buf, buf[:n])
			nn, err := dst.Write(buf[:n])
			if err != nil {
				return 0, err
			}
			nw += nn
		}
		if err == io.EOF{
			break
		}
		if err != nil {
			return 0, err
		}
	}

	return nw, nil
}

func copyDecrypt(key []byte, src io.Reader, dst io.Writer) (int, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	// Read the IV from the given io.Reader which in our case should be the block.BlockSize() byes we read
	iv := make([]byte, block.BlockSize())
	if _, err := src.Read(iv); err != nil {
		return 0, err
	}

	stream := cipher.NewCTR(block, iv)
	return copyStream(stream, block.BlockSize(), src, dst)
}

func copyEncrypt(key []byte, src io.Reader, dst io.Writer) (int, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return 0, err
	}

	iv := make([]byte, block.BlockSize()) //16 bytes ig
	if _, err := io.ReadFull(rand.Reader, iv); err != nil{
		return 0, err
	}

	// prepend the IV to the file.
	if _,err := dst.Write(iv); err != nil{
		return 0, err
	}

	stream := cipher.NewCTR(block, iv)
	return copyStream(stream, block.BlockSize(), src, dst)
	}
	
