package main

import(

	"os"
	"fmt"
	"bytes"
	"encoding/json"
	"strings"
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
	Date string`json:"created_at"`
}

type item struct{
	//the last tweet in json is refencing the one in the archive tweet.js file
	Tweet tweet`json:"tweet"`
}

type criteria struct{
	Words []string`json:"forbidden_words"`
	Rule []string`json:"judgment_rules"`
}

type config struct{
	Criteria criteria`json:"criteria"`
	APIkey string`json:"api_key"`
}


func loadConfig() (config, error){
	file, err := os.ReadFile("config.json")
	if err != nil{
		fmt.Println(err)
		return config{}, err
	}

	var rules config

	err = json.Unmarshal(file, &rules)
	if err != nil {
		fmt.Println(err)
		return config{},err
	}
	return rules, nil
}

//function that return true or false based on sensitive words, just one tweet
func flagged(text string, words []string) bool{
	//takes the the tweet values turn them to lowercase
	lowerText:= strings.ToLower(text)
	//check for the flagged words

	for _, value := range words{
		if strings.Contains(lowerText, strings.ToLower(value)){
			return true
		}
		
	}

	return false

	
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
	
	data = data[line:]

	var items []item

	err = json.Unmarshal(data, &items)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	//loop the whole tweet and  pass it to flagged one at a time
	value := 0
	for _, it := range items{
		if flagged(it.Tweet.Text, cfg.Criteria.Words) == true{
			value++
		}
	}
	
	fmt.Println(value)
	fmt.Println(len(items), items[0].Tweet.Id, items[0].Tweet.Text, items[0].Tweet.Date)
}