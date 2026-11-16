package main

import(

	"os"
	"fmt"
)

/**
create a struct for the things to return
read the file in memory
theres nothing about returning struct just bytes
wtf do i do with the bytes

*/

type  tweet{

	id_str  id_str
	full_text full_text
}

func main(){
//read file inn memory
	file, err := os.ReadFile()
	if err != nil {
		fmt.Println(err)
	}

}