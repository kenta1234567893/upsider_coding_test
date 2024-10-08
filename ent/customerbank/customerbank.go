// Code generated by ent, DO NOT EDIT.

package customerbank

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the customerbank type in the database.
	Label = "customer_bank"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldBankName holds the string denoting the bank_name field in the database.
	FieldBankName = "bank_name"
	// FieldBranchName holds the string denoting the branch_name field in the database.
	FieldBranchName = "branch_name"
	// FieldBankAccountNumber holds the string denoting the bank_account_number field in the database.
	FieldBankAccountNumber = "bank_account_number"
	// FieldBankAccountName holds the string denoting the bank_account_name field in the database.
	FieldBankAccountName = "bank_account_name"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// Table holds the table name of the customerbank in the database.
	Table = "customer_banks"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "customer_banks"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_customer_banks"
)

// Columns holds all SQL columns for customerbank fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldBankName,
	FieldBranchName,
	FieldBankAccountNumber,
	FieldBankAccountName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "customer_banks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"customer_customer_banks",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the CustomerBank queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBankName orders the results by the bank_name field.
func ByBankName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankName, opts...).ToFunc()
}

// ByBranchName orders the results by the branch_name field.
func ByBranchName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBranchName, opts...).ToFunc()
}

// ByBankAccountNumber orders the results by the bank_account_number field.
func ByBankAccountNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankAccountNumber, opts...).ToFunc()
}

// ByBankAccountName orders the results by the bank_account_name field.
func ByBankAccountName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankAccountName, opts...).ToFunc()
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}
