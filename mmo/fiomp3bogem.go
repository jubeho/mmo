package mmo

import (
	"fmt"

	id3v2bog "github.com/bogem/id3v2/v2"
)

func (mmo *MyMusicOrga) ReadMP3Bogem(fp string) (*AudioFile, error) {
	tag, err := id3v2bog.Open(fp, id3v2bog.Options{Parse: true})
	if err != nil {
		return nil, fmt.Errorf("could not open mp3 file with bogem %s: %v", fp, err)
	}
	defer tag.Close()

	af := NewAudioFile(fp)
	tagVersion := fmt.Sprintf("ID3v2.%d", tag.Version())
	af.TagType = tagVersion
	af.FileType = "MP3"

	for frameName, frms := range tag.AllFrames() {
		tagname := mmo.TagMapping.GetTagName(af.TagType, frameName, "MMO")
		for _, frm := range frms {
			// TextFrames:
			textFrm, ok := frm.(id3v2bog.TextFrame)
			if ok {
				af.SetTagMMO(tagname, textFrm.Text, textFrm.Encoding.Name)
				continue
			}

			// CommentFrames
			comFrm, ok := frm.(id3v2bog.CommentFrame)
			if ok {
				af.SetTagMMO(tagname, comFrm.Text, "")
				continue
			}

			//UserDefinedTextFrame
			udtFrm, ok := frm.(id3v2bog.UserDefinedTextFrame)
			if ok {
				af.SetTagMMO(udtFrm.Description, udtFrm.Value, udtFrm.Encoding.Name)
				continue
			}

			// PictureFrame
			picFrm, ok := frm.(id3v2bog.PictureFrame)
			if ok {
				pic := &Picture{
					TagName:  frameName,
					MimeType: picFrm.MimeType,
					Data:     picFrm.Picture,
				}
				af.Pics = append(af.Pics, pic)
				continue
			}

			// ChapterFrame
			chapterFrm, ok := frm.(id3v2bog.ChapterFrame)
			if ok {
				fmt.Print(chapterFrm.UniqueIdentifier(), chapterFrm.Description, "\n")
				continue
			}

			// PopularimeterFrame
			popmFrm, ok := frm.(id3v2bog.PopularimeterFrame)
			if ok {
				af.SetTagMMO(
					"POPM",
					fmt.Sprintf("Email:%s§§Rating:%v§§Counter:%v", popmFrm.Email, popmFrm.Rating, popmFrm.Counter),
					"",
				)
				continue
			}

			// UFID-Frame
			ufidFrm, ok := frm.(id3v2bog.UFIDFrame)
			if ok {
				af.SetTagMMO(
					ufidFrm.UniqueIdentifier(),
					string(ufidFrm.Identifier),
					"",
				)
				continue
			}

			// Lyrics Frame (unsyncronised)
			uslFrm, ok := frm.(id3v2bog.UnsynchronisedLyricsFrame)
			if ok {
				af.SetTagMMO("LYRICS_UNSYNC", uslFrm.Lyrics, uslFrm.Encoding.Name)
				continue
			}
			/*
				unknownFrm, ok := frm.(id3v2bog.UnknownFrame)
				if ok {
					fmt.Println("unknonw Frame:")
					fmt.Printf("\tID: %v\n", unknownFrm.UniqueIdentifier())
					fmt.Printf("\tsize: %d\n", unknownFrm.Size())
					continue
				}
			*/
		}
	}
	return af, nil
}
