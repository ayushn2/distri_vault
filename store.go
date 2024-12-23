package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// filename => clown.jpg
// pat =>

const defaultRootFolderName = "nainwalnetwork"

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

func (p  PathKey) FirstPathName() string{
	paths := strings.Split(p.Pathname, "/")
	if len(paths) == 0{
		return ""
	}
	return paths[0]
}

func (p PathKey) FullPath() string{
	return fmt.Sprintf("%s/%s",p.Pathname,p.Filename)
}

type StoreOpts struct{
	// Root is the folder name of the root , containing all the folders/files of the system 
	Root string
	PathTransformFunc PathTransformFunc
}

var DefaultTransformFunc = func(key string) PathKey{
	return PathKey{
		Pathname: key,
		Filename: key,
	}
}

type Store struct{
	StoreOpts
}

func NewStore(opts StoreOpts) *Store{
	if opts.PathTransformFunc == nil{
		opts.PathTransformFunc = DefaultTransformFunc
	}
	if len (opts.Root) == 0{
		opts.Root = defaultRootFolderName
	}

	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Has(id string,key string) bool{
	pathKey := s.PathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s/%s",s.Root,id,pathKey.FullPath())

	_, err := os.Stat(fullPathWithRoot)

	return !errors.Is(err,os.ErrNotExist)
}

func (s *Store) Clear() error{
	return os.RemoveAll(s.Root)
}

func (s *Store) Delete(id string, key string) error{
	pathKey := s.PathTransformFunc(key)
	
	defer func(){
		log.Printf("deleted [%s] from disk",pathKey)
	}()

	firstPathNameWithRoot := fmt.Sprintf("%s/%s/%s",s.Root,id,pathKey.FirstPathName())

	return os.RemoveAll(firstPathNameWithRoot)
}

func (s *Store)Write (id string,key string,r io.Reader) (int64, error){
	return s.writeStream(id,key,r)
}

func (s *Store) WriteDecrypt(id string,encKey []byte, key string, r io.Reader)(int64, error){
	f, err := s.openFileForWriting(id, key)
	if err != nil{
		return 0, err
	}

	n, err := copyDecrypt(encKey, r, f)
	return int64(n), nil
}

func (s *Store) openFileForWriting(id string, key string) (*os.File, error){
	pathKey := s.PathTransformFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s/%s",s.Root,id,pathKey.Pathname)

	if err := os.MkdirAll(pathNameWithRoot,os.ModePerm); err != nil{
		return nil,err
	}

	fullPathWithRoot := fmt.Sprintf("%s/%s/%s",s.Root,id,pathKey.FullPath())
	return os.Create(fullPathWithRoot)
}

func (s *Store) writeStream(id string, key string,r io.Reader) (int64, error){
	f, err := s.openFileForWriting(id, key)
	if err != nil{
		return 0, err
	}
	return io.Copy(f, r)

}

// FIXME: Instead of copying directly to a reader , we first copy this into a buffer. Maybe just return the file from the readstream? (Fixed)
func (s *Store) Read(id string, key string)(int64, io.Reader, error){
	return s.readStream(id, key)
}

func (s *Store) readStream (id string, key string) (int64, io.ReadCloser, error){
	pathKey := s.PathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s/%s",s.Root,id,pathKey.FullPath())

	file, err := os.Open(fullPathWithRoot)
	if err != nil{
		return 0, nil, err
	}
	fi, err := file.Stat()
	if err != nil{
		return 0, nil, err
	}
	
	return fi.Size(), file, nil
}

