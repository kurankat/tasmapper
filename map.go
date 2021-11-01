package tasmapper

import (
	"strconv"
	"strings"
	"time"
)

func drawGridMap() {
	canvas.Gid("gridAndNumbers")
	makeGrid()
	setNumbers()
	canvas.Gend()
	addCompactMap()
}

func drawPlainMap() {
	addCompactMap()
}

func makeGrid() {
	// Draw King Island grid
	for i, j := leftMargin, 29; i < leftMargin+5*cellSide; i++ {
		if (i-leftMargin)%cellSide == 0 {
			if j < 34 {
				canvas.Line(i, topMargin, i, topMargin+(8*cellSide), "stroke:black")
			}
			j++
		}
	}
	for i, j := topMargin, 62; i < topMargin+(9*cellSide); i++ {
		if (i-topMargin)%cellSide == 0 {
			if j > 53 {
				canvas.Line(leftMargin, i, leftMargin+(4*cellSide), i, "stroke:black")
			}
			j--
		}
	}

	// Draw rest of Tasmania grid
	for i, j := leftMargin, 29; i < mapWidth+leftMargin+1; i++ {
		if (i-leftMargin)%cellSide == 0 {
			if j < 34 {
				canvas.Line(i, topMargin+(9*cellSide), i, mapHeight+topMargin, "stroke:black")
			} else {
				canvas.Line(i, topMargin, i, mapHeight+topMargin, "stroke:black")
			}

			j++
		}
	}
	for i, j := topMargin, 62; i < mapHeight+topMargin+1; i++ {
		if (i-topMargin)%cellSide == 0 {
			if j > 53 {
				canvas.Line(leftMargin+(5*cellSide), i, mapWidth+leftMargin, i, "stroke:black")
			} else {
				canvas.Line(leftMargin, i, mapWidth+leftMargin, i, "stroke:black")
			}
			j--
		}
	}
}

