package panic

import (
	"fmt"
	"os"
	"testing"
)

var (
	myerror = MyError{
		arg:  1,
		prob: "custom error",
	}
)

type MyError struct {
	arg  int
	prob string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//goland:noinspection GoUnreachableCode
func TestPanic(t *testing.T) {
	panic(&myerror)

	if f, err := os.Create("/tmp/file"); err != nil {
		panic(err)
	} else {
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				//
			}
		}(f)
	}

}

func TestPanic2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if me, ok := r.(*MyError); ok {
				t.Log("myError prob:", me.prob, "arg:", me.arg)
			} else {
				t.Log("other Error:", r)
			}

		}
	}()

	//panic("a problem")
	//panic(errors.New("a error"))
	panic(&myerror)
	t.Log("unreasonable !")
}
