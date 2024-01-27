package Tools

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"mime/multipart"
	"strings"
)

//go:embed DefaultPalWorldSettings.json
var IniJsonFile []byte

var DefaultSetting = InitPalWorldSettings()

type Item struct {
	Key   string
	Value interface{}
}

type InI struct {
	Inidata map[string]any `json:"inidata"`
}

func CreateIni(inifile map[string]any) (string, error) {
	var builder strings.Builder
	for i, v := range DefaultSetting {
		if _, ok := inifile[v.Key]; ok {
			v.Value = inifile[v.Key]
		}
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(v.Key)
		builder.WriteString("=")
		builder.WriteString(fmt.Sprintf("%v", v.Value))
	}
	return "[/Script/Pal.PalGameWorldSettings]" + "\n" + "OptionSettings=(" + builder.String() + ")", nil
	//return ResFile{IniFile: cfg}, nil
}

func InitPalWorldSettings() []Item {
	var res []Item
	err := json.Unmarshal(IniJsonFile, &res)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func AnalyzeInifile(inifile multipart.File) (map[string]any, error) {
	cfg, err := ini.Load(inifile)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("文件解析失败")
	}

	result := make(map[string]string)
	// 遍历每个 section
	for _, section := range cfg.Sections() {
		// 遍历每个 key-value
		for _, key := range section.Keys() {
			result[key.Name()] = key.String()
		}
	}
	res := make(map[string]any)
	if len(result["OptionSettings"]) > 2 {
		temp := strings.Split(result["OptionSettings"][1:len(result["OptionSettings"])-1], ",")
		if temp == nil {
			return nil, nil
		}
		for _, i := range temp {
			if len(strings.Split(i, "=")) > 1 {
				res[strings.Split(i, "=")[0]] = strings.Split(i, "=")[1]
			} else {
				res[strings.Split(i, "=")[0]] = ""
			}
		}
	}
	return res, nil
}
