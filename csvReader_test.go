package reader

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCsvFileReader_ReadCsvFile(t *testing.T) {
	var cr CsvFileReader
	sourceFile, err := ioutil.ReadFile("../full_file.csv")
	fmt.Println("readFile err: ", err)
	rd := cr.GetNew()
	rec := rd.ReadCsvFile(sourceFile)
	fmt.Println("csv read err: ", rec.CsvReadErr)
	fmt.Println("csv len: ", len(rec.CsvFileList))
	fmt.Println("csv first row: ", rec.CsvFileList[0])
	fmt.Println("csv last row: ", rec.CsvFileList[len(rec.CsvFileList)-2])
	if len(rec.CsvFileList) == 0 {
		t.Fail()
	}
}

//go mod init github.com/Ulbora/FileReader
