package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	apiURL    = "https://neocities.org/api/"
	userAgent = "neocities (Go package using net/http)"
)

// Credentials contains the username and password
type Credentials struct {
	User string
	Pass string
}

// UploadData contains the filename and content
type UploadData struct {
	FileName string
	Content  []byte
}

// UploadFiles takes a set of credentials and
// a list of filename paths to upload to Neocities.
func UploadFiles(cred *Credentials, paths []string) (Response, error) {
	req, err := newUploadRequest(cred, paths)
	check(err)

	return performHTTPRequest(req)
}

func Upload(cred *Credentials, data []UploadData) (Response, error) {
	req, err := newUploadDataRequest(cred, data)
	check(err)

	return performHTTPRequest(req)
}

// DeleteFiles deletes the given filenames
func DeleteFiles(cred *Credentials, filenames []string) (Response, error) {
	req, err := newDeleteRequest(cred, filenames)
	check(err)

	return performHTTPRequest(req)
}

// SiteInfo returns a site info response
func SiteInfo(cred *Credentials, site string) (Response, error) {
	req, err := newInfoRequest(cred, site)
	check(err)

	return performHTTPRequest(req)
}

// Create a new info request
func newInfoRequest(cred *Credentials, site string) (*http.Request, error) {
	endpoint := "info"

	if site != "" {
		endpoint = endpoint + "?sitename=" + site
	}

	req, err := http.NewRequest("GET", apiURL+endpoint, nil)
	if err != nil {
		return req, err
	}

	if cred != nil {
		// Authenticate using the user and pass
		req.SetBasicAuth(cred.User, cred.Pass)
	}

	return req, nil
}

// Create a new delete request
func newDeleteRequest(cred *Credentials, filenames []string) (*http.Request, error) {
	data := url.Values{}

	for _, file := range filenames {
		data.Add("filenames[]", file)
	}

	req, err := http.NewRequest("POST", apiURL+"delete", strings.NewReader(data.Encode()))
	if err != nil {
		return req, err
	}

	// Authenticate using the user and pass
	req.SetBasicAuth(cred.User, cred.Pass)

	// Set the content type of the form data
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

// Create a new upload request
func newUploadRequest(cred *Credentials, paths []string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the contents of each file to the multipart body
	for _, path := range paths {
		filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			file, err := os.Open(p)

			defer file.Close()

			if err != nil {
				return err
			}

			part, err := writer.CreateFormFile(p, p)

			if err != nil {
				return err
			}

			_, err = io.Copy(part, file)
			check(err)
			return nil
		})
	}

	err := writer.Close()
	check(err)

	req, err := http.NewRequest("POST", apiURL+"upload", body)
	if err != nil {
		return req, err
	}

	// Authenticate using the user and pass
	req.SetBasicAuth(cred.User, cred.Pass)

	// Set the content type of the form data
	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}

// Create a new upload data request
func newUploadDataRequest(cred *Credentials, data []UploadData) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the contents of each file to the multipart body
	for _, d := range data {
		part, err := writer.CreateFormFile(d.FileName, d.FileName)
		check(err)

		_, err = part.Write(d.Content)
		check(err)
	}

	err := writer.Close()
	check(err)

	req, err := http.NewRequest("POST", apiURL+"upload", body)
	if err != nil {
		return req, err
	}

	// Authenticate using the user and pass
	req.SetBasicAuth(cred.User, cred.Pass)

	// Set the content type of the form data
	req.Header.Add("Content-Type", writer.FormDataContentType())

	return req, nil
}

func performHTTPRequest(req *http.Request) (Response, error) {
	res, err := sendHTTPRequest(req)
	check(err)

	defer res.Body.Close()

	var r Response

	r.PopulateFromHTTPResponse(res)

	if res.StatusCode == 200 {
		return r, nil
	}

	return r, errors.New("unsuccessful")
}

// Send HTTP request
func sendHTTPRequest(req *http.Request) (*http.Response, error) {
	// Create a HTTP client
	client := &http.Client{}

	// Make sure that the correct User-Agent is set
	req.Header.Add("User-Agent", userAgent)

	// Send the request
	res, err := client.Do(req)
	check(err)

	// Return the response
	return res, nil
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)

		os.Exit(1)
	}
}
