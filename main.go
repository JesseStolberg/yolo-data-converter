package main

import (
	"encoding/json"
	"os"
)

const imgSz = 1025

type c struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type i struct {
	Id          int    `json:"id"`
	FileName    string `json:"file_name"`
	Annotations []a
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
func transform(annotation *a) {
	annotation.Bbox[0] = annotation.Bbox[0] + annotation.Bbox[2]/2
	annotation.Bbox[1] = annotation.Bbox[1] + annotation.Bbox[3]/2
	annotation.Bbox[0] /= imgSz
	annotation.Bbox[1] /= imgSz
	annotation.Bbox[2] /= imgSz
	annotation.Bbox[3] /= imgSz
	annotation.CategoryId -= 1

}

func main() {
	dat, err := os.ReadFile("C:/Users/Jesse/PycharmProjects/PythonProject/DocLayNet_core/COCO/val.json")
	check(err)
	var data jdoc
	err = json.Unmarshal(dat, &data)
	check(err)
	for idx := range data.Categories {
		data.Categories[idx].Id -= 1
	}
	j := 0
	for _, annotation := range data.Annotations {
		for annotation.ImageId > data.Images[j].Id {
			j++
		}
		if annotation.ImageId != data.Images[j].Id {
			panic(err)
		}
		transform(&annotation)
		data.Images[j].Annotations = append(data.Images[j].Annotations, annotation)

	}
	// TODO: create file structure
	// creat YAML file
	// genorate annotation text files
	// move image files (or symlink to files with new names)

}
