package ffmpego

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/grafov/m3u8"
)

func readChunks(input io.Reader) ([]string, error) {
	playlist, listType, err := m3u8.DecodeFrom(input, true)
	if err != nil || listType != m3u8.MEDIA {
		return nil, err
	}
	p := playlist.(*m3u8.MediaPlaylist)

	var chunks []string
	for _, v := range p.Segments {
		if v != nil {
			chunks = append(chunks, v.URI)
		}
	}
	return chunks, nil
}

func downloadChunks(m3u8URI string) ([]string, error) {
	resp, err := http.Get(m3u8URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	chunks, err := readChunks(resp.Body)
	if err != nil {
		return nil, err
	}

	// 相対パスを考慮
	u, err := url.Parse(m3u8URI)
	if err != nil {
		return nil, err
	}
	for i := range chunks {
		if !strings.HasPrefix(chunks[i], "http") {
			chunkURI := url.URL{
				Scheme: u.Scheme,
				Host:   u.Host,
				Path:   path.Join(u.Path[:strings.LastIndex(u.Path, "/")+1], chunks[i])}
			chunks[i] = chunkURI.String()
		}
	}
	return chunks, nil
}
