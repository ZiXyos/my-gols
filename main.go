package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"github.com/Zixyos/my-gols/store"
)

func isError(err error) {

	if err != nil {
		panic(err);
	}
}

func main() {

	args := os.Args[1:];
	log.Println(args);

	files, err := ioutil.ReadDir("./");
	isError(err);
	items := store.List[fs.FileInfo]{ Head: nil, Tail: nil};
	for _, v := range files {
		
		ok := items.Push(v); if ok != nil {

			panic(ok);
		}
		log.Println("[Log]: ", items.Head);
	}
}
