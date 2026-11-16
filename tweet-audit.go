package main

import(

	"os"
	"fmt"
)

/**
create a struct for the things to return
read the file in memory
theres nothing about returning struct just bytes
wth do i do with the bytes

*/

type tweet struct{

	Id  string `json: id_str`
	Text string `json: full_text`
}

func main(){
//read file inn memory
	data, err := os.ReadFile("archive/tweets.js")
	if err != nil {
		fmt.Println(err)
	}return

}