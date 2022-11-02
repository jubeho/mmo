package tui

import (
	"fmt"
	"strconv"

	"beckx.online/mmo/mmo"
	"github.com/AlecAivazis/survey/v2"
)

type EditValues struct {
	AlbArtists     map[string][]*mmo.AudioFile // key:AlbArtistName,val:audiofile
	AlbArtistOrder []string                    //AlbArtistName
	AlbArtistToSet string

	Albums     map[string][]*mmo.AudioFile
	AlbumOrder []string
	AlbumToSet string

	CDNums     map[string][]*mmo.AudioFile
	CDNumOrder []string
	CDNumToSet string

	Years     map[string][]*mmo.AudioFile
	YearOrder []string
	YearToSet string

	Genres     map[string][]*mmo.AudioFile
	GenreOrder []string
	GenreToSet string
}

func EditTags(mymo *mmo.MyMusicOrga) error {
	ev := &EditValues{
		AlbArtists: make(map[string][]*mmo.AudioFile),
		Albums:     make(map[string][]*mmo.AudioFile),
		CDNums:     make(map[string][]*mmo.AudioFile),
		Years:      make(map[string][]*mmo.AudioFile),
		Genres:     make(map[string][]*mmo.AudioFile),
	}

	for _, af := range mymo.AudioFiles {
		val := af.GetAlbArtist()
		_, ok := ev.AlbArtists[val]
		if !ok {
			ev.AlbArtistOrder = append(ev.AlbArtistOrder, val)
		}
		ev.AlbArtists[val] = append(ev.AlbArtists[val], af)

		val = af.GetTagValue("ALBUM")
		_, ok = ev.Albums[val]
		if !ok {
			ev.AlbumOrder = append(ev.AlbumOrder, val)
		}
		ev.Albums[val] = append(ev.Albums[val], af)

		val = af.GetTagValue("CDNUM")
		_, ok = ev.CDNums[val]
		if !ok {
			ev.CDNumOrder = append(ev.CDNumOrder, val)
		}
		ev.CDNums[val] = append(ev.CDNums[val], af)

		val = af.GetTagValue("YEAR")
		_, ok = ev.Years[val]
		if !ok {
			ev.YearOrder = append(ev.YearOrder, val)
		}
		ev.Years[val] = append(ev.Years[val], af)

		val = af.GetTagValue("GENRE")
		_, ok = ev.Genres[val]
		if !ok {
			ev.GenreOrder = append(ev.GenreOrder, val)
		}
		ev.Genres[val] = append(ev.Genres[val], af)
	}

	fmt.Printf("read in %d files...\n", len(mymo.AudioFiles))

	ev.AlbArtistToSet = userIntAct(ev.AlbArtistOrder, ev.AlbArtists, "Album Artis")
	ev.AlbumToSet = userIntAct(ev.AlbumOrder, ev.Albums, "Album")
	ev.CDNumToSet = userIntAct(ev.CDNumOrder, ev.CDNums, "CD-Number")
	ev.YearToSet = userIntAct(ev.YearOrder, ev.Years, "Year")
	ev.GenreToSet = userIntAct(ev.GenreOrder, ev.Genres, "Genre")

	fmt.Println("Summary for Tags to write:")
	fmt.Println("--------------------------")
	fmt.Printf("Album Artist: %s\n", ev.AlbArtistToSet)
	fmt.Printf("Album: %s\n", ev.AlbumToSet)
	fmt.Printf("CD-Number: %s\n", ev.CDNumToSet)
	fmt.Printf("Year: %s\n", ev.YearToSet)
	fmt.Printf("Genre: %s\n", ev.GenreToSet)

	return nil
}

func userIntAct(selectionList []string, selectionMap map[string][]*mmo.AudioFile, tagLabel string) string {
	fmt.Printf("%s (%d):\n", tagLabel, len(selectionList))
	i := 1
	for _, v := range selectionList {
		fmt.Printf("\t%2d: '%s' (%d)\n", i, v, len(selectionMap[v]))
		i++
	}

	tagValue := ""
	sel := -1
	stringSel := ""
	var err error
	prompt := &survey.Input{
		Message: fmt.Sprintf("Select a %s to write to all files or ENTER to choose an own one...:", tagLabel),
	}
	survey.AskOne(prompt, &stringSel)

	if stringSel == "" {
		prompt := &survey.Input{
			Message: tagLabel,
		}
		survey.AskOne(prompt, &tagValue)
		return tagValue
	}

	sel, err = strconv.Atoi(stringSel)
	if err != nil || sel-1 >= len(selectionList) || sel == 0 {
		fmt.Printf("sie müssen einen gültige Zahl eingeben! %v\n", stringSel)
		return ""
	}
	return selectionList[sel-1]

}
