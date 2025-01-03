package main

import (
	"encoding/json"
	"os"
)

type c struct {
	SuperCategory string `json:"supercategory"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
}

type i struct {
	Id          int    `json:"id"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	FileName    string `json:"file_name"`
	Collection  string `json:"collection"`
	DocName     string `json:"doc_name"`
	PageNum     int    `json:"page_no"`
	Precedence  int    `json:"precedence"`
	DocCategory string `json:"doc_category"`
}

type a struct {
	Id         int         `json:"id"`
	ImageId    int         `json:"image_id"`
	CategoryId int         `json:"category_id"`
	Bbox       []float64   `json:"bbox"`
	Segs       [][]float64 `json:"segmentation"`
	Area       float64     `json:"area"`
	Crowd      int         `json:"iscrowd"`
	Precedence int         `json:"precedence"`
}

type doc struct {
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
	var data doc
	err = json.Unmarshal(dat, &data)
	check(err)

}
