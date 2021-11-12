package tasmapper

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

const (
	lat     = -42.23342
	badLat  = 42.23342
	lon     = 147.75223
	dotH    = 717
	dotV    = 767
	gridH   = 27
	gridV   = 29
	h       = 710
	v       = 769
	testrad = 9
	utmE    = 562069
	utmN    = 5324033
)

// // assert fails the test if the condition is false.
// func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
// 	if !condition {
// 		_, file, line, _ := runtime.Caller(1)
// 		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
// 		tb.FailNow()
// 	}
// }

// // ok fails the test if an err is not nil.
// func ok(tb testing.TB, err error) {
// 	if err != nil {
// 		_, file, line, _ := runtime.Caller(1)
// 		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
// 		tb.FailNow()
// 	}
// }

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestUnvoucheredRecord(t *testing.T) {

	record := newRecord(lat, lon)
	badRecord := newRecord(badLat, lon)

	equals(t, lat, record.lat)
	equals(t, lat, badRecord.lat)
	equals(t, lon, record.lon)
	equals(t, dotH, record.dotH)
	equals(t, dotV, record.dotV)
	equals(t, gridH, record.gridH)
	equals(t, gridV, record.gridV)
	equals(t, h, record.h)
	equals(t, v, record.v)
	equals(t, testrad, record.rad)
	equals(t, utmE, record.utmE)
	equals(t, utmN, record.utmN)
}

func TestVoucheredRecord(t *testing.T) {

	record := newVoucherRecord(-41.34221, 145.43442, true)

	equals(t, -41.34221, record.lat)
	equals(t, 145.43442, record.lon)
	equals(t, true, record.voucher)

	record = newVoucherRecord(-41.34221, 145.43442, false)
	equals(t, false, record.voucher)
}

var mixedVouchered string = `-42.63341,147.75224,v
-42.23342,147.75223,a
41,24,0,145,24,0,1
-41.34221,145.43442,a
41,24,,145,24,,v`

func TestNewRecordListVouchered(t *testing.T) {
	ml := NewRecordList(mixedVouchered, "Test taxon")

	testLat := (41.0 + (24.0 / 60) + 0.0/3600) * -1
	testLong := 145 + 24.0/60 + 0.0/3600

	equals(t, testLat, ml.recordSlice[4].lat)
	equals(t, testLong, ml.recordSlice[4].lon)
	equals(t, testLat, ml.recordSlice[2].lat)
	equals(t, testLong, ml.recordSlice[2].lon)
	equals(t, true, ml.recordSlice[0].voucher)
	equals(t, false, ml.recordSlice[1].voucher)

	equals(t, true, ml.recordSlice[2].voucher)
	equals(t, false, ml.recordSlice[3].voucher)
	equals(t, true, ml.recordSlice[4].voucher)

	equals(t, 5, ml.numRecs)
	equals(t, 5, ml.RecordNumber())
	equals(t, "Test taxon", ml.name)
	equals(t, "Test_taxon.svg", ml.FileName())
	equals(t, "Test taxon", ml.GetName())
	equals(t, 4, ml.numCells)
	equals(t, 2, ml.numVouchered)
	equals(t, 2, ml.numAnecdotal)
	equals(t, lat, ml.recordSlice[1].lat)
	equals(t, lon, ml.recordSlice[1].lon)
	equals(t, dotH, ml.recordSlice[1].dotH)
	equals(t, dotV, ml.recordSlice[1].dotV)
	equals(t, gridH, ml.recordSlice[1].gridH)
	equals(t, gridV, ml.recordSlice[1].gridV)
	equals(t, h, ml.recordSlice[1].h)
	equals(t, v, ml.recordSlice[1].v)
	equals(t, testrad, ml.recordSlice[1].rad)
	equals(t, utmE, ml.recordSlice[1].utmE)
	equals(t, utmN, ml.recordSlice[1].utmN)
}

var mixedNonVoucher string = `41,23,45.4,145,24,54.2
-41.34221,145.43442
43,12,,146,23,
-42.63341,147.75224
-42.23342,147.75223`

func TestNewRecordListUnvouchered(t *testing.T) {
	ml := NewRecordList(mixedNonVoucher, "Test taxon")
	testLat := (41.0 + (23.0 / 60) + 45.4/3600) * -1
	testLong := 145 + 24.0/60 + 54.2/3600

	equals(t, testLat, ml.recordSlice[0].lat)
	equals(t, testLong, ml.recordSlice[0].lon)

	equals(t, 5, ml.numRecs)
	equals(t, 5, ml.RecordNumber())
	equals(t, "Test taxon", ml.name)
	equals(t, "Test_taxon.svg", ml.FileName())
	equals(t, "Test taxon", ml.GetName())
	equals(t, lat, ml.recordSlice[4].lat)
	equals(t, lon, ml.recordSlice[4].lon)
	equals(t, dotH, ml.recordSlice[4].dotH)
	equals(t, dotV, ml.recordSlice[4].dotV)
	equals(t, gridH, ml.recordSlice[4].gridH)
	equals(t, gridV, ml.recordSlice[4].gridV)
	equals(t, h, ml.recordSlice[4].h)
	equals(t, v, ml.recordSlice[4].v)
	equals(t, testrad, ml.recordSlice[4].rad)
	equals(t, utmE, ml.recordSlice[4].utmE)
	equals(t, utmN, ml.recordSlice[4].utmN)
}

