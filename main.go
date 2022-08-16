package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"github.com/brentahughes/jmerge/internal/junit"
)

func main() {
	sources := os.Args[1:]
	if len(sources) == 0 {
		handleErr(errors.New("at least one source xml must be given"))
	}

	sourceFiles := make([]*junit.XML, 0)

	for _, source := range sources {
		parsed, err := junit.ReadXML(source)
		handleErr(err)

		sourceFiles = append(sourceFiles, parsed)
	}

	merged := junit.MergeXMLs(sourceFiles)

	mergedXML, err := xml.MarshalIndent(merged, "", "  ")
	handleErr(err)

	fmt.Fprintln(os.Stdout, string(mergedXML))

}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
