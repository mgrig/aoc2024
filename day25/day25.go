package day25

func Part1(lines []string) int {
	keys := make([]KeyLock, 0)
	locks := make([]KeyLock, 0)

	keylockLines := make([]string, 0)
	for r, line := range lines {
		if r == len(lines)-1 {
			keylockLines = append(keylockLines, line)
		}
		if line == "" || r == len(lines)-1 {
			kl := ParseKeyLock(keylockLines)
			if kl.isLock {
				locks = append(locks, kl)
			} else {
				keys = append(keys, kl)
			}
			keylockLines = make([]string, 0)
		} else {
			keylockLines = append(keylockLines, line)
		}
	}

	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			//fmt.Printf("key: %v, lock: %v ", key, lock)
			if match(key, lock) {
				//fmt.Println("match")
				count++
			} else {
				//fmt.Println("overlap")
			}
		}
	}

	return count
}

func match(key KeyLock, lock KeyLock) bool {
	if len(key.pins) != len(lock.pins) {
		panic("pins length mismatch")
	}
	for i := 0; i < len(key.pins); i++ {
		if key.pins[i]+lock.pins[i] > 5 {
			return false
		}
	}
	return true
}
