package semver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	numbers = "0123456789"
)

const (
	MAJOR = iota
	MINOR
	PATCH
	LEN
)

type version struct {
	major uint16
	minor uint16
	patch uint16
}

// 要求非空字符串
func justContain(s, set string) bool {
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(set, s[i]) == -1 {
			return false
		}
	}

	return true
}

// 要求非空字符串
func isZeroLeftFill(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return true
	}

	return false
}

// 解析字符串，构造version对象
func parse(elems []string) (*version, error) {

	var tem [LEN]uint16
	for i := 0; i < LEN; i++ {
		if !justContain(elems[i], numbers) {
			return nil, fmt.Errorf("%s should be made up of 0-9", elems[i])
		}

		if isZeroLeftFill(elems[i]) {
			return nil, fmt.Errorf("%s starts with 0", elems[i])
		}

		val, err := strconv.ParseUint(elems[i], 10, 16)
		if err != nil {
			return nil, err
		}
		tem[i] = uint16(val)
	}

	return &version{major: tem[MAJOR], minor: tem[MINOR], patch: tem[PATCH]}, nil
}

func New(v string) (*version, error) {
	if len(v) == 0 {
		return nil, errors.New("empty version")
	}

	// 错误格式
	elems := strings.Split(v, ".")
	if len(elems) != LEN {
		return nil, errors.New("unexcepted format")
	}

	return parse(elems)
}

func (v *version) MakeMajor() {
	v.major++
	v.minor, v.patch = 0, 0
}

func (v *version) MakeMinor() {
	v.minor, v.patch = v.minor+1, 0
}

func (v *version) MakePatch() {
	v.patch++
}

func (v version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}
