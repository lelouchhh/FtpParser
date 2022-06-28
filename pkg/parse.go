package pkg

import (
	"github.com/secsy/goftp"
	"strconv"
	"strings"
)

func GetFolders(client *goftp.Client, path string) []string {
	files, err := client.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var list []string
	for _, file := range files {
		list = append(list, file.Name())
	}
	return list
}

func MatchingFolder(list []string, pattern string) []string{
	var outputList []string
	for _, item := range list {

		if strings.ToLower(item[0:len(pattern)-1]) == strings.ToLower(pattern[:len(pattern)-1]){
			outputList = append(outputList, item)
		}
	}
	return outputList
}
func GetNextDay(s string) string{
	num, _ := strconv.Atoi(s[6:8])
	num++
	if num >= 10{
		return s[0:6] + strconv.Itoa(num) + s[8:10]
	}else {
		return s[0:6] + "0" + strconv.Itoa(num) + s[8:10]
	}
}

