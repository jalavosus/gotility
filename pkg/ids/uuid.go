package ids

import (
	"github.com/google/uuid"
)

// UUID returns a uuidv4 string.
func UUID() string {
	return uuid.NewString()
}
