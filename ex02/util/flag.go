package util

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// ValidateOneRequired validates one required flag
func ValidateOneRequired(cmd cobra.Command, flagNames ...string) error {
	count := 0
	for _, name := range flagNames {
		currentFlag := cmd.Flags().Lookup(name)
		if currentFlag == nil {
			return errors.New(fmt.Sprintf("failed to find flag %s", name))
		}
		if currentFlag.Changed {
			count++
		}
	}
	if count == 0 {
		return InputTypeMissingError
	}
	return nil

}
