package error

import "errors"

var testError = errors.New("Test error!")

const TEST_STR = ""

func TEST(s string) {

}

type Error struct {
  a, b int
}
