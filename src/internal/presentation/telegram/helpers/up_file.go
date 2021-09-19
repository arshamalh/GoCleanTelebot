package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
	tb "gopkg.in/tucnak/telebot.v2"
)

type filePathResultStruct struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileID       string `json:"file_id"`
		FileUniqueID string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		FilePath     string `json:"file_path"`
	} `json:"result"`
}

// UpFile take a message, determine the file id, download the file and return downloaded file path
func UpFile(m *tb.Message) (string, error) {
	filePath, e := makeTelegramFilePath(m.Document.FileID)
	if e != nil {
		log.Fatalf("Error happened")
	}
	fmt.Println("Url:", filePath.GetFileURL())
	pathName := filepath.Join("src", "statics", filepath.Base(m.Document.FileName))
	if e := DownloadFile(filePath.GetFileURL(), pathName); e != nil {
		fmt.Println("Error reading file:", e)
		os.Exit(1)
	}

	return pathName, e
}

func makeTelegramFilePath(fileID string) (filePathResultStruct, error) {
	resp, err := http.Get("https://api.telegram.org/bot" + os.Getenv("TOKEN") + "/getFile?file_id=" + fileID)
	if err != nil {
		return filePathResultStruct{}, err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return filePathResultStruct{}, err
	}
	jres := filePathResultStruct{}
	if e := json.Unmarshal(res, &jres); e != nil {
		return filePathResultStruct{}, e
	}
	return jres, nil
}

func (f *filePathResultStruct) GetFileURL() string {
	return "https://api.telegram.org/file/bot" + os.Getenv("TOKEN") + "/" + f.Result.FilePath
}

// WriteCounter will help to show download progress in the terminal
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

// PrintProgress prints the progress of a file write
func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 50))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

// DownloadFile will download from a url to a specified path.
func DownloadFile(url string, filepath string) error {
	// Create the file with .tmp extension, so that we won't overwrite a
	// file until it's downloaded fully
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our bytes counter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Println("\n" + filepath + " downloaded.")
	out.Close() // close the file after writing data to it

	// Rename the tmp file back to the original file
	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}
