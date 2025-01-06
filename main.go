package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const imgSz = 1025
const filePath = "C:/Users/Jesse/PycharmProjects/PythonProject/DocLayNet_core/COCO/val.json"

type class struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type image struct {
	Id          int    `json:"id"`
	FileName    string `json:"file_name"`
	Annotations []annotation
}
type annotation struct {
	ImageId    int       `json:"image_id"`
	CategoryId int       `json:"category_id"`
	Bbox       []float64 `json:"bbox"`
}
type jdoc struct {
	Categories  []class      `json:"categories"`
	Images      []image      `json:"images"`
	Annotations []annotation `json:"annotations"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func transform(a *annotation) {
	a.Bbox[0] = a.Bbox[0] + a.Bbox[2]/2
	a.Bbox[1] = a.Bbox[1] + a.Bbox[3]/2
	a.Bbox[0] /= imgSz
	a.Bbox[1] /= imgSz
	a.Bbox[2] /= imgSz
	a.Bbox[3] /= imgSz
	a.CategoryId -= 1

}
func main() {
	b, err := os.ReadFile(filePath)
	check(err)
	var d jdoc
	err = json.Unmarshal(b, &d)
	check(err)
	for i := range d.Categories {
		d.Categories[i].Id -= 1
	}
	j := 0
	for _, a := range d.Annotations {
		for a.ImageId > d.Images[j].Id {
			j++
		}
		if a.ImageId != d.Images[j].Id {
			panic(err)
		}
		transform(&a)
		d.Images[j].Annotations = append(d.Images[j].Annotations, a)

	}
	err = os.MkdirAll("labels/train", 0777)
	check(err)
	err = os.Chdir("labels/train")
	check(err)
	for _, i := range d.Images {
		f, err := os.Create(strings.Split(i.FileName, ".")[0] + ".txt")
		check(err)
		defer f.Close()
		for _, a := range i.Annotations {
			fmt.Fprintf(f, "%d %6f %6f %6f %6f\n", a.CategoryId, a.Bbox[0], a.Bbox[1], a.Bbox[2], a.Bbox[3])
			check(err)
		}
		f.Close()
	}
}
