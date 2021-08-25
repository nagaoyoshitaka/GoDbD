package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyLoginWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
	label    *walk.Label
}

func showLoginMenu() {
	lw := &MyLoginWindow{}
	LW := MainWindow{
		AssignTo: &lw.MainWindow, // Widgetを実体に割り当て
		Title:    "Login",
		Size:     Size{Width: 300, Height: 50}, //width×height
		Layout:   VBox{},                       // ウィジェットを垂直方向に並べる
		Children: []Widget{
			Label{
				Text:     "password",
				AssignTo: &lw.label,
			},
			TextEdit{
				AssignTo: &lw.textArea,
				MaxSize:  Size{Width: 15, Height: 20},
			},
			PushButton{
				Text:      "Login",
				OnClicked: lw.pbClicked, // ログイン試行イベント
			},
		},
	}

	if _, err := LW.Run(); err != nil {
		log.Fatal(err)
	}
}

// ログイン試行
func (lw *MyLoginWindow) pbClicked() {
	if lw.textArea.Text() == "dbd" {
		lw.login()
	} else {
		showMessage("error", "incorrect password")
	}
}

func (lw *MyLoginWindow) login() {

	lw.SetVisible(false)
	showMainMenu()
	lw.SetVisible(true)
}
