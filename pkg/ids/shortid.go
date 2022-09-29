package ids

import (
	"time"

	"github.com/teris-io/shortid"
)

var (
	sidGen *shortid.Shortid
)

func init() {
	sidGen, _ = shortid.New(1, shortid.DefaultABC, uint64(time.Now().Unix()))
}

// ShortID returns the string representation of a newly generated shortid.ShortID
func ShortID() string {
	id, _ := sidGen.Generate()
	return id
}
