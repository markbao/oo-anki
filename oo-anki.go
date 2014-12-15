package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Outline struct {
	XMLName xml.Name `xml:"outline"`
	Text    string   `xml:"text,attr"`
	Answer  string   `xml:"Answer,attr"`
}

type Body struct {
	XMLName  xml.Name  `xml:"body"`
	Outlines []Outline `xml:"outline"`
}

type OPML struct {
	XMLName xml.Name `xml:"opml"`
	Body    Body     `xml:"body"`
}

func main() {
	// get arg
	path := os.Args[1]
	tag := os.Args[2]
	fmt.Println("Reading from " + path)

	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)

	opml := &OPML{}
	xmlErr := xml.Unmarshal([]byte(XMLdata), opml)

	if xmlErr != nil {
		log.Fatal(xmlErr)
	}

	// begin new csv
	// define csv path
	name := filepath.Base(path)
	newName := strings.Replace(name, filepath.Ext(path), ".csv", -1)
	newPath := filepath.Dir(path) + "/" + newName
	fmt.Println("Saving to " + newPath)

	csvfile, csvErr := os.Create(newPath)
	if csvErr != nil {
		fmt.Println("Error:", csvErr)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)

	for i := 0; i < len(opml.Body.Outlines); i++ {
		log.Printf("%#v", opml.Body.Outlines[i].Text)
		log.Printf("%#v", opml.Body.Outlines[i].Answer)

		answ := strings.Replace(opml.Body.Outlines[i].Answer, "\n", "<br>", -1)
		wstr := []string{opml.Body.Outlines[i].Text, answ, tag}

		err := writer.Write(wstr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		} else {
			fmt.Println("wrote line")
		}
	}

	writer.Flush()
}
