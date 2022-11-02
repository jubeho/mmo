package mmo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-flac/flacpicture"
	"github.com/go-flac/flacvorbis"
	"github.com/go-flac/go-flac"
)

func (mmo *MyMusicOrga) ReadFlacGoFlac(fp string) (*AudioFile, error) {
	file, err := flac.ParseFile(fp)
	if err != nil {
		return nil, fmt.Errorf("could not parse flac file %v: %v", fp, err)
	}

	af := NewAudioFile(fp)
	af.TagType = "VORBISCOMMENT"
	af.FileType = "FLAC"

	var vcmt *flacvorbis.MetaDataBlockVorbisComment
	// var cmtIdx int

	for _, meta := range file.Meta {
		switch meta.Type {
		case flac.StreamInfo:
		case flac.Padding:
		case flac.Application:
		case flac.SeekTable:
		case flac.VorbisComment:
			af.FlacVorbisCommentCount++
			vcmt, err = flacvorbis.ParseFromMetaDataBlock(*meta)
			if err != nil {
				return nil, fmt.Errorf("could not parse metadatablock from %v: %v", fp, err)
			}
			for _, cmt := range vcmt.Comments {
				sepIdx := strings.Index(cmt, "=")
				if sepIdx < 0 {
					fmt.Printf("equalindex < 0 for %v\n", cmt)
					continue
				}
				frameName := cmt[:sepIdx]
				value := cmt[sepIdx+1:]
				mmoTagName := mmo.TagMapping.GetTagName(af.TagType, frameName, "MMO")
				af.SetTagMMO(mmoTagName, value, "")
			}
		case flac.CueSheet:
		case flac.Picture:
			pic, err := flacpicture.ParseFromMetaDataBlock(*meta)
			if err != nil {
				return nil, fmt.Errorf("could not parse picture meta data block %v: %v", fp, err)
			}
			afPic := &Picture{
				TagName:  pic.Description,
				MimeType: pic.MIME,
				Data:     pic.ImageData,
			}
			af.Pics = append(af.Pics, afPic)
		case flac.Reserved:
		case flac.Invalid:
		}
	}

	return af, nil
}

func (mmo *MyMusicOrga) DeleteAllVorbisComments(fp string) error {
	file, err := flac.ParseFile(fp)
	if err != nil {
		return fmt.Errorf("could not parse flac file %v: %v", fp, err)
	}

	vorbisCommentIdxs := []int{}
	for i, meta := range file.Meta {
		if meta.Type == flac.VorbisComment {
			vorbisCommentIdxs = append(vorbisCommentIdxs, i)
		}
	}

	fmt.Println(vorbisCommentIdxs)
	for i := len(vorbisCommentIdxs) - 1; i >= 0; i-- {
		idx := vorbisCommentIdxs[i]
		fmt.Println(idx)
		if idx == len(file.Meta)-1 {
			file.Meta = file.Meta[:idx]
		} else {
			file.Meta = append(file.Meta[:idx], file.Meta[idx+1:]...)
		}
	}

	return file.Save(fp)
}

func (mmo *MyMusicOrga) DeleteVorbisComment(fp string, sIdx string) error {
	idx, err := strconv.Atoi(sIdx)
	if err != nil {
		return fmt.Errorf("could not convert string-Index to int %v: %v", sIdx, err)
	}

	file, err := flac.ParseFile(fp)
	if err != nil {
		return fmt.Errorf("could not parse flac file %v: %v", fp, err)
	}

	if idx >= len(file.Meta) {
		return fmt.Errorf("vorbis commentindex out of range %d", idx)
	}
	fmt.Printf("%v\n", file.Meta[idx].Type)
	if file.Meta[idx].Type == flac.VorbisComment {
		fmt.Printf("index %d is vorbis-comment... deleting frame...\n", idx)
		if idx == len(file.Meta)-1 {
			file.Meta = file.Meta[:idx]
		} else {
			file.Meta = append(file.Meta[:idx], file.Meta[idx+1:]...)
		}
	}

	return file.Save(fp)
}

func (mmo *MyMusicOrga) FlacInfo(fp string) error {
	file, err := flac.ParseFile(fp)
	if err != nil {
		return fmt.Errorf("could not parse flac file %v: %v", fp, err)
	}

	fmt.Println(fp)
	for i, meta := range file.Meta {
		metaType := ""
		switch meta.Type {
		case 0:
			metaType = "StreamInfo"
			fmt.Printf("%d: %v\n", i, metaType)
		case 1:
			metaType = "Padding"
			fmt.Printf("%d: %v\n", i, metaType)
		case 2:
			metaType = "Application"
			fmt.Printf("%d: %v\n", i, metaType)
		case 3:
			metaType = "SeekTable"
			fmt.Printf("%d: %v\n", i, metaType)
		case 4:
			metaType = "VorbisComment"
			fmt.Printf("%d: %v\n", i, metaType)

			mdb, err := flacvorbis.ParseFromMetaDataBlock(*meta)
			if err != nil {
				return fmt.Errorf("could not parse metadatablock from %v: %v", fp, err)
			}
			fmt.Printf("Vendor: %s\nComments#: %d\n", mdb.Vendor, len(mdb.Comments))
			fmt.Println("------------------------------------------")
		case 5:
			metaType = "CueSheet"
			fmt.Printf("%d: %v\n", i, metaType)
		case 6:
			metaType = "Picture"

			pic, err := flacpicture.ParseFromMetaDataBlock(*meta)
			if err != nil {
				panic(err)
			}
			fmt.Printf("PictureType: %v\n", pic.PictureType)
		case 7:
			metaType = "Reserved"
			fmt.Printf("%d: %v\n", i, metaType)
		default:
			metaType = "Invalid"
			fmt.Printf("%d: %v\n", i, metaType)
		}
	}

	return nil
}
