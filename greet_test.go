package honour_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive // ok
	. "github.com/onsi/gomega"    //nolint:revive // ok
	"github.com/snivilised/honour"
)

var _ = Describe("Greet", func() {
	Context("foo", func() {
		It("should: display", func() {
			Expect(honour.Greet("carrie")).NotTo(BeEmpty())
		})
	})
})
