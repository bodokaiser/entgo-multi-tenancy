// Code generated by ent, DO NOT EDIT.

package team

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "members"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "members"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "members"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "team_id"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"team_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/bodokaiser/entgo-multi-tenancy/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
)
