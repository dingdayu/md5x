package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

func md51(filePath string) (string, time.Duration, error) {
	start_time := time.Now() // get current time
	file, err := os.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	md5h := md5.New()
	io.Copy(md5h, file)

	// 文件大小
	//fileInfo, err := os.Stat(filePath)
	//fileSize := fileInfo.Size() //获取size
	//fmt.Println("File Size: " , fileSize / 1024 / 1024, "M")

	elapsed := time.Since(start_time)
	return fmt.Sprintf("%x", md5h.Sum([]byte(""))), elapsed, nil
}

func md52(filePath string) (string, time.Duration, error) {
	startTime := time.Now() // get current time
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	defer file.Close()

	// 文件大小
	fileInfo, err := os.Stat(filePath)
	fileSize := fileInfo.Size() //获取size
	//fmt.Println("File Size: " , fileSize / 1024 / 1024, "M")

	blocks := uint64(math.Ceil(float64(fileSize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(fileSize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	elapsed := time.Since(startTime)
	return fmt.Sprintf("%x", hash.Sum(nil)), elapsed, nil
}
