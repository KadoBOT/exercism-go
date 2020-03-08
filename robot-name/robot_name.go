package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Robot struct {
	name string
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var nums = []rune("0123456789")
var usedNames string
var max = 26 * 26 * 10 * 10 * 10
var i int

func createName() (string, error) {
	if i > max {
		return "", errors.New("max names exhausted")
	}

	i++
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	name := make([]rune, 5)
	name[0] = letters[r.Intn(len(letters))]
	name[1] = letters[r.Intn(len(letters))]
	name[2] = nums[r.Intn(len(nums))]
	name[3] = nums[r.Intn(len(nums))]
	name[4] = nums[r.Intn(len(nums))]
	n := string(name)

	if strings.Contains(usedNames, n) {
		return createName()
	}

	usedNames = fmt.Sprintf("%s%s", usedNames, n)
	return string(name), nil
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name, err := createName()
	r.name = name
	return r.name, err
}

func (r *Robot) Reset() {
	r.name = ""
}

