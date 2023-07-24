package test

import (
	"GoBlog/utils"
	"testing"
)

func TestPasswordGenerate(t *testing.T) {
	//5c6e69a99f90bb00282001f70cec9f85
	str := utils.Md5Crypt("87c62335d3e28d89189ee83e1fac8747", "jmycool")
	t.Log(str)
}
