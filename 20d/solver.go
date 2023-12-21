package main

import (
	"fmt"
	"slices"
	"strings"
)

type Dest = []string

type Flip struct {
	state bool // on:true,off:false
	dest  Dest
}
type Flips = map[string]Flip

type Con struct {
	state map[string]bool
	dest  Dest
}
type Cons = map[string]Con
type Ope struct {
	from  string
	pulse bool // high:true low: false
	to    string
}

const (
	BROADCASTER = "broadcaster"
)

var (
	OFF  = false
	ON   = true
	LOW  = false
	HIGH = true
)

// type Bcast struct {
// 	dest []string
// }

func main() {
	ans1() // ans:
	// ans2() // ans:
}

func ans1() {
	// file, scanner := getScanner("./example.txt")

	cons, flips, bcasts := getActualInput()
	// cons, flips, bcasts := getExampleInput(1)
	// fmt.Println(cons, flips, bcasts)
	fmt.Println("bcasts", bcasts)
	mapPrint(flips, "flips")
	mapPrint(cons, "cons")

	opes := make([]Ope, 0) // this is supposed to be queue
	ans := []int{0, 0}     // low, high
	loopCnt := 0

Loop:
	for {
		loopCnt++
		opes = append(opes, Ope{from: "button", pulse: LOW, to: BROADCASTER})

		for {
			for _, ope := range slices.Clone(opes) {
				if ope.pulse == LOW {
					ans[0]++
				} else if ope.pulse == HIGH {
					ans[1]++
				}
				// fmt.Println("ope", ope)
				// fmt.Println("flips", flips)
				// fmt.Println("cons", cons)

				if ope.to == BROADCASTER {
					for _, dest := range bcasts {
						opes = append(opes, Ope{from: BROADCASTER, pulse: ope.pulse, to: dest})
					}
				} else if hasKey(cons, ope.to) {
					// update state here?
					con := cons[ope.to]
					con.state[ope.from] = ope.pulse
					cons[ope.to] = con
					allhigh := every(getMapValues(cons[ope.to].state), func(s bool) bool {
						return s == HIGH
					})
					for _, dest := range cons[ope.to].dest {
						if allhigh == true {
							opes = append(opes, Ope{from: ope.to, pulse: LOW, to: dest})
						} else {
							opes = append(opes, Ope{from: ope.to, pulse: HIGH, to: dest})
						}
					}
				} else if hasKey(flips, ope.to) {
					if ope.pulse == HIGH {
						// do nothing
					} else {
						flip := flips[ope.to]
						flip.state = !flip.state
						flips[ope.to] = flip
						for _, dest := range flips[ope.to].dest {
							if flip.state == ON {
								opes = append(opes, Ope{from: ope.to, pulse: HIGH, to: dest})
							} else if flip.state == OFF {
								opes = append(opes, Ope{from: ope.to, pulse: LOW, to: dest})
							}
						}
					}
				}

				opes = opes[1:]
			}

			if len(opes) == 0 {
				break
			}
		}

		isDefault := true
		for _, c := range cons {
			for _, v := range c.state {
				if v == ON {
					isDefault = false
					continue Loop
				}
			}
		}
		for _, f := range flips {
			if f.state == ON {
				isDefault = false
				continue Loop
			}
		}
		if isDefault {
			break
		}
	}

	fmt.Println(loopCnt, ans)
}

func getActualInput() (Cons, Flips, Dest) {
	file, scanner := getScanner("./input.txt")
	defer file.Close()

	flips := make(Flips, 0)
	cons := make(Cons, 0)
	bcasts := make(Dest, 0)
	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()

		// **for actual input
		dests := make(Dest, 0)

		for i := strings.Index(line, "->") + 3; i < len(line)-1; {
			dests = append(dests, line[i:i+2])
			i += 4
		}

		switch line[0] {
		case '&':
			cons[line[1:3]] = Con{make(map[string]bool, 0), dests}
		case '%':
			flips[line[1:3]] = Flip{false, dests}
		case 'b':
			bcasts = dests
		}
	}

	for k, c := range cons {
		for fk, f := range flips {
			for _, fd := range f.dest {
				// fmt.Println("deb", k, fk, fd)
				if k == fd {
					c.state[fk] = false
				}
			}
		}
		for k2, c2 := range cons {
			for _, cd := range c2.dest {
				if k != k2 && k == cd {
					c.state[k2] = false
				}
			}
		}
	}

	return cons, flips, bcasts
}

func getExampleInput(v int) (cons Cons, flips Flips, bcasts Dest) {
	if v == 1 {
		cons = Cons{"inv": Con{state: map[string]bool{"c": false}, dest: []string{"a"}}}
		flips = Flips{
			"a": Flip{state: false, dest: []string{"b"}},
			"b": Flip{state: false, dest: []string{"c"}},
			"c": Flip{state: false, dest: []string{"inv"}},
		}
		bcasts = []string{"a", "b", "c"}
	}
	if v == 2 {
		cons = Cons{
			"inv": Con{state: map[string]bool{"a": false}, dest: []string{"b"}},
			"con": Con{state: map[string]bool{"a": false, "b": false}, dest: []string{"output"}},
		}
		flips = Flips{
			"a": Flip{state: false, dest: []string{"inv", "con"}},
			"b": Flip{state: false, dest: []string{"con"}},
		}
		bcasts = []string{"a"}
	}
	return
}
