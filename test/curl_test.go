package test

import (
	"encoding/json"
	"jupiter/application/library"
	"strconv"
	"testing"
)

type HttpCourseRes struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}

func TestPost(t *testing.T) {
	classId, err := library.DecryptDES(`db6ddfa1c88b7075`)
	t.Log(classId, err)
	host := `http://course13.eeo.dom:61170`
	path := `/classes/get`
	params := `class_ids=` + strconv.Itoa(int(classId))
	res, err := library.CurlPost(host, path, params)
	t.Log(string(res), err)
	var UnRes HttpCourseRes
	err = json.Unmarshal(res, &UnRes)
	t.Log(err)
	t.Log(UnRes.Code)
	if len(UnRes.Data) > 0 {
		t.Log(int64(UnRes.Data[0]["class_etime"].(float64)))
	}
}
