package day25

type KeyLock struct {
	isLock bool
	pins   []int
}

func ParseKeyLock(lines []string) KeyLock {
	kl := KeyLock{}
	kl.pins = make([]int, len(lines[0]))
	for r, line := range lines {
		if r == 0 {
			if line[0] == '#' {
				kl.isLock = true
			} else {
				kl.isLock = false
			}
		} else if r < len(lines)-1 {
			for c := 0; c < len(line); c++ {
				if line[c] == '#' {
					kl.pins[c]++
				}
			}
		}
	}
	return kl
}
