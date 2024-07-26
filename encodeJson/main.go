package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("create json from struct");

	parseIntoJson();

}
 
type User struct{
	Name string `json:"username"`
	Age int `json:"userage" `
	Email string `json:"useremail"`
	Password string `json:"-"`
	Hobby []string  `json:"hobby,omitempty"`
}

func parseIntoJson(){

	users:=[]User{
		{"Vipin",21,"vipin@gmail.com","abc11",[]string{"music","circket"}},
		{"Srijan",22,"srijan@gmail.com","fgc22",nil},
		{"Synder",20,"synder@gmail.com","synd56",[]string{"hacking","football"}},
	}
	jsonData,err:=json.MarshalIndent(users,"","\t");
	if err!=nil{
		panic(err);
	}
	// fmt.Println(jsonData) // ? byteData

	fmt.Printf("%s",jsonData)
}