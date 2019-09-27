package xlsl

import (
	"github.com/extrame/xls"
)

type processorXls struct {

}

func (p *processorXls) Read (name string, sheetIndex int) (data [][]string, err error){
	workBook, err := xls.Open(name, "utf-8")

	if err != nil {
		return nil, err
	}
	workSheet := workBook.GetSheet(sheetIndex)
	maxRow := int(workSheet.MaxRow)

	for i := 0; i < maxRow; i++ {
		row := workSheet.Row(i)
		firstCol := row.FirstCol()
		lastCol := row.LastCol()
		var cols []string
		for j := firstCol; j < lastCol; j++ {
			cols = append(cols,row.Col(j))
		}
		data = append(data, cols)
	}
	return
}

