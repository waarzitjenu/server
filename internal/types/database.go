package types

// Entry contains the database entry. The ID field is automatically incremented and
// Timestamp is indexed to allow fast queries based on time ranges
type Entry struct {
	ID        int    `storm:"id,increment"`
	Timestamp uint64 `storm:"index"`
	Data      LocationUpdate
}
