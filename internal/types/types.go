package types

// DatabaseEntry contains the database entry. The ID field is automatically incremented and
// Timestamp is indexed to allow fast queries based on time ranges
type DatabaseEntry struct {
	ID        int    `storm:"id,increment"`
	Timestamp uint64 `storm:"index"`
	Data      LocationUpdate
}
