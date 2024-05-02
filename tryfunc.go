package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//type Folder string

	type Pc struct {
		Dir     string `json:"dir"`
		Files   string `json:"files"`
		Folders string `json:"folders"`
	}

	//jsonString := `{"dir": "root","files": ["cfrdytregmkhnmlop", "nive", "wmfxknb", "qgqesetq", "weakbeumihbdjtzde", "xiyniqcinfpwogunz", "mtrfqmewza", "bpozqgjmisanumxnjqqv"]}`
	jsonString := `{"dir": "root", "files": [".zshrc"], "folders": [{"dir": "desktop","files": ["config.yaml"]}]}`

	var byt = []byte(jsonString)

	var dat map[string]interface{}

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		fmt.Println("Не распаковался!!!!")
	}

	fmt.Printf("%T\n", dat["dir"])
	fmt.Printf("%T\n", dat["files"])
	fmt.Printf("%T\n", dat["folders"])

	ff, _ := dat["files"].([]string)
	fol, _ := dat["folders"].(string)

	fmt.Println("а теперь")
	fmt.Printf("%T\n", ff)
	fmt.Printf("%T\n", fol)
}

/*
func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
*/
