package main

import (
	"encoding/json"
	ty "shiva/types"
	"sync"
	"testing"
)

var pool = sync.Pool{
	New: func() any {
		return &ty.Menu{}
	},
}

func BenchmarkJsonUnMarshallerWOPool1(b *testing.B) {
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
	menu := &ty.Menu{}
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &menu)
	}
}
func BenchmarkJsonUnMarshallerWOPool(b *testing.B) {
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
		menu := &ty.Menu{}
		_ = json.Unmarshal(data, &menu)
	}
}

func BenchmarkJsonUnMarshallerPool(b *testing.B) {
	m := pool.Get().(*ty.Menu)
	defer pool.Put(m)
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
		_ = json.Unmarshal(data, m)
	}

}

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
