package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type folder struct {
	backet   string `json:"backet"`
	s3folder string `json:"s3folder"`
	lfolder  string `json:"localfolder"`
}

func main() {

	content, err := ioutil.ReadFile("folders.json")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s\n", content)
	}

	var folders []folder
	// json.Unmarshal(content, &folders)
	json.Unmarshal(content, &folders)
	// if err2 != nil {
	// 	fmt.Println("Error JSON Unmarshalling")
	// 	fmt.Println(err2.Error())
	// }
	//	d := 0
	for _, x := range folders {
		//d++
		//	fmt.Println(d)
		fmt.Printf("%s \n", x)
	}
}
