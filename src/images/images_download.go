package images

import (
	"domain"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	fileName    string
	fullUrlFile string
)

func Imagedown(imgurl string) domain.ResultModel {
	fullUrlFile = imgurl
	fileName = buildFilename()
	file := createFile()
	// Put content on file
	putFile(file, httpClient())
	rm := domain.ResultModel{}
	return rm
}

func buildFilename() string {
	return "2.gif"
}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullUrlFile)

	checkError(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	checkError(err)

	fmt.Println("Just Downloaded a file %s with size %d", fileName, size)
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)
	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
