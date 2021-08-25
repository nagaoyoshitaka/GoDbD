package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyInformWindow struct {
	*walk.MainWindow
	informLabel      *walk.Label
	killerLabel      *walk.Label
	killLabel        *walk.Label
	mapAndStageLabel *walk.Label
}

func showInformDialog(killer string, kill string, mp string, stage string) {
	id := &MyInformWindow{}
	ID := MainWindow{
		AssignTo: &id.MainWindow, // Widgetを実体に割り当て
		Title:    "Success",
		Size:     Size{Width: 300, Height: 50}, //width×height
		Layout:   VBox{},                       // ウィジェットを垂直方向に並べる
		Children: []Widget{
			Label{
				Text:     "Successful Registration!!",
				AssignTo: &id.informLabel,
			},
			Label{
				Text:     "■ キラー: " + killer,
				AssignTo: &id.killerLabel,
			},
			Label{
				Text:     "■ キル数: " + kill,
				AssignTo: &id.killLabel,
			},
			Label{
				Text:     "■ マップ: " + mp + "/" + stage,
				AssignTo: &id.mapAndStageLabel,
			},
		},
	}

	if _, err := ID.Run(); err != nil {
		log.Fatal(err)
	}
}

func showMessage(title string, message string) {
	id := &MyInformWindow{}
	ID := MainWindow{
		AssignTo: &id.MainWindow, // Widgetを実体に割り当て
		Title:    title,
		Size:     Size{Width: 300, Height: 50}, //width×height
		Layout:   VBox{},                       // ウィジェットを垂直方向に並べる
		Children: []Widget{
			Label{
				Text:     message,
				AssignTo: &id.informLabel,
			},
		},
	}

	if _, err := ID.Run(); err != nil {
		log.Fatal(err)
	}
}
