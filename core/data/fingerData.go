package data

import (
	_ "embed"
	"encoding/json"
	"gofinger/core/module"
	"gofinger/core/options"
	"log"
	"strings"
)

// finger.json => goby + icon_hash => 3499
// finger_big.json goby + chunsou => 10379

//go:embed fingers/finger_big.json
var high string

//go:embed fingers/finger.json
var low string

// GetFingerData 加载指纹数据
func GetFingerData(options *options.Options) []module.FingerData {
	var fingerString string
	switch options.Level {
	case 2:
		fingerString = high
	default:
		fingerString = low
	}
	reader := strings.NewReader(fingerString)
	var fingers []module.FingerData
	data := json.NewDecoder(reader)
	err := data.Decode(&fingers)
	if err != nil {
		log.Fatalln(err)
	}
	return fingers
}
