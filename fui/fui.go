package fui

import (
	"fmt"

	"beckx.online/mmo/mmo"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AuFi struct {
	File     string
	Selected bool
}

type AuFiList struct {
	AufIs []AuFi
}

type FuiData struct {
	app     fyne.App
	win     fyne.Window
	tagList *widget.List

	mymo       *mmo.MyMusicOrga
	tagsToShow []string // all tags shown in tag-list
	unionTags  []string // collection of all tags from given files

	datalist AuFiList
}

func InitFui(givenMMO *mmo.MyMusicOrga) error {
	fd := &FuiData{
		mymo: givenMMO,
	}

	fd.initValues()

	fd.app = app.New()
	fd.win = fd.app.NewWindow("MyMusicOrganizer")

	fd.win.SetContent(
		container.NewHSplit(
			fd.makeFileList(),
			fd.makeTagList(),
			//fd.makeValueList(),
		),
	)

	fd.win.Resize(fyne.NewSize(800, 600))
	fd.win.CenterOnScreen()
	fd.win.ShowAndRun()

	return nil
}

func (fd *FuiData) initValues() {
	for k := range fd.mymo.TagNames {
		fd.unionTags = append(fd.unionTags, k)
	}

	fd.datalist = AuFiList{
		AufIs: []AuFi{
			AuFi{File: "Sepp", Selected: false},
			AuFi{File: "Sepp", Selected: false},
		},
	}
}

func (fd *FuiData) updateShowTags(af *mmo.AudioFile) {
	for _, tag := range af.Tags {
		addUniqueToList(tag[0], &fd.tagsToShow)
	}
}

func (fd *FuiData) makeFileList() *fyne.Container {
	list := widget.NewList(
		func() int {
			return len(fd.mymo.AudioFiles)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewCheck("", func(b bool) {
			}), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(fd.mymo.AudioFiles[id].Filepath)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
	}
	/*
		list.OnUnselected = func(id widget.ListItemID) {
			label.SetText("Select An Item From The List")
			icon.SetResource(nil)
		}
		list.Select(125)
	*/

	return container.NewHBox(list)

}

func (fd *FuiData) makeTagList() *fyne.Container {
	fd.tagList = widget.NewList(
		func() int {
			return len(fd.unionTags)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(fd.tagsToShow[id])
		},
	)

	fd.tagList.OnSelected = func(id widget.ListItemID) {
		fmt.Println(fd.tagsToShow[id])
	}
	/*
		list.OnUnselected = func(id widget.ListItemID) {
			label.SetText("Select An Item From The List")
			icon.SetResource(nil)
		}
		list.Select(125)
	*/

	return container.NewMax(fd.tagList)

}

func addUniqueToList(s string, list *[]string) {
	for _, le := range *list {
		if le == s {
			return
		}
	}
	*list = append(*list, s)
}
