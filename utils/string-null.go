package utils

import "encoding/json"

type StringNull struct {
	IsNull bool
	Value  string
}

func (s StringNull) MarshalJSON() ([]byte, error) {
	if s.IsNull {
		return json.Marshal(nil)
	}
	return json.Marshal(s.Value)
}

func (s *StringNull) UnmarshalJSON(data []byte) error {
	var r string
	if err := json.Unmarshal(data, &r); err != nil {
		s.IsNull = true
		s.Value = ""
	} else {
		s.Value = r
		s.IsNull = false
	}
	return nil
}
