package data

import (
	"bufio"
	"log"
	"os"
)

type FileReader interface {
	NextLine() string
	HasNext() bool
	Close()
}

type FileReaderImpl struct {
	dataPath string
	file     *os.File
	scanner  *bufio.Scanner
	done     bool
	buf      string
}

func (impl *FileReaderImpl) NextLine() string {
	t := impl.buf
	impl.done = !impl.scanner.Scan()

	if err := impl.scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		v := impl.scanner.Text()
		impl.buf = v
	}

	return t
}

func (impl *FileReaderImpl) HasNext() bool {
	return !impl.done
}

func (impl *FileReaderImpl) Close() {
	err := impl.file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func NewFileReader(dp string) FileReader {
	file, err := os.Open(dp)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	sr := &FileReaderImpl{dataPath: dp, file: file, scanner: scanner, done: false}
	// read one line, for HasNext
	sr.NextLine()
	return sr
}
