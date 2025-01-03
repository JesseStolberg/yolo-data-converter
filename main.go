package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type c struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type i struct {
	Id       int    `json:"id"`
	FileName string `json:"file_name"`
}

type a struct {
	ImageId    int       `json:"image_id"`
	CategoryId int       `json:"category_id"`
	Bbox       []float64 `json:"bbox"`
}

type jdoc struct {
	Categories  []c `json:"categories"`
	Images      []i `json:"images"`
	Annotations []a `json:"annotations"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("C:/Users/Jesse/PycharmProjects/PythonProject/DocLayNet_core/COCO/train.json")
	check(err)
	var data jdoc
	err = json.Unmarshal(dat, &data)
	check(err)
	fmt.Println(data.Images[len(data.Images)-1])
	fmt.Println(data.Annotations[len(data.Annotations)-1])

}
