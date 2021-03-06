// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
)

// ModuleVersion is the model entity for the ModuleVersion schema.
type ModuleVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Major holds the value of the "major" field.
	Major int `json:"major,omitempty"`
	// Minor holds the value of the "minor" field.
	Minor int `json:"minor,omitempty"`
	// Patch holds the value of the "patch" field.
	Patch int `json:"patch,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
}

// FromRows scans the sql response data into ModuleVersion.
func (mv *ModuleVersion) FromRows(rows *sql.Rows) error {
	var vmv struct {
		ID    int
		Major sql.NullInt64
		Minor sql.NullInt64
		Patch sql.NullInt64
		Tag   sql.NullString
	}
	// the order here should be the same as in the `moduleversion.Columns`.
	if err := rows.Scan(
		&vmv.ID,
		&vmv.Major,
		&vmv.Minor,
		&vmv.Patch,
		&vmv.Tag,
	); err != nil {
		return err
	}
	mv.ID = vmv.ID
	mv.Major = int(vmv.Major.Int64)
	mv.Minor = int(vmv.Minor.Int64)
	mv.Patch = int(vmv.Patch.Int64)
	mv.Tag = vmv.Tag.String
	return nil
}

// QueryModule queries the module edge of the ModuleVersion.
func (mv *ModuleVersion) QueryModule() *ModuleQuery {
	return (&ModuleVersionClient{mv.config}).QueryModule(mv)
}

// Update returns a builder for updating this ModuleVersion.
// Note that, you need to call ModuleVersion.Unwrap() before calling this method, if this ModuleVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (mv *ModuleVersion) Update() *ModuleVersionUpdateOne {
	return (&ModuleVersionClient{mv.config}).UpdateOne(mv)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (mv *ModuleVersion) Unwrap() *ModuleVersion {
	tx, ok := mv.config.driver.(*txDriver)
	if !ok {
		panic("ent: ModuleVersion is not a transactional entity")
	}
	mv.config.driver = tx.drv
	return mv
}

// String implements the fmt.Stringer.
func (mv *ModuleVersion) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("ModuleVersion(")
	buf.WriteString(fmt.Sprintf("id=%v", mv.ID))
	buf.WriteString(fmt.Sprintf(", major=%v", mv.Major))
	buf.WriteString(fmt.Sprintf(", minor=%v", mv.Minor))
	buf.WriteString(fmt.Sprintf(", patch=%v", mv.Patch))
	buf.WriteString(fmt.Sprintf(", tag=%v", mv.Tag))
	buf.WriteString(")")
	return buf.String()
}

// ModuleVersions is a parsable slice of ModuleVersion.
type ModuleVersions []*ModuleVersion

// FromRows scans the sql response data into ModuleVersions.
func (mv *ModuleVersions) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vmv := &ModuleVersion{}
		if err := vmv.FromRows(rows); err != nil {
			return err
		}
		*mv = append(*mv, vmv)
	}
	return nil
}

func (mv ModuleVersions) config(cfg config) {
	for _i := range mv {
		mv[_i].config = cfg
	}
}
