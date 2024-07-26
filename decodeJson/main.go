package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("create json from struct");

	// parseIntoJson();
	decodeJson()

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

func decodeJson(){
	jsonData:=[]byte(`
	{
        "username": "Vipin",
        "userage": 21,
        "useremail": "vipin@gmail.com",
        "hobby": ["music","circket"]
    }
	`)
	// jsonData:=[]byte(`
	// {
    //     "username": "Vipin",
    //     "userage": 21,
    //     "useremail": "vipin@gmail.com",
    //     "hobby": ["music","circket"]
    // },
	// {
	
    //     "username": "Synder",
    //     "userage": 20,
    //     "useremail": "synder@gmail.com",
    //     "hobby": ["hacking","football"]
        
	// }
	// `)
	var userData User;
	isVaild:=json.Valid(jsonData);
	if !isVaild{
		fmt.Println("invalid json")
	}else{
		fmt.Println("json is vaild");
		json.Unmarshal(jsonData,&userData);
		fmt.Printf("%#v\n",userData)
	}

	var userDataMap map[string]interface{};
	json.Unmarshal(jsonData,&userDataMap);
	// fmt.Printf("%+v\n",userDataMap)

	for k,v:=range userDataMap{
		fmt.Printf("%v : %v\n",k,v)
	}

}