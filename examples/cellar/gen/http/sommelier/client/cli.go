// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/cellar/design -o
// $(GOPATH)/src/goa.design/goa/examples/cellar

package client

import (
	"encoding/json"
	"fmt"

	sommelier "goa.design/goa/examples/cellar/gen/sommelier"
)

// BuildPickCriteria builds the payload for the sommelier pick endpoint from
// CLI flags.
func BuildPickCriteria(sommelierPickBody string) (*sommelier.Criteria, error) {
	var err error
	var body PickRequestBody
	{
		err = json.Unmarshal([]byte(sommelierPickBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"name\": \"Blue\\'s Cuvee\",\n      \"varietal\": [\n         \"pinot noir\",\n         \"merlot\",\n         \"cabernet franc\"\n      ],\n      \"winery\": \"longoria\"\n   }'")
		}
	}
	if err != nil {
		return nil, err
	}
	v := &sommelier.Criteria{
		Name:   body.Name,
		Winery: body.Winery,
	}
	if body.Varietal != nil {
		v.Varietal = make([]string, len(body.Varietal))
		for j, val := range body.Varietal {
			v.Varietal[j] = val
		}
	}
	return v, nil
}