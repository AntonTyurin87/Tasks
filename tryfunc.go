package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Folder struct {
	Dir   string   `json:"dir"`
	Files []string `json:"files"`
}

type Target struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders,omitempty"` // Пропускается, если значение 0
}

func main() {

	//jsonString := `{"dir": "root","files": ["cfrdytregmkhnmlop", "nive", "wmfxknb", "qgqesetq", "weakbeumihbdjtzde", "xiyniqcinfpwogunz", "mtrfqmewza", "bpozqgjmisanumxnjqqv"]}`
	jsonString := `{"dir": "root", "files": [".zshrc"], "folders": [{"dir": "desktop","files": ["config.yaml"]}]}`

	//jsonString := `{"dir": "root","files": ["cfrdytregmkhnmlop", "nive", "wmfxknb", "qgqesetq", "weakbeumihbdjtzde", "xiyniqcinfpwogunz", "mtrfqmewza", "bpozqgjmisanumxnjqqv"]}`
	CountFiles := 0

	Flag := true

	CountFiles, Flag = fileCounter(jsonString, CountFiles, Flag)

	fmt.Println(CountFiles, Flag)

}

func fileCounter(jsonString string, CountFiles int, Flag bool) (int, bool) {

	stringForParse := fmt.Sprint("[", jsonString, "]")

	targets := []Target{}

	err := json.Unmarshal([]byte(stringForParse), &targets)
	if err != nil {
		fmt.Println(err)
	}

	for _, t := range targets {
		//fmt.Println("Файлов в папке", len(t.Files))

		if !Flag {
			CountFiles += len(t.Files)
			break
		}

		for _, values := range t.Files {
			matched, _ := regexp.MatchString(`.hack\z`, string(values))
			if matched {
				Flag = false
				CountFiles += len(t.Files)
			}
		}

		//CountFiles += len(t.Files)

		if t.Folders != nil {
			//fmt.Println("Папка есть!")
			//fmt.Println(t.Folders)

			b, _ := json.Marshal(t.Folders)

			fmt.Println(string(b[1 : len(b)-1]))

			prov := string(b[1 : len(b)-1])

			CountFiles, Flag = fileCounter(prov, CountFiles, Flag)
		}

		//fmt.Printf("%T\n", t.Folders)
		//fmt.Println([]byte(t.Folders))

	}
	return CountFiles, Flag
}

/*
	var byt = []byte(jsonString)

	var dat map[string]string

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		fmt.Println("Не распаковался!!!!")
	}

	fmt.Println(dat["dir"])
	fmt.Println(dat["files"])
	fmt.Println(dat["folders"])

	dir, _ := dat["dir"]
	files, _ := dat["files"]

	fmt.Println(dir)
	fmt.Println(files)

	/*
		fmt.Printf("%T\n", dat["dir"])
		fmt.Printf("%T\n", dat["files"])
		fmt.Printf("%T\n", dat["folders"])

		dir, _ := dat["files"]
		files, _ := dat["files"]
		folders, _ := dat["folders"]

		fmt.Printf("%T\n", dir)
		fmt.Printf("%T\n", files)
		fmt.Printf("%T\n", folders)

		fmt.Println(dir)
		fmt.Println(files)
		fmt.Println(folders)

		fmt.Println(dat["dir"])
		fmt.Println(dat["files"])
		fmt.Println(dat["folders"])
*/

//CountFiles := 0
//Flag := true

//CountFiles, Flag = fileCounter(jsonString, CountFiles, Flag)

//fmt.Println(CountFiles, Flag)

/*
func fileCounter(jsonString string, CountFiles int, Flag bool) (int, bool) {

	var byt = []byte(jsonString)
	var dat map[string]string

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		fmt.Println("Не распаковался!!!!")
	}

	//files, _ := dat["files"].([]string)

	//fmt.Println(files)

	if folders, ok := dat["folders"]; ok {
		fileCounter(folders, CountFiles, Flag)
	}

	return (CountFiles + len(files)), Flag
}

*/

//var folderString []string
/*
	if folderString, ok := dat["folders"]; ok {

		fmt.Println("перешли к папке")

		fS, _ := folderString.(string)

		fmt.Printf("%T\n", fS)
		fmt.Println(fS)
	}

	//fmt.Println(folderString)

	//fmt.Printf("%T\n", dat["dir"])
	//fmt.Printf("%T\n", dat["files"])
	//fmt.Printf("%T\n", dat["folders"])

	//ff, _ := dat["files"].([]string)
	//fol, _ := dat["folders"].(string)

	//fmt.Println("а теперь")
	//fmt.Printf("%T\n", ff)
	//fmt.Printf("%T\n", fol)
}

func fileCounter(jsonFloor map[string]interface{}, CountFiles int, Flag bool) (int, bool) {

	fmt.Println("Здесь работает fileCounter")

	if folderFiles, ok := jsonFloor["files"]; ok {
		files, _ := folderFiles.(string)

		fmt.Println(files)

		for _, values := range files {
			matched, _ := regexp.MatchString(`.hack\z`, string(values))
			if matched && !Flag {
				Flag = true
				CountFiles = len(files)
			}
		}
		CountFiles += len(files)
		fmt.Println("Здесь работает")
		fmt.Println("Файлов ", CountFiles)
		fmt.Println("Флаг ", Flag)
	}

	if _, ok := jsonFloor["folders"]; ok {

		folderIn, _ := jsonFloor["folders"].(string)

		var byt = []byte(folderIn)

		var dat map[string]interface{}

		err := json.Unmarshal(byt, &dat)
		if err != nil {
			fmt.Println("Не распаковался!!!!")
		}

		CountFiles, Flag = fileCounter(dat, CountFiles, Flag)
	}
	//fmt.Println("Файлов ", CountFiles)
	//fmt.Println("Флаг ", Flag)
	return CountFiles, Flag
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
