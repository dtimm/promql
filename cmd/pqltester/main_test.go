package main

import (
	"os"
	"os/exec"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
)

func TestBuild(t *testing.T) {
	var binPath string
	var err error

	t.Run("build", func(t *testing.T) {
		g := NewGomegaWithT(t)

		binPath, err = gexec.Build("./main.go")
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("returns 1 with no args", func(t *testing.T) {
		g := NewGomegaWithT(t)

		cmd := exec.Command(binPath)

		sess, err := gexec.Start(cmd, os.Stdout, os.Stderr)
		g.Expect(err).ToNot(HaveOccurred())

		g.Eventually(sess).Should(gexec.Exit(1))
	})

	t.Run("returns 0 with passing test and sample data", func(t *testing.T) {
		g := NewGomegaWithT(t)

		cmd := exec.Command(binPath, "--yml", "fixtures/test.yml", "--tests", "fixtures/test.pql")

		sess, err := gexec.Start(cmd, os.Stdout, os.Stderr)
		g.Expect(err).ToNot(HaveOccurred())

		g.Eventually(sess).Should(gexec.Exit(0))
	})

	gexec.CleanupBuildArtifacts()
}
