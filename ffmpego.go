package ffmpego

import (
	"log"
	"os"
)

const (
	defaultWorkDirPath   = "./tmp"
	defaultMaxAttempts   = 4
	defaultMaxGoroutines = 16
)

// FfMPEGo is ffmpeg wrapper
type FfMPEGo struct {
	workDirPath           string
	downloadMaxAttempts   int
	downloadMaxGoroutines int
}

// New create FfMPEGo
func New() *FfMPEGo {
	return &FfMPEGo{
		workDirPath:           defaultWorkDirPath,
		downloadMaxAttempts:   defaultMaxAttempts,
		downloadMaxGoroutines: defaultMaxGoroutines,
	}
}

// M3U8ConvertMP3 m3u8 convert to mp3
func (f *FfMPEGo) M3U8ConvertMP3(uri, outputFilePath string) error {
	chunks, err := downloadChunks(uri)
	if err != nil {
		return err
	}

	defer os.RemoveAll(f.workDirPath)
	if err := os.MkdirAll(f.workDirPath, 0700); err != nil {
		log.Println(err)
	}

	if err := bulkDownload(f.downloadMaxAttempts, f.downloadMaxGoroutines, chunks, f.workDirPath); err != nil {
		return err
	}

	if err := convertTsToMP3(f.workDirPath, outputFilePath); err != nil {
		return err
	}

	return nil
}
