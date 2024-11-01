package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) PathKey{
	hash := sha1.Sum([]byte(key))// [20]byte => []byte => [:]
	hashStr := hex.EncodeToString(hash[:])
	blocksize := 5
	sliceLen := len(hashStr) / blocksize

	paths := make([]string,sliceLen)

	for i :=0 ; i < sliceLen ; i++{
		from, to := i * blocksize, (i * blocksize) + blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname : strings.Join(paths,"/"),
		Filename :  hashStr,
	}
}

type PathTransformFunc func(string) PathKey

type PathKey struct{
	Pathname string
	Filename string
}


func (p PathKey) FullPath() string{
	return fmt.Sprintf("%s/%s",p.Pathname,p.Filename)
}

type StoreOpts struct{
	PathTransformFunc PathTransformFunc
}

var DefaultTransformFunc = func(key string) string{
	return key
}

type Store struct{
	StoreOpts
}

func NewStore(opts StoreOpts) *Store{
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Read(key string)(io.Reader, error){
	f,err := s.readStream(key)
	if err!=nil{
		return nil,err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_,err = io.Copy(buf, f)
	f.Close()

	return buf, nil
}

func (s *Store) readStream (key string) (io.ReadCloser, error){
	pathKey := s.PathTransformFunc(key)
	f,err := os.Open(pathKey.FullPath())

	if err!=nil{
		return  nil,err
	}

	return f,nil
}

func (s *Store) writeStream(key string,r io.Reader) error{
	pathKey := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathKey.Pathname,os.ModePerm); err != nil{
		return err
	}



	fullPath := pathKey.FullPath()

	

	f, err := os.Create(fullPath)
	if err!=nil{
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil{
		return err
	}

	log.Printf("written (%d) bytes to disk: %s",n,fullPath)

	return nil
}

