package cmd

import (
	"fmt"
	"github.com/rovn208/df-go/ex02/util"
	"github.com/spf13/cobra"
)

var (
	intType    bool
	floatType  bool
	stringType bool
	mixType    bool
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "shorter",
		Short: "Sort elements CLI application",
		Long:  "Sort elements CLI application based on provided input type.",
		Args:  expectedArguments,
		RunE:  run,
	}

	rootCmd.Flags().BoolVarP(&intType, string(util.INTEGER), "i", false, "Type of input array is integer")
	rootCmd.Flags().BoolVarP(&floatType, string(util.FLOAT), "f", false, "Type of input array is float")
	rootCmd.Flags().BoolVarP(&stringType, string(util.STRING), "s", false, "Type of input array is string")
	rootCmd.Flags().BoolVarP(&mixType, string(util.MIX), "m", false, "Type of input array is a mix of primitive types")
	rootCmd.MarkFlagsMutuallyExclusive(string(util.INTEGER), string(util.FLOAT), string(util.STRING), string(util.MIX))

	return rootCmd
}

func run(cmd *cobra.Command, args []string) error {
	var sortedString string
	var err error
	inputType := util.GetInputType(cmd)

	switch inputType {
	case util.INTEGER:
		sortedString, err = util.SortIntArr(args)
		if err != nil {
			return err
		}
	case util.FLOAT:
		sortedString, err = util.SortFloatArr(args)
		if err != nil {
			return err
		}
	case util.STRING:
		sortedString, err = util.SortStringArr(args)
		if err != nil {
			return err
		}
	case util.MIX:
		sortedString, err = util.SortMix(args)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(cmd.OutOrStdout(), sortedString)
	return nil
}

func expectedArguments(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}
	if err := util.ValidateOneRequired(*cmd, string(util.INTEGER), string(util.FLOAT), string(util.STRING), string(util.MIX)); err != nil {
		return err
	}
	return nil
}
