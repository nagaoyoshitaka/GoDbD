package main

import (
	"log"

	"github.com/lxn/walk"
	//. "github.com/lxn/walk/declarative"
)

func main() {
	showLoginMenu()
}

type MyStatsWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
}

//統計データウィンドウ表示
func (mw *MyMenuWindow) statsMenuClicked() {
}

//設定ウィンドウ表示
func (mw *MyMenuWindow) optionMenuClicked() {
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
