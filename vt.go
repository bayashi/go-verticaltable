package verticaltable

import (
	"fmt"
	"io"
	"strings"
)

type row struct {
	key   string
	value string
}

type table struct {
	header string
	rows   []*row
	maxLen int
}

// Vertical Table
type VT struct {
	out    io.Writer
	opt    *VTOptions
	tables []*table
}

type VTOptions struct {
	HeaderFormat  string
	ShowCount     bool
	CountFormat   string
	KvSeparator   string
	KeyAlignRight bool
}

func NewTable(writer io.Writer, opts ...*VTOptions) *VT {
	var opt *VTOptions
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = &VTOptions{
			HeaderFormat:  "********** %s **********",
			ShowCount:     true,
			CountFormat:   "%d. ",
			KvSeparator:   ": ",
			KeyAlignRight: true,
		}
	}

	v := &VT{
		out: writer,
		opt: opt,
	}

	return v
}

func (v *VT) Header(header string) *VT {
	tb := &table{
		header: header,
		maxLen: 0,
	}

	v.tables = append(v.tables, tb)

	return v
}

func (v *VT) Row(k string, val string) *VT {
	l := len(v.tables)

	if len(k) > v.tables[l-1].maxLen {
		v.tables[l-1].maxLen = len(k)
	}

	rows := v.tables[l-1].rows
	v.tables[l-1].rows = append(rows, &row{key: k, value: val})

	return v
}

func (v *VT) Render() error {
	_, err := fmt.Fprint(v.out, v.render())

	return err
}

func (v *VT) render() string {
	table := ""

	for i, tb := range v.tables {
		table += v.buildHeader(i, tb)
		for _, row := range tb.rows {
			table += v.buildRow(tb, row)
		}
	}

	return table
}

func (v *VT) buildHeader(i int, tb *table) string {
	count := ""
	if v.opt.ShowCount {
		count = fmt.Sprintf(v.opt.CountFormat, i+1)
	}

	return fmt.Sprintf(v.opt.HeaderFormat, count+tb.header) + "\n"
}

func (v *VT) buildRow(tb *table, row *row) string {
	key := ""
	if v.opt.KeyAlignRight {
		key = strings.Repeat(" ", tb.maxLen-len(row.key)) + row.key
	} else {
		key = row.key + strings.Repeat(" ", tb.maxLen-len(row.key))
	}

	return key + v.opt.KvSeparator + row.value + "\n"
}
