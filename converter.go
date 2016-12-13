package ffmpego

import (
	"fmt"
	"path"
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
	concatFileNames, err := ConcatFileNames(inputDirPath)
	if err != nil {
		return err
	}

	concatArg := fmt.Sprintf("concat:%s", concatFileNames)
	f, err := newFFMPEG(concatArg)
	if err != nil {
		return err
	}

	f.setArgs("-c", "copy")
	// TODO: console出力内容をlogに吐く
	return f.run(outputFilePath)
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
	// TODO: console出力内容をlogに吐く
	return f.run(outputFilePath)
}
