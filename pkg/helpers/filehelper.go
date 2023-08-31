package helpers

import (
	"avitoTask/pkg/models"
	"bytes"
	"fmt"
	"os"

	"github.com/dnlo/struct2csv"
)

type FileHelper struct{}

func (u *FileHelper) CreateCSVFile(operations []models.Operation, filename string) {
	buff := &bytes.Buffer{}
	writerCSV := struct2csv.NewWriter(buff)
	headers := []string{"id", "user_id", "segment_name", "action", "time"}

	writerCSV.WriteColNames(headers)
	writerCSV.WriteStructs(operations)

	writerCSV.Flush()

	file, err := os.Create(fmt.Sprintf("./historyfiles/%s.csv", filename))

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(buff.Bytes())
}
