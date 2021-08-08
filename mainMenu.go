package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMenuWindow struct {
	*walk.MainWindow
}

func showMainMenu() {
	mw := &MyMenuWindow{}
	MW := MainWindow{
		AssignTo: &mw.MainWindow, // Widgetを実体に割り当て
		Title:    "StatsWindow",
		Size:     Size{Width: 400, Height: 150},
		Layout:   VBox{}, // ウィジェットを垂直方向に並べる
		Children: []Widget{
			PushButton{
				Text:      "データ記録",
				MaxSize:   Size{Width: 80, Height: 30},
				MinSize:   Size{Width: 80, Height: 30},
				OnClicked: mw.inputMenuClicked,
			},
			PushButton{
				Text:      "統計データ",
				MaxSize:   Size{Width: 80, Height: 30},
				MinSize:   Size{Width: 80, Height: 30},
				OnClicked: mw.statsMenuClicked,
			},
			PushButton{
				Text:      "設定",
				MaxSize:   Size{Width: 80, Height: 30},
				MinSize:   Size{Width: 80, Height: 30},
				OnClicked: mw.optionMenuClicked,
			},
		},
	}
	if _, err := MW.Run(); err != nil {
		log.Fatal(err)
	}
}
