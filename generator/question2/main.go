package main

import (
	"fmt"
	"math"
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

func countL(seqA string, seqB string) float64 {
	ans := 0.0
	strLen := len(seqA)
	for idx := 0; idx < strLen; idx++ {
		if seqA[idx:strLen] == seqB[0:strLen-idx] {
			ans += math.Pow(2, float64(strLen-idx-1))
		}
	}
	return ans
}

func countConway(seqA string, seqB string) float64 {
	if seqA == seqB {
		return float64(1)
	}
	a := countL(seqB, seqB) - countL(seqB, seqA)
	b := countL(seqA, seqA) - countL(seqA, seqB)
	return a / (a + b)
}

func countWinningRate(seqA string, seqB string) Result {
	return Result{seqA + " vs " + seqB, countConway(seqA, seqB)}
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
			if strCurSeq == strings.ToUpper(seqA) {
				aWiningCnt++
				break
			}
			if strCurSeq == strings.ToUpper(seqB) {
				break
			}
		}
	}

	return Result{seqA + " vs " + seqB, (float64(aWiningCnt) / float64(repTime))}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	seqLen := 4

	if len(os.Args) == 2 {
		seqLen, _ = strconv.Atoi(os.Args[1])
	}

	combArray := make([]string, 0)
	permutation("", &combArray, seqLen)

	ans := make([]Result, 0)
	sortAns := make([]Result, 0)

	for _, seqA := range combArray {
		curAns := make([]Result, 0)
		for _, seqB := range combArray {
			r := countWinningRate(seqA, seqB)
			curAns = append(curAns, r)
		}
		sort.Slice(curAns, func(i, j int) bool { return curAns[i].winRate < curAns[j].winRate })
		ans = append(ans, curAns[0])

		counterA := make([]Result, 0)

		for i := 0; i < seqLen; i++ {
			var seqAChange string
			if seqA[i] == 'H' {
				seqAChange = seqA[:i] + "t" + seqA[i+1:]
			} else {
				seqAChange = seqA[:i] + "h" + seqA[i+1:]
			}
			r := countWinningRate(seqAChange, curAns[0].matchComb[seqLen+4:])
			counterA = append(counterA, r)
		}
		sort.Slice(counterA, func(i, j int) bool { return counterA[i].winRate > counterA[j].winRate })
		displayStr := fmt.Sprintf("%s | %f => %s | %f\n", curAns[0].matchComb, curAns[0].winRate, counterA[0].matchComb, counterA[0].winRate)

		sortAns = append(sortAns, Result{displayStr, counterA[0].winRate})
	}
	sort.Slice(sortAns, func(i, j int) bool { return sortAns[i].winRate < sortAns[j].winRate })
	for _, r := range sortAns {
		fmt.Print(r.matchComb)
	}
}
