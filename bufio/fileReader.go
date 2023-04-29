package bufio

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	filePath = "largefile.txt"
)

func TestReadFile() {
	// Create a large file with random data for testing
	createLargeFile()

	// Benchmark reading the file without using bufio.Reader
	start := time.Now()
	readFileWithoutBuffer()
	end := time.Now()
	fmt.Printf("Time taken without buffer: %v\n", end.Sub(start))

	// Benchmark reading the file using bufio.Reader
	start = time.Now()
	readFileWithBuffer()
	end = time.Now()
	fmt.Printf("Time taken with buffer: %v\n", end.Sub(start))
}

func createLargeFile() {
	f, _ := os.Create(filePath)
	defer f.Close()

	for i := 0; i < 1000000; i++ {
		f.WriteString("abcdefghijklmnopqrstuvwxyz\n")
	}
}

func readFileWithoutBuffer() {
	f, _ := os.Open(filePath)
	defer f.Close()

	var b []byte
	for {
		buf := make([]byte, 1024)
		n, err := f.Read(buf)
		if err != nil {
			break
		}
		b = append(b, buf[:n]...)
	}
}

func readFileWithBuffer() {
	f, _ := os.Open(filePath)
	defer f.Close()

	r := bufio.NewReader(f)

	var b []byte
	for {
		buf := make([]byte, 1024)
		n, err := r.Read(buf)
		if err != nil {
			break
		}
		b = append(b, buf[:n]...)
	}
}
