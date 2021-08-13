package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// PepinoDB is a database handle
type PepinoDB struct {
	url      string
	password string
	dbName   string
}

func (p *PepinoDB) buildURLForEntry(entryName string) string {
	return p.url + "/?password=" + p.password +
		"&db=" + p.dbName + "&entry=" + entryName
}

// SaveEntry lets the user save bytes on the Pepino Database
// Returns: HTTP StatusCode, Response as String, Error.
// When Error is not nil, it should be prioritized over anything else.
func (p *PepinoDB) SaveEntry(entryName string, entryValue []byte) (int, string, error) {
	valueReader := bytes.NewReader(entryValue)
	res, err := http.Post(p.buildURLForEntry(entryName), "application/octet-stream", valueReader)
	if err != nil {
		return -99, "", err
	}
	if res.StatusCode == 200 {
		return res.StatusCode, "", nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, "", err
	}
	return res.StatusCode, string(body), nil
}

// GetEntry the user remove an entry from the Pepino Database
// Returns: HTTP StatusCode, Response as []byte, Error.
// When Error is not nil, it should be prioritized over anything else.
// If Error is nil and StatusCode is different from 200, Description from
// server can be get by treating the response as an UTF-8 String
func (p *PepinoDB) GetEntry(entryName string) (int, []byte, error) {
	res, err := http.Get(p.buildURLForEntry(entryName))
	if err != nil {
		return -99, nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}
	return res.StatusCode, body, nil
}

// DeleteEntry lets the user get an entry from the Pepino Database
// Returns: HTTP StatusCode, Response as String, Error.
// When Error is not nil, it should be prioritized over anything else.
// If Error is nil and StatusCode is different from 200, Description from
// server can be get by treating the response as an UTF-8 String
func (p *PepinoDB) DeleteEntry(entryName string) (int, string, error) {
	req, err := http.NewRequest(http.MethodDelete, p.buildURLForEntry(entryName), nil)
	if err != nil {
		return -99, "", err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return -99, "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, "", err
	}
	if len(body) > 0 {
		return res.StatusCode, string(body), nil
	}
	return res.StatusCode, string(body), nil
}
