package mmo

import (
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func (mmo *MyMusicOrga) ReadAudiofile(fp string) (*AudioFile, error) {
	if strings.ToLower(filepath.Ext(fp)) == ".mp3" {
		af, err := mmo.ReadMP3Bogem(fp)
		if err != nil {
			logrus.Error(err)
			//READ HOWDEN
			return nil, err
		}
		return af, nil

	}
	if strings.ToLower(filepath.Ext(fp)) == ".flac" {
		return mmo.ReadFlacGoFlac(fp)
		// return mmo.ReadFlacFrolovo(fp)

	}
	return nil, nil
}
