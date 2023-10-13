package main

import "fmt"

func Solve(f1 func(float64) float64, f2 func(float64) float64, b1 int, b2 int) float64 {
	var values [][]float64

	for i := b1; i <= b2; i++ {
		var equals []float64

		equals = append(equals, f1(float64(i)))
		equals = append(equals, f2(float64(i)))
		equals = append(equals, float64(i))

		values = append(values, equals)
	}

	values = sort(values, func(a []float64, b []float64) int {
		d1 := abs(a[0] - a[1])
		d2 := abs(b[0] - b[1])

		if d1 > d2 {
			return 1
		} else if d1 < d2 {
			return -1
		} else {
			return 0
		}
	})

	least := values[0]
	accuracy := 0.00000001
	diff := 0.1
	upperSide := calc_diff(f1, f2, least[2]+accuracy)
	lowerSide := calc_diff(f1, f2, least[2]-accuracy)

	useUpper := upperSide < lowerSide
	ogNumber := least[2]
	newNumber := least[2]
	var triedNumbers []float64
	triedNumbers = append(triedNumbers, ogNumber)
	i := 0

	for calc_diff(f1, f2, newNumber) > accuracy {
		fmt.Println(newNumber)
		ogNumber = newNumber
		if useUpper {
			newNumber += diff
		} else {
			newNumber -= diff
		}

		if calc_diff(f1, f2, newNumber) > calc_diff(f1, f2, ogNumber) {
			upperSide = calc_diff(f1, f2, ogNumber+accuracy)
			lowerSide = calc_diff(f1, f2, ogNumber-accuracy)
			useUpper = upperSide < lowerSide

			diff *= 0.1
		}
		i++
	}

	return newNumber
}

func calc_diff(f1 func(f float64) float64, f2 func(f float64) float64, i float64) float64 {
	val1 := f1(i)
	val2 := f2(i)

	return abs(val1 - val2)
}

// Sort sorts an array
func sort(a [][]float64, f func(a []float64, b []float64) int) [][]float64 {
	for i, elemA := range a {
		if i != 0 {
			i2 := 0
			for f(elemA, a[i2]) > 0 && i2 < len(a) {
				i2++
			}

			var i3, length int
			if i > i2 {
				length = i
				i3 = i2
			} else {
				length = i2
				i3 = i
			}
			for ; i3 < length; i3++ {
				swap(a, i3, length)
			}
		}
	}
	return a
}

func swap(fz [][]float64, i1, i2 int) [][]float64 {
	elem1 := fz[i1]
	elem2 := fz[i2]
	fz[i1] = elem2
	fz[i2] = elem1
	return fz
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	} else {
		return f
	}
}
