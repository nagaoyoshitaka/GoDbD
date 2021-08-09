package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type MyDatabaseWindow struct {
	*walk.MainWindow
	tv      *walk.TableView
	dbModel *databaseModel
}

type databaseModel struct {
	walk.TableModelBase
	walk.SorterBase
	items []*databaseItem
}

type databaseItem struct {
	killer string
	kill   string
	mp     string
	stage  string
	date   string
}

func (m *databaseModel) RowCount() int {
	return len(m.items)
}

func (m *databaseModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.killer

	case 1:
		return item.kill

	case 2:
		return item.mp

	case 3:
		return item.stage

	case 4:
		return item.date
	}

	panic("unexpected col")
}

func showDatabaseMenu() {
	dw := &MyDatabaseWindow{dbModel: NewDatabaseModel()}
	DW := MainWindow{
		AssignTo: &dw.MainWindow, // Widgetを実体に割り当て
		Title:    "DatabaseWindow",
		Size:     Size{Width: 1000, Height: 500},
		Layout:   VBox{}, // ウィジェットを垂直方向に並べる
		Children: []Widget{
			TableView{
				AssignTo: &dw.tv,
				Columns: []TableViewColumn{
					{Name: "killer", Title: "killer"},
					{Name: "kill", Title: "kill"},
					{Name: "mp", Title: "mp"},
					{Name: "stage", Title: "stage"},
					{Name: "date", Title: "date"},
				},
				Model: dw.dbModel,
			},
		},
	}
	if _, err := DW.Run(); err != nil {
		log.Fatal(err)
	}
}

func NewDatabaseModel() *databaseModel {
	file1, err := os.Open("matchLog.csv")
	failOnError(err)
	defer file1.Close()
	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))

	matchLogList, err := reader.ReadAll()
	failOnError(err)
	l := len(matchLogList)
	m := &databaseModel{items: make([]*databaseItem, l)}
	for i := 0; i < l; i++ {
		killer := matchLogList[i][0]
		kill := matchLogList[i][1]
		mp := matchLogList[i][2]
		stage := matchLogList[i][3]
		date := matchLogList[i][4]
		m.items[i] = &databaseItem{killer: killer, kill: kill, mp: mp, stage: stage, date: date}
	}
	return m
}
