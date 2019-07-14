package parser

import (
	"strconv"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (int, error) {
	code, ok := tokens[input]
	if !ok {
		val, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return 0, err
		}

		return int(val), nil
	}

	return code, nil
}
