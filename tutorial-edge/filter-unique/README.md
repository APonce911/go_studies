In this challenge, we are going to be looking at how you can effectively filter out the duplicate entries from a slice in Go.

The task will be to implement the FilterDuplicates function so that it returns a slice of type string which contains only the unique names of developers retrieved from the inputted list.

Example:
```
// input
[]Developers{
  Developer{Name: "Elliot"},
  Developer{Name: "Alan"},
  Developer{Name: "Jennifer"},
  Developer{Name: "Graham"},
  Developer{Name: "Paul"},
  Developer{Name: "Alan"},
}

// output
[]string{
  "Elliot",
  "Alan",
  "Jennifer",
  "Graham",
}
```