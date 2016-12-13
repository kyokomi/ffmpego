package ffmpego

import (
	"log"
	"os"
)

func Download(uri, outputFilePath string) error {
	chunks, err := GetChunklistFromM3U8(uri)
	if err != nil {
		return err
	}

	// TODO: 設定とかで変更できるようにする
	workDirPath := "./tmp"

	if err := os.MkdirAll(workDirPath, 0700); err != nil {
		log.Println(err)
	}
	defer os.RemoveAll(workDirPath)

	if err := BulkDownload(chunks, workDirPath); err != nil {
		return err
	}

	if err := convertTsToMP3(workDirPath, outputFilePath); err != nil {
		return err
	}

	return nil
}
