package ws

import (
	"math/big"
	"testing"
)

func TestTokenString(t *testing.T) {
	arg := big.NewInt(123)
	tests := []struct {
		Token  *Token
		String string
	}{
		{&Token{Type: Push, Arg: arg}, "push 123"},
		{&Token{Type: Push, Arg: nil}, "push <nil>"},
		{&Token{Type: Add, Arg: arg}, "add"},
		{&Token{Type: Add, Arg: nil}, "add"},
		{&Token{Type: Label, Arg: arg}, "label_123"},
		{&Token{Type: Label, Arg: nil}, "label_<nil>"},
	}

	for i, test := range tests {
		if str := test.Token.String(); str != test.String {
			t.Errorf("test %d: String() = %q, want %q", i+1, str, test.String)
		}
	}
}

func TestTypeGroups(t *testing.T) {
	tests := []struct {
		IsStack, IsArith, IsHeap, IsFlow, IsIO bool
		Types                                  []Type
	}{
		{true, false, false, false, false, []Type{Push, Dup, Copy, Swap, Drop, Slide}},
		{false, true, false, false, false, []Type{Add, Sub, Mul, Div, Mod}},
		{false, false, true, false, false, []Type{Store, Retrieve}},
		{false, false, false, true, false, []Type{Label, Call, Jmp, Jz, Jn, Ret, End}},
		{false, false, false, false, true, []Type{Printc, Printi, Readc, Readi}},
	}

	for _, test := range tests {
		for _, typ := range test.Types {
			if typ.IsStack() != test.IsStack {
				t.Errorf("(%s).IsStack() = %t, want %t", typ, typ.IsStack(), test.IsStack)
			}
			if typ.IsArith() != test.IsArith {
				t.Errorf("(%s).IsArith() = %t, want %t", typ, typ.IsArith(), test.IsArith)
			}
			if typ.IsHeap() != test.IsHeap {
				t.Errorf("(%s).IsHeap() = %t, want %t", typ, typ.IsHeap(), test.IsHeap)
			}
			if typ.IsFlow() != test.IsFlow {
				t.Errorf("(%s).IsFlow() = %t, want %t", typ, typ.IsFlow(), test.IsFlow)
			}
			if typ.IsIO() != test.IsIO {
				t.Errorf("(%s).IsIO() = %t, want %t", typ, typ.IsIO(), test.IsIO)
			}
		}
	}
}

func TestInstrTypeString(t *testing.T) {
	tests := []struct {
		Type   Type
		String string
	}{
		{Illegal, "token(0)"},
		{Push, "push"},
		{Dup, "dup"},
		{Copy, "copy"},
		{Swap, "swap"},
		{Drop, "drop"},
		{Slide, "slide"},
		{Add, "add"},
		{Sub, "sub"},
		{Mul, "mul"},
		{Div, "div"},
		{Mod, "mod"},
		{Store, "store"},
		{Retrieve, "retrieve"},
		{Label, "label"},
		{Call, "call"},
		{Jmp, "jmp"},
		{Jz, "jz"},
		{Jn, "jn"},
		{Ret, "ret"},
		{End, "end"},
		{Printc, "printc"},
		{Printi, "printi"},
		{Readc, "readc"},
		{Readi, "readi"},
		{100, "token(100)"},
	}

	for i, test := range tests {
		if str := test.Type.String(); str != test.String {
			t.Errorf("test %d: String() = %q, want %q", i+1, str, test.String)
		}
	}
}
