package main

import (
	ty "shiva/types"
)

func populate() *ty.Menu {
	return &ty.Menu{
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
}
func main() {

}
