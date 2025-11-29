package initializr

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var Url = "https://start.spring.io/starter.zip?type=maven-project&language=java&bootVersion=4.0.0&baseDir=demo&groupId=com.example&artifactId=demo&name=demo&description=Demo%20project%20for%20Spring%20Boot&packageName=com.example.demo&packaging=jar&javaVersion=17"

func StarterZip() error {
	zipFile := "demo.zip"

	// 1️⃣ Download file
	err := downloadFile(zipFile, Url)
	if err != nil {
		return err
	}
	fmt.Println("Downloaded:", zipFile)

	// 2️⃣ Extract
	err = unzip(zipFile, "./")
	if err != nil {
		return err
	}
	fmt.Println("Extracted successfully")

	// 3️⃣ Delete ZIP
	err = os.Remove(zipFile)
	if err != nil {
		return err
	}

	fmt.Println("Zip file deleted")
	return nil
}

func downloadFile(filepath, url string) error {
	// prepare form data (e.g. dependencies=web)
	data := strings.NewReader("dependencies=web")
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Go-http-client/1.1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.StatusCode)
	if resp.StatusCode != 200 {
		return fmt.Errorf("non-200 response: %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		filePath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer destFile.Close()

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(destFile, rc)
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
