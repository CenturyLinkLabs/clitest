package main

import (
	"testing"

	"github.com/CenturyLinkLabs/clitest"
)

var b clitest.BuildTester

func init() {
	b = clitest.NewBuild()
}

func TestNoArguments(t *testing.T) {
	r := b.Run(t)
	r.AssertSuccessful()
	r.AssertStdout("No Arguments Passed")
	r.AssertStderr("")
}

func TestArguments(t *testing.T) {
	r := b.Run(t, "-test")
	r.AssertSuccessful()
	r.AssertStdout("You set the test flag")
	r.AssertStderr("")
}

func TestEnvironmentVariables(t *testing.T) {
	r := b.RunWithOptions(t,
		clitest.RunOptions{
			Environment: map[string]string{
				"CLITEST_TEST_VAR": "testing123",
			},
		},
	)
	r.AssertSuccessful()
	r.AssertStdout("CLITEST_TEST_VAR is testing123")
	r.AssertStderr("")
}

func TestBadExit(t *testing.T) {
	r := b.Run(t, "-explode")
	r.AssertExitCode(19)
	r.AssertStderr("I exploded")
	r.AssertStdout("")
}
