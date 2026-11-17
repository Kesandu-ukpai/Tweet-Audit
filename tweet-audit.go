package main

import(

	"os"
	"fmt"
	"bytes"
	"encoding/json"
)

/**
create a struct for the things to return
read the file in memory
theres nothing about returning struct just bytes
wth do i do with the bytes

*/

type tweet struct{

	Id  string`json:"id_str"`
	Text string`json:"full_text"`
}

type item struct{
	//the last tweet in json is refencing the one in the archive tweet.js file
	Tweet tweet`json:"tweet"`
}

func main(){
//read file inn memory
	data, err := os.ReadFile("archive/tweets.js")
	if err != nil {
		fmt.Println(err)
			return
	}

//cut of the unecessary part that is causing problem
	line := bytes.IndexByte(data, '[')
	if line < 0 {
		fmt.Println("lower than 0")
		return
	}
	//i coulddnt think of anything better so for now
	data = data[line:]

	var items []item

	err = json.Unmarshal(data, &items)
	if err != nil{
		fmt.Println(err)
		return
	}
	

	//probably writing rubbish at this point
	fmt.Println(len(items), items[0].Tweet.Id, items[0].Tweet.Text)
}