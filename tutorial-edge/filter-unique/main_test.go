package main

import "testing"

func TestFilterUnique(t *testing.T) {
	devs := []Developer{
	Developer{Name: "Elliot"},
	Developer{Name: "Alan"},
	Developer{Name: "Jennifer"},
	Developer{Name: "Graham"},
	Developer{Name: "Paul"},
	Developer{Name: "Alan"},
	}

	filteredDevNames := filterUnique(devs)
	uniqueDevNames := []string{
		"Elliot",
		"Alan",
		"Jennifer",
		"Graham",
		"Paul",
	  }

	if len(filteredDevNames) != len(uniqueDevNames) {
		t.Errorf("Expected %v names, got %v", len(uniqueDevNames), len(filteredDevNames))
	}
}