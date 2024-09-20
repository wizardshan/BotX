package validate

import (
	"golang.org/x/exp/constraints"
	"regexp"
)

const (
	Alpha        string = "^[a-zA-Z]+$"
	Alphanumeric string = "^[a-zA-Z0-9]+$"
	Numeric      string = "^[0-9]+$"
	Int          string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float        string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
)

const (
	IsNotEmptyMsg     = "%s不能为空"
	IsPositiveMsg     = "%s必须为正数"
	FormatMsg         = "%s格式不正确"
	MinLenMsg         = "%s最小长度为%d"
	MaxLenMsg         = "%s最大长度为%d"
	LenMsg            = "%s长度必须为%d"
	MinMsg            = "%s最小值为%d"
	MaxMsg            = "%s最大值为%d"
	IsAlphaMsg        = "%s必须由字母组成"
	IsAlphanumericMsg = "%s必须由数字字母组成"
	IsNumericMsg      = "%s必须由数字组成"
	IsIntMsg          = "%s必须是整型"
	IsFloatMsg        = "%s必须是浮点型"
)

func IsPositive[T constraints.Integer | constraints.Float](v T) bool {
	return v > 0
}

func Min[T constraints.Integer | constraints.Float](v T, n T) bool {
	return v >= n
}

func Max[T constraints.Integer | constraints.Float](v T, n T) bool {
	return v <= n
}

// IsEmpty returns true if argument is a zero value.
func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

// IsNotEmpty returns true if argument is not a zero value.
func IsNotEmpty[T comparable](v T) bool {
	var zero T
	return zero != v
}

func IsEqual[T comparable](v1 T, v2 T) bool {
	return v1 == v2
}

func IsNotEqual[T comparable](v1 T, v2 T) bool {
	return v1 != v2
}

func MinLen(v string, n int) bool {
	return len(v) >= n
}

func MaxLen(v string, n int) bool {
	return len(v) <= n
}

func Len(v string, n int) bool {
	return len(v) == n
}

func IsAlpha(v string) bool {
	return regexp.MustCompile(Alpha).MatchString(v)
}

func IsAlphanumeric(v string) bool {
	return regexp.MustCompile(Alphanumeric).MatchString(v)
}

func IsNumeric(v string) bool {
	return regexp.MustCompile(Numeric).MatchString(v)
}

func IsInt(v string) bool {
	return regexp.MustCompile(Int).MatchString(v)
}

func IsFloat(v string) bool {
	return regexp.MustCompile(Float).MatchString(v)
}
