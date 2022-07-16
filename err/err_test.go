package err

import (
	"fmt"
	"testing"
)

func BaseTestError(t *testing.T) {
	e1 := NewErr(E_NONE, "")
	e2 := FromMsgf("err: %s", "some bad happened")
	fmt.Println(e1, e2)
}
