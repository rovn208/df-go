package cmd

import (
	"errors"
	"flag"
	"github.com/rovn208/df-go/ex02/util"
)

var (
	intType    bool
	floatType  bool
	stringType bool
	mixType    bool
)

// Execute returns the sorted string based on input type
func Execute() (string, error) {
	flag.Parse()
	args := flag.Args()

	if err := ValidateArguments(args); err != nil {
		return "", err
	}

	sortedString, err := getSortedString(args)
	if err != nil {
		return "", err
	}
	return sortedString, nil
}

func init() {
	flag.BoolVar(&intType, string(util.INTEGER), false, "Type of input array is integer")
	flag.BoolVar(&floatType, string(util.FLOAT), false, "Type of input array is float")
	flag.BoolVar(&stringType, string(util.STRING), false, "Type of input array is string")
	flag.BoolVar(&mixType, string(util.MIX), false, "Type of input array is a mix of primitive types")
}

// ValidateArguments validates arguments before CLI is executed
func ValidateArguments(args []string) error {
	if len(args) < 1 {
		return errors.New("requires at least 1 arg(s), only received 0")
	}

	if err := util.ValidateOneRequired(); err != nil {
		return err
	}

	return nil
}

func getSortedString(args []string) (string, error) {
	var sortedString string
	var err error
	inputType := util.GetInputType()
	switch inputType {
	case util.INTEGER:
		sortedString, err = util.SortIntArr(args)
		if err != nil {
			return "", err
		}
	case util.FLOAT:
		sortedString, err = util.SortFloatArr(args)
		if err != nil {
			return "", err
		}
	case util.STRING:
		sortedString, err = util.SortStringArr(args)
		if err != nil {
			return "", err
		}
	case util.MIX:
		sortedString, err = util.SortMix(args)
		if err != nil {
			return "", err
		}
	}
	return sortedString, nil
}
