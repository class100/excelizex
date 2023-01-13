package excelizex

import (
	"github.com/xuri/excelize/v2"
	"reflect"
	"testing"
)

type testStream struct {
	sector   int
	TestData []testStruct
}

func (t *testStream) Next() bool {
	return len(t.TestData) >= t.sector+1
}

func (t *testStream) DataRow() (data []excelize.Cell, excelFunc ExcelFunc) {
	data = append(data, excelize.Cell{Value: t.TestData[t.sector].Name})
	data = append(data, excelize.Cell{Value: t.TestData[t.sector].Sex})
	data = append(data, excelize.Cell{Value: t.TestData[t.sector].HelloWorld})

	t.sector++

	return
}

func (t *testStream) Close() error {
	t.sector = 0

	return nil
}

func TestFile_StreamWriteIn(t *testing.T) {
	testData := []testStruct{
		{"测试人员1", "男", "123123123"},
		{"测试人员2", "男", "helloWorld"},
		{"测试人员3", "男", "&sad1231w2"},
	}

	test := testStream{TestData: testData}

	testFile := New()

	if err := testFile.AddSheetByStream(&test, NewSheet(Name(testName), Notice(testNotice))); err != nil {
		t.Fatal("TestFile_StreamWriteIn", "写入数据表错误:", err)
	}

	rows, err := testFile.excel().GetRows(testName)
	if err != nil {
		t.Fatal("TestFile_StreamWriteIn:", "获取数据表行:", err)
	}

	expectData := testStructs(testData).ToExpectData()

	if !reflect.DeepEqual(expectData, rows) {
		t.Fatalf("Expect:%+v,but%+v", expectData, rows)
	}

}
