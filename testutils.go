// Copyright (c) 2014 AKUALAB INC., All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gjoa

import (
	"testing"
)

func Comparef64(f1, f2, epsilon float64) bool {
	err := f2 - f1
	if err < 0 {
		err = -err
	}
	if err < epsilon {
		return true
	}
	return false
}

func CompareSliceFloat(t *testing.T, expected []float64, actual []float64, message string, epsilon float64) {
	for i, _ := range expected {
		if !Comparef64(expected[i], actual[i], epsilon) {
			t.Errorf("[%s]. Expected: [%f], Got: [%f]",
				message, expected[i], actual[i])
		}
	}
}

func CompareFloats(t *testing.T, expected float64, actual float64, message string, epsilon float64) {
	if !Comparef64(expected, actual, epsilon) {
		t.Errorf("[%s]. Expected: [%f], Got: [%f]",
			message, expected, actual)
	}
}

func CompareSliceInt(t *testing.T, expected []int, actual []int, message string) {
	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("[%s]. Expected: [%d], Got: [%d]",
				message, expected[i], actual[i])
		}
	}
}

func CheckError(t *testing.T, e error) {

	if e != nil {
		t.Fatal(e)
	}
}