func setNumbers() {
	// Write King Island numbers
	for i, j := leftMargin, 22; j < 26; i++ {
		if (i-leftMargin)%cellSide == 0 {
			if j > 22 {
				canvas.Text(i, topMargin-10, strconv.Itoa(j), "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:18px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
			}
			j++
		}
	}

	// Write rest of Tasmania numbers
	// Horizontal axes
	for i, j := leftMargin, 29; i < mapWidth+leftMargin; i++ {
		if (i-leftMargin)%cellSide == 0 {
			if j > 34 {
				canvas.Text(i, topMargin-10, strconv.Itoa(j), "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:18px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
			}
			if j > 29 {
				canvas.Text(i, canvasHeight-bottomMargin+22, strconv.Itoa(j), "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:18px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
			}
			j++
		}
	}

	// Vertical axes
	for i, j := topMargin, 62; j > 15; i++ {
		if (i-topMargin)%cellSide == 0 {
			if j < 62 {
				canvas.Text(leftMargin-15, i+5, strconv.Itoa(j), "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:18px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
				canvas.Text(canvasWidth-rightMargin+15, i+5, strconv.Itoa(j), "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:18px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
			}
			j--
		}
	}
	canvas.Text(canvasWidth-55, canvasHeight-14, "Grid: MGA94", "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:18px;line-height:125%;font-family:Arial;-inkscape-font-specification:Arial;text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
}

func gridBox(r *RecordList) {
	recString := "Total records: " + strconv.Itoa(r.numRecs)
	canvas.Gid("infoBox")
	canvas.Rect(180, 80, 500, 250, "opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000")
	if len(r.name) < 28 {
		canvas.Text(430, 175, r.name, "font-style:italic;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:30px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	} else {
		slice := strings.Split(r.name, " ")
		top := ""
		bottom := ""
		for i, subst := range slice {
			if i < 2 {
				top = top + subst + " "
			} else {
				bottom = bottom + subst + " "
			}
		}
		canvas.Text(430, 140, top, "font-style:italic;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:30px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
		canvas.Text(430, 175, bottom, "font-style:italic;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:30px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	}
	canvas.Text(430, 220, strconv.Itoa(r.numCells)+" cells", "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:24px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	canvas.Text(210, 315, recString, "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:24px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:left;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:left;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	canvas.Text(665, 315, time.Now().Format("02.01.2006"), "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:24px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:end;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:end;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	canvas.Gend()
}

func plainBox(r *RecordList) {
	recString := "Total records: " + strconv.Itoa(r.numRecs)
	canvas.Gid("infoBox")
	canvas.Rect(180, 80, 500, 250, "opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000")
	if len(r.name) < 28 {
		canvas.Text(430, 175, r.name, "font-style:italic;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:30px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	} else {
		slice := strings.Split(r.name, " ")
		top := ""
		bottom := ""
		for i, subst := range slice {
			if i < 2 {
				top = top + subst + " "
			} else {
				bottom = bottom + subst + " "
			}
		}
		canvas.Text(430, 160, top, "font-style:italic;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:30px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
		canvas.Text(430, 195, bottom, "font-style:italic;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:30px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:center;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:middle;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	}
	canvas.Text(210, 315, recString, "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:24px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:left;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:left;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	canvas.Text(665, 315, time.Now().Format("02.01.2006"), "font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:24px;line-height:125%;font-family:'Arial';-inkscape-font-specification:'Times New Roman, ';text-align:end;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;text-anchor:end;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	canvas.Gend()
}

func webBox(r *RecordList) {
	recString := "Total records: " + strconv.Itoa(r.numRecs)
	canvas.Gstyle("font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;line-height:125%;font-family:Arial;-inkscape-font-specification:Arial;letter-spacing:0px;word-spacing:0px;writing-mode:lr-tb;fill:#000000;fill-opacity:1;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1")
	canvas.Text(420, 60, "Distribution map from records held in the", "font-size:32px;text-align:center;text-anchor:middle")
	canvas.Text(417, 100, "Tasmanian Herbarium collection (HO)", "font-size:32px;text-align:center;text-anchor:middle")
	if len(r.name) < 28 {
		canvas.Text(420, 260, r.name, "font-style:italic;font-size:37px;text-align:center;text-anchor:middle")
	} else {
		slice := strings.Split(r.name, " ")
		var top, bottom string
		for i, subst := range slice {
			if i < 2 {
				top = top + subst + " "
			} else {
				bottom = bottom + subst + " "
			}
		}
		canvas.Text(430, 220, top, "font-style:italic;font-size:37px;text-align:center;text-anchor:middle")
		canvas.Text(430, 265, bottom, "font-style:italic;font-size:37px;text-align:center;text-anchor:middle")
	}
	canvas.Text(40, 1220, recString, "font-size:30px; text-align:left;text-anchor:left")
	canvas.Text(870, 1220, time.Now().Format("02.01.2006"), "font-size:30px;text-align:end;text-anchor:end")
	canvas.Gend()
}

func addCompactMap() {
	canvas.Path("M519 1162l3-3h1l1 1v-1h2l3 1v-4h2l1-2h1l4-3 2-1v-1l-1-1v-2l2-2h-1v-2l-1-1-1 1h-1v1l-1 1-1 1h-1l-1-1 1-1-1-3 1-1 1 1h1v-1l-2-1v-1h-1v-1l2-3-1-1 2-4 2 3-1 1-1 1h1l1 1h2v-2l1-1 1 1 4-1h1v-5c1-3 5-4 5-7h-1l-4 6h-2l-1-1 1-1-2-1-1-1 1-2v-1l1-1h1l1-1h2v-2h4l1 2-1 1 1 2 2-2h2l-1-1v-2l-1-1 1-1h-2l-1-2h-3l-2-2-3 1-1-1v-1h1v-1h-3l-2 1h-1v-2l-2-2h2l1 2h2l-2-6 3 4 2 1v2l1 1h2l1-2 2-2 2 1v2h1l1-2h1l1 1 1-1 2 3 1-2h1l-1-3-1-1 2-3-2-1v-2l3-2v-1l1-1c3 0 4-6 4-8h1v-1c0-2-4-3-5-3l-3-3h-1l-1 2-1 1v-1h-2l-2-2h-1v1l-1-2h-1l2-1c1 1 2 3 4 2l1-2 2-3 1 1 1-1 4-1c2 1 3 4 4 6h1l1-2-1-1v-2l1-1 2-6 3 1v-3l-1-1-1-1-5-5v-3h-2l-2-1h-1c0-2-2-2-3-3l-1-2-4-3h-2l1-2-2-1-2 1-1-1v-1l-1-2h-1l2-2v-1l-1-1 1-2-3-3 2-1 2 1v-4l-1-1v-1h1l1-2h2l2-6v-1l3-6 3-7 1-1 2-1-5 12v3l-3 6c1 1-1 4-2 5 1 2 0 7-1 8l-1 1v3l3 1h2l2 2h1l1 3 3 1v1l2 1 2 1v-2l4-3h-2v-1l3-3-1-2v-5l1 1h2l-1 2 1 2h1l-1 1 1 1v1l1 1v2h1l1 1v1l-4 1-1 1 1 4v2h-1v1h1l1 2 1-1 2-1h1v-1h1l2 2h2v6l2 1h1l2 2c2 1 2-3 7-1l3 1c3 0 4-3 5-3l1-1-1-3 2-4 2-2 1-2-1-8-1-2v-1l-2-2-1-1v-1l2-2-2-1v-2l2-1-1-1 1-1v-1l2-1-2-1h3l-1-1-2-1v-1h4l1-1v-1l-2-1v-1h1l1-1 3 2c0-2 2-4 3-5v-2l-4-2-5-2 1-1 1-4v-2c2-1 2-4 2-5v-1h-1c0-2 1-4 3-4 1 4 3 10 6 14h1l5-2v-2l-1-1-2-7v-1l-1-2c1-1 2-3 1-5l-1-1 1-1 2-1 2-6 2-1 1-7v-3h-1v-1h-2l-3-3v-1c2-1 2-2 2-4l-1-2v-3h-1l-3-2 1-1h1v-2l-2 1 1-3-1-1-3-1v1l-1-1v-1l1-1-1-1h-3l-1 2-3-2-2-2v-1l1-3v-1l2 1 2-2h-2l-1-1-2 1-1-1 1-2-1-3h-1l-3-4-6-3c-1-1-8 2-9 2l3-3h1l1-2 4-1 2 2 2 1 2 1c2 0 4 2 4 3l1-1 2 2v6l3 1 1-1 2 3-1 1-2 1 1 2 2 2h2l2 2h2l-1 1c0 2 3 4 3 5h2l-1 2 2 1 1 1 1-1-1 3v1h1l-1 2 1 1 2-1-1 1 1 2 1-1h5l2 1v1l3 5v2l-2 1v6h2l3-4v-4l1-1 4 1 5-1 1 2v3h1v1h-1l-1-3h-1l-2 2-1 2-1 2v2l-1 1v3l-1 1v3l1-1h1l2-1 3 4-1 1v1l-1 3v6c-2 3-5 3-8 1-2-1-1-2-1-4l-2-3v-4l-2-2-1-1 1-1v-2h-1l-1 2-1 1h-1l2 2 1 2-1 1v3l2 2c1 0 2 3 1 4l-2 2v1l1 1v1l1 1h1l1-2 7-2h3l5-3-1-1v-1h1v-1l1 1v2l3-1 1 1 1-1-1-1 2-3-1-3 5-2h1l-1-2 1-3-2-2-1 1h-1l-1-1v2l2 1v1l-2 3-1-1-2-3-1-2 2-1-1-2h2l3 1h-1l1-4-3-5h-2l-1-2h-2l-1-1v-2l3-8-1-3 5-5 6-2h7l4 2v-1l-1-1-2-2-3-2h-6l-6 1c-2-1-2-3-2-4h-2l-1 2-2 2-2-1-2-2 4-1-1-1h-5l-3-5-2-2v-1h2v-1l2-1-2-5h1v2l1 1 2 1v1l-2 2 2 3h2l3 3 1-2 4-3h1l2 5 2 1v-3l-1-1v-1l-1-1 1-2h1v2h3v3h2v-1h1v1l4 1 1-1 1-3 1 2v3l-1 1 2 4v2l3 3v3l-1 2h-1v3c1-1 6 1 8 2l3-1 1-1h1l3-1h1l1 2 1-1v1l-3-2c-1 2-3 3-5 3l-1 1h-3l-1 1 1 2 1 1 1-2 2-1 1 1v1l1 3 3-1 1-1 1-1 2 1 1-1 2-1h1l1 1v2h1v-1l4 1v1h2l1 1h1v-2h1l3 1v-3l2 1 2 1h1l1 2 1 3-1 3h-1v3l1-2h1l1 1 1 1v2l4-2 1-1 2 3h-3l-1 1v1l1 1-3 1-1 1h-3l-1-1-1 1 3 3 1-1h1l1 1v1l4 1-4 3 1 1-1 2 2 1v1h4l1 1 5 1h5l-9 2-1-1v1l-1 1-1 1v6l-1-2v-2l-3-2h-1v1l-1 1 2 3-1 2h-2l-1-1-2 1-2-3h-2v1h-1l-1-1-1 1h-1v1h-1v-2l-1-1-1-3v1l-1 2h-1l-1-1v-1l-1-1c1-1 0-4-1-5l-2 3-1-2-1-2h-2v-4h1l1-3h1l2-5h-1v-3l-1-2h-1l-1 2h-1l-2-1v-2l-4-3-3 1v2l2 3v1h-2l4 4c0 3-2 5-4 7l-3-2h-2c-2 0-4 2-5 3v3l1 2-1 1 1 1v10l1 2h2l2 2v2h5l1 3 1-1v2l4-1h1l1 1 1-1h1l3-2v1h1l1 4h-1v-1l-2-2-3 2h1l3 2v2l-1 2-2-2h-3l-4 5 1 1h1v3l1 2-1 2 2 3 1-1 1 1-1 4 2 1 1-1 2 1h2v2l2 1 1-1h1l4 4-1 3 1 2h2l1 2v-7l1-1 1-1 1-1h1l1-1 1-2 1 1 1-2h1v-1h2l2 3 1 2h2v-1l-1-1v-2l1-1-1-2-1 1h-1l-1-2 1-2h2v-6l-1 4h-1l-1-1-2-2 3-3v-2h1l-2-3v-2l2-2v1l1 5 3-2h1l1 4-1 1 2 3v10h2c1 0 5 1 6 3l1-1 2 3-1 1 1 1c2 1 5 1 7 3l2-1 1 2 3-2-1-1c-2 0-4-2-4-3-2 0-5-3-5-4 0-5 2-4 3-8l1-3h1l3-3h1v-1h-2l-1 1-2-2h-2l-1 2h-1l-1-1-1-1 1-1-1-1 2-1v1l1-1 1-1 1 1v-5l-1-1-2-4v-1l-2-5-2-2 2-2-1-1 1-1-2-4 1 1-1 1h-2l-1-1v-6l1-1h2v-1l2-1 3-5h2l2-1v-3l1 1 1-1v-3h-1l-1-4c1-3 3-6 2-9h-1l-2-5-2-1 1-1h-4v-1l3-1v-2l-1-1-2 1v1l-2 1-1 1-2-1c-1 0-3-2-3-4l1-1-1-2v-1h-2l-1-1h-3v3l-2-1v1l1 2h-1l-2 1v2h-1l1 1 1-1 1 1 5 2-1 5-2-2h-1l-2-2-1 2-6-1h-1l-4 1-3-1 4-3 2-4v1l1 2h1v-1l-1-2v-1l1-1v-3h2v1l1-1-1-2 2-5 4 7-3-9v-4c1-2 3-11 5-12v-1l2-2h1l1-1h4l3-3v-4l1-1-1-1-1 1 1-1-1-1 2-3c1-1 3-2 3-4l-1-1 1-1-2-3 1-1-3-1-2-2 5 2v-1c-1-1-5-3-5-5l2-4h-1l2-2-2-2-1-2h-1c-3-1-4-5-4-7l-2-1v1l-4-1-1-2v-2h2v-1h6l1-1-1-3-3-2v-1l1-1 1-1 2-4 1 2-2 2v1l3 3v3l1 2v1l2-1 1 2 2-1v-4l1-1 2-2-1-1 2-2h2l1 3h1l1-1v-1l2-2-4-4c1-2 2-9 4-9l1-4-1-1 1-1-1-1h-2v-2l-1-4-2-1v-2c0-3 1-6 3-8h2v-1c0-3 2-5 4-7h-1v-2c-1-2-2-5-1-7v-1l-1 1-4-1h-1l1 2-1 1h-2v4l-1-1-1 2-1-1-1 1v4l-1-3h-5v-2h1l7-2h1l-1-1 2-5h4l4-7-1-2 2-4h1l-2-3-1-2 4-5v-2l3-3v-1c0-2 3-2 4-3 2-1 1-4 2-5l1-2h1v-1l-1-3 1-3h-1v-3l1-1v-1h-1v-1l-1-2 1-2c3-3 8-4 12-4 7 0 13 0 20 2l1-1-1-1c-2-2-7-3-10-3l-2 1h-1l-4-3-4 2-1-1v-1h1l-1-1h-1v-1l2-1h4l3-3-1-1 1-2-1-1-2-1v4l-1 1h-1l1-5v-1l1-1-1-3 2-1 2-4 1-1 1 1v3l3 1v2h2l1-1 3-3 1-1v-1c-1 0-2 2-3-1v-1h-1v-2c-2-2-1-4-2-6l3-2 1-1h1v1h-1v1l-1 2 1 2v1l2 2v2l3-2h1l1 2v3l-2 2-1 3-2 1-1 3-3 1-2-1-2 2h-2v4l-2 1-1 1v4h6l2-3 5 1h1v1h-1l-2-1v2l7 4-2 1 3 6 5-1 2 2v2h1l1-1h1l1 3c-1 2-4 3-6 3l-3 5 1 1 2 1 2 2v-1h1l2 1v2h-1v1c0 2-1 4-3 5l-1 2v2l-1 1-3 1-1 4 2 2 3 2 1-2h1v2l2 1v5l1 1 2-1 1-1-1-1 1-2h2c-1-2 1-4 1-5v-4l2-1h1l1-2v-3l2-4v-2-3h-1l-1 1-2-2-4-1h-1l-1 3-2-4 1-2c2 1 4-1 5-3l-1-4 1-1h2v-1l2-1v-1l-1-2 1-2-1-1-2-1-1-1c1-4-1-4-2-6v-4l-4-1-3-4h-2l2-1v-1-9l1-4 1-2 2-5 1 1 1-2-1-1v-6l2-1v-1h1l1-3-1-1v-2h-1l-2-3v-3h1v-2h1l-2-1h-1l-2-1h-1v-2c-3-1-4-5-4-7v-3l-1-3 1-4c0-3 3-10 5-13h3v-3h-2l-1-1h-1l-2-2v-2h1l1 3c-1-2 0-8 1-9l1-1-2-2 3-11 3-3v-3l1-1-1-14h-2l-1-3h-1l-1-1-2-2 1-2-1-4v-1l-1-5h-1l-1-1h-1v-3l-1-1h-1v-2h1v1l2 5h1l-1-11-3-1-2 2-2-1 2-1 2-1 3 1 2-6 1-6 2-4c1-3 0-5 2-8h-2l-1 2h-1v-2h1l1-1-1-1h1l2 1c3-4 4-11 5-15l5-7s2 0 3-2v-1l-1-1-3 2-2 1c-1 5-3 4-4 6v3c-2 2-3 1-5 2l-4 4h-1l-2-2-3-1-1-2h1l1 1c1 0 6-1 6-3l-1-2 1-4 1-1v-2h3v2l1 1-2 3h3l2-1v-3l2-2h1l2-5v-2l-1-2-1 1h-1l-1 1-1-2-1 1-2-1v2l-3 1 1-2h-1l2-1-1-3-1-2-1-5 1-1h-2l-1-6-3 1 5-4v-2h1v-2l-2-2-1-7c0-5 2-6 4-11 1-3 0-4 2-7v-2h1l-1-1-2-2-3-1v1l-2 1h-1v-1l1-1 1-1 5-2 1-1 1-1-1 1v6h1l-1-4 3-4 1-1v-2c2-2 4-5 7-5l-1-1h-2v1h-1l-3-4v-1l-2-2c-1-2-2-5-1-7l-3-2v-1h-1c-1-2-4-5-3-9h-1l-1-2-2-1h-2c-2-1-4-5-4-7l1-4h-1c-3 0-5-2-8-5v-1l-1-1-1 2-1 1h-5l-2-2c-3-2-6-6-7-9v-1l-2-2-3-4-1 1h-4l-3-1-2-3v-2l-1 1h-2v-1l-2-2-1 1-5-1-1-1v1l-4 1-3-2 1 4h1l2 1v1h-1l1 2-2 2h-1v4c-1 5-6 12-9 17l-2 2-2 3v1l1 1v1h-3l1-2v-1l-5 3c-5 2-5 2-10 1l-5-1-4-4-2-2-3-1-2 1-4 1c-3-2-4-6-6-8l-3-2v-2l-1 1-3 3h-1v1l-5 2-5 1-1 4h-2v1h1v1l-2 4-3 5-2 3a89 89 0 0 1-15 19l-7 5h-1c-2 0-7 2-8 1l-1 1 4 1h-4c-1-2-3-6-2-8-2-2-5-3-7-2v-1l-1-4v-1l-2-5-3 3-2 1c-2 1-4 2-6 1l-2-2v-2h-1l-1 2h-1l-4 1-1 3c-2 4-5 8-9 10h-1c-2 0-3 2-5 3l-3-1-1 5h-1v2l-2 1c2-2 2-5 4-6h-2l-1-1h-1l-1-1-5-3h-5l-1 1h-3l-1-2-3-2-4 1-2-2h-1l-1 1h-4c-2 6-6 7-9 11l-6 2-3-1h-3l-2-2-2-1-2 3c-1 1-5 4-7 3-2 2-3 6-6 7l-3-3v2l1 1h1v1h1v3c2 1 2 3 2 4l1 4h2l-1 3-1 1 1 2h2v2h1l1 1 3-1 4-2c2 0 4 2 5 3v-1 2h2c0 3 4 6 6 8l4 2 2 3h-1c-1-2-2-2-4-2h-2l-1 1-1 2 1 2h1l-2 2-1-1-1 2v1l2 1 3-1 1-2 1 6 2 1h4v2l4 3 4 7v1c-3 1-6 3-8 6l1 1 1 2c2 2 6 0 8 0l3 1 2 1-1 1 1 3h1l2-1h3l1 4-1 2c-1 4 4 6 4 9l6 6-3-2h-1l-3-1-3-4v-2l-1-6 1-3-3-2-4 3h-1l-1 1-2-3-3-3v-1l-1-1h-4l-3-4v-2l-1-1h2v-6c-2-1-2-4 0-5l-2-1-2 3-2-1v-1c-1-1-3-4-2-6l-7-7-1-2 2 1v-1l1-1 1-2h-2l1-1 1-1 2-1v-1l-1-2h-1l-3-3-5-2h-1l-3 6h-1l-2-1-2-2v1l1 3v1l2 2h-1l-2 2-1-2 1-3-1-1-2-1v-5l-1-2h-2l-1 1-3-1-4 5v-3h2l2-3h2l1-1h1l1 1 1-1c-2-2-2-2-2-5l-1-2v1l-3-1 2-2-2-3-3-3-3-1-2 1h-2v-1l-2-1-1-1-2-1v-1l-1-1-1 1 1 2c-1 3-5 7-8 8l-4-2h-3l1 4-2 2v2c-4 4-10 7-15 7l-1 3 1 1 2-1 1-1 2 1 3-1-2 3h-1l-1 1-2 4 1 1v2l2 2v2l1 1h1l1-1h1v1h-1v5l-1 1c1 1 0 3-1 4l1-4h-1v-5l-2-1c-1 1 0 3-3 5v3h-1v-1h-1l1-3h1l1-3h-1l-1-4v-3l1-1v-3l-1-1-3-1h-1 3l-1-3-3-1v-1l1-1-1-6-1-1-1-1-1 1v1c-1 2-6 6-9 7l-5 2-8 1c-1 1-5 3-6 2l-2-1c-1 0-4 1-4 3l1 4-1 2 2 2 3 5 1 1h-1l-2 1v-3l-1 1-1-2h-1v-2h-2l-1 1h-1l2-2v-2l2-1v-9h-1l-1-1h-1l-2 1-4-3-4 2h-5c-2-1-2-2-5-2l-1 6-1 1-1-1 1-1v-1-3c-2 1-4 1-6-1h-3v-1h-7v2l-5 2-1-1h-1l-1-1h-1v1h-1v-1h5l1 1 1-1h2v-3l-1-1-7-3v-2l-2-1h-3l-1-1-5-1-2-3-2-1-1 1-1-1v-1h-3v-1l-4-1-2-3-5-1-1-2h-1l-2-1h-1c-4 1-7 2-9-2v-1l-1-1h-1l-1-1-3-1-2-1h-7-1l-6-3-1-1-3-2c-2-1-3-4-4-6h-2l-4-2-3-1 1-1-1-4v-2l1-1v-3h-4c-2-2-9-2-12-1-2-2-5-3-7-3l-1-1h-3c-6-3-1-4-9-3l-3-3v-1h-1l-1-3h-1l-2-2v-2l-2-4 1-1-2-2-1 1v1l-5 3-5 1h-1l-1 2 1 1-1 2-1-1h1v-1-2-2l-5-1-1-1-2-1-1-1-1 1-2-1v-1h-1c-2 0-9 1-11-1l-1-1h-1l1 2v1l-1-1h-1l-1-2 1-1c0-2-3-4-3-5l-2-3-1-4c0-2 2-2 1-5l1-3h1v1l2-1-2-2h-1l-1-2 1-2-1-2h-2l-3-4-2-3-1 1-1 2 1 1-1 1 2 1v4c1 1 2 2 1 3l-2 1 1 6-2 1c-4 2-8 4-11 4l-12-1-6-2v2h1l4 1 1 1 4 2v1h1v1l-4-1h-3l-2 5-1-2h-1v-3l-2-2-1 1v1l1 1-1 1-2-1v-3l-2-1-1 2c-1-1-3-7-5-7v-1l-2-1h-1v1h-1l-1 2-1-2 1-2-1-1 1-1-3 1h-1l3-3-1-1-2 2-1 1v-3l-2-2-1-1-1-1-1-1-2 3h-4l-1 1h-1l-2 1-1 2-2-4-6-1-4-1-4-2-1 1-1 1v-3c1-1 0-4-1-5l-4 1-1 2-1-1 1-2-1-1-4-1h-1c-1 1 0 6-2 8h-1l1-3-1-2 1-4-2-2h-2l-1-1h-1l-2-4h1v1l2 2-2-6c-1-3-3-3-2-7l-1-1-2 4-1 1-2-1v2h-2l-2 3 1 1v2h2v3l-2 1v1l-1 1 1 1h-2l-1 2h1l1 1v1l1 1c-1 1-2 4-1 6h2v1l1-1h1c2 2 2 7-1 9l1 1 1 2-2 4v1h1v3c2 2 2 3 2 6v3l-1 2-1-1-1 1 2 1-3 8-3 3-6-2v1l1 1v2l-1 1-2 1v-1l-1 1h-1v1l-1 2h-1v1h-2v2l2 1 2 6-1 3 1 6-3 2h-1v1l2 3 1 1 1 2 2 3 2 1 3 2-1 1 1 2 2 5v9l-2 1 1 2v2l1 1h1c1 2 0 5-1 6l1 2v1l1 1h-1l1 2h1l1 5-1 5-1 3 1 1-1 1 1 2 2 3h1c2 2 1 5 2 6v2l2 1v4l1 1 3 2-1 2 1 1-1 1h1v1c6 3 6 9 8 10l1 10-4 3h-2l-2 2 2 2h1l1 1v1c4 0 3 3 5 4l2 6 1-1 1 3h1l1 5v1l2 1 1 2 10 19 1 2 1 3v1l2 2v3l1 2-1 4 1 1v1l1 1c-1 2 3 5 4 6v1l-3 2-1 2 1 3 1 1v1l3 4h1l2-2h1l5 4v2c2 2 3 2 4 5v1l3 3v3l1 3 1 1 1 2h2v1l-1 1 1 1 1-1 4 4 1 2h1c1 6 2 6 6 10h1l4 5 1 2 6 5 1 2 3 1 5 8 10 21 2-3-2 5c2 6 5 20 4 26s-1 13-4 18l-5 5 1 1h4l3 3h5l1-1-1-3c1-1 2-3 1-5v3l-2 1-1-1-1-1v-1l4-5h1v-1h3l1-1h1v-1l-1-3 1-1-1-1 2 1 2-1h1l-2 2 1 3h1v-1l5 2c2 3-2 7-3 9l1 1h1l1-2h2v1h-1c1 1 2 4 1 5l-2 2c2 3 1 2 1 5h2l3 3 3 5 3 3h1l2 7c1-1 6 2 7 2l1 7c1 0 3 0 4-2h1v-5l1 3 2-1 2 2v2l2 5-1 1h-4l2 2h4v-3l2-1 1-2 2 2v1l-2 2-1 1 1 1-1 2 1 1h-2v3c0 2-2 5-4 6l2 1-3 3h-1l-1-1h-2v-2h-1l-6 3-1 2h4-2l2 9h1v4c-1 1-2 2-1 4l1 1v6h-1l-1-2c-2-5-2-6-1-10l-1-1v-7c-1-2-3-4-3-7l1-7-1-1c-2-1-1-2-1-3l-2-2-1-1-1-1-2-2h-1l-2 1h-1c-1-2-4-3-5-6l-1 1-3-4h-1l-1-1v2l-1-1-1-1v-2c-2 0-3-2-4-3s0-6 1-8h-1l-2 1h-2l-1-3-2-1-2 1v2l-3-1-1-1 1-5-3-1-1-2v-1l-1-1 1-1-1-1-1 1-1-2 1-2-2-3v-2h-1l-1-1h-2v-3l-4-1-1 1h-1l1 2v1h3l1 1-1 3 1 1-1 4 4 4v3l1 4v2h-1l1 3-1 2-2 1v5l3 6 2 2v2h-1v1h2c1 1 0 3-1 5l3 2v7l2 4-1 1 1 1h1v4h-1l-1 3-1 1v1l2 3v1h1l1 1v1l1 3v4h-1l2 3 2 6h1v3l1 1 1 1h3v1l-1 1 1 3-1 2v2h-1l-1 1-2 3-3-1v3l1 1-1 1h1l1 1 1-2 1 2h1l-1-1 1-1v-1l4-1 3 3h2l1 2h-1v6l1-2h2l4 2v9l1 1 1 3 1 4 1 1v1l2-1v3h-1v2l-2 2 1 6h1l1 1h1v1l2-1 1 5 1 2v6l1 1v3c1 1 1 4 3 5v5l1 1 2 5v3l1 1v-2l3 1 1 2h1l1 4-1 1h2l1 2v9l1 1h1l-2 4 3-3h6v-1c2 0 5-2 7-1l1-1 1 1h2l2 5h1v2l3 5h1c3 1 3 4 4 6l1 3h1l2-1 3 2 1-1 1 5h-1l-1 1v1l-1 1 1 3-1 1h-1v1l1 2 2-1v1l3 3 3 1v-2h3l1 1 1 1-2 1v3l-1 1 2 1 1-1 2 4v1l3 3v2l1 2-1 2 1 1h1v-1h3l3 1-1 2 2 2 1 1-1 3v2h2v2h-1v2h-1l1 1v1h2v9h3v2l2 1h2l1 2-1 2-1-1v2h-1v1h1l1-2h1l1 3 3 1 1-3 3 1h1v-2l1-2h1l3 1 1-1-1-1 1-6v-2h-5l1 1-3 4h-1l-1 1-2-1-2-4h3v-1h-1l-1-1 1-1 4 1h1l2-4h4l1-3-2-6h-1l1-3-2-3 1-1 4-1v1h-3l-1 1 1 1 1 2-1 2h2l2 3 1 1 1 1h2l1-1h1l2 3-1 2-1 1 1 1v1l-1 3 1 1-1 2h-3v1l2 1v1l-1 1 1 1h1v2l1 1 1-1 1 1v4h5l-1 3-1-1v1h3v1l3 1 1 1 2-2 2 1 2-3v-1l-2-1 1-3 1-1-1-1 1-2h1l-1 2 2 2 1 5 3 3v-3c0 2 1 4 3 5h2l1 1h3l2-1 1-7 2-1 1-1-1-3c0-1 2-3 4-3v7l1 2h1v-2h3l3 4 1-2h3l-1 2 1 7-2 2h-1v1l1 1v1h1l1 2-2 1c-2-2-3-4-2-6h-5l-1-1h-1l-4 1 2 2-1 1h-1v-1h-1l1 2 2-1v1h2v1l-4 1v1h-1l-2-3-1-1-1 1 2 4v4l-1-5h-1c0-1-2-4-1-5l1 1v-1l-1-2 1-1-2-2h-2l-2-2h-1l-1 1-2-1v1l-1-1-2 1-1-2v1h-1c0 2 6 4 6 6l-1 1h-2v-1h-1v-1l1-1c-2-3-5-2-6-4h-9l-1-1 1-1h-2l-2-2h-1c0 1-2 4-1 5v1l2-1h1v4h1v4l-1 4h-1c-1-1 0-4 1-5l-2-5h-2l-1 1h1v2h-2v-1l-2-1-2 3h-3v2h-1l-1 1 3 2h1l2-2h3c1 1 3 3 3 5h-1v1h5v4l1 5 2 3v1l1 1-2 3h1l1-2 2 1 1 2v6l1 1v2l2 2-2 3 1 4-1 2v2l-1 2 1 1-1 3 3-5h5l1-1v-1l1-1h1c2 0 2 2 2 4l2-1 1 1h1v-1-8l2 1v-1h1l1-1h3l1 2v-1l-1-5v-1h3v3l1 1 2-2 3 6 1 1v1l1 1 3-3v-6l2-4-2-1v-1h1l2 1h3l1 1v-1l6 3v3l-1 1v3l2 1v1l1 1 1 2c3-1 6-1 7-4v-2h1l1-1 3-1 5 1 1 4 4-1 1 1 2 2 1-2 1 1 1-1h6l2 3 2-1 1-1h2l1 1 1-2h1l2-3h1l1 1v1h3l2-2 10 5v-1l-5-4v-1l-2-2-1-3 1-1v-1l-1-1v-2l1-4h2l3 5v9l-1 1 3 4v2l3 2-1 2 1 1 1-1v1l-1 2h1l2-2c2 0 3 3 4 4l2 1-1 2 1-2 2 1-1 4 2 3 2-4 3 1 1 2 2 5 2 2 1-1 2 1 4-4 1-1 2-1 2 2 1-2h2l1-1v-1h2v1l1-1 2 1 2 1 3 2c1 1 2 4 1 5l-2 2zM772 60l-1-2h-1l-2 2h-2l-1-2-1 1 2 1-1 2-3 1v2h1l3-1h1l5-4zm23 69h-2l1 1h2c2 0 2 2 3 3 2 3 5 4 7 8l1 3 1-1h5l9 5 4 1 2-2 3 1c0 2-3 7-2 8l-2 1v6c1 4 2 13 1 17l-1 1v1l-3 2-5 3-4 5 2 2 2-1 1 1 1-1 1-1 4-1 2-1v-9l2 9 3 11 2 7 3 9-4 3c-1 1-7 4-9 2v-3l-1-1-2-1h-2l-1 1-2-1h-5l-4 4v2l-1 3v2h-1l-1 1-1-1-2 2v1l-1 1-1-1-1 1h-2v1h-2l-1 1h-2l-1-2-3-1v1l-1 1-1-2-3-1h-1l-1-3-3 1-1-1-1-5h-2l-1-5 1 1h1l2-2c0-5-3-6-5-10v-9l-2-8-1-2h-2l-1-3h-2l-1-2h-1l-2-3-1 2v1h1l1 1-3 2v-4c0-3-2-12-4-13l-3-2-2 1-2 2-5-5h-4l2-1v-2h1l1 1h1l1-3 2-1v-1l3-2v-6l-2-11-1-2c-2-4-4-3-6-7h-2c-4 1-7 2-9 5h-1l-1-1-2-3c-2 0-3-1-5-3h-2l-1-3-1-1 4-3 3-3 2-1 2-3 3-1h5l2-5-2-1 1-1 3-3 1-3 1-1-1-1 1-1-2-3-1-1-1-2v-1h2l1-1h1c0 2 2 4 3 4l3-3 4-2 1-2h2l1-1 1-3 1-1 2 2h1l1 1-1 1v4l1 1c-1 1 1 5 1 6h1v-3l1 1 1 2 1 1h1l-2-5h-2c0-1-1-4-3-5v-1h1l6 9 4 3 3 3 4 7 6 6c2 2 11 12 11 16zm74 150v-1h-2l-3-3c0-1 1-4-3-4h-1v-2c-2-1-1-5 0-5h1v2l1 1v-1-1-3h-1-4l-3 1c0-2-4-7-5-8l-2-1-1-1h-1l-1-5-3-1v-4l-3-3-1 2h-1l-1 3h-2l-2 3-2-1h-1l-3 4-2 1h-1l-2-1-1-2-2 1h-1l-1-2h-1l-1-1h-1v6l-2 3-2 1-2-3h-1v1h-2l-2-2h-3v-1l-2 1-2-1-4 2v-1l-4 1-1-2-1 3h-2l-1 3h-2v-1l-4 2-1 1-1 1v-1l-4 3-1-1-3 4 1 1-2 3h2l-1 4 2 1 1-1 3-1 4 3-1 1v1h3v3l2 2h1l1 1h4l1 1 1-2 5-1 2-2h1l3 4v1l2-1-1-1 1-1 2 1 4-1 1-1 2-1 2 1 2 4v2h2l3-1v-1h2l5-6 1-1 2-1 1 2 2-1 1-1h1l1 1 5 10c0 2-1 4 1 6l2 1v-1l2-3h5l1 3h2l1-7 3-1 1 1 1 1 1-3v-1l-1-1 2-3h7l1 1 1-2zm-238 773l-2-1v-1h1l1-2-1-1h-1l-1-3h1v-1l3-1 2 2 1-1v-1h1l1-3-2-6h-3l-1 2-1 1h-1l-1-1h-3l3-3-3-1-6 3v-3l-3 1-2-2 2-2h2l-1-1 1-1-2-2-1 1v-1l6-2v1h1l3 3v-2c1-2 4-3 5-3v2c2-1 2-3 2-5h-1v-2l-1 4h-1l-1-2h-1l-2 3-1-1v-2l1-2-3 1h-1v-2h1l2-6v-1l2-1 4-5 2 1 1 1v1 4h2v3l2 1v1l1 4h1l1 1v3h-1v4l-1 1 1 3 2 1 2 3-1 2 1 3 1 2h2c-1 1 0 3 1 5l-1 8h-2l-1-3c-1-1-7 2-7 2l3-1-1-1h-1l-3 3h1l-4 3-8 12 1 6-1 4 2 5 4 2v-1l2-1 2-3 1 1c2 2 1 2 2 5l-1 5-3 1-3 3-1 3-1 1v3h1l-1 3-1-1-1 2-2 7v1l1 5-1 4-1-1-3 4h-1l-1-3h-2l-4-4h-4v-3h-1l2-1-1-1h-1v-2l1-1 1 1 1-1v-2l-2-4-6-1h-2l-1-1 1-1 1 1h3l4-3v-1l-2-2-7 2v11h-1l-1 2h-1l-1 1v1l1 2-1 1-1 2-1-1h-1l-1-1h-2l-1 3-2-1-2 2 1-10v-2l-2-1-2-2-1-2-2-1-2-4 1-1 1-1c5 0 5 3 7 7l2 3c1 1 2 1 2 3h1v1l1 1 4-3h-1l-1-1 1-2 1-1-2-1 3-2h1l-1-3h-1l-1 1-1-1 1-1-2-2 1-1v-3l3-1-1-5-1-5 2-1 1 1-1 1 2 1 1 2v3l2 1 2-2v-3h3l-1-2-1-2v-4l3-2 1-6 2-2 3-1 1-6 2-2 2-6 2 5-1 2 2 7h2l1-1 7-9 1-4zm70-980h1l1-1v-1l-2 1h-1zm45 4h2l2-1h1l1 1h3l3-1-4-4c-1 1-3 2-4 1l-2-2-1 1h-2l1 5zm94 69h2l1-1-4-6-1 4h-2l1 4 3-1zm-120 28h1l2-3-2-3v-2l-2 1-2 7-1 3h-1v2l1 1v1l-5 3v3h1l2-1h1l2-1 1-3-1-4 3-4zm24 13v2h1v-3l-1-1v2zm5 23v1h1v-1h-1v-1l-1-1-2 2-1 2v1h1l1-2zm13 2l1 1 2-2-3-4 1 4zm59 9v1h-1l-1 1v1h4l-1-2zm-4 9l-1-4 7 1 1 3v1l-2-1v1l-2-1-2-1zm-6 1l-1-2 2-1 1 3zm-60 7l2-1c2-3 2-2 1-6h-1l-1 1h-1l-1 1v2l-1 1 1 1 1-1zm81 3l3-2h1v-1l-2-5-4-2-1 3-1 4-1-1v3h4zm-35 2c1-1 3-2 5-1l1-1-1-1h-2v1l-2-1-2 2v1zm-5-2c1 1 0 3-1 5h-2v-1zm-69 0h-1v-1l1 1 2-1v1l1 5-1-1-1-3h-1zm13-1l-1 1-1-1v3l3 1 1 1h-1v1l-1 2v1c1-1 3-2 4-1l1 1 1-1 1 1 3-1v-1h2v-3-4h-1l-1 2-1 1-4 1-1-2h-2l-2-2zm30 20l-1-1-1 1v1l1-1 2 1 1-1h1l4-2v-1h-2-2zm-635 15l-2-4v-2h1l-2-4-1 1v1h-1l-4-1-2 1-3-1-1-1h-3v3h-1v1l2 1v2l-5 5h-2l-3-1v1h1l-1 1-1 2h1l4 5v2l4 1 4 3 4-1c0-2 4-3 5-4v-1l1-1h2v-1l-1-1v-3l1-2 1 1 2-3zm649 14l-1 1h-1l3 3 1-1 2 3h1l-2-4-3-2zm59 10l1 4 1 1v-4l-2-3h-1l-1 1zm-7 5h2v-4l1-2-1-1h-1l-1 3 1 1-1 2zm-24-9c-2 0-2 1-3 2l-3 1c-2 1-4-1-7 1l-1 2h-1l-1-1v2l-2 1-1 1v1l-1 1v2h2l2-1 1 1-1 3h1l-1 4 2 3 1-2h1l1 1 2-1 3 3 3 4 3-1h1l1-1-2-1c0-4 3-5 4-8l3-2v-1l2-2-1-1v-1l1-7 2-1 1-5-4 2c-1 1-2 2-1 3l-3 2h-1zM93 267l1 1v4c-2 4 1 9 0 13h-1v1l1 3v3l-1 3 1 2c-1 0-2 2-1 3l-2 2-1 1v3l-1 4 1 2 1 4c0 2-2 6-4 7l-5-3-1-4-1-1 1-7-1-4v-1l-1-2v-2l-3-2v-1h2v-3l-1-2v-1l3-1 1 2h4c2 0 4-4 4-6h1v-2l-1-2 2-7 1-2v-1l1-1-1-1zm31 63l2-1 1-1-3-7 1-2-2-3-1 2-1 7 1 1 2 1zm-50 2v-1l-2-1-2-1 2 4zm28 18l2-1v-1l-1-1h-1l-1 1v1zm35 2l-3-1-1 1-3 1-4 5h-1l-3-1-3 1-1 2h-2l-5-3-1-2 2-3-1-9c2-2 7-5 8-8l1-2 1-4 2-1v3l2 2-1 2 1 1 4 1 1 1h1v-1l-3-6 5 8 9 6 6-1-1 2c1 2 3 6 5 7l-4 1v-1l-1 1-2-2h1v-1c-3 0-6 0-7 2zm650 11l1-1h1l1 1 1-1h1v-1h2v-2h-2l-2 2h-3l-1 1zm-665 0v-1l-2 1h-1l1 1zm25 9l-1-2 1-7 1 1c3 5 4 6 9 9l2 2v1h-5c-2 0-4-4-7-4zm545 1l1 2-2 5-4 2v-1l1-1 1-3 3-4zm141 426v-5l-2-2-2-1h-2v1l-2 1-1-1-4-2-3 3-1 2 1 3 2 2c1 1 5 1 6-1l2 3h1l1 1 3 2h6l1-1c-2 0-3-1-4-3v-1l1-1-1-1zm-595 5l-1 1 1 3 1-1-1-1 1-1zm542 90l-3-3h-2l-1 2-3 1-2-2 5-2v-2l-2-4v-2l2-1v-1l2-2v-4l2-2-1-1h1l1-1h1v-1l1-2h1l3 4c5-1 7 0 9 4l1 4v4l6 5-1 2-2 1-1 3h-1l-2-2-2 2h-1l-1-2-4 1h-3l-2 3v3h2l2 2-1 1 2 2-1 5h-1v1h-1l-3-2v2h-2l-1-1-1 1-4 3-2 2-1-6 3-3 1-4 1-3 2-1 2 1h1l1-5zm-59 77h1l-2-3-1 2zm-29 2v-3l-1-1-2 1v1l3 2zm-34 30l1-1-1-5h-3v1zm-96 2l2-5h-1c-1 1-2 4-1 5zm-6 12h1v1l4-13-5 12zm210 32l-1 2v2l2-1 1-2zm-183 8l1-3h-1v3zm-6 40l-2-3 2-3 1 3-1 3zm-237 1l-1-1-1 3 1 1-1 1h1l1-4zm114 46h1v-2h1l1-2v-1l-1 1zm-36 8l2 1h3l4-3-1-3 1-2h-1l-4 2-2-1h-2l-1 2 1 1zM65 262s2-3 0-4v4zm337 904l2-5v-3l1 1v3l1 3-4 1zm5-10h2c0-2-1-3-2-2v2zm-8 10c-1 0-2 1 1 2 2 0 1-2-1-2zm26-5l-1 2c1 1 1-2 1-2zm-1 26v2c2 0 1-2 0-2zm-94-87v-1 1zm10-28l-1 5 1-5zm216 57c-1 0-2 2-1 3l1-3zm46-58l1 2c2 0 0-2-1-2zm-17-6c1 1 2 0 2-2-1-1-2 1-2 2zm113-45l-1 4 1-4zm141-491c-1 0 0 3 2 2 1-2-1-3-2-2zM710 194l2-1c-1-1-2 1-2 1zm36 69c1 1 2-1 2-1-1-1-3 1-2 1zm-24-132l-1 2c1 1 2-2 1-2zm-3 15c1 1 4-1 3-2s-3 1-3 2zm4-10c-1 0-2 1-1 2l1-2zM79 324c2 0 1-2 0-2v2zm-9-13c0 1 1 2 2 1l-2-1zm20 14s3-3 2-4-2 4-2 4zM47 64l1-1 1-1 1 1v-1l-2-1-2 3zm34 121l-3 2v2l-2 1v-1l-2-1-3 2c0 2-1 2-2 2l-3 2c-1 1 0 2 1 3v5l-1-1-1-2-1-2-2-1v-3-1c0-2-1-1-2-1h-3l-1-1v-5l-1-2v-2l-1-2-1-3c-2-2 0-1 0-2 1-1 2 1 3 1l3-1c1-1 0-2-1-5-1-2 0-6 1-8v-2l-1-2 2-1-1-4-3-5-1-2c-1-1-6-5-5-7l1-1-1-1h-1v-2-4l1-1 1-2-1-2 4-7-1-2 2-1c2-6-1-4-1-9v-4l1-3-3-5V84l-1-1v-5c2-1 4 1 6-3l5-2c5-4 9-10 7-17l-1-7v-3l1-1 1 1 2-4h1l2 2 3 1 1 2 4 1 1-1 3 4 2 1 6 2 2 4 9 8v9l-1 2v9c0 7-1 11 3 17l-1 3-1 1v1h1l1-2h1l1 3-2 7-2 3c-1 2 0 13 1 14l6 5-2 9c1 2 1 5-1 7v2l-2 1-2 3 1 1v1l-3 2v6h-1l-1 1h-5l-2 4h-1l-2 1 1 2-3 1-2 4h-3l-2 2zM56 56v3c2 0 1-3 0-3z", "fill:none;stroke:#000000;stroke-opacity:1")
}
