package main

import (
	"encoding/json"
	"io/ioutil"
)

func ExportMD5(v interface{}, o string) error {
	// 写入文件
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(o, b, 0644)
	return err
}

func ExportRepeat(v interface{}, o string) error {
	// 写入文件
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(o, b, 0644)
	return err
}
