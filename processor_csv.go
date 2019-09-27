package xlsl

import (
	"encoding/csv"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os"
)

type processCsv struct {

}

func (p *processCsv) Read (name string, sheetIndex int) (data [][]string, err error) {
	f, err := os.Open(name)
	if err != nil {
		return nil,err
	}
	defer f.Close()

	decoder := simplifiedchinese.GBK.NewDecoder()
	r := transform.NewReader(f, decoder)

	cf := csv.NewReader(r)
	return cf.ReadAll()
}
