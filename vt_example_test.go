package verticaltable

import (
	"os"
)

func ExampleVT() {
	vt := NewTable(os.Stdout)
	buildTables(vt).Render()
	// Output:
	// ********** 1. foo **********
	//          ID: 123
	// Select Type: SIMPLE
	//       Table: foo
	// ********** 2. bar **********
	//          ID: 456
	// Select Type: UNIQUE
	//       Table: bar
}

func ExampleVTOptions() {
	vt := NewTable(os.Stdout, &VTOptions{
		HeaderFormat:  "--- %s",
		ShowCount:     false,
		KvSeparator:   " = ",
		KeyAlignRight: false,
	})
	buildTables(vt).Render()
	// Output:
	// --- foo
	// ID          = 123
	// Select Type = SIMPLE
	// Table       = foo
	// --- bar
	// ID          = 456
	// Select Type = UNIQUE
	// Table       = bar
}

func buildTables(vt *VT) *VT {
	vt.Header("foo")
	vt.Row("ID", "123")
	vt.Row("Select Type", "SIMPLE")
	vt.Row("Table", "foo")

	vt.Header("bar")
	vt.Row("ID", "456")
	vt.Row("Select Type", "UNIQUE")
	vt.Row("Table", "bar")

	return vt
}
