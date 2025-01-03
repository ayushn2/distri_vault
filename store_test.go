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
	id := generateID()
	key := "momspecials"

	data := []byte("some jpg bytes")

	if _, err := s.writeStream(id, key,bytes.NewReader(data)); err !=nil{
		t.Errorf(err.Error())
	}

	if err := s.Delete(id,key);err != nil{
		t.Errorf(err.Error())
	}
}

func TestStore(t *testing.T){
	s := newStore()
	id := generateID()
	defer tearDown(t,s)

	for i := 0 ; i< 50 ; i++{

		key := fmt.Sprintf("pizz%d",i)
		data := []byte("some jpg bytes")

		if _, err := s.writeStream(id, key,bytes.NewReader(data)); err !=nil{
			t.Errorf(err.Error())
		}

		if ok := s.Has(id, key); !ok{
			t.Errorf("Expected to have key %s",key)
		}

		_, r, err := s.Read(id, key)
		if err!=nil{
			t.Errorf(err.Error())
		}

		b, _ := io.ReadAll(r)

		if string(b) != string(data){
			t.Errorf("want %s have %s",data,b)
		}

		if err := s.Delete(id, key); err != nil{
			t.Errorf(err.Error())
		}

		if ok := s.Has(id, key); ok{
			t.Errorf("expected to NOT hav key %s",key)
		}
	}

	
}

func newStore() *Store{
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func tearDown(t *testing.T,s *Store){
	if err := s.Clear();err != nil{
		t.Errorf(err.Error())
	}
}