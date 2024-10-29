package test

import (
	"GoBlog/utils"
	"testing"
)

func TestPasswordGenerate(t *testing.T) {
	//fe92645546e4bfd7a1f9c204042c9f76
	str := utils.Md5Crypt("root", "jmycool")
	t.Log(str)
}
