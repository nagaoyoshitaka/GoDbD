package main

import (
	"log"
	//. "github.com/lxn/walk/declarative"
)

func main() {
	showLoginMenu()
}

//設定ウィンドウ表示
func (mw *MyMenuWindow) optionMenuClicked() {
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
