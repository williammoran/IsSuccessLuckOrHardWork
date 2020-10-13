package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println("Results of many candidate to job ratios:")
	runRatios()
	fmt.Println()
	fmt.Println("Results of duplicating Derek's claims:")
	singleRun(18300, 100)
	fmt.Println()
	fmt.Println("Results of Derek's simulation with reasonable ranges:")
	singleRun(18300, 100000)
	fmt.Println()
	fmt.Println("Result of multiple executions to average the randomness:")
	multiRun()
	fmt.Println()
}

func runRatios() {
	fmt.Println("# Candidates, # jobs, Percent properly chosen")
	for candidates := 1000; candidates < 10000; candidates += 500 {
		pSortCombined, pSortSkill := oneGroup(candidates, 100000)
		for i := 10; i < 500; i += 10 {
			v := properlyChosen(pSortSkill, pSortCombined, i)
			percent := v * 100 / i
			fmt.Printf("%d, %d, %d\n", candidates, i, percent)
		}
	}
}

func singleRun(candidates, scorerange int) {
	pSortCombined, pSortSkill := oneGroup(candidates, scorerange)
	printCompare(pSortCombined, pSortSkill)
	for i := 1; i < 30; i++ {
		v := properlyChosen(pSortSkill, pSortCombined, i)
		percent := v * 100 / i
		fmt.Printf("Same hirings at %d jobs is %d or %d%%\n", i, v, percent)
	}
}

func multiRun() {
	const runs = 1000
	sum := 0
	for i := 0; i < runs; i++ {
		pCombined, pSkill := oneGroup(1000, 100)
		v := properlyChosen(pCombined, pSkill, 10)
		sum = sum + v
	}
	fmt.Printf("Average of %d runs: %d\n", runs, sum/runs)
}

func oneGroup(size, scorerange int) ([]person, []person) {
	pool := buildData(size, scorerange)
	pSortCombined := clonePool(pool)
	sortByCombined(pSortCombined, 95, 5)
	pSortSkill := clonePool(pool)
	sortByCombined(pSortSkill, 100, 0)
	return pSortCombined, pSortSkill
}

type person struct {
	id    int
	skill int
	luck  int
}

func buildData(size, scorerange int) []person {
	var rv []person
	for x := 0; x < size; x++ {
		p := person{}
		p.id = x
		p.skill = int(rand.Int31n(int32(scorerange))) + 1
		p.luck = int(rand.Int31n(int32(scorerange))) + 1
		rv = append(rv, p)
	}
	return rv
}

func sortByCombined(pool []person, skill, luck int) {
	sort.Slice(pool, func(i, j int) bool {
		vi := pool[i].skill*skill + pool[i].luck*luck
		vj := pool[j].skill*skill + pool[j].luck*luck
		return vi > vj
	})
}

func properlyChosen(pool1, pool2 []person, top int) int {
	same := 0
	for x := 0; x < top; x++ {
		p1 := pool1[x]
		for y := 0; y < top; y++ {
			if p1.id == pool2[y].id {
				same++
				break
			}
		}
	}
	return same
}

func printCompare(pool1, pool2 []person) {
	for index, p1 := range pool1 {
		fmt.Printf("%4d: %s %s\n", index, sPrintPerson(p1), sPrintPerson(pool2[index]))
	}
}

func sPrintPerson(p person) string {
	return fmt.Sprintf("ID: %4X Skill: %3d Luck: %3d", p.id, p.skill, p.luck)
}

func clonePool(pIn []person) []person {
	var out []person
	for _, p := range pIn {
		out = append(out, p)
	}
	return out
}
