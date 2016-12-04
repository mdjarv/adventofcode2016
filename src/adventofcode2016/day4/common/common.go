package common

import (
	"regexp"
	"sort"
	"bytes"
	"strconv"
)

type Room struct {
	Id       string
	Sector   int
	Checksum string
}

func generateChecksum(id string) string {
	checksum := make(map[string]int)

	for _, c := range id {
		if c >= 'a' && c <= 'z' {
			if _, ok := checksum[string(c)]; ok {
				checksum[string(c)]++
			} else {
				checksum[string(c)] = 1
			}

		}
	}

	pairList := rankByValueCount(checksum)
	var buffer bytes.Buffer
	for _, pair := range pairList[:5] {
		buffer.WriteString(pair.Key)
	}

	return buffer.String()
}

func rankByValueCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value || (p[i].Value == p[j].Value && p[i].Key > p[j].Key)
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func RoomFromString(s string) Room {
	r := regexp.MustCompile(`^(?P<Id>[a-zA-Z-]+)-(?P<Sector>[0-9]+)\[(?P<Checksum>\w+)]`)
	values := r.FindStringSubmatch(s)
	sector, _ := strconv.Atoi(values[2])
	return Room{values[1], sector, values[3]}
}

func IsRoomReal(room Room) bool {
	return generateChecksum(room.Id) == room.Checksum
}

func GetValidRooms(lines []string) []Room {
	rooms := []Room{}
	for _, line := range lines {
		room := RoomFromString(line)
		if IsRoomReal(room) {
			rooms = append(rooms, room)
		}
	}
	return rooms
}