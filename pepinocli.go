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
	return res.StatusCode, string(body), err
}
