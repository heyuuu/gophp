package phpparse

import "gophp/php/ast"

func decodeData(data any) (any, error) {
	switch value := data.(type) {
	case []any:
		for i, item := range value {
			if newItem, err := decodeData(item); err == nil {
				value[i] = newItem
			} else {
				return nil, err
			}
		}
	case map[string]any:
		if _, ok := value["nodeType"]; ok {
			return decodeNode(value)
		}
		for key, item := range value {
			if newItem, err := decodeData(item); err == nil {
				value[key] = newItem
			} else {
				return nil, err
			}
		}
	}

	return data, nil
}
