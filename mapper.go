package tasmapper

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
	utm "github.com/kurankat/tasutm"
)

/*
MAP Coordinates & basic assumptions
Map W starts line 29											144.5034 -- W of Sandy Cape
Map E ends line 63						(34 cells wide)		  	148.5325 -- E of Cape Barren Is
Map N starts line 62											-39.5210 -- S of Pedra Branca
Map S ends line 15						(47 cells high)			-43.8028 -- N of Hogan Group
*/

const (
	leftMargin   = 30      // Left margin of map in pixels
	rightMargin  = 30      // Right margin of map in pixels
	topMargin    = 30      // Top margin of map in pixels
	bottomMargin = 55      // Bottom margin of map in pixels (wider to accomodate legend)
	cellSide     = 25      // Pixel dimesions of each side of a cell in grid map
	gridCellsW   = 34      // Number of cells in each row across the map
	gridCellsV   = 47      // Number of cells in each column
	tasWestLine  = 290000  // Westernmost easting in map
	tasEastLine  = 630000  // Easternmost easting in map
	tasNorthLine = 5620000 // Northernmost northing in map
	tasSouthLine = 5150000 // Southernmost northing in map
	kingWestLine = 220000  // Westernmost easting on King Island submap
	mapWidth     = cellSide * gridCellsW
	mapHeight    = cellSide * gridCellsV
	canvasWidth  = mapWidth + (leftMargin + rightMargin)
	canvasHeight = mapHeight + (topMargin + bottomMargin)
	pixelWidth   = (tasEastLine - tasWestLine) / mapWidth
	pixelHeight  = (tasNorthLine - tasSouthLine) / mapHeight
)

var (
	canvas  *svg.SVG
	viewBox string = `viewBox="0 0 ` + fmt.Sprint(canvasWidth) + " " + fmt.Sprint(canvasHeight) + `"`
)

func VoucherMap(records *RecordList, w io.Writer) {
	canvas = svg.New(w)

	canvas.Startraw(viewBox)

	canvas.Gid("theLot")
	drawGridMap()
	gridBox(records)

	canvas.Gid("dots")

	for i, hor := range records.voucherGrid {
		for j, ver := range hor {
			dotXCoord := i*25 + 13 + leftMargin
			dotYCoord := j*25 + 13 + topMargin
			if ver > 0 {
				var fill string
				switch ver {
				case 1:
					fill = "fill:black"
				case 2:
					fill = "fill:white"
				}
				style := fill + ";stroke-width:3px;stroke:black"
				canvas.Circle(dotXCoord, dotYCoord, 9, style)
			}
		}
	}

	canvas.Gend()
	canvas.Gend()

	canvas.End()
}

func WebMap(records *RecordList, w io.Writer) {
	canvas = svg.New(w)

	canvas.Startraw(viewBox)

	canvas.Gid("theLot")
	drawPlainMap()
	webBox(records)

	canvas.Gid("dots")
	for _, l := range records.recordSlice {
		canvas.Circle(l.h, l.v, l.rad, "fill:black")
	}
	canvas.Gend()
	canvas.Gend()
	canvas.End()
}

func GridMap(records *RecordList, w io.Writer) {
	gridList := records.GetGridRecords()
	canvas = svg.New(w)

	canvas.Startraw(viewBox)

	canvas.Gid("theLot")
	drawGridMap()
	gridBox(gridList)

	canvas.Gid("dots")

	// Make grid, test whether square is populated, if so, fill the circle black
	for i, hor := range gridList.voucherGrid {
		for j, ver := range hor {
			dotXCoord := i*25 + 12 + leftMargin
			dotYCoord := j*25 + 12 + topMargin
			if ver > 0 {
				style := "fill:black;stroke-width:2px;stroke:black"
				canvas.Circle(dotXCoord, dotYCoord, 9, style)
			}
		}
	}

	canvas.Gend()
	canvas.Gend()

	canvas.End()
}

func ExactMap(records *RecordList, w io.Writer) {
	canvas = svg.New(w)

	canvas.Startraw(viewBox)

	canvas.Gid("theLot")
	drawPlainMap()
	plainBox(records)

	canvas.Gid("dots")
	for _, l := range records.recordSlice {
		canvas.Circle(l.h, l.v, l.rad, "fill:black")
	}
	canvas.Gend()
	canvas.Gend()

	canvas.End()
}

type RecordList struct {
	numRecs      int
	numCells     int
	numVouchered int
	numAnecdotal int
	name         string
	recordSlice  []record
	voucherGrid  [50][50]int
}

func (r *RecordList) RecordNumber() (n int) {
	return r.numRecs
}

func (r *RecordList) FileName() string {
	fn := strings.Replace(r.name, " ", "_", -1) + ".svg"
	return fn
}

func (r *RecordList) GetName() string {
	return r.name
}

func (or *RecordList) GetPlainRecords() *RecordList {
	return or
}

