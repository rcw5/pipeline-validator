package utils

import "strings"

type ValidationError struct {
	MissingVarsError error
	ExtraVarsError   error
}

func (v ValidationError) Error() string {
	errors := []string{}
	if v.MissingVarsError != nil {
		errors = append(errors, v.MissingVarsError.Error())
	}
	if v.ExtraVarsError != nil {
		errors = append(errors, v.ExtraVarsError.Error())
	}
	return strings.Join(errors, "\n")
}

func (v ValidationError) IsSuccessful() bool {
	return v.MissingVarsError == nil && v.ExtraVarsError == nil
}

func CompareArrays(arr1, arr2 []string) ([]string, []string, []string) {
	if len(arr1) == 0 {
		return []string{}, arr2, []string{}
	}
	if len(arr2) == 0 {
		return arr1, []string{}, []string{}
	}

	var inLeft, inRight, inBoth []string
	for _, v := range arr1 {
		if Contains(arr2, v) {
			inBoth = appendIfNotPresent(inBoth, v)
		} else {
			inLeft = appendIfNotPresent(inLeft, v)
		}
	}

	for _, v := range arr2 {
		if !Contains(arr1, v) {
			inRight = appendIfNotPresent(inRight, v)
		}
	}
	return inLeft, inRight, inBoth
}

func Contains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func appendIfNotPresent(arr []string, val string) []string {
	if !Contains(arr, val) {
		return append(arr, val)
	}
	return arr
}
