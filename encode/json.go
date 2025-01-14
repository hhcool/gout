package encode

import (
	"encoding/json"
	"errors"
	"github.com/guonaihong/gout/core"
	"io"
	"strings"
)

var ErrNotJSON = errors.New("Not json data")

// JSONEncode json encoder structure
type JSONEncode struct {
	obj interface{}
}

// NewJSONEncode create a new json encoder
func NewJSONEncode(obj interface{}) Encoder {
	if obj == nil {
		return nil
	}

	return &JSONEncode{obj: obj}
}

// Encode json encoder
func (j *JSONEncode) Encode(w io.Writer) (err error) {
	if v, ok := core.GetBytes(j.obj); ok {
		if b := json.Valid(v); !b {
			return ErrNotJSON
		}
		_, err = w.Write(v)
		return err
	}

	//encode := json.NewEncoder(w)
	all, err := json.Marshal(j.obj)
	if err != nil {
		return err
	}

	// 不使用Encode函数的原因，encode结束之后会自作聪明的加'\n'
	//return encode.Encode(j.obj)

	// json.Marshal会自作聪明把 < > & 这三个符号换成unicode编码，此处做个替换
	content := string(all)
	content = strings.ReplaceAll(content, "\\u0026", "&")
	content = strings.ReplaceAll(content, "\\u003c", "<")
	content = strings.ReplaceAll(content, "\\u003e", ">")
	_, err = w.Write([]byte(content))

	return err
}

// Name json Encoder name
func (j *JSONEncode) Name() string {
	return "json"
}
