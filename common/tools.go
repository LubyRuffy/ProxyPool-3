package common

import (
	"regexp"
	"fmt"
)

var re *regexp.Regexp
//区分大陆地址
func CheckArea(s string) bool {
	return re.MatchString(s)
}

func init() {
	var err error
	re, err = regexp.Compile(`辽宁|吉林|黑龙江|河北|山西|陕西|甘肃|青海|山东|安徽|江苏|浙江|河南|湖北|湖南|江西|台湾|福建|云南|海南|四川|贵州|广东|内蒙古|新疆|广西|西藏|宁夏北京|上海|天津|重庆`)
	if err != nil {
		fmt.Println("确定代理区域正则有误")
	}

}
