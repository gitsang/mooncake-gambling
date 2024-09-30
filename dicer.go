package mooncake_gambling

import (
	"math/rand"
)

var dice6Map = map[int]string{
	1: "⚀",
	2: "⚁",
	3: "⚂",
	4: "⚃",
	5: "⚄",
	6: "⚅",
}

func Gamble() ([]int, []string) {
	var (
		keyCnts          = make(map[int]int)
		dResult []int    = make([]int, 0, 6)
		results []string = make([]string, 0, 7)
	)

	for i := 0; i < 6; i++ {
		d := rand.Intn(6) + 1
		keyCnts[d]++
		dResult = append(dResult, d)
		results = append(results, dice6Map[d])
	}

	var hall bool
	for _, keyCnt := range keyCnts {
		if keyCnt != 1 {
			hall = false
			break
		}
		hall = true
	}
	if hall {
		results = append(results, "对堂")
		return dResult, results
	}

	for key, keyCnt := range keyCnts {
		if key == 4 {
			if keyCnt == 1 {
				results = append(results, "一秀")
			} else if keyCnt == 2 {
				results = append(results, "二举")
			} else if keyCnt == 3 {
				results = append(results, "三红")
			}
		}
		if keyCnt == 4 {
			if key == 4 {
				results = append(results, "状元")
			} else {
				results = append(results, "四进")
			}
		}
		if keyCnt == 5 {
			results = append(results, "状元")
		}
	}

	return dResult, results
}
