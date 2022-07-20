package pirate

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Title    string `xml:"title"`
	Hash     string `xml:"infoHash"`
	Seeders  string `xml:"seeders"`
	Category string `xml:"category"`
	Size     string `xml:"size"`
}

type Nyaa struct {
	Items []Item `xml:"channel>item"`
}

type NyaaSearch struct{}

func (n NyaaSearch) Search(search string) ([]Metadata, error) {

	searchUrl := fmt.Sprintf(nyaaUrlTemplate, search)
	fmt.Println(searchUrl)

	response, err := http.Get(searchUrl)
	if err != nil {
		return nil, fmt.Errorf("api get: %w", err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("api get read response: %w", err)
	}

	var nyaa Nyaa
	err = xml.Unmarshal(bytes, &nyaa)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	var metadata []Metadata
	for i := range nyaa.Items {

		item := nyaa.Items[i]
		metadata = append(metadata, Metadata{Name: item.Title, Hash: item.Hash, Seeders: item.Seeders, Size: item.Size})
	}

	return metadata, nil
}

func (n NyaaSearch) GetMagnet(metadata Metadata) string {

	return getMagnet(metadata, nyaaTrackers())
}

func (p NyaaSearch) GetName() string {
	return "Nyaa"
}

func nyaaTrackers() []string {

	return []string{
		"http://nyaa.tracker.wf:7777/announce",
		"udp://open.stealth.si:80/announce",
		"udp://tracker.opentrackr.org:1337/announce",
		"udp://exodus.desync.com:6969/announce",
		"udp://tracker.torrent.eu.org:451/announce",
	}
}
