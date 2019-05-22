package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/dolchi21/cobranzas-hash/lib"
)

var configFilename = flag.String("c", "config.yml", "Specify a custom config.yml")
var inputFilename = flag.String("i", "input.txt", "txt file to compare")
var outputFilename = flag.String("o", "output.txt", "txt file to write result")

func main() {
	flag.Parse()
	lib.LoadConfig(configFilename)
	//hashes := lib.GetSourceHashes(1, "DAI")

	file, err := os.Open(*inputFilename)
	if nil != err {
		panic("os.Open failed.")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	//fmt.Println(len(hashes))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		hasher := md5.New()
		data := []byte(line)
		hasher.Write(data)
		hash := hasher.Sum(nil)
		fmt.Println(hash, line, hex.EncodeToString(hash))
	}
}
