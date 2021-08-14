package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// PepinoDB is a database handle
type PepinoDB struct {
	url      string
	dbName   string
	password string
}

// Initialize sets-ups the initial values for a PepinoDB instance
func (p *PepinoDB) Initialize(url string, dbName string, password string) {
	p.url = url
	p.dbName = dbName
	p.password = password
}

func (p *PepinoDB) buildURLForEntry(entryName string) string {
	return p.url + "/?password=" + p.password +
		"&db=" + p.dbName + "&entry=" + entryName
}

// SaveEntry lets the user save bytes on the Pepino Database
func (p *PepinoDB) SaveEntry(entryName string, entryValue []byte) error {
	valueReader := bytes.NewReader(entryValue)
	url := p.buildURLForEntry(entryName)
	res, err := http.Post(url, "application/octet-stream", valueReader)
	if err != nil {
		return fmt.Errorf("error saving entry:\n\t%s", err.Error())
	}
	if res.StatusCode == http.StatusOK {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	errDesc := strings.Builder{}
	errDesc.WriteString("HTTP Error " + strconv.Itoa(res.StatusCode))
	if err == nil && body != nil && len(body) > 1 {
		errDesc.WriteString(" : " + string(body))
	}
	return fmt.Errorf("error saving entry:\n\t%s", errDesc.String())
}

// GetEntry the user remove an entry from the Pepino Database
func (p *PepinoDB) GetEntry(entryName string) ([]byte, error) {
	res, err := http.Get(p.buildURLForEntry(entryName))
	if err != nil {
		return nil, fmt.Errorf("error getting entry:\n\t%s", err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error getting entry:\n\t%s", err.Error())
	}
	if res.StatusCode == http.StatusOK {
		return body, nil
	}
	errDesc := strings.Builder{}
	errDesc.WriteString("HTTP Error " + strconv.Itoa(res.StatusCode))
	if body != nil && len(body) > 1 {
		errDesc.WriteString(" : " + string(body))
	}
	return nil, fmt.Errorf("error getting entry:\n\t%s", errDesc.String())
}

// DeleteEntry lets the user get an entry from the Pepino Database
func (p *PepinoDB) DeleteEntry(entryName string) error {
	url := p.buildURLForEntry(entryName)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("error deleting entry:\n\t%s", err.Error())
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error deleting entry:\n\t%s", err.Error())
	}
	if res.StatusCode == http.StatusOK {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	errDesc := strings.Builder{}
	errDesc.WriteString("HTTP Error " + strconv.Itoa(res.StatusCode))
	if err == nil && body != nil && len(body) > 1 {
		errDesc.WriteString(" : " + string(body))
	}
	return fmt.Errorf("error deleting entry:\n\t%s", errDesc.String())
}
