package json

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/MarcoVitangeli/go-type-generator/types"
)

func getJsonData(jsonData map[string]json.RawMessage) (types.TypeData, error) {
	var t types.TypeData
	if jsonData["struct"] == nil || jsonData["interface"] == nil {
		return t, errors.New("missing struct or interface keys in json file")
	}

	structs, err := getStructs(jsonData["struct"])

	if err != nil {
		return t, err
	}
	t.Structs = structs

	interfs, err := getInterfaces(jsonData["interface"])

	if err != nil {
		return t, err
	}
	t.Interfaces = interfs

	return t, nil
}

func ParseJson(path string) (types.TypeData, error) {
	var t types.TypeData
	file, err := os.Open(path)

	if err != nil {
		return t, err
	}
	defer file.Close()

	var jsonData map[string]json.RawMessage

	err = json.NewDecoder(file).Decode(&jsonData)

	if err != nil {
		return t, err
	}

	return getJsonData(jsonData)
}

func getStructs(j json.RawMessage) ([]types.Struct, error) {
	var s []types.Struct

	err := json.Unmarshal(j, &s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func getInterfaces(j json.RawMessage) ([]types.Interface, error) {
	var i []types.Interface

	err := json.Unmarshal(j, &i)

	if err != nil {
		return nil, err
	}

	return i, nil
}
