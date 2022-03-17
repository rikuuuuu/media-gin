package utils

import "encoding/json"

func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

func StructToMap(val interface{}) (map[string]interface{}, error) {
	tmp, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var mapData map[string]interface{}
	err = json.Unmarshal(tmp, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}
