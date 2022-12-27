package excelizex

import (
	"reflect"
	"testing"
)

var (
	testName   = "test_sheet"
	testNotice = "test_sheet_notice"
	testHeader = []string{"test1", "test2", "test3"}
)

type testStruct struct {
	Name       string `excel:"名称" json:"sheet"`
	Sex        string `excel:"性别" json:"sex"`
	HelloWorld string `excel:"测试" json:"helloWorld"`
}

type testStructs []testStruct

func (t testStructs) ToExpectStrings() [][]string {
	var ss [][]string

	ss = append(ss, []string{testNotice})
	ss = append(ss, []string{"名称", "性别", "测试"})
	for _, ts := range t {
		ss = append(ss, []string{ts.Name, ts.Sex, ts.HelloWorld})
	}

	return ss
}

func TestGen(t *testing.T) {
	t.Run("TestGen", func(t *testing.T) {
		var ttt testStruct
		var expectSheet = Sheet{
			Header: []string{"名称", "性别", "测试"},
		}
		sheet := genSheet(ttt)

		if !reflect.DeepEqual(expectSheet, sheet) {
			t.Fatalf("expect %+v,but %+v", expectSheet, sheet)
		}
	})
}

func TestSliceGen(t *testing.T) {
	t.Run("TestGen", func(t *testing.T) {
		ttt := []testStruct{
			{"123", "男", "456"},
			{"456", "女", "213"},
		}

		var expectSheet = Sheet{
			Header: []string{"名称", "性别", "测试"},
			Data: [][]any{
				{"123", "男", "456"},
				{"456", "女", "213"},
			},
		}
		sheet := Gen(ttt)

		if !reflect.DeepEqual(expectSheet, sheet) {
			t.Fatalf("expect %+v,but %+v", expectSheet, sheet)
		}
	})
}
