package adapters_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/carpeliam/git-story-branch/adapters"
)

var _ = Describe("Git", func() {
	path := os.Getenv("PATH")

	BeforeEach(func() {
		os.Setenv("PATH", fmt.Sprintf("./fixtures:%s", path))
	})
	AfterEach(func() {
		os.Setenv("PATH", path)
	})

	It("should know the current branch", func() {
		branchName := adapters.NewRepository().GetBranchName()
		Expect(branchName).To(Equal("current-branch-123456789"))
	})
})
