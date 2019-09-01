package error

import "errors"

const TEST_STR = ""

var testError = errors.New("Test error!")

type Error struct {
  a, b int
}

func TEST(s string) {

}
