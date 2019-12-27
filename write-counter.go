package main

import "fmt"

type WriteCounter struct {
	Total   uint64
	Current uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Current += uint64(n)
	wc.printProgress()
	return n, nil
}

func (wc WriteCounter) printProgress() {
	p := wc.progress()
	fmt.Printf("\rDownloading...  complete %2.2f%% %d/%dmb", p, toMb(wc.Current), toMb(wc.Total))
}

func (wc WriteCounter) progress() float64 {
	return float64(wc.Current) / float64(wc.Total) * 100
}

func toMb(bytes uint64) uint64 {
	return bytes / 1048576
}
