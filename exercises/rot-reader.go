package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) rot13(b byte, basis byte) byte {
	return (b-basis+13)%26 + basis
}

func (rot rot13Reader) cryptByte(b byte) byte {

	if b >= 'A' && b <= 'Z' {
		return rot.rot13(b, 'A')
	}

	if b >= 'a' && b <= 'z' {
		return rot.rot13(b, 'a')
	}

	return b
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rot.cryptByte(b[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
