package challenges

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

type Challenge20150401 struct{}

func (Challenge20150401) GetYear() int {
	return 2015
}

func (Challenge20150401) GetDay() int {
	return 04
}

func (Challenge20150401) GetChallenge() int {
	return 01
}

func (Challenge20150401) Execute(rawFile string) error {
	i := 0
	for {
		sum := md5.Sum([]byte(fmt.Sprintf("%v%v", rawFile, i)))
		by := sum[:5]
		if strings.HasPrefix(hex.EncodeToString(by), "00000") {
			log.Println(i)
			return nil
		}
		i++
	}
}
