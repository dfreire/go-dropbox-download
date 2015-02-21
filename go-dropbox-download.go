package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println()

	if len(os.Args) != 4 {
		fmt.Println("Usage: go-dropbox-download <dropbox_folder_url> <local_download_folder> <math_string>")
		return
	}

	dropboxFolder := os.Args[1]
	downloadFolder := os.Args[2]
	matchFileName := os.Args[3]

	fmt.Println("Dropbox folder:", dropboxFolder)
	fmt.Println("Download folder:", downloadFolder)
	fmt.Println("Match file names with:", matchFileName)
	fmt.Println()

	doc, err := goquery.NewDocument(dropboxFolder)
	if err != nil {
		fmt.Println("Dropbox folder cannot be accessed.")
		return
	}

	dir, err := os.Stat(downloadFolder)
	if err != nil {
		fmt.Println("Download folder does not exist.")
		return
	}
	if !dir.IsDir() {
		fmt.Println("Download folder is not a directory.")
		return
	}

	fileNames := make(map[string]string)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		if !strings.Contains(href, "?dl=0") {
			return
		}

		tokens := strings.Split(href, "/")
		fileName := tokens[len(tokens)-1]
		fileName = strings.Split(fileName, "?dl=")[0]
		fileName, _ = url.QueryUnescape(fileName)
		if !strings.Contains(fileName, matchFileName) {
			return
		}

		href = strings.Replace(href, "?dl=0", "?dl=1", -1)

		fileNames[fileName] = href
	})

	for fileName, href := range fileNames {
		downloadFile(downloadFolder, href, fileName)
	}
}

func downloadFile(downloadFolder, href, fileName string) {
	filePath := path.Join(downloadFolder, fileName)
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		// file exists
		fmt.Println("SKIP", fileName)
		return
	} else {
		fmt.Println("DOWNLOAD", fileName)
	}

	tmpFilePath := path.Join(downloadFolder, strings.Join([]string{fileName, "tmp"}, "."))

	out, err := os.Create(tmpFilePath)
	defer out.Close()
	if err != nil {
		fmt.Print(err)
		return
	}

	resp, err := http.Get(href)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	err = os.Rename(tmpFilePath, filePath)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}
