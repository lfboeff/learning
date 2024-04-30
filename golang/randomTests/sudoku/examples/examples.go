package examples

import (
	"encoding/json"
	"log"
	"os"
)

type Example struct {
	Board    [][]int
	Solution [][]int
}

// A
func CreateExamples() []Example {

	var examples []Example

	examplesBytes, err := os.ReadFile("examples/examples.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(examplesBytes, &examples); err != nil {
		log.Fatal(err)
	}

	return examples

}
