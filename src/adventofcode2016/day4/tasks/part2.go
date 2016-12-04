package tasks

import (
	"adventofcode2016/day4/common"
	"bytes"
	"fmt"
	"strings"
)

func DecryptRoom(room common.Room) common.Room {
	decryptedRoom := room
	var buffer bytes.Buffer
	for _, c := range decryptedRoom.Id {
		if (c == '-') {
			buffer.WriteRune(' ')
		} else {
			temp := c - 'a'
			temp += rune(decryptedRoom.Sector)
			temp = temp % ('z' - 'a' + 1)
			temp += 'a'
			buffer.WriteRune(temp)
		}
	}
	decryptedRoom.Id = buffer.String()
	return decryptedRoom
}

func Part2PrintValidRooms(lines []string) {
	for _, room := range common.GetValidRooms(lines) {
		room = DecryptRoom(room)
		fmt.Printf("%s [%d]\n", room.Id, room.Sector)
	}
}

func Part2GetNorthPoleObjectRoom(lines []string) common.Room {
	for _, room := range common.GetValidRooms(lines) {
		room = DecryptRoom(room)
		if strings.Contains(room.Id, "north") {
			return room
		}
		//fmt.Printf("%s [%d]\n", room.Id, room.Sector)
	}
	return common.Room{}
}