package parser

type Parser interface {
	Parse(strToParse string, operation *string, num *float64) error
}
