package fortune_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive // ok
	. "github.com/onsi/gomega"    //nolint:revive // ok
	"github.com/snivilised/fortune"
)

var _ = Describe("Greet", func() {
	Context("foo", func() {
		It("should: display", func() {
			Expect(fortune.Greet("carrie")).NotTo(BeEmpty())
		})
	})
})
