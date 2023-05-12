package junit

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"unicode"
)

func ReadXML(path string) (*XML, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file, %v", err)
	}

	// Clean the file data to remove any unicode characters that will not print properly.
	scrubbedFileData := bytes.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, fileData)

	out := new(XML)
	if err := xml.Unmarshal(scrubbedFileData, out); err != nil {
		return nil, fmt.Errorf("error unmarshalling xml, %v", err)
	}

	return out, nil
}
