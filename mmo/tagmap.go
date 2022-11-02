package mmo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var MapTemplate [][]string
var TagTypeCols map[string]int // key:TagType,val:colIdx

type TagMap struct {
	rawCSVMap     [][]string     // expects csv-file with header as file-types and mmo
	TagNameRowMMO map[string]int // key:TagNameMMO,val:rowIdx
}

func NewTagMap(fp string) (*TagMap, error) {
	tm := &TagMap{}
	TagTypeCols = make(map[string]int)

	f, err := os.OpenFile(fp, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("could not open csv-file to read %v: %v", fp, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read csv data from %v: %v", fp, err)
		}
		row := []string{}
		for _, v := range rec {
			row = append(row, strings.ToUpper(v))
		}
		tm.rawCSVMap = append(tm.rawCSVMap, row)
	}

	// populate variants: expecting filetypes in first row:
	// row 0: MMO;VORBIS;ID3v2.3;mediainfo;ID3v2.2
	for k, val := range tm.rawCSVMap[0] {
		TagTypeCols[val] = k
	}
	MapTemplate = tm.rawCSVMap
	return tm, nil
}

func (tm *TagMap) GetTagName(srcTagType string, srcTagName string, wantedTagName string) string {
	srcTagType = strings.ToUpper(srcTagType)
	srcTagName = strings.ToUpper(srcTagName)
	wantedTagName = strings.ToUpper(wantedTagName)
	// ceck if srcVariant and wantedVariant exists
	srcTagTypeColIdx, ok := TagTypeCols[srcTagType]
	if !ok {
		logrus.Errorf("could not find column Index for source variant: %s", srcTagType)
		return ""
	}
	wntdTagTypeColIdx, ok := TagTypeCols[wantedTagName]
	if !ok {
		logrus.Errorf("could not find column Index for wanted tag name: %s", wantedTagName)
		return ""
	}

	// search for "line" with srcVariant and srcTagname
	mapLineIdx := -1
	for i, row := range tm.rawCSVMap {
		if row[srcTagTypeColIdx] == srcTagName {
			mapLineIdx = i
			break
		}
	}

	// if source Tagname not found:
	if mapLineIdx < 0 {
		return srcTagName
	}

	return tm.rawCSVMap[mapLineIdx][wntdTagTypeColIdx]
}
