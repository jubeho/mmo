package mmo

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	ColMMO = iota
	ColTAGVALUE
	ColENCODING
)

type AudioFile struct {
	Filepath string
	FileType string // MP3, FLAC
	TagType  string // ID3v1, ID3v22, ...VORBISCOMMENT

	FlacVorbisCommentCount int

	// MMO, TagValue,Encoding
	Tags [][]string

	Pics []*Picture
}

type Picture struct {
	TagName  string
	MimeType string
	Data     []byte
}

func NewAudioFile(fp string) *AudioFile {
	af := &AudioFile{}
	af.Filepath = fp

	// Default MMO Frames
	for _, mapRow := range MapTemplate[1:] {
		af.Tags = append(af.Tags, []string{
			mapRow[ColMMO], "", "",
		})
	}

	return af
}

func (af *AudioFile) SetTagMMO(tagname, tagvalue, tagencoding string) {
	tagname = strings.ToUpper(tagname)
	tagvalue = strings.ReplaceAll(tagvalue, "\n", "<NL>")
	tagvalue = strings.ReplaceAll(tagvalue, "\r", "<RL>")
	for i, tag := range af.Tags {
		if tag[ColMMO] == tagname {
			af.Tags[i][ColTAGVALUE] = tagvalue
			af.Tags[i][ColENCODING] = tagencoding
			return
		}
	}
	af.Tags = append(af.Tags, []string{tagname, tagvalue, tagencoding})
}

func (af *AudioFile) String() string {
	s := fmt.Sprintf("%s\n", af.Filepath)
	for _, tag := range af.Tags {
		s = fmt.Sprintf("%s%s:%s:%s\n", s, tag[0], tag[1], tag[2])
	}
	for _, pic := range af.Pics {
		s = fmt.Sprintf("%sPicture: %v (%v)\n", s, pic.MimeType, pic.TagName)
	}

	return s
}

func (af *AudioFile) CreateCSVRow() []string {
	row := []string{af.Filepath, af.FileType, af.TagType}

	for i, tagRow := range af.Tags {
		if i < len(TagTypeCols) {
			row = append(row, tagRow[ColTAGVALUE])
		} else {
			row = append(row, fmt.Sprintf("%s=%s", tagRow[ColMMO], tagRow[ColTAGVALUE]))
		}
	}
	//fmt.Println(row)
	return row
}

func AudiofileFromCSVRow(row []string) *AudioFile {
	af := NewAudioFile(row[0])
	af.FileType = row[1]
	af.TagType = row[2]

	if len(row) > 3 {
		fmt.Println(row)
		for i, v := range []string{"AlbArtist", "Album", "Year", "CDNum", "TrackNum", "Artist", "Title", "Genre", "Comment"} {
			af.Tags[i][0] = v
			af.Tags[i][1] = row[i+3]
		}
	}

	if len(row) > 12 {
		for _, v := range row[12:] {
			sep := strings.Index(v, "=")
			if sep == -1 {
				logrus.Warnf("no seperator in csv-row-field %v", v)
				continue
			}
			af.Tags = append(af.Tags, []string{v[:sep], v[sep+1:], ""})
		}
	}
	return af
}

func (af *AudioFile) GetAlbArtist() string {
	for _, tag := range af.Tags {
		if strings.ToUpper(tag[0]) == "ALBARTIST" {
			return tag[1]
		}
	}
	return ""
}

func (af *AudioFile) GetAlbum() string {
	for _, tag := range af.Tags {
		if strings.ToUpper(tag[0]) == "ALBUM" {
			return tag[1]
		}
	}
	return ""
}

func (af *AudioFile) GetTagValue(tagname string) string {
	for _, tag := range af.Tags {
		// if strings.ToUpper(tag[0]) == strings.ToUpper(tagname) {
		if strings.EqualFold(tag[0], tagname) {
			return tag[1]
		}
	}
	return ""
}
