package goini

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var config map[string]interface{}

func Init(name string) error {
	conf := map[string]interface{}{}
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		line, err := br.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}

		//替换掉当前行里面所有的空格.
		line = strings.Replace(line, " ", "", -1)

		//判断是否为注释.
		if strings.HasPrefix(line, "#") {
			continue
		}

		//去掉换行符.
		line = strings.Replace(line, "\n", "", -1)
		if len(line) < 1 {
			continue
		}

		key, value, err := parseLine(line)
		if err != nil {
			return err
		}
		conf[key] = value
	}
	config = conf
	return nil
}

//解析每一行的配置.
func parseLine(line string) (key string, value interface{}, err error) {
	sArr := strings.Split(line, "=")
	key = sArr[0]

	//判断是否为string.
	if isString(sArr[1]) {
		value = strings.Replace(sArr[1], `"`, "", -1)
	} else if isBool(sArr[1]) {
		if sArr[1] == "false" {
			value = false
		} else {
			value = true
		}
	} else {
		num, err := strconv.Atoi(sArr[1])
		if err != nil {
			return key, value, err
		}
		value = num
	}

	return key, value, nil
}

func isString(val string) bool {
	return strings.HasPrefix(val, `"`)
}

func isBool(val string) bool {
	if val == "true" {
		return true
	}

	if val == "false" {
		return true
	}

	return false
}

func Get(key string) interface{} {
	return config[key]
}
