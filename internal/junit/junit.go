package junit

import "encoding/xml"

// XML represents the xml used for a junit report.
type XML struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Failures   int         `xml:"failures,attr"`
	Errors     int         `xml:"errors,attr"`
	Tests      int         `xml:"tests,attr"`
	Skipped    int         `xml:"skipped,attr"`
	TestSuites []TestSuite `xml:"testsuite"`
}

func (x *XML) getFailures() int  { return x.Failures }
func (x *XML) setFailures(i int) { x.Failures = i }
func (x *XML) getErrors() int    { return x.Errors }
func (x *XML) setErrors(i int)   { x.Errors = i }
func (x *XML) getTests() int     { return x.Tests }
func (x *XML) setTests(i int)    { x.Tests = i }
func (x *XML) getSkipped() int   { return x.Skipped }
func (x *XML) setSkipped(i int)  { x.Skipped = i }

// TestSuite represents the xml used for a testsuite in a junit report.
type TestSuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Failures  int        `xml:"failures,attr"`
	Errors    int        `xml:"errors,attr"`
	Tests     int        `xml:"tests,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Time      string     `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
}

func (x *TestSuite) getFailures() int  { return x.Failures }
func (x *TestSuite) setFailures(i int) { x.Failures = i }
func (x *TestSuite) getErrors() int    { return x.Errors }
func (x *TestSuite) setErrors(i int)   { x.Errors = i }
func (x *TestSuite) getTests() int     { return x.Tests }
func (x *TestSuite) setTests(i int)    { x.Tests = i }
func (x *TestSuite) getSkipped() int   { return x.Skipped }
func (x *TestSuite) setSkipped(i int)  { x.Skipped = i }

// TestCase represents the xml used for a testcase in a junit report.
type TestCase struct {
	XMLName xml.Name `xml:"testcase"`
	Name    string   `xml:"name,attr"`
	Status  string   `xml:"status,attr"`
	Time    string   `xml:"time,attr"`
}

type stats interface {
	getFailures() int
	setFailures(int)

	getErrors() int
	setErrors(int)

	getTests() int
	setTests(int)

	getSkipped() int
	setSkipped(int)
}
