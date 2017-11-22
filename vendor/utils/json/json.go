package json

import (
	j "encoding/json"
)

func (j *Mjson) Unmarshal(data []byte, v interface{}) {
	err := j.Unmarshal(data, v)
	Log.Error(err, "json.Unmarshal error")
}
