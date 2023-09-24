package data

import (
	_ "embed"
	"encoding/json"
	"gofinger/core/module"
	"gofinger/core/options"
	"log"
	"strings"
)

// high => goby + chunsou
// medium => goby + chunsou 非单条的规则
// low => goby + chunsou 仅仅 icon_hash

//go:embed fingers/12830.json
var high string

//go:embed fingers/8541.json
var medium string

//go:embed fingers/7288.json
var low string

// GetFingerData 加载指纹数据
func GetFingerData(options *options.Options) module.FingerData {
	var fingerString string
	switch options.Level {
	case 3:
		fingerString = high
	case 2:
		fingerString = medium
	default:
		fingerString = low
	}
	reader := strings.NewReader(fingerString)
	var fingers module.FingerData
	data := json.NewDecoder(reader)
	err := data.Decode(&fingers)
	if err != nil {
		log.Fatalln(err)
	}
	return fingers
}
