package template

import "reflect"

type Props struct {
	Code               uint16 `token:"code"`          // http status code
	Message            string `token:"message"`       // status message
	Description        string `token:"description"`   // status description
	RequestID          string `token:"request_id"`    // unique request ID: {SERVER_ICAO}-{upstream_id} or {SERVER_ICAO}-{random}-{uuidv7}
	Host               string `token:"host"`          // the value of the `Host` header
	ShowRequestDetails bool   `token:"show_details"`  // (config) show request details?
	L10nDisabled       bool   `token:"l10n_disabled"` // (config) disable localization feature?
}

// Values convert the Props struct into a map where each key is a token associated with its corresponding value.
func (p Props) Values() map[string]any {
	var result = make(map[string]any, reflect.ValueOf(p).NumField())

	for i, v := 0, reflect.ValueOf(p); i < v.NumField(); i++ {
		if token, tagExists := v.Type().Field(i).Tag.Lookup("token"); tagExists {
			result[token] = v.Field(i).Interface()
		}
	}

	return result
}
