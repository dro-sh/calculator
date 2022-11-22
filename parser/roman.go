package parser

import (
	"fmt"
)

type RomanParser struct{}

func (p *RomanParser) Parse(strToParse string, operation *string, num *float64) error {
	if _, err := fmt.Sscanf(strToParse, "%s %s", operation, num); err != nil {
		return err
	}

	// logic

	return nil
}
