package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const argumentError = `Некорректный ввод оргументов, нужно передать URL`

type Wget struct {
	Url    string
	isFile bool
}

func (w Wget) createFile(p []byte) error {
	var fileName string
	if w.isFile {
		fileName = w.Url[strings.LastIndex(w.Url, "/")+1:]
	} else {
		fileName = "index.html"
	}
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0777) //"sudo"
	defer f.Close()
	if err != nil {
		return err
	}
	ioutil.WriteFile(f.Name(), p, 07777)
	return nil
}

func (w Wget) getPage(client *http.Client) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, w.Url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %d", resp.StatusCode)
	}
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (w Wget) checkUrl() error {
	_, err := url.ParseRequestURI(w.Url)
	if err != nil {
		return err
	}
	return nil
}

func (w *Wget) parseArgs() bool {
	if len(os.Args) == 2 {
		w.Url = os.Args[1]
		return true
	}
	if len(os.Args) == 3 {
		if os.Args[1] == "-m" {
			w.isFile = true
			w.Url = os.Args[2]
			return true
		}
	}
	return false
}

func main() {
	w := &Wget{}
	ok := w.parseArgs()
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", argumentError)
		os.Exit(1)
	}
	err := w.checkUrl()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	client := &http.Client{}
	p, err := w.getPage(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	err = w.createFile(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
