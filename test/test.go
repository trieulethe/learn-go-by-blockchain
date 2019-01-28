package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// const targetBits = 24

// func main() {
// 	target := big.NewInt(12345789)
// 	fmt.Println("target 1:", target)
// 	target.Lsh(target, uint(3))
// 	fmt.Println("target 2:", target)
// 	i := 3
// 	x := i << 4
// 	fmt.Println("x", x)
// }

// serialize block
func Serialize(b string) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// deserialize block
func DeserializeBlock(d []byte) {
	var b string

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&b)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("deserialize:", b)
}

func main() {

	serialize := Serialize("test 01")
	fmt.Println("serialize:", serialize)
	DeserializeBlock(serialize)

	// s := []byte("hello world")
	// // Open the my.db data file in your current directory.
	// // It will be created if it doesn't exist.
	// db, err := bolt.Open("my.db", 0600, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db.Update(func(tx *bolt.Tx) error {
	// 	b, err := tx.CreateBucket([]byte("blocks"))
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	err = b.Put([]byte("1"), s)
	// 	return nil
	// })

	// db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("blocks"))
	// 	v := b.Get([]byte("1"))
	// 	fmt.Printf("The answer is: %s\n", v)
	// 	return nil
	// })

	// defer db.Close()
}
