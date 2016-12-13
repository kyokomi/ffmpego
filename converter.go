package ffmpego

import (
	"fmt"
	"path"
	"log"
)

const tmpConcatAACFileName = "concat.aac"

func convertTsToMP3(aacDirPath, outputFilePath string) error {
	concatFilePath := path.Join(aacDirPath, tmpConcatAACFileName)
	if err := convertConcatAACFile(aacDirPath, concatFilePath); err != nil {
		return err
	}
	return convertAACToMP3(concatFilePath, outputFilePath)
}

func convertConcatAACFile(inputDirPath, outputFilePath string) error {
	concatFileNames, err := concatFileNames(inputDirPath)
	if err != nil {
		return err
	}

	concatArg := fmt.Sprintf("concat:%s", concatFileNames)
	f, err := newFFMPEG(concatArg)
	if err != nil {
		return err
	}

	f.setArgs("-c", "copy")
	result, err := f.execute(outputFilePath)
	if err != nil {
		return err
	}
	log.Println(string(result))
	return nil
}

func convertAACToMP3(inputFilePath, outputFilePath string) error {
	f, err := newFFMPEG(inputFilePath)
	if err != nil {
		return err
	}

	f.setArgs(
		"-c:a", "libmp3lame",
		"-ac", "2",
		"-q:a", "2",
	)
	result, err := f.execute(outputFilePath)
	if err != nil {
		return err
	}
	log.Println(string(result))
	return nil
}
