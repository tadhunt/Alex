package main

import(
	"fmt"
)

const (
	A = iota
	B
	C
	D
	E
	F
	G
	H
	I
	N_VARS

	GUESS_C = false // set to false to match the exercise, true to also guess C
)

var possible_values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
type Guess [N_VARS]int

type d_minus_a_equals_i struct {}

func (f d_minus_a_equals_i) Test(guess Guess) bool {
	if !ValuesOK(guess[D], guess[A], guess[I]) {
		return false
	}

	return (guess[D] - guess[A]) == guess[I]
}

func (f d_minus_a_equals_i) CalcGuesses() []Guess {
	possibilities := make([]Guess, 0)

	for _, dv := range possible_values {
		for _, av := range possible_values {
			guess := Guess{
				D: dv,
				A: av,
				I: dv - av,
			}
			if f.Test(guess) {
				possibilities = append(possibilities, guess)
			}
		}
	}

	return possibilities
}

type b_minus_c_equals_g struct {}

func (f b_minus_c_equals_g) Test(guess Guess) bool {
	if !ValuesOK(guess[B], guess[C], guess[G]) {
		return false
	}

	return (guess[B] - guess[C]) == guess[G]
}

func (f b_minus_c_equals_g) CalcGuesses() []Guess {
	possibilities := make([]Guess, 0)

	for _, bv := range possible_values {
		if GUESS_C {
			for _, cv := range possible_values {
				guess := Guess{
					B: bv,
					C: cv,
					G: bv - cv,
				}

				if f.Test(guess) {
					possibilities = append(possibilities, guess)
				}
			}
		} else {
			guess := Guess{
				B: bv,
				C: 4,
				G: bv - 4,
			}

			if f.Test(guess) {
				possibilities = append(possibilities, guess)
			}
		}
	}

	return possibilities
}

type ac_plus_ic_equals_de struct {}

func (f ac_plus_ic_equals_de) Test(guess Guess) bool {
	if !ValuesOK(guess[A], guess[C], guess[I], guess[D], guess[E]) {
		return false
	}

	e1 := (10 * guess[A]) + guess[C]
	e2 := (10 * guess[I]) + guess[C]
	result := (10 * guess[D]) + guess[E]

	return e1 + e2 == result
}

func (f ac_plus_ic_equals_de) CalcGuesses() []Guess {
	possibilities := make([]Guess, 0)

	for _, av := range possible_values {
		for _, iv := range possible_values {
			for _, dv := range possible_values {
				for _, ev := range possible_values {
					if GUESS_C {
						for _, cv := range possible_values {
							guess := Guess{
								A: av,
								C: cv,
								I: iv,
								D: dv,
								E: ev,
							}
							if f.Test(guess) {
								possibilities = append(possibilities, guess)
							}
						}
					} else {
						guess := Guess{
							A: av,
							C: 4,
							I: iv,
							D: dv,
							E: ev,
						}
						if f.Test(guess) {
							possibilities = append(possibilities, guess)
						}
					}
				}
			}
		}
	}

	return possibilities
}

type fg_minus_de_equals_hi struct {}

func (f fg_minus_de_equals_hi) Test(guess Guess) bool {
	if !ValuesOK(guess[F], guess[G], guess[D], guess[E], guess[H], guess[I]) {
		return false
	}

	e1 := (guess[F] * 10) + guess[G]
	e2 := (guess[D] * 10) + guess[E]
	result := (guess[H] * 10) + guess[I]

	return e1 - e2 == result
}

func (f fg_minus_de_equals_hi) CalcGuesses() []Guess {
	possibilities := make([]Guess, 0)

	for _, fv := range possible_values {
		for _, gv := range possible_values {
			for _, dv := range possible_values {
				for _, ev := range possible_values {
					for _, hv := range possible_values {
						for _, iv := range possible_values {
							guess := Guess{
								F: fv,
								G: gv,
								D: dv,
								E: ev,
								H: hv,
								I: iv,
							}

							if f.Test(guess) {
								possibilities = append(possibilities, guess)
							}
						}
					}
				}
			}
		}
	}

	return possibilities
}

func ValuesOK(values ...int) bool {
	alreadySeen := make(map[int]bool)
	for _, v := range values {
		if v < 1 {
			return false
		}

		_, seen := alreadySeen[v]
		if seen {
			return false
		}

		alreadySeen[v] = true
	}

	return true
}

func ValuesOKSkipZero(values ...int) bool {
	alreadySeen := make(map[int]bool)
	for _, v := range values {
		if v == 0 {
			continue
		}

		if v < 1 {
			return false
		}

		_, seen := alreadySeen[v]
		if seen {
			return false
		}
		alreadySeen[v] = true
	}

	return true
}

func MatchingGuesses(guess Guess, guesses []Guess) []Guess {
	results := make([]Guess, 0)

	for _, pg := range guesses {
		matched := true
		for i, v := range guess {
			if v == 0 {
				continue
			}
			if pg[i] == 0 {
				continue
			}
			if pg[i] != v {
				matched = false
				break
			}
		}

		if matched {
			results = append(results, pg)
		}
	}

	return results
}

// joins the known values from two overlapping guesses into a new guess
func Join(ag, bg Guess) (Guess, bool) {
	result := Guess{}

	for i, v := range ag {
		if v < 1 {
			continue
		}
		result[i] = v
	}

	for i, v := range bg {
		if v < 1 {
			continue
		}
		if result[i] == 0 {
			result[i] = v
		} else if result[i] != v {
			panic(fmt.Sprintf("result %c: expected %d got %d", 'A' + i, result[i], v))
		}
	}

	values := result[:]

	if !ValuesOKSkipZero(values...) {
		return result, false
	}

	return result, true
}

// returns all guesses in A joined with B
func Intersect(a []Guess, b []Guess) []Guess {
	results := make([]Guess, 0)

	for _, ag := range a {
		matches := MatchingGuesses(ag, b)
		if len(matches) < 1 {
			continue
		}

		for _, match := range matches {
			newGuess, ok := Join(ag, match)
			if ok {
				results = append(results, newGuess)
			}
		}
	}

	return results
}

func main() {
	f1 := d_minus_a_equals_i{}
	f2 := b_minus_c_equals_g{}
	f3 := ac_plus_ic_equals_de{}
	f4 := fg_minus_de_equals_hi{}

	f1Guesses := f1.CalcGuesses()
	f2Guesses := f2.CalcGuesses()
	f3Guesses := f3.CalcGuesses()
	f4Guesses := f4.CalcGuesses()

	solutions := Intersect(f1Guesses, f2Guesses)
	solutions = Intersect(solutions, f3Guesses)
	solutions = Intersect(solutions, f4Guesses)

	fmt.Println("TOTAL SOLUTIONS: ", len(solutions))
	fmt.Printf("---\n")
	for _, ag := range solutions {
		for i, v := range ag {
			fmt.Printf("%c = %d\n", 'a' + i, v)
		}
		fmt.Printf("---\n")
	}
}
