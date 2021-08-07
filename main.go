package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//    "fmt"

type MyMainWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
}

func main() {
	mw := &MyMainWindow{}

	// Windowのレイアウト
	//
	// MainWindow
	//     TextEdit
	//     PushButton
	//
	MW := MainWindow{
		AssignTo: &mw.MainWindow, // Widgetを実体に割り当て
		Title:    "Hello",
		MinSize:  Size{300, 400},
		Layout:   VBox{}, // ウィジェットを垂直方向に並べる
		Children: []Widget{ // ウィジェットを入れるスライス
			TextEdit{ // 最初のウィジェット（テキスト）
				AssignTo: &mw.textArea,
			},
			PushButton{ // 次のウィジェット（ボタン）
				Text: "Hello!", //     PushButtonは割り当てしない？
				//     内部にデータを持たない為？
				OnClicked: mw.pbClicked, // ボタンクリックイベント
			},
		},
	}

	if _, err := MW.Run(); err != nil {
		log.Fatal(err)
	}
}

// ボタンクリック時の処理
func (mw *MyMainWindow) pbClicked() {

	const s string = "Hello 世界！\r\n" //　TextEdit内の改行は\r\n

	text := mw.textArea
	text.AppendText(s) // 文字列追加
}
