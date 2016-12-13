package ffmpego

import (
	"io"
	"net/http"

	"github.com/grafov/m3u8"
)

// GetChunklist returns a slice of uri string.
func GetChunklist(input io.Reader) ([]string, error) {
	playlist, listType, err := m3u8.DecodeFrom(input, true)
	if err != nil || listType != m3u8.MEDIA {
		return nil, err
	}
	p := playlist.(*m3u8.MediaPlaylist)

	var chunklist []string
	for _, v := range p.Segments {
		if v != nil {
			chunklist = append(chunklist, v.URI)
		}
	}
	return chunklist, nil
}

// GetChunklistFromM3U8 returns a slice of url.
func GetChunklistFromM3U8(uri string) ([]string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return GetChunklist(resp.Body)
}
