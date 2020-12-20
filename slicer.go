package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadSeekCloser Reader+Seeker+Closer
type ReadSeekCloser interface {
	io.Reader
	io.Seeker
	io.Closer
}

// Slicer スライサー
type Slicer struct {
	filePath string
	in       ReadSeekCloser
}

// NewSlicer スライサーの生成
func NewSlicer(filePath string) (r Slicer, err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return
	}

	r = Slicer{
		filePath: filePath,
		in:       f,
	}
	return
}

// Do スライスの処理
func (s Slicer) Do(number int, slice string) (err error) {
	var start, stop int
	var out string
	start, stop, out, err = parseSlice(slice)
	if err != nil {
		return
	}

	outFile := fmt.Sprintf(out, number)
	err = writeFile(s.in, start, stop, outFile)
	if err == nil {
		fmt.Printf("%d-", start)
		if stop > 0 {
			fmt.Printf("%d", stop-1)
		} else {
			fmt.Printf("EOF")
		}
		fmt.Printf(" saved in %s\n", outFile)
	}

	return
}

// Close スライサーの終了
func (s Slicer) Close() (err error) {
	if s.in != nil {
		err = s.in.Close()
	}
	return
}

// slice
// len = 1
// n1           => start=stop=n1
// len = 2
// :            => start=stop=0
// n1:n2        => start=n1, stop=n2
// n1:          => start=n1, stop=0
// :n2          => start=0, stop=n2
// len = 3
// ::           => start=stop=0
// ::out
// n1::
// n1::out
// n1:n2:
// n1:n2:out
// :n2:
// :n2:out

func parseSlice(slice string) (start, stop int, out string, err error) {
	a := strings.Split(slice, ":")
	l := len(a)

	if l < 1 {
		err = fmt.Errorf("empty slice")
		return
	}
	if l > 3 {
		err = fmt.Errorf("abnormal slice: %s", slice)
		return
	}

	if a[0] != "" {
		start, err = strconv.Atoi(a[0])
		if err == nil {
			if start < 0 {
				// 負数の指定は禁止。ただし 0 はOK.
				err = fmt.Errorf("abnormal slice: start is negative: %s", slice)
			}
		}
	}

	if l == 1 {
		return
	}

	if a[1] != "" {
		relative := false
		if a[1][0] == '+' {
			relative = true
		}
		stop, err = strconv.Atoi(a[1])
		if err == nil {
			if relative {
				stop = start + stop
			}
			if stop <= 0 {
				// 負数の指定は禁止
				err = fmt.Errorf("abnormal slice: stop is negative: %s", slice)
			} else {
				if start > stop {
					err = fmt.Errorf("abnormal slice: start>stop: %s", slice)
				}
			}
		}
	}

	if l == 2 {
		return
	}

	out = a[2]

	if out == "" {
		out = "%d.dat"
	} else {
		if out[0] == '.' {
			out = "%d" + out
		}
	}
	return
}

func writeFile(in ReadSeekCloser, start, stop int, outFile string) (err error) {
	_, err = in.Seek(int64(start), 0)
	if err != nil {
		return
	}

	var out *os.File
	out, err = os.Create(outFile)
	if err != nil {
		return
	}

	if stop == 0 {
		_, err = io.Copy(out, in)
	} else {
		_, err = io.CopyN(out, in, int64(stop-start))
	}

	err2 := out.Close()
	if err == nil {
		err = err2
	} else {
		fmt.Fprintln(os.Stderr, err2)
	}

	return
}
