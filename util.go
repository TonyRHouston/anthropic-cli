package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func parseInt(string string) (int64, error) {
	integer, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		return 0, err
	}

	return integer, nil
}

func parseFloat(string string) (float64, error) {
	number, err := strconv.ParseFloat(string, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func jsonSet(json []byte, path string, value interface{}) ([]byte, error) {
	keys := strings.Split(path, ".")
	path = ""
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if key == "#" {
			key = strconv.Itoa(len(gjson.GetBytes(json, path).Array()) - 1)
		}

		if len(path) > 0 {
			path += "."
		}
		path += key
	}
	return sjson.SetBytes(json, path, value)
}

func getStdInput() []byte {
	if !isInputPiped() {
		return nil
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
