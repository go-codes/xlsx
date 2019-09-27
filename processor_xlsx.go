package xlsl

import (
	"github.com/tealeg/xlsx"
)

type processorXlsx struct {

}

func (p *processorXlsx) Read (name string, sheetIndex int) (data [][]string, err error){
	f, err := xlsx.OpenFile(name)
	if err != nil {
		return nil, err
	}

	workSheet := f.Sheets[sheetIndex]
	rows := workSheet.Rows

	for _, row := range rows {
		var cols []string
		for _, cell := range row.Cells {
			cols = append(cols, cell.Value)
		}
		data = append(data, cols)
	}
	return

}

