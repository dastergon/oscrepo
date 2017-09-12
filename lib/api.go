package lib

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.opensuse.org"

type Client struct {
	Username string
	Password string
}

type ProjMeta struct {
	XMLName xml.Name    `xml:"project"`
	Names   []RepoNames `xml:"repository"`
}

type RepoNames struct {
	Name string `xml:"name,attr"`
}

type ProjList struct {
	XMLName  xml.Name  `xml:"directory"`
	Projects []Project `xml:"entry"`
}
type Project struct {
	Name string `xml:"name,attr"`
}

// NewBasicAuthClient authenticates to the openSUSE API
func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

func (s *Client) sendRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// GetRepositories returns all the available repositories from build.opensuse.org
func (s *Client) GetRepositories() (*ProjList, error) {
	url := fmt.Sprintf(baseURL + "/source")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.sendRequest(req)
	if err != nil {
		return nil, err
	}
	var data ProjList
	err = xml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMeta returns the supported releases for a repository
func (s *Client) GetMeta(repository string) (*ProjMeta, error) {
	url := fmt.Sprintf(baseURL+"/source/%s/_meta", repository)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.sendRequest(req)
	if err != nil {
		return nil, err
	}
	var data ProjMeta
	err = xml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
