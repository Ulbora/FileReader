package reader

import (
	"encoding/csv"
	"strings"
)

//CsvFiles CsvFiles
type CsvFiles struct {
	CsvFileList [][]string
	CsvReadErr  error
}

//CsvFileReader CsvFileReader
type CsvFileReader struct {
}

//ReadCsvFile ReadCsvFile
func (f *CsvFileReader) ReadCsvFile(file []byte) *CsvFiles {
	var rtn CsvFiles
	r := csv.NewReader(strings.NewReader(string(file)))
	records, rerr := r.ReadAll()
	rtn.CsvReadErr = rerr
	if rerr == nil {
		rtn.CsvFileList = records
	}
	return &rtn
}

//GetNew GetNew
func (f *CsvFileReader) GetNew() FileReader {
	return f
}
