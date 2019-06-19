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

func main() {
	rand.Seed(time.Now().UnixNano())
	seqLen := 3

	if len(os.Args) == 2 {
		seqLen, _ = strconv.Atoi(os.Args[1])
	}

	combArray := make([]string, 0)
	permutation("", &combArray, seqLen)

	ans := make([]Result, 0)
	for _, seqA := range combArray {
		curAns := make([]Result, 0)
		for _, seqB := range combArray {
			r := countWinningRate(seqA, seqB)
			curAns = append(curAns, r)
		}
		sort.Slice(curAns, func(i, j int) bool { return curAns[i].winRate < curAns[j].winRate })
		ans = append(ans, curAns[0])
	}

	sort.Slice(ans, func(i, j int) bool { return ans[i].winRate > ans[j].winRate })

	fmt.Printf("%-*s vs %-*s | A's Winning Rate\n\n", seqLen, "A", seqLen, "B")
	for _, r := range ans {
		fmt.Print(r.matchComb + " | ")
		fmt.Printf("%f\n", r.winRate)
	}
}
