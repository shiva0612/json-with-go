package types

type MenuItem struct {
	Value   string `json:"value"`
	OnClick string `json:"onclick"`
}

type Popup struct {
	MenuItem []MenuItem `json:"menuitem"`
}

type Menu struct {
	Id    string `json:"id"`
	Value string `json:"value"`
	Popup Popup  `json:"popup"`
}
