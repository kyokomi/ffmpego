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

	fgo := ffmpego.New()
	if err := fgo.M3U8ConvertMP3(*uri, *outputPath); err != nil {
		log.Fatalln(err)
	}
}
