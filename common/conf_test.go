package common

import (
	"testing"
	"fmt"
)

func TestUtf8ToGbk(t *testing.T) {

}

func TestGbkToUtf8(t *testing.T) {
	c := Conf{}
	c.getConf()
	fmt.Printf("%v", c)
}

func TestGetTimeStamp(t *testing.T) {
	time := GetTimeStampNow()
	fmt.Println(time)
}
func TestGetTimeNow(t *testing.T) {
	println(GetTimeNow())
}
