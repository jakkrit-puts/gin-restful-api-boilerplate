package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/vincent-petithory/dataurl"
)

func UploadImage(fileName string) string {
  fmt.Println(fileName)
	b64data := fileName[strings.IndexByte(fileName, ',')+1:]
	ext, _ := dataurl.DecodeString(fileName) //หาประเภทไฟล์
	data, _ := base64.StdEncoding.DecodeString(b64data)
	r := string(data)
	newFilename := uuid.New().String()
	var extFilename string = ""
	if ext.ContentType() == "image/png" {
		extFilename = ".png"
		os.WriteFile("./public/images/"+newFilename+".png", []byte(r), 0644)
	} else if ext.ContentType() == "image/jpeg" {
		extFilename = ".jpg"
		os.WriteFile("./public/images/"+newFilename+".jpg", []byte(r), 0644)
	}

	filePath := newFilename + extFilename

	return filePath
}