func (or *RecordList) GetGridRecords() (nr *RecordList) {
	nr = &RecordList{name: or.name, numRecs: or.numRecs}
	nr.recordSlice = []record{}

	// Iterate through records in recordSlice. If a cell in map array is not already populated,
	// populate it by marking the array cell with voucher data
	for _, rec := range or.recordSlice {
		// If cell contains no record mark cell 1 if voucher, 2 if anecdotal
		if nr.voucherGrid[rec.gridH][rec.gridV] == 0 {
			if rec.voucher {
				nr.voucherGrid[rec.gridH][rec.gridV] = 1
			} else {
				nr.voucherGrid[rec.gridH][rec.gridV] = 2
			}
			// If the cell is only unvouchered and the new record in vouchered, mark it 1 - vouchered
		} else if nr.voucherGrid[rec.gridH][rec.gridV] == 2 {
			if rec.voucher {
				nr.voucherGrid[rec.gridH][rec.gridV] = 1
			} // There is no case for cells marked as vouchered already, as there is nothing to do.
		}
		nr.recordSlice = append(nr.recordSlice, rec)
	}

	// Add the total number of populated cells
	for i := range nr.voucherGrid {
		for _, j := range nr.voucherGrid[i] {
			if j > 0 {
				nr.numCells++
			}
			if j == 1 {
				nr.numVouchered++
			}
			if j == 2 {
				nr.numAnecdotal++
			}
		}
	}
	return nr
}

// Helper functions

// dmsToDD takes in a latitude and longitude in degrees, minutes and seconds, and returns the same
// coordinates in decimal degrees
func dmsToDD(latDeg, latMin, latSec, lonDeg, lonMin, lonSec float64) (ddLat, ddLon float64) {
	ddLat = latDeg + latMin/60 + latSec/3600
	ddLon = lonDeg + lonMin/60 + lonSec/3600
	return
}

// llParse parses a coordinate string into a floating point number and returns 0 is the string
// cannot be prsed
func llParse(coord string) (parsed float64) {
	parsed, err := strconv.ParseFloat(coord, 64)
	if err != nil {
		return 0
	}
	return
}

// voucherStatus takes in a voucher string and returns true if the voucher string is "a" or "1"
func voucherStatus(voucherString string) (voucherStatus bool) {
	switch voucherString {
	case "v":
		voucherStatus = true
	case "1":
		voucherStatus = true
	default:
		voucherStatus = false
	}
	return
}

func NewRecordList(coordData string, name string) (rl *RecordList) {
	firstLine := strings.TrimSpace(strings.Split(coordData, "\n")[0])
	dataReader := strings.NewReader(coordData)

	// Regular expressions that define whether the data is in decimal degrees or degrees,
	// minutes and seconds, and whether the data contains voucher information or not
	voucherData := regexp.MustCompile(`^(-?[34][90123](\.\d{0,10})?,14[45678](\.\d{0,10})?,[av01]|\-?[34][90123],([0123456])?\d,(([0123456])?\d(\.\d{1,2})?)?,14[5678],([0123456])?\d,(([0123456])?\d(\.\d{1,2})?)?,[av01])$`)
	nonVoucherData := regexp.MustCompile(`^(-?[34][90123],([0-5])?\d,([0-5]\d(\.\d{1,9})?)?,14[5678],([0-5])?\d,([0-5]\d(\.\d{1,9})?)?|-?\d{2}(\.\d{0,10})?,\d{3}(\.\d{0,10})?)$`)

	if voucherData.MatchString(firstLine) {
		rl = newVoucherRecordList(dataReader, name)
	} else if nonVoucherData.MatchString(firstLine) {
		rl = newNonVoucherRecordList(dataReader, name)
	}
	return rl
}

func newNonVoucherRecordList(data io.Reader, name string) (rl *RecordList) {
	rl = &RecordList{}
	rl.name = name

	dmsPattern := regexp.MustCompile(`^-?\d{2},([0-5]?\d),([0-5]?\d(\.\d{1,2})?)?,\d{3},([0-5]?\d),([0-5]?\d(\.\d{1,2})?)?$`)
	ddPattern := regexp.MustCompile(`^\-?\d{2}(\.\d{0,10})?,\d{3}(\.\d{0,10})?$`)

	dataScanner := bufio.NewScanner(data)

	for dataScanner.Scan() {
		line := strings.TrimSpace(dataScanner.Text())
		splitLine := strings.Split(line, ",")

		if dmsPattern.MatchString(line) {
			var lat, lon float64
			latDeg := llParse(splitLine[0])
			latMin := llParse(splitLine[1])
			latSec := llParse(splitLine[2])
			lonDeg := llParse(splitLine[3])
			lonMin := llParse(splitLine[4])
			lonSec := llParse(splitLine[5])
			lat, lon = dmsToDD(latDeg, latMin, latSec, lonDeg, lonMin, lonSec)
			rec := newRecord(lat, lon)
			if rec.gridH > 0 && rec.gridH < 50 && rec.gridV > 0 && rec.gridV < 50 {
				rl.numRecs++
				rl.recordSlice = append(rl.recordSlice, *rec)
			}
		} else if ddPattern.MatchString(line) {
			lat := llParse(splitLine[0])
			lon := llParse(splitLine[1])

			rec := newRecord(lat, lon)
			if rec.gridH > 0 && rec.gridH < 50 && rec.gridV > 0 && rec.gridV < 50 {
				rl.numRecs++
				rl.recordSlice = append(rl.recordSlice, *rec)
			}
		}
	}
	return rl
}

