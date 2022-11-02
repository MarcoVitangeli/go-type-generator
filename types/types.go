package types

type FunctionParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type InterfaceAttr struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
	Params []FunctionParam
}

type Interface struct {
	Name   string          `json:"name"`
	Public bool            `json:"public"`
	Attrs  []InterfaceAttr `json:"attrs"`
}

type StructAttr struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Public bool   `json:"public"`
}

type Struct struct {
	Name   string       `json:"name"`
	Public bool         `json:"public"`
	Attrs  []StructAttr `json:"attrs"`
}
