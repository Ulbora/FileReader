package reader

//FileReader FileReader
type FileReader interface {
	ReadCsvFile(file []byte) *CsvFiles
}

//go mod init github.com/Ulbora/FileReader
