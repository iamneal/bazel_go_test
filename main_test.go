package main_test

import (
	"fmt"
	"testing"

	main "github.com/iamneal/bazel_go_test"
)

func Test_main(t *testing.T) {
	fmt.Printf("works..\n")
	main.Main()
}
