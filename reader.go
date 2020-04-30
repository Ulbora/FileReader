package reader

//FileReader FileReader
type FileReader interface {
	ReadCsvFile(file []byte) *CsvFiles
}
