package store

import (
	"errors"
	"io/fs"
	"log"
)

type LStore interface {

	Pop(item fs.FileInfo) error;
	Push(item fs.FileInfo) error;
	Get(key string) (fs.FileInfo, error);
};

type LInfo struct {

	_data fs.FileInfo;
	Next *LInfo;
}

func (l* LInfo) Pop() error {

	if l == nil {
		return errors.New("no elem to remove");
	}

	tmp := (*l.Next)
	l = (&tmp);
	return nil;
}

func (l* LInfo) Push(item fs.FileInfo) error {

	if item == nil {
		return errors.New("no item send");
	}

	tmp := LInfo{_data: item};
	if l != nil {
		tmp.Next = l;
		l = (&tmp);
	}
	l = (&tmp);

	log.Println(l._data.Name());	
	return nil;
}

func (l* LInfo) Get(key string) (fs.FileInfo, error) {

	if len(key) == 0 || l == nil {
		return nil, errors.New("nothing to find here");
	}

	tmp := (*l);
	for ; tmp.Next != nil; tmp = (*tmp.Next) {

		if tmp._data.Name() == key {

			return tmp._data, nil;
		}
	} 
	return nil, errors.New("no file found");
}
