package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Result : MatchComb format as "seqA vs seqB"
type Result struct {
	matchComb string
	winRate   float64
}

func permutation(curComb string, combArray *[]string, len int) {
	if strings.Count(curComb, "")-1 == len {
		*combArray = append(*combArray, curComb)
		return
	}
	permutation(curComb+"H", combArray, len)
	permutation(curComb+"T", combArray, len)
}

func randomToss(headProb float64) string {
	r := rand.Float64()
	if r < headProb {
		return "H"
	}
	return "T"
}

func singleMatch(seqA string, seqB string, repTime int, headProb float64) Result {
	aWiningCnt := 0
	for i := 0; i < repTime; i++ {
		curSeq := make([]string, 0)
		for true {
			curSeq = append(curSeq, randomToss(headProb))
			if len(curSeq) > strings.Count(seqA, "")-1 {
				_, curSeq = curSeq[0], curSeq[1:]
			}
			strCurSeq := strings.Join(curSeq, "")
			if strCurSeq == seqA {
				aWiningCnt++
				break
			}
			if strCurSeq == seqB {
				break
			}
		}
	}

	return Result{seqA + " vs " + seqB, (float64(aWiningCnt) / float64(repTime))}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	seqLen := 5
	repTime := 1000
	headProb := 0.5

	if len(os.Args) == 2 {
		seqLen, _ = strconv.Atoi(os.Args[1])
	}

	combArray := make([]string, 0)
	permutation("", &combArray, seqLen)

	ans := make([]Result, 0)
	for _, seqA := range combArray {
		curAns := make([]Result, 0)
		for _, seqB := range combArray {
			r := singleMatch(seqA, seqB, repTime, headProb)
			curAns = append(curAns, r)
		}
		sort.Slice(curAns, func(i, j int) bool { return curAns[i].winRate < curAns[j].winRate })
		ans = append(ans, curAns[0])
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i].winRate < ans[j].winRate })

	for _, r := range ans {
		fmt.Print(r.matchComb + " | ")
		fmt.Printf("%f\n", r.winRate)
	}
}
