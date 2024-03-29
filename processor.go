package xlsx

import "errors"

var supportProcessor map[string]Processor


type Processor interface {
	Read(string, int) ([][]string, error)
	ReadBytes (content []byte, sheetIndex int)  (data [][]string, err error)
}
func ReadBytes (content []byte, ext string, sheetIndex int) ([][]string, error) {
	processor, found := supportProcessor[ext]
	if !found {
		checkError( errors.New("不支持的文件格式"))
	}
	return processor.ReadBytes(content, sheetIndex)
}

func Read (name string, sheetIndex int) ([][]string, error) {
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