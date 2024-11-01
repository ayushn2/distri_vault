package main

import (
	"bytes"
	"fmt"
	"io"

	"testing"
)

func TestPathTransformFunc(t *testing.T){
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)

	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathKey.Pathname != expectedPathName{
		t.Errorf("have %s want %s",pathKey.Pathname,expectedPathName)
	}
	if pathKey.Filename != expectedOriginalKey{
		t.Errorf("have %s want %s",pathKey.Filename,expectedOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T){
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momspecials"

	data := []byte("some jpg bytes")

	if err := s.writeStream(key,bytes.NewReader(data)); err !=nil{
		t.Errorf(err.Error())
	}

	if err := s.Delete(key);err != nil{
		t.Errorf(err.Error())
	}
}

func TestStore(t *testing.T){
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momspecials"

	data := []byte("some jpg bytes")

	if err := s.writeStream(key,bytes.NewReader(data)); err !=nil{
		t.Errorf(err.Error())
	}

	r, err := s.Read(key)
	if err!=nil{
		t.Errorf(err.Error())
	}

	b, _ := io.ReadAll(r)

	if string(b) != string(data){
		t.Errorf("want %s have %s",data,b)
	}

	fmt.Println(string(b))

	s.Delete(key)
}