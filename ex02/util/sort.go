package util

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

func SortIntArr(args []string) (string, error) {
	var err error
	intArr := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		if intArr[i], err = strconv.Atoi(args[i]); err != nil {
			return "", errors.New(fmt.Sprintf("invalid argument: %s", args[i]))
		}
	}

	return getSortedIntArrToString(intArr)
}

func SortFloatArr(args []string) (string, error) {
	var err error
	float64Arr := make([]float64, len(args))
	for i := 0; i < len(args); i++ {
		if float64Arr[i], err = strconv.ParseFloat(args[i], 64); err != nil {
			return "", errors.New(fmt.Sprintf("invalid argument: %s", args[i]))
		}
	}

	return getSortedFloatArrToString(float64Arr)
}

func SortStringArr(args []string) (string, error) {
	var err error
	float64Arr := make([]float64, len(args))
	for i := 0; i < len(args); i++ {
		if float64Arr[i], err = strconv.ParseFloat(args[i], 64); err == nil {
			return "", errors.New(fmt.Sprintf("invalid argument: %s", args[i]))
		}
	}

	return getSortedStringArrToString(args)
}

func SortMix(args []string) (string, error) {
	return getSortedStringArrToString(args)
}

func getSortedIntArrToString(args []int) (string, error) {
	sort.Ints(args)

	rs := ""
	for _, str := range args {
		rs += fmt.Sprintf("%d ", str)
	}

	return rs, nil
}

func getSortedFloatArrToString(args []float64) (string, error) {
	sort.Float64s(args)

	rs := ""
	for _, str := range args {
		rs += fmt.Sprintf("%v ", str)
	}

	return rs, nil
}

func getSortedStringArrToString(args []string) (string, error) {
	sort.Strings(args)

	rs := ""
	for _, str := range args {
		rs += fmt.Sprintf("%s ", str)
	}

	return rs, nil
}
