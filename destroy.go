package main



import (
"log"
"os"
"math/rand"
"strconv"
"time"
"fmt"
)



// SimpleFunction prints a message

func Shred(file string) {
	rand.Seed(time.Now().UnixNano())
	for i:=1; i < 4; i++ {
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {

			log.Fatal(err)

			}
		f.WriteString(strconv.Itoa(rand.Intn(10000000000)))

		if err := f.Close(); err != nil {

			log.Fatal(err)

		}
	}

	os.Remove(file)

}



func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your File Path: ")

	// var then variable name then variable type
	var file string

	fmt.Scanln(&file)

	Shred(file)

}
