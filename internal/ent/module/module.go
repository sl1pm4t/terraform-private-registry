// Code generated by entc, DO NOT EDIT.

package module

import (
	"github.com/meringu/terraform-private-registry/internal/ent/schema"
)

const (
	// Label holds the string label denoting the module type in the database.
	Label = "module"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwner holds the string denoting the owner vertex property in the database.
	FieldOwner = "owner"
	// FieldNamespace holds the string denoting the namespace vertex property in the database.
	FieldNamespace = "namespace"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldProvider holds the string denoting the provider vertex property in the database.
	FieldProvider = "provider"
	// FieldDescription holds the string denoting the description vertex property in the database.
	FieldDescription = "description"
	// FieldSource holds the string denoting the source vertex property in the database.
	FieldSource = "source"
	// FieldDownloads holds the string denoting the downloads vertex property in the database.
	FieldDownloads = "downloads"
	// FieldPublishedAt holds the string denoting the published_at vertex property in the database.
	FieldPublishedAt = "published_at"
	// FieldInstallationID holds the string denoting the installation_id vertex property in the database.
	FieldInstallationID = "installation_id"
	// FieldAppID holds the string denoting the app_id vertex property in the database.
	FieldAppID = "app_id"
	// FieldRepoName holds the string denoting the repo_name vertex property in the database.
	FieldRepoName = "repo_name"

	// Table holds the table name of the module in the database.
	Table = "modules"
	// VersionTable is the table the holds the version relation/edge.
	VersionTable = "module_versions"
	// VersionInverseTable is the table name for the ModuleVersion entity.
	// It exists in this package in order to avoid circular dependency with the "moduleversion" package.
	VersionInverseTable = "module_versions"
	// VersionColumn is the table column denoting the version relation/edge.
	VersionColumn = "module_id"
)

// Columns holds all SQL columns are module fields.
var Columns = []string{
	FieldID,
	FieldOwner,
	FieldNamespace,
	FieldName,
	FieldProvider,
	FieldDescription,
	FieldSource,
	FieldDownloads,
	FieldPublishedAt,
	FieldInstallationID,
	FieldAppID,
	FieldRepoName,
}

var (
	fields = schema.Module{}.Fields()

	// descDownloads is the schema descriptor for downloads field.
	descDownloads = fields[6].Descriptor()
	// DefaultDownloads holds the default value on creation for the downloads field.
	DefaultDownloads = descDownloads.Default.(int64)
)