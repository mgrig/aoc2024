package day14

import (
	"aoc2024/common"
	"regexp"
)

func Part1(lines []string, nx, ny int) int {
	reg := regexp.MustCompile(`p=(.+),(.+) v=(.+),(.+)`)

	repeats := 100
	size := NewPoint(nx, ny)
	midx, midy := (nx-1)/2, (ny-1)/2

	q1, q2, q3, q4 := 0, 0, 0, 0
	for i := range lines {
		matches := reg.FindStringSubmatch(lines[i])
		px, py := common.StringToInt(matches[1]), common.StringToInt(matches[2])
		vx, vy := common.StringToInt(matches[3]), common.StringToInt(matches[4])

		end := propagate(NewPoint(px, py), NewPoint(vx, vy), size, repeats)
		endx, endy := end.x, end.y
		//fmt.Println(end)

		if endx < midx {
			if endy < midy {
				q1++
			} else if endy > midy {
				q3++
			}
		} else if endx > midx {
			if endy < midy {
				q2++
			} else if endy > midy {
				q4++
			}
		}
	}
	//fmt.Println(q1, q2, q3, q4)

	return q1 * q2 * q3 * q4
}

func propagate(p, v, n Point, repeats int) (end Point) {
	end.x = (p.x + repeats*v.x) % n.x
	if end.x < 0 {
		end.x += n.x
	}

	end.y = (p.y + repeats*v.y) % n.y
	if end.y < 0 {
		end.y += n.y
	}

	return
}
