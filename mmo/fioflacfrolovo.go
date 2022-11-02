package mmo

import (
	"fmt"
	"os"

	"github.com/frolovo22/tag"
)

func (mmo *MyMusicOrga) ReadFlacFrolovo(fp string) (*AudioFile, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	af := NewAudioFile(fp)
	af.TagType = "VorbisComment"
	af.FileType = "FLAC"

	flacfile, err := tag.ReadFLAC(file)
	if err != nil {
		return nil, err
	}

	for framename, value := range flacfile.Tags {
		fmt.Printf("%s: %s\n", framename, value)
	}
	return af, nil
}
