package resper

import "github.com/tidwall/resp"

func RespValuesToAny(values []resp.Value) []any {
	var result []any
	for _, v := range values {
		switch v.Type() {
		case resp.BulkString:
			result = append(result, v.String())
		case resp.SimpleString:
			result = append(result, v.String())
		case resp.Integer:
			result = append(result, v.Integer())
		case resp.Error:
			result = append(result, v.Error())
		case resp.Array:
			result = append(result, v.Array())
		default:
			result = append(result, nil)
		}
	}
	return result
}
