package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func fileExists(name string) (b bool, err error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
	}
	// Propagates the error if the error is not FileNotExist error.
	return true, err
}

func TestFile(t *testing.T) {
	// if not exists
	exists, err := fileExists("/tmp/dat")
	check(err)
	if !exists {
		err = ioutil.WriteFile("/tmp/dat", []byte("hello\n go"), 0644)
		check(err)
		fmt.Println("文件不存在，初始化文件完成")
	}

	// 也许大部分基本的文件读取任务是将文件内容读取到
	// 内存中。
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 你经常会想对于一个文件是怎么读并且读取到哪一部分
	// 进行更多的控制。对于这个任务，从使用 `os.Open`
	// 打开一个文件获取一个 `os.File` 值开始。
	f, err := os.Open("/tmp/dat")
	check(err)

	// 从文件开始位置读取一些字节。这里最多读取 5 个字
	// 节，并且这也是我们实际读取的字节数。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 你也可以 `Seek` 到一个文件中已知的位置并从这个位置开
	// 始进行读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// `io` 包提供了一些可以帮助我们进行文件读取的函数。
	// 例如，上面的读取可以使用 `ReadAtLeast` 得到一个更
	// 健壮的实现。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内置的回转支持，但是使用 `Seek(0, 0)` 实现。
	_, err = f.Seek(0, 0)
	check(err)

	// `bufio` 包实现了带缓冲的读取，这不仅对有很多小的读
	// 取操作的能提升性能，也提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 任务结束后要关闭这个文件（通常这个操作应该在 `Open`
	// 操作后立即使用 `defer` 来完成）。
	f.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
