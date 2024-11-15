package locale

// CLIENT-TODO: Should be updated to use url of the implementing project,
// so should not be left as astrolib. (this should be set by auto-check)
const AstrolibSourceID = "github.com/snivilised/fortune"

type honourTemplData struct{}

func (td honourTemplData) SourceID() string {
	return AstrolibSourceID
}
