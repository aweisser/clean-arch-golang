package flagcli_test

import (
	"flag"
	"os"
	"testing"

	"github.com/aweisser/clean-arch-golang/io/flagcli"
	"github.com/stretchr/testify/assert"
)

func TestCallbackFunctionReturnsNameParameter(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}()
	os.Args = []string{"cmd", "--name", "John"}

	clbckExecuted := 0
	flagcli.Execute(flagcli.Definition{
		Callback: func(name string) {
			clbckExecuted++
			assert.Equal(t, "John", name)
		},
	})
	assert.Equal(t, 1, clbckExecuted)
}

func TestCallbackFunctionReturnsDefaultName(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}()
	os.Args = []string{"cmd"}

	clbckExecuted := 0
	flagcli.Execute(flagcli.Definition{
		Callback: func(name string) {
			clbckExecuted++
			assert.Equal(t, "John Doe", name)
		},
	})
	assert.Equal(t, 1, clbckExecuted)
}
