package ffmpego

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
)

func bulkDownload(maxAttempts, maxGoroutines int, list []string, output string) error {
	if maxAttempts <= 0 {
		maxAttempts = defaultMaxAttempts
	}

	if maxGoroutines <= 0 {
		maxGoroutines = defaultMaxGoroutines
	}

	var sem = make(chan struct{}, maxGoroutines)
	var errFlag bool
	var wg sync.WaitGroup

	for _, v := range list {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()

			var err error
			for i := 0; i < maxAttempts; i++ {
				sem <- struct{}{}
				err = download(link, output)
				<-sem
				if err == nil {
					break
				}
			}
			if err != nil {
				log.Printf("Failed to download: %s", err)
				errFlag = true
			}
		}(v)
	}
	wg.Wait()

	if errFlag {
		log.Println("error")
		return errors.New("Lack of aac files")
	}
	return nil
}

func download(link, output string) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, fileName := path.Split(link)
	file, err := os.Create(path.Join(output, fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}