func newVoucherRecordList(data io.Reader, name string) (rl *RecordList) {
	tempList := new(RecordList)
	tempList.name = name
	dmsPattern := regexp.MustCompile(`^-?[34]\d,[12345]?\d,([12345]?\d(\.\d{1,3})?)?,14[45678],[12345]?\d,([12345]?\d(\.\d{1,3})?)?,[av01]$`)
	ddPattern := regexp.MustCompile(`^-?[34]\d(\.\d{1,9})?,14[45678](\.\d{1,9})?,[av01]$`)

	dataScanner := bufio.NewScanner(data)

	for dataScanner.Scan() {
		line := strings.TrimSpace(dataScanner.Text())
		splitLine := strings.Split(line, ",")
		var voucher bool

		if dmsPattern.MatchString(line) {
			var lat, lon float64

			latDeg := llParse(splitLine[0])
			latMin := llParse(splitLine[1])
			latSec := llParse(splitLine[2])
			lonDeg := llParse(splitLine[3])
			lonMin := llParse(splitLine[4])
			lonSec := llParse(splitLine[5])

			voucher = voucherStatus(splitLine[6])

			lat, lon = dmsToDD(latDeg, latMin, latSec, lonDeg, lonMin, lonSec)
			rec := newVoucherRecord(lat, lon, voucher)
			if rec.gridH > 0 && rec.gridH < 50 && rec.gridV > 0 && rec.gridV < 50 {
				tempList.numRecs++
				tempList.recordSlice = append(tempList.recordSlice, *rec)
			}
		} else if ddPattern.MatchString(line) {
			lat := llParse(splitLine[0])
			lon := llParse(splitLine[1])

			voucher = voucherStatus(splitLine[2])

			rec := newVoucherRecord(lat, lon, voucher)
			if rec.gridH > 0 && rec.gridH < 50 && rec.gridV > 0 && rec.gridV < 50 {
				tempList.numRecs++
				tempList.recordSlice = append(tempList.recordSlice, *rec)
			}
		}
		rl = tempList.GetGridRecords()
	}
	return
}

type record struct {
	lat, lon     float64
	h, v         int
	utmE, utmN   int
	dotH, dotV   int
	gridH, gridV int
	rad          int
	voucher      bool
}

func newRecord(lat, lon float64) (r *record) {
	r = new(record)
	r.lat, r.lon = lat, lon

	if r.lat > 0 {
		r.lat = lat * -1 // In case hemispheres are switched. This program is only for Tasmanian data
	}

	r.rad = 9
	easting, northing, _, _, err := utm.FromLatLonZone(r.lat, r.lon, false, 55)
	if err != nil {
		r.utmE, r.utmN = 0, 0
	} else {
		r.utmE, r.utmN = int(easting), int(northing)
	}

	kingIs := r.utmE > 220000 && r.utmE < 260000 && r.utmN > 5540000 && r.utmN < 5620000

	var h, v int
	if kingIs { // There may still be problems with the placement of KI specimens
		h = ((r.utmE - kingWestLine) / pixelWidth)
		v = (((tasNorthLine - 1) - r.utmN) / pixelHeight)
	} else {
		h = ((r.utmE - tasWestLine) / pixelWidth)
		v = (((tasNorthLine - 1) - r.utmN) / pixelHeight)
	}

	r.dotH = (int(h)/25)*25 + 12 + leftMargin
	r.dotV = (int(v)/25)*25 + 12 + topMargin
	r.gridH = h / 25
	r.gridV = v / 25
	r.h = h + leftMargin
	r.v = v + topMargin

	return r
}

func newVoucherRecord(lat, lon float64, voucher bool) (r *record) {
	r = new(record)
	r.lat, r.lon = lat, lon
	r.voucher = voucher

	if r.lat > 0 {
		r.lat = lat * -1 // In case hemispheres are switched. This program is only for Tasmanian data
	}

	r.rad = 9
	easting, northing, _, _, err := utm.FromLatLonZone(r.lat, r.lon, false, 55)
	if err != nil {
		r.utmE, r.utmN = 0, 0
	} else {
		r.utmE, r.utmN = int(easting), int(northing)
	}

	kingIs := r.utmE > 220000 && r.utmE < 260000 && r.utmN > 5540000 && r.utmN < 5620000

	var h, v int
	if kingIs { // There may still be problems with the placement of KI specimens
		h = ((r.utmE - kingWestLine) / pixelWidth)
		v = (((tasNorthLine - 1) - r.utmN) / pixelHeight)
	} else {
		h = ((r.utmE - tasWestLine) / pixelWidth)
		v = (((tasNorthLine - 1) - r.utmN) / pixelHeight)
	}

	r.dotH = (int(h)/25)*25 + 12 + leftMargin
	r.dotV = (int(v)/25)*25 + 12 + topMargin
	r.gridH = h / 25
	r.gridV = v / 25
	r.h = h + leftMargin
	r.v = v + topMargin

	return r
}
