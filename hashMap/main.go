package main

import (
	"fmt"

	"github.com/jorgeAM/structures/hashMap/hash"
)

func main() {
	hashTable := hash.Init()

	hashTable.Insert("Jorge")
	hashTable.Insert("Lessly")
	hashTable.Insert("Lili")
	hashTable.Insert("Basti")
	hashTable.Insert("Panchito")

	fmt.Println(hashTable.Search("Panchito"))

	hashTable.Delete("Jorge")

	fmt.Println(hashTable)
}
