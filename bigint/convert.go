package bigint // import "github.com/andrewarchi/nebula/bigint"

import (
	"math"
	"math/big"
	"strings"
	"unicode/utf8"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

// ToInt converts a *big.Int x to an int and returns whether x can be
// contained within int.
func ToInt(x *big.Int) (int, bool) {
	if !x.IsInt64() {
		return 0, false
	}
	i64 := x.Int64()
	if i64 > int64(maxInt) {
		return 0, false
	}
	return int(i64), true
}

// ToUint converts a *big.Int x to a uint and returns whether x can be
// contained within uint.
func ToUint(x *big.Int) (uint, bool) {
	if !x.IsUint64() {
		return 0, false
	}
	u64 := x.Uint64()
	if u64 > uint64(maxUint) {
		return 0, false
	}
	return uint(u64), true
}

// ToInt64 converts a *big.Int to an int64 and returns whether x can be
// contained within int64.
func ToInt64(x *big.Int) (int64, bool) {
	if !x.IsInt64() {
		return 0, false
	}
	return x.Int64(), true
}

// ToUint64 converts a *big.Int to a uint64 and returns whether x can be
// contained within uint64.
func ToUint64(x *big.Int) (uint64, bool) {
	if !x.IsUint64() {
		return 0, false
	}
	return x.Uint64(), true
}

// ToInt32 converts a *big.Int to an int32 and returns whether x can be
// contained within int32.
func ToInt32(x *big.Int) (int32, bool) {
	if !x.IsInt64() {
		return 0, false
	}
	i64 := x.Int64()
	if i64 > math.MaxInt32 {
		return 0, false
	}
	return int32(i64), true
}

// ToUint32 converts a *big.Int to a uint32 and returns whether x can be
// contained within uint32.
func ToUint32(x *big.Int) (uint32, bool) {
	if !x.IsUint64() {
		return 0, false
	}
	i64 := x.Uint64()
	if i64 > math.MaxUint32 {
		return 0, false
	}
	return uint32(i64), true
}

// ToRune converts a *big.Int x to a rune. When x is not a valid UTF-8
// codepoint, U+FFFD � replacement character is returned.
func ToRune(x *big.Int) rune {
	i32, ok := ToInt32(x)
	if !ok || !utf8.ValidRune(i32) {
		return '\uFFFD'
	}
	return i32
}

// FormatSlice formats a slice of *big.Int to a space separated string.
func FormatSlice(s []*big.Int) string {
	var b strings.Builder
	b.WriteByte('[')
	for i, x := range s {
		if i != 0 {
			b.WriteByte(' ')
		}
		b.WriteString(x.String())
	}
	b.WriteByte(']')
	return b.String()
}
