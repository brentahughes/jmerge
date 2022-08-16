package junit

import (
	"encoding/xml"
	"os"
)

func ReadXML(path string) (*XML, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	out := new(XML)
	if err := xml.Unmarshal(fileData, out); err != nil {
		return nil, err
	}

	return out, nil
}
