package radix

import (
	"sort"
)

const (
	MinSize      = 256
	radix   uint = 8
	bitSize uint = 64
)

func Sort(x []uint64) {
	if len(x) < MinSize {
		sort.Slice(x, func(i, j int) bool { return x[i] < x[j] })
	} else {
		buffer := make([]uint64, len(x))
		SortBYOB(x, buffer)
	}
}

func SortBYOB(x, buffer []uint64) {
	if len(x) > len(buffer) {
		panic("Buffer so small")
	}
	if len(x) < 2 {
		return
	}

	from := x
	to := buffer[:len(x)]
	var key uint8

	for keyOffset := uint(0); keyOffset < bitSize; keyOffset += radix {
		var offset [512]int
		sorted := true
		var prev uint64 = 0

		for _, elem := range from {

			key = uint8(elem >> keyOffset)
			offset[key]++
			if sorted {
				sorted = elem >= prev
				prev = elem
			}
		}

		if sorted {
			if (keyOffset/radix)%2 == 1 {
				copy(to, from)
			}
			return
		}
		watermark := 0
		if keyOffset == bitSize-radix {
			for i := 128; i < len(offset); i++ {
				count := offset[i]
				offset[i] = watermark
				watermark += count
			}
			for i := 0; i < 128; i++ {
				count := offset[i]
				offset[i] = watermark
				watermark += count
			}
		} else {
			for i, count := range offset {
				offset[i] = watermark
				watermark += count
			}
		}

		for _, elem := range from {
			key = uint8(elem >> keyOffset)
			to[offset[key]] = elem
		}

		from, to = to, from
	}
}
