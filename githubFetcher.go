package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// SourceFetcher implements the methods needed to fetch a files.
type SourceFetcher interface {
	fetchLicense(licenseName string) license
	fetchLicenseList() []licenseSummary
	fetchGitIgnore(language string) gitignore // /gitignore/templates/go
}

type githubSourceFetcher struct {
}

type licenseSummary struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

type license struct {
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	SpdxID         string   `json:"spdx_id"`
	URL            string   `json:"url"`
	HTMLURL        string   `json:"html_url"`
	Description    string   `json:"description"`
	Implementation string   `json:"implementation"`
	Permissions    []string `json:"permissions"`
	Conditions     []string `json:"conditions"`
	Limitations    []string `json:"limitations"`
	Body           string   `json:"body"`
	Featured       bool     `json:"featured"`
}

type gitignore struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

func (l githubSourceFetcher) fetchLicense(licenseName string) license {
	url := "https://api.github.com/licenses/" + licenseName
	body := fetchFromURL(url)

	license := license{}

	if body == nil {
		return license
	}

	err := json.Unmarshal(body, &license)

	if err != nil {
		log.Fatal(err)
	}

	return license
}

func (l githubSourceFetcher) fetchLicenseList() []licenseSummary {
	url := "https://api.github.com/licenses"
	body := fetchFromURL(url)

	licenses := []licenseSummary{}

	if body == nil {
		return licenses
	}

	err := json.Unmarshal(body, &licenses)

	if err != nil {
		log.Fatal(err)
	}

	return licenses
}

func (l githubSourceFetcher) fetchGitIgnore(language string) gitignore {
	url := "https://api.github.com/gitignore/templates/" + language
	body := fetchFromURL(url)

	gitignore := gitignore{}

	err := json.Unmarshal(body, &gitignore)

	if err != nil {
		log.Fatal(err)
	}

	return gitignore
}

func fetchFromURL(url string) []byte {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
