package mmo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	//Log.SetLevel(logrus.DebugLevel)
	Log.SetLevel(logrus.InfoLevel)
}

type MyMusicOrga struct {
	AudioFiles []*AudioFile
	TagMapping *TagMap

	FileTypes map[string][]*AudioFile //key: flac, mp3
	TagTypes  map[string][]*AudioFile
	TagNames  map[string][]*AudioFile // key: mmo framevariant like Title, Album
	Genres    map[string][]*AudioFile
}

func NewMMORaw() (*MyMusicOrga, error) {
	mmo := &MyMusicOrga{}
	mmo.TagTypes = make(map[string][]*AudioFile)
	mmo.TagNames = make(map[string][]*AudioFile)
	mmo.FileTypes = make(map[string][]*AudioFile)
	mmo.Genres = make(map[string][]*AudioFile)

	var err error
	mmo.TagMapping, err = NewTagMap("/home/juergen/opi/beckx.online/mmo/tagmap.csv")
	if err != nil {
		return nil, err
	}
	return mmo, nil
}

func NewMMO(args []string) (*MyMusicOrga, error) {
	mmo := &MyMusicOrga{}
	mmo.TagTypes = make(map[string][]*AudioFile)
	mmo.TagNames = make(map[string][]*AudioFile)
	mmo.FileTypes = make(map[string][]*AudioFile)
	mmo.Genres = make(map[string][]*AudioFile)

	var err error
	mmo.TagMapping, err = NewTagMap("/home/juergen/opi/beckx.online/mmo/tagmap.csv")
	if err != nil {
		return nil, err
	}

	if len(args) == 0 {
		err = mmo.ReadLib("/home/juergen/mmo.csv")
		if err != nil {
			return nil, err
		}
	} else {
		mmo.ImportFiles(args)
	}

	return mmo, nil
}

func (mmo *MyMusicOrga) ReadLib(fp string) error {
	f, err := os.OpenFile(fp, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open feil to write csv-values %v: %v", fp, err)
	}
	defer f.Close()

	c := csv.NewReader(f)
	c.Comma = ';'
	c.FieldsPerRecord = -1

	rowCount := 0
	for {
		rec, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if rowCount > 0 {
			mmo.AudioFiles = append(mmo.AudioFiles, AudiofileFromCSVRow(rec))
		}
		rowCount++
	}
	return nil
}

// ImportFiles reads in all valid audiofiles from args and
// populates mmo.AudioFiles
func (mmo *MyMusicOrga) ImportFiles(args []string) error {
	// collecting files
	files := []string{}
	// err := FindFiles(args, []string{".mp3", ".flac"}, &files)
	err := FindFiles(args, []string{".mp3", ".flac"}, &files)
	if err != nil {
		return err
	}
	for _, fp := range files {
		af, err := mmo.ReadAudiofile(fp)
		if err != nil {
			Log.Error(err)
			continue
		}
		mmo.AudioFiles = append(mmo.AudioFiles, af)
		mmo.FileTypes[af.FileType] = append(mmo.FileTypes[af.FileType], af)
		mmo.TagTypes[af.TagType] = append(mmo.TagTypes[af.TagType], af)
		for _, tag := range af.Tags {
			mmo.TagNames[tag[0]] = append(mmo.TagNames[tag[0]], af)
		}
	}
	return nil
}

func FindFiles(args []string, validExts []string, fileList *[]string) error {
	for _, arg := range args {
		fi, err := os.Stat(arg)
		if err != nil {
			return fmt.Errorf("could not stat file/path %v: %v", arg, err)
		}
		if fi.IsDir() {
			Log.Debugln("arg is a directory:", arg)
			fis, err := os.ReadDir(arg)
			if err != nil {
				return fmt.Errorf("could not read dir %v: %v", arg, err)
			}
			dirContent := []string{}

			for _, thisFI := range fis {
				dirContent = append(dirContent, path.Join(arg, thisFI.Name()))
			}
			err = FindFiles(dirContent, validExts, fileList)
			if err != nil {
				return err
			}
		}
		ext := filepath.Ext(arg)
		if !hasValidExtension(ext, validExts) {
			Log.Debugf("skip non audio file: %v", arg)
			continue
		}
		*fileList = append(*fileList, arg)
	}
	return nil
}

func hasValidExtension(ext string, validExts []string) bool {
	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}
	return false
}

func (mmo *MyMusicOrga) AudiofilesToLib(fp string) error {
	rec := [][]string{{"Filepath", "FileType", "TagType", "AlbArtist", "Album", "Year", "CDNum", "TrackNum", "Artist",
		"Title", "Genre", "Comments"}}
	for _, af := range mmo.AudioFiles {
		rec = append(rec, af.CreateCSVRow())
	}
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file to write csv-values %v: %v", fp, err)
	}
	defer f.Close()

	c := csv.NewWriter(f)
	c.Comma = ';'
	err = c.WriteAll(rec)
	if err != nil {
		return fmt.Errorf("could not write csv-records to file %v: %v", fp, err)
	}
	return nil
}

func (mmo *MyMusicOrga) GetAudiofileSummary() string {
	s := fmt.Sprintf("File_Count: %d\n", len(mmo.AudioFiles))
	for k, v := range mmo.FileTypes {
		s = fmt.Sprintf("%sFileType_Count %s: %d\n", s, k, len(v))
	}
	for k, v := range mmo.TagTypes {
		s = fmt.Sprintf("%sTagTypes_count %s: %d\n", s, k, len(v))
	}
	for k, v := range mmo.TagNames {
		s = fmt.Sprintf("%sTagNames_count %s: %d\n", s, k, len(v))
	}
	for k, v := range mmo.Genres {
		s = fmt.Sprintf("%sGenre_Count %s: %d\n", s, k, len(v))
	}
	return s
}

// WriteTags write Audiofile-Tags to file for all Audiofiles in mmo.Audiofiles
func (mmo *MyMusicOrga) WriteTags() error {

	return nil
}
