package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func readFilesName() (dirs, files []string) {
	items, err := ioutil.ReadDir("./share")
	if err != nil {
		fmt.Println("read dir error")
	}
	for _, item := range items {
		if item.IsDir() {
			dirs = append(dirs, item.Name())
		} else {
			files = append(files, item.Name())
		}
	}
	return dirs, files
}

func fileSize(file *os.File) string {
	fInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		return "None"
	}
	return byteSize(fInfo.Size())
}

func byteSize(n int64) string {
	size := float64(n)
	switch {
	case size >= PB:
		return fmt.Sprintf("%.2fPB", size/PB)
	case size >= TB:
		return fmt.Sprintf("%.2fTB", size/TB)
	case size >= GB:
		return fmt.Sprintf("%.2fGB", size/GB)
	case size >= MB:
		return fmt.Sprintf("%.2fMB", size/MB)
	case size >= KB:
		return fmt.Sprintf("%.2fKB", size/KB)
	}
	return fmt.Sprintf("%.2fB", size)
}
