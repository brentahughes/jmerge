package junit

import (
	"fmt"
	"sort"
	"time"
)

func MergeXMLs(sources []*XML) *XML {
	merged := new(XML)
	for _, xml := range sources {
		merged.merge(xml)
	}

	return merged
}

func (d *XML) merge(source *XML) {
	mergeStats(d, source)

	for _, testSuite := range source.TestSuites {
		index, existing := d.getTestSuiteByName(testSuite.Name)

		mergeStats(&existing, &testSuite)

		existing.TestCases = append(existing.TestCases, testSuite.TestCases...)
		for i, testCase := range existing.TestCases {
			caseTime := convertDuration(testCase.Time)
			testCase.Time = fmt.Sprintf("%.2f", caseTime.Seconds())
			existing.TestCases[i] = testCase
		}

		if testSuite.Time != nil && existing.Time != nil {
			suiteTime := convertDuration(*testSuite.Time)
			existingTime := convertDuration(*existing.Time)
			t := fmt.Sprintf("%.2f", (existingTime + suiteTime).Seconds())
			existing.Time = &t
		}

		if index < 0 {
			d.TestSuites = append(d.TestSuites, existing)
		} else {
			d.TestSuites[index] = existing
		}
	}

	// Check all testcases and increment duplicate names
	for i, s := range d.TestSuites {
		sort.Slice(s.TestCases, func(a, b int) bool {
			return !sort.StringsAreSorted([]string{s.TestCases[a].Name, s.TestCases[b].Name})
		})

		var last string
		count := 1
		for j, c := range s.TestCases {
			if c.Name != last {
				last = c.Name
				count = 1

				continue
			}

			last = c.Name
			count++

			c.Name += fmt.Sprintf(" %d", count)
			s.TestCases[j] = c
		}

		d.TestSuites[i] = s
	}
}

func (d *XML) getTestSuiteByName(name string) (int, TestSuite) {
	for i, suite := range d.TestSuites {
		if suite.Name == name {
			return i, suite
		}
	}

	return -1, TestSuite{
		Name: name,
	}
}

func convertDuration(dur string) time.Duration {
	parsed, err := time.ParseDuration(dur)
	if err != nil {
		parsed, err = time.ParseDuration(dur + "s")
		if err != nil {
			return 0
		}
	}

	return parsed
}

func mergeStats(src, dest stats) {
	src.setTests(src.getTests() + dest.getTests())
	src.setErrors(src.getErrors() + dest.getErrors())
	src.setFailures(src.getFailures() + dest.getFailures())
	src.setSkipped(src.getSkipped() + dest.getSkipped())
}
