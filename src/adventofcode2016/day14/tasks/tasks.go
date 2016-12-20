package tasks

import (
	"crypto/md5"
	"fmt"
	"strings"
	"encoding/hex"
	"runtime"
)

func IsKey(salt string, index int) bool {
	sum := GenerateHash(salt, index)
	found := ' '

	for i := 0; i < len(sum) - 2; i++ {
		if sum[i] == sum[i + 1] && sum[i] == sum[i + 2] {
			found = rune(sum[i])
			break
		}
	}

	if found != ' ' {
		search := fmt.Sprintf("%c%c%c%c%c", found, found, found, found, found)

		for i := index + 1; i < index + 1000; i++ {
			sum := GenerateHash(salt, i)
			if strings.Contains(sum, search) {
				return true
			}
		}
	}

	return false
}

func Day14Part1(salt string) int {
	index := -1
	keys := []string{}
	for len(keys) < 64 {
		index++
		if IsKey(salt, index) {
			keys = append(keys, GenerateHash(salt, index))
		}
	}
	//fmt.Println(index, " | ", keys[63])
	return index
}

func GenerateHash(salt string, index int) string {
	data := fmt.Sprintf("%s%d", salt, index)
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func GenerateStretchHash(salt string, index int) string {
	data := fmt.Sprintf("%s%d", salt, index)
	for i := 0; i <= 2016; i++ {
		d := md5.Sum([]byte(data))
		data = hex.EncodeToString(d[:])
	}
	return data
}

func IsKey2(salt string, index int) bool {
	sum := GenerateStretchHash(salt, index)
	found := ' '

	for i := 0; i < len(sum) - 2; i++ {
		if sum[i] == sum[i + 1] && sum[i] == sum[i + 2] {
			found = rune(sum[i])
			break
		}
	}

	if found != ' ' {
		search := fmt.Sprintf("%c%c%c%c%c", found, found, found, found, found)
		for i := index + 1; i < index + 1000; i++ {
			sum := GenerateStretchHash(salt, i)
			if strings.Contains(sum, search) {
				//fmt.Println("  IsKey2(",salt,",",index," returns",true)
				return true
			}
		}
	}

	//fmt.Println("  IsKey2(",salt,",",index," returns",false)
	return false
}

func Day14Part2(salt string) int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var ch [](chan bool)
	for i := 0; i < 1000; i++ {
		ch = append(ch, make(chan bool))
	}

	index := -len(ch)
	keys := []string{}
	for len(keys) < 64 {
		index += len(ch)

		for i := 0; i < len(ch); i++ {
			go func(i int) {
				ch[i] <- IsKey2(salt, index + i)
			}(i)
		}

		for i := 0; i < len(ch); i++ {
			if <-ch[i] {
				keys = append(keys, GenerateStretchHash(salt, index + i))
				fmt.Println("At", index + i, "found keys:", len(keys))
				if len(keys) >= 64 {
					return index + i
				}
			}
		}
	}
	//fmt.Println(index, " | ", keys[63])
	return index
}
