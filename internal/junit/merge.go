package junit

import (
	"fmt"
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
		if index < 0 {
			existing = &TestSuite{}
		}

		mergeStats(existing, &testSuite)

		existing.TestCases = append(existing.TestCases, testSuite.TestCases...)
		for i, testCase := range existing.TestCases {
			caseTime := convertDuration(testCase.Time)
			testCase.Time = fmt.Sprintf("%f", caseTime.Seconds())
			existing.TestCases[i] = testCase
		}

		suiteTime := convertDuration(testSuite.Time)
		existingTime := convertDuration(existing.Time)
		existing.Time = fmt.Sprintf("%f", (existingTime + suiteTime).Seconds())

		if index < 0 {
			d.TestSuites = append(d.TestSuites, *existing)
		} else {
			d.TestSuites[index] = *existing
		}
	}
}

func (d *XML) getTestSuiteByName(name string) (int, *TestSuite) {
	for i, suite := range d.TestSuites {
		if suite.Name == name {
			return i, &suite
		}
	}

	return -1, nil
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
