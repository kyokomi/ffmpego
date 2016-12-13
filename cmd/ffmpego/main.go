package main

import (
	"flag"
	"log"

	"github.com/kyokomi/ffmpego"
)

func main() {
	uri := flag.String("i", "", "convert m3u8 url")
	outputPath := flag.String("o", "result.mp3", "output file path")
	flag.Parse()

	if err := ffmpego.Download(*uri, *outputPath); err != nil {
		log.Fatalln(err)
	}
}
