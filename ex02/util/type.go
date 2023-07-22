package util

import "github.com/spf13/cobra"

type InputType string

const (
	INTEGER InputType = "int"
	FLOAT   InputType = "float"
	STRING  InputType = "string"
	MIX     InputType = "mix"
)

// GetInputType returns the input type
func GetInputType(cmd *cobra.Command, inputTypes []InputType) InputType {
	for _, t := range inputTypes {
		if cmd.Flags().Changed(string(t)) {
			return t
		}
	}
	return MIX
}
