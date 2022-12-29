package store

import (
	"errors"
	"io/fs"
)

type LStore interface {

	Pop(item fs.FileInfo) error;
	Push(item fs.FileInfo) error;
	Get(key string) (fs.FileInfo, error);
};

type List[T interface{}] struct {

	Head *LInfo[T];
	Tail *LInfo[T]
}

type LInfo[T interface{}] struct {

	_data T;
	_next *LInfo[T];
	_prev *LInfo[T];
}

func (l* LInfo[T]) IsEmpty() bool {

	return l == nil;
}

func (l* List[T]) GetSize() int {

	var i int = 0;
	if l == nil {
		return i;
	}

	tmp := l;
	node := tmp.Head;

	for ; node._next.IsEmpty(); i++ {
		node = node._next;
	}

	return i + 1;
}

func (l* List[T]) Pop() error {

	return nil;
}

func (l* List[T]) Push(item T) error {

	if l == nil {
		return errors.New("error at init of list");
	}

	tmp := LInfo[T]{_data: item};
	if l.Head != nil {

		tmp._next = l.Head;
		tmp._prev = l.Tail;
		l.Tail = (&tmp);
	}
		
	l.Head = (&tmp);
	return nil;
}

func (l* LInfo[T]) Get(key string) (fs.FileInfo, error) {

	return nil, nil;
}
