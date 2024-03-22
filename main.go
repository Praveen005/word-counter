package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"

	"os"
	"runtime/pprof"
	"unicode"
)

var cpuprofile = flag.String("cpu-profile", "", "write cpu profile to `file`")

// var buf [1]byte
func readByte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func wc() int {
	// f, err := os.Open(os.Args[1])
	f, err := os.Open("ego.txt")

	if err != nil {
		// log.Fatalf("Couldn't open the file %q: %v", os.Args[1], err)
		log.Fatalf("Couldn't open the file %s: %v", "ego.txt", err)
	}
	defer f.Close()

	words := 0
	inword := false
	spaceAtEOF := false

	// The default buffer size used by bufio.NewReader() is 4096 bytes (4 KB)
	b := bufio.NewReader(f) // Optimization 1

	for {
		// r, err := readByte(f)  // Baseline code
		// Instaed of reading byte by byte, we now read from a buffered Reader.
		// Subsequent reads from b often get served from this buffer, reducing the number of direct interactions with the file system leading to faster input operations.
		r, err := readByte(b) // Optimization 1

		if err == io.EOF {
			if !spaceAtEOF { // the file ended with a non-space char, and hence it was not accounted for below, so add it here.
				words++
			}
			break
		}

		if err != nil {
			// log.Fatalf("Couldn't open the file %q: %v", os.Args[1], err)
			log.Fatalf("Couldn't open the file %s: %v", "ego.txt", err)
		}

		if unicode.IsSpace(r) && inword {
			words++
			inword = false
			spaceAtEOF = true // the text file is ending with a space(inc. \n) character and hence the last word has been accounted for.
		}
		if unicode.IsPunct(r) || unicode.IsLetter(r) {
			inword = true
			spaceAtEOF = false
		}
	}
	// fmt.Printf("%q: %v\n", "ego.txt", words)
	return words

}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	words := wc()
	fmt.Printf("%q: %v\n", "ego.txt", words)
}
