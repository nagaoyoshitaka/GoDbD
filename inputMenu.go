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

type MyInputWindow struct {
	*walk.MainWindow
	killerModel      *Model
	mapModel         *Model
	stageModel       *Model
	killerLabel      *walk.Label
	killLabel        *walk.Label
	mapLabel         *walk.Label
	stageLabel       *walk.Label
	killerComboBox   *walk.ComboBox
	mapComboBox      *walk.ComboBox
	stageComboBox    *walk.ComboBox
	killRadioButton0 *walk.RadioButton
	killRadioButton1 *walk.RadioButton
	killRadioButton2 *walk.RadioButton
	killRadioButton3 *walk.RadioButton
	killRadioButton4 *walk.RadioButton
	kill             string
}

func showInputMenu() {
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
						AssignTo:  &iw.killRadioButton0,
						Name:      "0kill",
						Text:      "0",
						Value:     "0",
						OnClicked: iw.kill0Clicked,
					},
					RadioButton{
						AssignTo:  &iw.killRadioButton1,
						Name:      "1kill",
						Text:      "1",
						Value:     "1",
						OnClicked: iw.kill1Clicked,
					},
					RadioButton{
						AssignTo:  &iw.killRadioButton2,
						Name:      "2kill",
						Text:      "2",
						Value:     "2",
						OnClicked: iw.kill2Clicked,
					},
					RadioButton{
						AssignTo:  &iw.killRadioButton3,
						Name:      "3kill",
						Text:      "3",
						Value:     "3",
						OnClicked: iw.kill3Clicked,
					},
					RadioButton{
						AssignTo:  &iw.killRadioButton4,
						Name:      "4kill",
						Text:      "4",
						Value:     "4",
						OnClicked: iw.kill4Clicked,
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
				OnClicked: iw.registerClicked,
			},
		},
	}
	iw.kill = "null"
	if _, err := IW.Run(); err != nil {
		log.Fatal(err)
	}
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

	mapAndStageList, err := reader.ReadAll()
	failOnError(err)
	l := len(mapAndStageList)
	mapList := make([]string, l)
	for i := 0; i < l; i++ {
		mapList[i] = mapAndStageList[i][0]
	}
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
	mapAndStageList, err := reader.ReadAll()
	failOnError(err)
	cnt := len(mapAndStageList)
	m := &Model{items: make([]Item, cnt)}
	for i := 0; i < cnt; i++ {
		name := mapAndStageList[i][1]
		value := mapAndStageList[i][1]
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

	mapAndStageList, err := reader.ReadAll()
	failOnError(err)
	cnt := len(mapAndStageList)
	//まずは個数を確かめる
	num := 0
	for i := 0; i < cnt; i++ {
		if mp == mapAndStageList[i][0] {
			num++
		}
	}

	m := &Model{items: make([]Item, num)}
	index := 0
	for i := 0; i < cnt; i++ {
		if mp == mapAndStageList[i][0] {
			name := mapAndStageList[i][1]
			value := mapAndStageList[i][1]
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
	record, err := reader.ReadAll()
	failOnError(err)

	m := &Model{items: make([]Item, len(record))}
	for i, v := range record {
		name := v[0]
		value := v[0]
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

//マップ情報を入力すると
//ステージが対応するものに絞り込み
func (iw *MyInputWindow) mapChanged() {
	mp := iw.mapModel.items[iw.mapComboBox.CurrentIndex()].value
	iw.stageComboBox.SetModel(restrictStageModel(mp))
}

//試合情報の登録
func (iw *MyInputWindow) registerClicked() {
	killer := iw.killerModel.items[iw.killerComboBox.CurrentIndex()].value
	kill := iw.getKill()
	mp := iw.mapModel.items[iw.mapComboBox.CurrentIndex()].value
	stage := iw.stageModel.items[iw.stageComboBox.CurrentIndex()].value
	showInformDialog(killer, kill, mp, stage)
}

func (iw *MyInputWindow) kill0Clicked() {
	iw.kill = "0"
}
func (iw *MyInputWindow) kill1Clicked() {
	iw.kill = "1"
}
func (iw *MyInputWindow) kill2Clicked() {
	iw.kill = "2"
}
func (iw *MyInputWindow) kill3Clicked() {
	iw.kill = "3"
}
func (iw *MyInputWindow) kill4Clicked() {
	iw.kill = "4"
}

func (iw *MyInputWindow) getKill() string {
	return iw.kill
}
