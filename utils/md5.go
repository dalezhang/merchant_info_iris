package utils

import (
	"crypto/md5"
	"fmt"

	//	"github.com/emirpasic/gods/lists/arraylist"
	"strings"

	"github.com/emirpasic/gods/utils"
)

func GetMd5(str string) string {
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return sign
}

func GetMab(hash map[string]string) string {
	var mab string

	nameArr := []string{"sign", "controller", "action", "key", ""}
	newArr := []interface{}{}
	for k, _ := range hash {
		for _, name := range nameArr {
			if k == name {
				goto Here
			}
		}
		newArr = append(newArr, k)
	Here:
	}
	utils.Sort(newArr, utils.StringComparator)
	for index, value := range newArr {

		valueStr := fmt.Sprintf("%s", value)
		if hash[valueStr] == "" {
			continue
		}
		mab = mab + valueStr + "=" + hash[valueStr]
		if index < (len(newArr) - 1) {
			mab = mab + "&"
		}
	}
	return mab
}

func GetMac(js map[string]string, key string) string {
	sign := GetMd5(GetMab(js) + "&key=" + key)
	sign = strings.ToUpper(sign)
	return sign

}
