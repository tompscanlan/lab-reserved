package labreserved

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	"fmt"
	"io/ioutil"
	"net/http"
)

type BlobOut struct {
	Id      int    `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
	Tag     string `json:"tag,omitempty"`
}
type BlobIn struct {
	Id      int    `json:"id,omitempty,string"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
	Tag     string `json:"tag,omitempty"`
}

func PostBlob(id int, data string) error {
	url := fmt.Sprintf("%s", BlobEndpoint)
	//log.Println("URL:>", url)

	blob := new(BlobOut)
	blob.Id = id
	blob.Content = base64.StdEncoding.EncodeToString([]byte(data))
	blob.Tag = "fromLabReserved"

	jsonStr, err := json.Marshal(blob)
	if err != nil {
		return err
	}
	log.Println("json: ", string(jsonStr[:]))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func GetBlob(id int) (string, error) {

	url := fmt.Sprintf("%s/%d", BlobEndpoint, id)
	//log.Println("URL:>", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	client.Timeout = 5 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	blob := new(BlobIn)
	err = json.Unmarshal(body, blob)
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(blob.Content)
	if err != nil {
		return "", err
	}
	return string(decoded[:]), nil
}
