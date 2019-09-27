package xlsl

import "strings"

func getFileExt (name string) string {
	fileExt := strings.Split(name, ".")
	
	if len(fileExt) == 2 {
		return fileExt[1]
	}

	return ""
}

func checkError (err error) {
	if err != nil {
		panic(err)
	}
}