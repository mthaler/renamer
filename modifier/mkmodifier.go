package modifier

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func Mkmodifier(cmd string) (func(string) string, error) {
	parts := strings.Split(cmd, "/")
}