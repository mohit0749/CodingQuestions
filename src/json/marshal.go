package json

import (
	utils2 "codes/src/utils"
	"fmt"
	"strings"
)

func Marshal(data interface{}) string {
	if data == nil {
		return "null"
	}
	var jsn string

	if slice, ok := data.([]interface{}); ok {
		jsn = marshalSlice(slice)
	} else if mp, ok := data.(map[string]interface{}); ok {
		jsn = marshalMap(mp)
	} else if subPart, ok := data.(int); ok {
		jsn = jsn + fmt.Sprintf("%d,", subPart)
	} else if subPart, ok := data.(string); ok {
		jsn = jsn + fmt.Sprintf("\"%s\",", subPart)
	} else if subPart, ok := data.(float64); ok {
		jsn = jsn + fmt.Sprintf("%f,", subPart)
	} else if subPart, ok := data.([]int); ok {
		jsn = jsn + fmt.Sprintf("%s,", utils2.IntSliceToString(subPart))
	} else if subPart, ok := data.([]string); ok {
		jsn = jsn + fmt.Sprintf("%s,", utils2.StringSliceToString(subPart))
	} else if subPart, ok := data.([]float64); ok {
		jsn = jsn + fmt.Sprintf("%s,", utils2.FloatSliceToString(subPart))
	}
	return jsn
}

func marshalSlice(data []interface{}) string {
	if data == nil {
		return "null"
	}
	jsn := "["
	for _, value := range data {
		if value == nil {
			jsn = jsn + "null,"
		} else if subPart, ok := value.(map[string]interface{}); ok {
			jsn = jsn + fmt.Sprintf("%s,", marshalMap(subPart))
		} else if subPart, ok := value.([]interface{}); ok {
			jsn = jsn + fmt.Sprintf("%s,", marshalSlice(subPart))
		} else if subPart, ok := value.(int); ok {
			jsn = jsn + fmt.Sprintf("%d,", subPart)
		} else if subPart, ok := value.(string); ok {
			jsn = jsn + fmt.Sprintf("\"%s\",", subPart)
		} else if subPart, ok := value.(float64); ok {
			jsn = jsn + fmt.Sprintf("%f,", subPart)
		} else if subPart, ok := value.([]int); ok {
			jsn = jsn + fmt.Sprintf("%s,", utils2.IntSliceToString(subPart))
		} else if subPart, ok := value.([]string); ok {
			jsn = jsn + fmt.Sprintf("%s,", utils2.StringSliceToString(subPart))
		} else if subPart, ok := value.([]float64); ok {
			jsn = jsn + fmt.Sprintf("%s,", utils2.FloatSliceToString(subPart))
		}
	}
	jsn = strings.Trim(jsn, ",")
	jsn += "]"
	return jsn
}

func marshalMap(mp map[string]interface{}) string {
	if mp == nil {
		return "null"
	}
	jsn := "{"
	for key, value := range mp {
		if value == nil {
			jsn = jsn + fmt.Sprintf("\"%s\":null,", key)
		} else if subPart, ok := value.(map[string]interface{}); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":%s,", key, marshalMap(subPart))
		} else if subPart, ok := value.([]interface{}); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":%s,", key, marshalSlice(subPart))
		} else if subPart, ok := value.(int); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":%d,", key, subPart)
		} else if subPart, ok := value.(string); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":\"%s\",", key, subPart)
		} else if subPart, ok := value.(float64); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":%f,", key, subPart)
		} else if subPart, ok := value.([]int); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":%s,", key, utils2.IntSliceToString(subPart))
		} else if subPart, ok := value.([]string); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":%s,", key, utils2.StringSliceToString(subPart))
		} else if subPart, ok := value.([]float64); ok {
			jsn = jsn + fmt.Sprintf("\"%s\":%s,", key, utils2.FloatSliceToString(subPart))
		}
	}
	jsn = strings.Trim(jsn, ",")
	jsn += "}"
	return jsn
}
