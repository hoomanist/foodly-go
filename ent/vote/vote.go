// Code generated by entc, DO NOT EDIT.

package vote

const (
	// Label holds the string label denoting the vote type in the database.
	Label = "vote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldRestaurant holds the string denoting the restaurant field in the database.
	FieldRestaurant = "restaurant"
	// FieldFood holds the string denoting the food field in the database.
	FieldFood = "food"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"

	// Table holds the table name of the vote in the database.
	Table = "votes"
)

// Columns holds all SQL columns for vote fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldRestaurant,
	FieldFood,
	FieldUsername,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
