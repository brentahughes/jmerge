package junit

import "encoding/xml"

// XML represents the xml used for a junit report.
type XML struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Failures   int         `xml:"failures,attr"`
	Errors     int         `xml:"errors,attr"`
	Tests      int         `xml:"tests,attr"`
	Skipped    int         `xml:"skipped,attr"`
	Skip       int         `xml:"skip,attr,omitempty"`
	TestSuites []TestSuite `xml:"testsuite"`
}

func (x *XML) getFailures() int  { return x.Failures }
func (x *XML) setFailures(i int) { x.Failures = i }
func (x *XML) getErrors() int    { return x.Errors }
func (x *XML) setErrors(i int)   { x.Errors = i }
func (x *XML) getTests() int     { return x.Tests }
func (x *XML) setTests(i int)    { x.Tests = i }
func (x *XML) getSkipped() int   { return x.Skipped + x.Skip }
func (x *XML) setSkipped(i int) {
	x.Skipped = i
	x.Skip = 0
}

// TestSuite represents the xml used for a testsuite in a junit report.
type TestSuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Failures  int        `xml:"failures,attr"`
	Errors    int        `xml:"errors,attr"`
	Tests     int        `xml:"tests,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Skip      int        `xml:"skip,attr,omitempty"`
	Time      *string    `xml:"time,attr,omitempty"`
	TestCases []TestCase `xml:"testcase"`
}

func (x *TestSuite) getFailures() int  { return x.Failures }
func (x *TestSuite) setFailures(i int) { x.Failures = i }
func (x *TestSuite) getErrors() int    { return x.Errors }
func (x *TestSuite) setErrors(i int)   { x.Errors = i }
func (x *TestSuite) getTests() int     { return x.Tests }
func (x *TestSuite) setTests(i int)    { x.Tests = i }
func (x *TestSuite) getSkipped() int   { return x.Skipped + x.Skip }
func (x *TestSuite) setSkipped(i int) {
	x.Skipped = i
	x.Skip = 0
}

// TestCase represents the xml used for a testcase in a junit report.
type TestCase struct {
	XMLName   xml.Name `xml:"testcase"`
	ClassName string   `xml:"classname,attr,omitempty"`
	Name      string   `xml:"name,attr"`
	Status    string   `xml:"status,attr,omitempty"`
	Time      string   `xml:"time,attr"`
	Failure   *Failure `xml:"failure,omitempty"`
	Skipped   *Skipped `xml:"skipped,omitempty"`
}

type Failure struct {
	XMLName xml.Name `xml:"failure"`
	Type    string   `xml:"type,attr"`
	Message string   `xml:"message,attr"`
	Value   string   `xml:",cdata"`
}

type Skipped struct {
	XMLName xml.Name `xml:"skipped"`
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
