package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot/vg/draw"

	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

type MyStatsWindow struct {
	*walk.MainWindow
	playerLabel         *walk.Label
	spanLabel           *walk.Label
	graphTypeLabel      *walk.Label
	killerRadioButton   *walk.RadioButton
	survivorRadioButton *walk.RadioButton
	dayRadioButton      *walk.RadioButton
	monthRadioButton    *walk.RadioButton
	yearRadioButton     *walk.RadioButton
	killRateRadioButton *walk.RadioButton
	playRateRadioButton *walk.RadioButton
	imageView           *walk.ImageView
	period              string
}

type Mode struct {
	Name  string
	Value ImageViewMode
}

func showStatsMenu() {
	sw := &MyStatsWindow{}
	SW := MainWindow{
		AssignTo: &sw.MainWindow, // Widgetを実体に割り当て
		Title:    "StatsWindow",
		Size:     Size{Width: 400, Height: 150},
		Layout:   VBox{}, // ウィジェットを垂直方向に並べる
		Children: []Widget{
			Label{
				Text:     "player",
				AssignTo: &sw.playerLabel,
			},
			RadioButtonGroupBox{
				DataMember: "Bar",
				Layout:     HBox{},
				Buttons: []RadioButton{
					RadioButton{
						AssignTo:  &sw.killerRadioButton,
						Name:      "killer",
						Text:      "killer",
						Value:     "killer",
						OnClicked: sw.killerClicked,
					},
					RadioButton{
						AssignTo:  &sw.survivorRadioButton,
						Name:      "survivor",
						Text:      "survivor",
						Value:     "survivor",
						OnClicked: sw.survivorClicked,
					},
				},
			},
			Label{
				Text:     "span",
				AssignTo: &sw.spanLabel,
			},
			RadioButtonGroupBox{
				DataMember: "Bar",
				Layout:     HBox{},
				Buttons: []RadioButton{
					RadioButton{
						AssignTo:  &sw.dayRadioButton,
						Name:      "day",
						Text:      "day",
						Value:     "day",
						OnClicked: sw.dayClicked,
					},
					RadioButton{
						AssignTo:  &sw.monthRadioButton,
						Name:      "month",
						Text:      "month",
						Value:     "month",
						OnClicked: sw.monthClicked,
					},
					RadioButton{
						AssignTo:  &sw.yearRadioButton,
						Name:      "year",
						Text:      "year",
						Value:     "year",
						OnClicked: sw.yearClicked,
					},
				},
			},
			Label{
				Text:     "graphType",
				AssignTo: &sw.graphTypeLabel,
			},
			RadioButtonGroupBox{
				DataMember: "Bar",
				Layout:     HBox{},
				Buttons: []RadioButton{
					RadioButton{
						AssignTo:  &sw.killRateRadioButton,
						Name:      "kill rate",
						Text:      "kill rate",
						Value:     "kill rate",
						OnClicked: sw.killRateClicked,
					},
					RadioButton{
						AssignTo:  &sw.playRateRadioButton,
						Name:      "play rate",
						Text:      "play rate",
						Value:     "play rate",
						OnClicked: sw.playRateClicked,
					},
				},
			},
			ImageView{
				Background: SolidColorBrush{Color: walk.RGB(255, 191, 0)},
				Image:      "sample1.png",
				Margin:     1,
				Mode:       ImageViewModeIdeal,
				AssignTo:   &sw.imageView,
			},
		},
	}
	if _, err := SW.Run(); err != nil {
		log.Fatal(err)
	}
}

func (sw *MyStatsWindow) killerClicked() {
}

func (sw *MyStatsWindow) survivorClicked() {
}

func (sw *MyStatsWindow) dayClicked() {
	sw.period = "day"
}

func (sw *MyStatsWindow) monthClicked() {
	sw.period = "month"
}

func (sw *MyStatsWindow) yearClicked() {
	sw.period = "year"
}

func (sw *MyStatsWindow) killRateClicked() {
	//キルレのデータを取得
	X, Y := sw.makeKillRate()
	//キルレのグラフを保存
	sw.showKillRate(X, Y)
	//キルレのグラフを表示
	img, _ := walk.NewImageFromFile("sample2.png")
	img_walk, _ := walk.ImageFrom(img)
	sw.imageView.SetImage(img_walk)
}

func (sw *MyStatsWindow) showKillRate(labelX []string, dataY []float64) {
	//new plot
	p := plot.New()
	p.Title.Text = "only english title"
	//X
	p.NominalX(labelX...)
	p.X.Tick.Label.Rotation = math.Pi / 2.5
	p.X.Tick.Label.XAlign = draw.XRight
	p.X.Tick.Label.YAlign = draw.YCenter
	//Y
	labelY := plotter.Values{}
	for _, y := range dataY {
		labelY = append(dataY, y)
	}
	p.Y.Min = 0.0
	p.Y.Max = 4.0
	//new bar chart
	breadth := vg.Points(15)
	bar, err := plotter.NewBarChart(labelY, breadth)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(2)
	bar.Offset = 0
	bar.Horizontal = false
	p.Add(bar)
	p.Title.Text = "kill rate"

	if err := p.Save(10*vg.Inch, 5*vg.Inch, "sample2.png"); err != nil {
		panic(err)
	}
}

func (sw *MyStatsWindow) makeKillRate() ([]string, []float64) {
	p := sw.period
	if p == "day" {
		Xarray, Yarray := makeDayArray()
		return Xarray, Yarray
	} else if p == "month" {
		Xarray, Yarray := makeMonthArray()
		return Xarray, Yarray
	} else if p == "year" {
		Xarray, Yarray := makeYearArray()
		return Xarray, Yarray
	}
	return make([]string, 0), make([]float64, 0)
}

func makeDayArray() ([]string, []float64) {
	num := 30
	days := make([]string, num)
	const layout = "2006-01-02"
	for i := 0; i < num; i++ {
		date := time.Now().AddDate(0, 0, -i).Format(layout)
		days[i] = date
	}
	kills := calcMeanArray(days, num)
	return days, kills
}

func makeMonthArray() ([]string, []float64) {
	num := 12
	months := make([]string, num)
	const layout = "2006-01"
	for i := 0; i < num; i++ {
		month := time.Now().AddDate(0, -i, 0).Format(layout)
		months[i] = month
	}
	kills := calcMeanArray(months, num)

	return months, kills
}

func makeYearArray() ([]string, []float64) {
	num := 5
	years := make([]string, num)
	const layout = "2006"
	for i := num - 1; i >= 0; i-- {
		year := time.Now().AddDate(-i, 0, 0).Format(layout)
		years[i] = year
	}
	kills := calcMeanArray(years, num)
	return years, kills
}

func calcMeanArray(periodSet []string, num int) []float64 {
	kills := make([]float64, num)
	//統計データのcsv読み込み
	file1, err := os.Open("matchLog.csv")
	failOnError(err)
	defer file1.Close()
	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))
	record, err := reader.ReadAll()
	failOnError(err)
	for i, period := range periodSet {
		c := calcMean(period, record)
		kills[i] = c
	}
	return kills
}

func calcMean(period string, record [][]string) float64 {
	sum := 0
	cnt := 0
	for _, rec := range record {
		if strings.Contains(rec[4], period) {
			s, _ := strconv.Atoi(rec[1])
			sum += s
			cnt += 1
		}
	}
	if cnt == 0 {
		return 0.0
	}
	return float64(sum) / float64(cnt)
}

func (sw *MyStatsWindow) playRateClicked() {
}
