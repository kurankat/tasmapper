package tasmapper

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
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

// two anecdotal records and two vouchered, but the first anecdotal is
// overlapped by a vouchered record (2 voucher cells 1 anecdotal)
var testRecords string = `-42.23342,147.75223
-41.34221,145.43442
-42.63341,147.75224
-42.23343,147.75222`

var fernRecords string = `-42.23342,147.75223,0
-41.34221,145.43442,1
-42.63341,147.75224,0
-42.23343,147.75222,1`

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestNewRecord(t *testing.T) {

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

func TestNewRecordList(t *testing.T) {
	coordReader := strings.NewReader(testRecords)
	recordList := NewRecordList(coordReader, "Test taxon")

	equals(t, 4, recordList.numRecs)
	equals(t, 4, recordList.RecordNumber())
	equals(t, "Test taxon", recordList.name)
	equals(t, "Test_taxon.svg", recordList.FileName())
	equals(t, "Test taxon", recordList.GetName())
	equals(t, lat, recordList.recordSlice[0].lat)
	equals(t, lon, recordList.recordSlice[0].lon)
	equals(t, dotH, recordList.recordSlice[0].dotH)
	equals(t, dotV, recordList.recordSlice[0].dotV)
	equals(t, gridH, recordList.recordSlice[0].gridH)
	equals(t, gridV, recordList.recordSlice[0].gridV)
	equals(t, h, recordList.recordSlice[0].h)
	equals(t, v, recordList.recordSlice[0].v)
	equals(t, testrad, recordList.recordSlice[0].rad)
	equals(t, utmE, recordList.recordSlice[0].utmE)
	equals(t, utmN, recordList.recordSlice[0].utmN)
}

func TestFernRecord(t *testing.T) {

	record := newVoucherRecord(-41.34221, 145.43442, 1)

	equals(t, -41.34221, record.lat)
	equals(t, 145.43442, record.lon)
	equals(t, 1, record.voucher)
}

func TestNewFernRecordList(t *testing.T) {
	coordReader := strings.NewReader(fernRecords)
	recordList := NewVoucherRecordList(coordReader, "Test taxon")

	equals(t, 0, recordList.recordSlice[0].voucher)
	equals(t, 1, recordList.recordSlice[1].voucher)

	equals(t, 4, recordList.numRecs)
	equals(t, 4, recordList.RecordNumber())
	equals(t, "Test taxon", recordList.name)
	equals(t, "Test_taxon.svg", recordList.FileName())
	equals(t, "Test taxon", recordList.GetName())
	equals(t, 3, recordList.numCells)
	equals(t, 2, recordList.numVouchered)
	equals(t, 1, recordList.numAnecdotal)
	equals(t, lat, recordList.recordSlice[0].lat)
	equals(t, lon, recordList.recordSlice[0].lon)
	equals(t, dotH, recordList.recordSlice[0].dotH)
	equals(t, dotV, recordList.recordSlice[0].dotV)
	equals(t, gridH, recordList.recordSlice[0].gridH)
	equals(t, gridV, recordList.recordSlice[0].gridV)
	equals(t, h, recordList.recordSlice[0].h)
	equals(t, v, recordList.recordSlice[0].v)
	equals(t, testrad, recordList.recordSlice[0].rad)
	equals(t, utmE, recordList.recordSlice[0].utmE)
	equals(t, utmN, recordList.recordSlice[0].utmN)
}
