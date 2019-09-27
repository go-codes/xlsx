package xlsx

import "errors"

var supportProcessor map[string]Processor


type Processor interface {
	Read (string, int) ([][]string, error)
}

func ReadAll (name string, sheetIndex int) ([][]string, error) {
	fileExt := getFileExt(name)
	processor, found := supportProcessor[fileExt]
	if !found {
		checkError( errors.New("不支持的文件格式"))
	}
	return processor.Read(name, sheetIndex)
}

func init () {
	supportProcessor = make(map[string]Processor)
	AddProcessor("xlsx", newProcessorXlsx())
	AddProcessor("xls", newProcessorXls())
	AddProcessor("csv", newProcessorCsv())
}

func AddProcessor(name string, processor Processor) {
	supportProcessor[name] = processor
}