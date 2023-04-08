package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int                //s is a slice of int
	fmt.Println("len", len(s)) // len is "nil safe"
	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %v\n", s2)

	s3 := s2[0:4] //slicing operation half-open range
	fmt.Printf("s3 = %v\n", s3)

	// fmt.Println(s2[:100]) //Panic
	s3 = append(s3, 100)
	fmt.Printf("s3 (append) = %v\n", s3)
	fmt.Printf("s2 (append) = %v\n", s2) // s2 is changed as well

	fmt.Printf("s2: len%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("s3: len%d, cap=%d\n", len(s3), cap(s3))

	var s4 []int
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}

	fmt.Println("s4", len(s4), cap(s4))

	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"})) // A B C D E

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))
	vs4 := []float64{2, 1, 3, 4}
	fmt.Println(median(vs4))

	fmt.Println(median(nil))
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) { //enought space in underlying array
		s = s[:len(s)+1]
	} else { // need to re-allocate and copy
		fmt.Printf("rellocate: %d -> %d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	s[i] = v
	return s
}

func concat(s1, s2 []string) []string {
	// [TODO] Restrictions: No "for" loops
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func median(values []float64) (float64, error) {

	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	sort.Float64s(values)

	i := len(values) / 2
	if len(values)%2 == 1 {
		return values[i], nil
	}

	v := (values[i-1] + values[i]) / 2
	return v, nil
}
