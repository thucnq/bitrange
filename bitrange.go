package bitrange

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Range is time range in day
// for example: {"08:30"; "10:00"}
type Range struct {
	Start string
	End   string
}

// BitRange is presentation of time ranges in some unsigned 64-bits integer
type BitRange []uint64

func (br BitRange) String() string {
	str := ""
	for _, item := range br {
		str += fmt.Sprintf("%064b", item)
	}
	return str
}

// GetBitRanges transform ranges to BitRange
// minuteBlockRange is duration of a block in minute
func GetBitRanges(
	ranges []Range, minuteBlockRange int,
) (BitRange, error) {
	if minuteBlockRange%5 != 0 {
		return nil, errors.New("blockRange must be divisible by 5")
	}

	var blockCount int = 1440 / minuteBlockRange

	var storeCount int = blockCount / 64
	if blockCount%64 != 0 {
		storeCount++
	}

	hs := make([]uint64, 0)
	for idx := 0; idx < storeCount; idx++ {
		hs = append(hs, 0)
	}

	for _, shift := range ranges {
		s := getBitIndex(shift.Start, minuteBlockRange)
		e := getBitIndex(shift.End, minuteBlockRange)
		scIdx := s / 64
		var b uint64 = 1 << (63 - s%64)
		for idx := s; idx < e; idx++ {
			if idx%64 == 0 && idx >= 64 {
				scIdx++
				b = 1 << 63
			}
			hs[scIdx] += b
			b >>= 1
		}
	}

	return hs, nil
}

// HaveOverlappedBit check some BitRange have overlapped bit or not
func HaveOverlappedBit(arr []BitRange) bool {
	if arr == nil || len(arr) <= 1 {
		return false
	}
	tmp := arr[0]
	for idx := 1; idx < len(arr); idx++ {
		for i := 0; i < len(arr[idx]); i++ {
			if arr[idx][i]&tmp[i] != 0 {
				return true
			}
			tmp[i] |= arr[idx][i]
		}
	}

	return false
}

// getBitIndex transform normal time to index in BitRange each minuteBlockRange
func getBitIndex(timeStr string, minuteBlockRange int) int {
	t := strings.Split(timeStr, ":")
	hour, _ := strconv.Atoi(t[0])
	min, _ := strconv.Atoi(t[1])
	return hour*(60/minuteBlockRange) + min/minuteBlockRange
}
