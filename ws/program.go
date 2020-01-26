package ws // import "github.com/andrewarchi/nebula/ws"

import (
	"strings"

	"github.com/andrewarchi/nebula/bigint"
)

// Program is a sequence of tokens with source map information.
type Program struct {
	Name       string
	Tokens     []Token
	LabelNames *bigint.Map // map[*big.Int]string
}

// Dump formats a program.
func (p *Program) Dump(indent string) string {
	var b strings.Builder
	for _, tok := range p.Tokens {
		if tok.Type != Label {
			b.WriteString(indent)
		}
		b.WriteString(tok.Format(p.LabelNames))
		b.WriteByte('\n')
	}
	return b.String()
}

func (p *Program) String() string {
	return p.Dump("    ")
}
