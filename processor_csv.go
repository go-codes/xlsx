package xlsx

import (
	"bytes"
	"encoding/csv"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

type processorCsv struct {

}

func newProcessorCsv () *processorCsv {
	return &processorCsv{}
}

func (p *processorCsv) Read (name string, sheetIndex int) (data [][]string, err error) {
	byteStr, err := ioutil.ReadFile(name)
	if err != nil {
		return nil,err
	}

	detector := chardet.NewTextDetector()
	charset, err := detector.DetectBest(byteStr)
	if err != nil {
		panic(err)
	}
	byteStr = bytes.TrimPrefix(byteStr, []byte{239, 187, 191})
	r := bytes.NewReader(byteStr)

	var decoder *encoding.Decoder
	if strings.Contains(charset.Charset, "UTF") {
		decoder = unicode.UTF8.NewDecoder()
	} else {
		decoder = simplifiedchinese.GBK.NewDecoder()

	}
	rd := transform.NewReader(r, decoder)

	cf := csv.NewReader(rd)
	return cf.ReadAll()
}

func (p *processorCsv) ReadBytes (content []byte, sheetIndex int) (data [][]string, err error) {
	byteStr := content
	byteStr = bytes.TrimPrefix(byteStr, []byte{239, 187, 191})
	detector := chardet.NewTextDetector()
	charset, err := detector.DetectBest(byteStr)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(byteStr)
	var decoder *encoding.Decoder
	if strings.Contains(charset.Charset, "UTF") {
		decoder = unicode.UTF8.NewDecoder()
	} else {
		decoder = simplifiedchinese.GBK.NewDecoder()

	}
	rd := transform.NewReader(r, decoder)

	cf := csv.NewReader(rd)
	return cf.ReadAll()
}
