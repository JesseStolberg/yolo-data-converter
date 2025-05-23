package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const imgSz = 1025

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
func readAndReformat(metadata string, imageDir string) jdoc {
	dir := strings.Split(metadata, ".")[0]
	slice := strings.Split(dir, "/")
	dir = slice[len(slice)-1]
	b, err := os.ReadFile(metadata)
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

	err = os.MkdirAll("data/labels/"+dir, 0777)
	check(err)
	err = os.MkdirAll("data/images/"+dir, 0777)
	check(err)
	for _, i := range d.Images {
		f, err := os.Create("data/labels/" + dir + "/" + strings.Split(i.FileName, ".")[0] + ".txt")
		check(err)
		os.Rename(imageDir+i.FileName, "data/images/"+dir+"/"+i.FileName)
		for _, a := range i.Annotations {
			_, err = fmt.Fprintf(f, "%d %6f %6f %6f %6f\n", a.CategoryId, a.Bbox[0], a.Bbox[1], a.Bbox[2], a.Bbox[3])
			check(err)
		}
		f.Close()
	}
	return d
}
func main() {
	readAndReformat("doclaynet/COCO/train.json", "doclaynet/PNG/")
	readAndReformat("doclaynet/COCO/val.json", "doclaynet/PNG/")
	d := readAndReformat("doclaynet/COCO/test.json", "doclaynet/PNG/")
	err := os.MkdirAll("data", 0777)
	check(err)
	err = os.Chdir("data")
	check(err)
	f, err := os.Create("data.yaml")
	check(err)
	_, err = fmt.Fprint(f, "path: ../data\n")
	check(err)
	_, err = fmt.Fprint(f, "train: images/train\n")
	check(err)
	_, err = fmt.Fprint(f, "val: images/val\n")
	check(err)
	_, err = fmt.Fprint(f, "test: images/test\n")
	check(err)
	_, err = fmt.Fprint(f, "names:\n")
	check(err)
	for _, i := range d.Categories {
		_, err = fmt.Fprintf(f, "    %d: %s\n", i.Id, i.Name)
		check(err)
	}
	f.Close()

}
