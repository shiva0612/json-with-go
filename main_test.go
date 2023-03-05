package main

import (
	"encoding/json"
	ty "shiva/types"
	"testing"
)

func BenchmarkJsonMarshaller(b *testing.B) {
	menu := ty.Menu{
		Id:    "file",
		Value: "File",
		Popup: ty.Popup{
			MenuItem: []ty.MenuItem{
				{Value: "New", OnClick: "CreateDoc()"},
				{Value: "Open", OnClick: "OpenDoc()"},
				{Value: "Save", OnClick: "SaveDoc()"},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&menu)
	}
}

func BenchmarkJsonUnMarshaller(b *testing.B) {
	data := []byte(`{
        "menu": {
            "id": "file",
            "value": "File",
            "popup": {
                "menuitem": [
                    {"value": "New", "onclick": "CreateDoc()"},
                    {"value": "Open", "onclick": "OpenDoc()"},
                    {"value": "Save", "onclick": "SaveDoc()"}
                ]
            }
        }
    }`)
	for i := 0; i < b.N; i++ {
		var menu ty.Menu
		_ = json.Unmarshal(data, &menu)
	}
}

func BenchmarkEasyJsonMarshaller(b *testing.B) {
	menu := ty.Menu{
		Id:    "file",
		Value: "File",
		Popup: ty.Popup{
			MenuItem: []ty.MenuItem{
				{Value: "New", OnClick: "CreateDoc()"},
				{Value: "Open", OnClick: "OpenDoc()"},
				{Value: "Save", OnClick: "SaveDoc()"},
			},
		},
	}
	for i := 0; i < b.N; i++ {
		_, _ = menu.MarshalJSON()
	}
}

func BenchmarkEasyJsonUnMarshaller(b *testing.B) {
	data := []byte(`{
        "menu": {
            "id": "file",
            "value": "File",
            "popup": {
                "menuitem": [
                    {"value": "New", "onclick": "CreateDoc()"},
                    {"value": "Open", "onclick": "OpenDoc()"},
                    {"value": "Save", "onclick": "SaveDoc()"}
                ]
            }
        }
    }`)
	for i := 0; i < b.N; i++ {
		var menu ty.Menu
		_ = menu.UnmarshalJSON(data)
	}
}
