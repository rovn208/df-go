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
func GetInputType(cmd *cobra.Command) InputType {
	// Assume InputType has been validated at Args phase (one required type)
	for _, t := range []InputType{INTEGER, FLOAT, STRING, MIX} {
		if cmd.Flags().Changed(string(t)) {
			return t
		}
	}
	return MIX
}
