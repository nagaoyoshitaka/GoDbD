package main

import (
	"log"

	"encoding/csv"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type MyLoginWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
	label    *walk.Label
}

type MyMenuWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
}

type MyInputWindow struct {
	*walk.MainWindow
	killerModel    *Model
	mapModel       *Model
	stageModel     *Model
	killerLabel    *walk.Label
	killLabel      *walk.Label
	mapLabel       *walk.Label
	stageLabel     *walk.Label
	killerComboBox *walk.ComboBox
	mapComboBox    *walk.ComboBox
	stageComboBox  *walk.ComboBox
}

type MyStatsWindow struct {
	*walk.MainWindow
	textArea *walk.TextEdit
}

type Model struct {
	walk.ListModelBase
	items []Item
}

func NewMapModel() *Model {
	file1, err := os.Open("map_name.csv")
	failOnError(err)
	defer file1.Close()
	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))

	mapList, err := reader.Read() // 1行読み出す
	failOnError(err)
	mapList = distinct(mapList)
	cnt := len(mapList)
	m := &Model{items: make([]Item, cnt)}
	for i := 0; i < cnt; i++ {
		name := mapList[i]
		value := mapList[i]
		m.items[i] = Item{name, value}
	}
	return m
}

func NewStageModel() *Model {
	file1, err := os.Open("map_name.csv")
	failOnError(err)
	defer file1.Close()
	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))

	mapList, err := reader.Read() // 1行読み出す
	failOnError(err)
	stageList, err := reader.Read() // 1行読み出す
	failOnError(err)
	cnt := len(mapList)
	m := &Model{items: make([]Item, cnt)}
	for i := 0; i < cnt; i++ {
		name := stageList[i]
		value := stageList[i]
		m.items[i] = Item{name, value}
	}
	return m
}

//絞り込みステージの表示
func restrictStageModel(mp string) *Model {
	file1, err := os.Open("map_name.csv")
	failOnError(err)
	defer file1.Close()
	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))

	mapList, err := reader.Read() // 1行読み出す
	failOnError(err)
	stageList, err := reader.Read() // 1行読み出す
	failOnError(err)
	cnt := len(mapList)
	//まずは個数を確かめる
	num := 0
	for i := 0; i < cnt; i++ {
		if mp == mapList[i] {
			num++
		}
	}

	m := &Model{items: make([]Item, num)}
	index := 0
	for i := 0; i < cnt; i++ {
		if mp == mapList[i] {
			name := stageList[i]
			value := stageList[i]
			m.items[index] = Item{name, value}
			index++
		}
	}

	return m
}

func NewKillerModel() *Model {
	//キラー名のcsv読み込み
	file1, err := os.Open("killer_name.csv")
	failOnError(err)
	defer file1.Close()
	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))
	record, err := reader.Read() // 1行読み出す
	failOnError(err)

	m := &Model{items: make([]Item, len(record))}
	for i, v := range record {
		name := v
		value := v
		m.items[i] = Item{name, value}
	}

	return m
}

func (m *Model) ItemCount() int {
	return len(m.items)
}

func (m *Model) Value(index int) interface{} {
	return m.items[index].name
}

type Item struct {
	name  string
	value string
}

func main() {
	lw := &MyLoginWindow{}
	LW := MainWindow{
		AssignTo: &lw.MainWindow, // Widgetを実体に割り当て
		Title:    "Login",
		Size:     Size{300, 50}, //width×height
		Layout:   VBox{},        // ウィジェットを垂直方向に並べる
		Children: []Widget{
			Label{
				Text:     "password",
				AssignTo: &lw.label,
			},
			TextEdit{
				AssignTo: &lw.textArea,
				MaxSize:  Size{15, 20},
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
	if true {

	}
	mw := &MyMenuWindow{}
	MW := MainWindow{
		AssignTo: &mw.MainWindow, // Widgetを実体に割り当て
		Title:    "StatsWindow",
		Size:     Size{400, 150},
		Layout:   VBox{}, // ウィジェットを垂直方向に並べる
		Children: []Widget{
			PushButton{
				Text:      "データ記録",
				MaxSize:   Size{80, 30},
				MinSize:   Size{80, 30},
				OnClicked: mw.inputMenuClicked,
			},
			PushButton{
				Text:      "統計データ",
				MaxSize:   Size{80, 30},
				MinSize:   Size{80, 30},
				OnClicked: mw.statsMenuClicked,
			},
			PushButton{
				Text:      "設定",
				MaxSize:   Size{80, 30},
				MinSize:   Size{80, 30},
				OnClicked: mw.optionMenuClicked,
			},
		},
	}
	//ログインウィンドウを非表示に
	lw.SetVisible(false)
	if _, err := MW.Run(); err != nil {
		log.Fatal(err)
	}
}

//データ入力ウィンドウ表示
func (mw *MyMenuWindow) inputMenuClicked() {
	iw := &MyInputWindow{killerModel: NewKillerModel(), mapModel: NewMapModel(), stageModel: NewStageModel()}
	IW := MainWindow{
		AssignTo: &iw.MainWindow, // Widgetを実体に割り当て
		Title:    "InputWindow",
		Size:     Size{400, 150},
		Layout:   VBox{}, // ウィジェットを垂直方向に並べる
		Children: []Widget{
			Label{
				Text:     "キラー",
				AssignTo: &iw.killerLabel,
			},
			ComboBox{
				Model:    iw.killerModel,
				AssignTo: &iw.killerComboBox,
				MaxSize:  Size{50, 50},
			},
			Label{
				Text:     "サク数",
				AssignTo: &iw.killLabel,
			},
			RadioButtonGroupBox{
				DataMember: "Bar",
				Layout:     HBox{},
				Buttons: []RadioButton{
					RadioButton{
						Name:  "0kill",
						Text:  "0",
						Value: "0",
					},
					RadioButton{
						Name:  "1kill",
						Text:  "1",
						Value: "0",
					},
					RadioButton{
						Name:  "2kill",
						Text:  "2",
						Value: "2",
					},
					RadioButton{
						Name:  "3kill",
						Text:  "3",
						Value: "3",
					},
					RadioButton{
						Name:  "4kill",
						Text:  "4",
						Value: "4",
					},
				},
			},
			Label{
				Text:     "マップ",
				AssignTo: &iw.mapLabel,
			},
			ComboBox{
				Model:                 iw.mapModel,
				AssignTo:              &iw.mapComboBox,
				MaxSize:               Size{50, 50},
				OnCurrentIndexChanged: iw.mapChanged,
			},
			Label{
				Text:     "ステージ",
				AssignTo: &iw.stageLabel,
			},
			ComboBox{
				Model:    iw.stageModel,
				AssignTo: &iw.stageComboBox,
				MaxSize:  Size{50, 50},
			},
			PushButton{
				Text:      "登録",
				OnClicked: mw.registerClicked,
			},
		},
	}

	if _, err := IW.Run(); err != nil {
		log.Fatal(err)
	}
}

//マップ情報を入力すると
//ステージが対応するものに絞り込み
func (iw *MyInputWindow) mapChanged() {
	mp := iw.mapModel.items[iw.mapComboBox.CurrentIndex()].value
	iw.stageComboBox.SetModel(restrictStageModel(mp))
}

//統計データウィンドウ表示
func (mw *MyMenuWindow) statsMenuClicked() {
}

//設定ウィンドウ表示
func (mw *MyMenuWindow) optionMenuClicked() {
}

//試合情報の登録
func (mw *MyMenuWindow) registerClicked() {

}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}
