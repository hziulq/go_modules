package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Prompt(recommend string) (string, error) {
	fmt.Print(recommend)
	return InputString()
}

func InputString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	//	test
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("ERROR InputString")
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil

}