var newTestRecords string = `42,147,0
42.0,147.0,1
41,146,a
41.1,145.1,v
`

func TestRecordListDebugging(t *testing.T) {
	tl := NewRecordList(newTestRecords, "Test record")

	equals(t, 4, len(tl.recordSlice))
	equals(t, 4, tl.RecordNumber())
	equals(t, true, tl.recordSlice[1].voucher)
	equals(t, true, tl.recordSlice[3].voucher)
	equals(t, false, tl.recordSlice[0].voucher)
	equals(t, false, tl.recordSlice[2].voucher)

}

func TestVoucherMap(t *testing.T) {
	name := "Test"
	// cells := 4
	records := 4
	anecdotal := 1
	vouchered := 2

	mapRegex := fmt.Sprintf(`^(?s).*<svg\n *viewBox.*infoBox.*%s.*\d cells.*Total records: %d.*<\/svg>.*$`, name, records)
	gridMapPattern := regexp.MustCompile(mapRegex)

	tl := NewRecordList(newTestRecords, "Test")

	voucherDotRegex := regexp.MustCompile(`(?m)^<circle cx="\d*" cy="\d*" r="\d" style="fill:black.*" \/>$`)
	nonVoucherDotRegex := regexp.MustCompile(`(?m)^<circle cx="\d*" cy="\d*" r="\d" style="fill:white.*" \/>$`)

	mapBuffer := new(bytes.Buffer)
	VoucherMap(tl, mapBuffer)

	mapString := mapBuffer.String()

	equals(t, true, gridMapPattern.MatchString(mapString))
	equals(t, vouchered, len(voucherDotRegex.FindAllString(mapString, -1)))
	equals(t, anecdotal, len(nonVoucherDotRegex.FindAllString(mapString, -1)))

}

func TestNonVoucherGridMap(t *testing.T) {
	name := "Test"
	cells := 5
	records := 5

	mapRegex := fmt.Sprintf(`^(?s).*<svg\n *viewBox.*infoBox.*%s.*%d cells.*Total records: %d.*<\/svg>.*$`, name, cells, records)
	gridMapPattern := regexp.MustCompile(mapRegex)

	tl := NewRecordList(mixedNonVoucher, "Test")

	dotRegex := regexp.MustCompile(`(?m)^<circle cx="\d*" cy="\d*" r="\d" style="fill:black.*" \/>$`)

	mapBuffer := new(bytes.Buffer)
	GridMap(tl, mapBuffer)

	mapString := mapBuffer.String()

	equals(t, true, gridMapPattern.MatchString(mapString))
	equals(t, records, len(dotRegex.FindAllString(mapString, -1)))
}

func TestPlainMap(t *testing.T) {
	name := "Test"
	records := 5

	tl := NewRecordList(mixedNonVoucher, "Test")

	mapRegex := fmt.Sprintf(`^(?s).*<svg\n *viewBox.*infoBox.*>%s*<.*>Total records: %d<.*<\/svg>.*$`, name, records)

	dotRegex := regexp.MustCompile(`(?m)^<circle cx="\d*" cy="\d*" r="\d" style="fill:black" \/>$`)

	exactMapPattern := regexp.MustCompile(mapRegex)

	mapBuffer := new(bytes.Buffer)
	ExactMap(tl, mapBuffer)

	mapString := mapBuffer.String()

	equals(t, true, exactMapPattern.MatchString(mapString))
	equals(t, records, len(dotRegex.FindAllString(mapString, -1)))

}

func TestWebMap(t *testing.T) {
	name := "Test"
	records := 5

	tl := NewRecordList(mixedNonVoucher, "Test")

	mapRegex := fmt.Sprintf(`^(?sm)<\?xml version.*<svg\n *viewBox.*<g style.*Tasmanian Herbarium.*>%s<.*>Total records: %d<.*<\/svg>$`, name, records)

	dotRegex := regexp.MustCompile(`(?m)^<circle cx="\d*" cy="\d*" r="\d" style="fill:black" \/>$`)

	exactMapPattern := regexp.MustCompile(mapRegex)

	mapBuffer := new(bytes.Buffer)
	WebMap(tl, mapBuffer)

	mapString := mapBuffer.String()

	equals(t, true, exactMapPattern.MatchString(mapString))
	equals(t, records, len(dotRegex.FindAllString(mapString, -1)))

}
