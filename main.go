package main

import (
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
	
	items := new(store.LInfo);
	for _, v := range files {
		
		ok := items.Push(v); if ok != nil {

			panic(ok);
		}
		log.Println(items.Get(v.Name()));
	}
}
