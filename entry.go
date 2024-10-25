package ascii_table

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Table struct {
	headersOrder  []string
	headersMaxLen *sync.Map
	items         map[string][]string
}

func New() *Table {
	return &Table{
		headersOrder:  nil,
		headersMaxLen: &sync.Map{},
		items:         make(map[string][]string),
	}
}

func (t *Table) AddHeader(header string) *Table {
	t.addHeader(header, 0)
	return t
}

func (t *Table) AddHeaders(headers ...string) *Table {
	for _, header := range headers {
		t.addHeader(header, 0)
	}
	return t
}

func (t *Table) AddItem(header string, item string) *Table {
	t.addHeader(header, len(item))
	t.items[header] = append(t.items[header], item)
	return t
}

func (t *Table) AddRow(row map[string]string) *Table {
	for header, item := range row {
		t.AddItem(header, item)
	}
	return t
}

func (t *Table) addHeader(header string, currlen int) {
	preLen, loaded := t.headersMaxLen.LoadOrStore(header, len(header))
	if !loaded {
		t.headersOrder = append(t.headersOrder, header)
	} else if currlen > preLen.(int) {
		t.headersMaxLen.Store(header, currlen)
	}
}

const (
	minHeaderLen = 8
	colSpacing   = 2
)

func (t *Table) String() string {
	if len(t.headersOrder) == 0 {
		return ""
	}
	sb := strings.Builder{}
	headersOrder := t.headersOrder

	var cellFormat = make(map[string]string)
	t.headersMaxLen.Range(func(key, value any) bool {
		l := int64(value.(int))
		if l < minHeaderLen {
			l = minHeaderLen
		}
		d := strconv.FormatInt(l+colSpacing, 10)
		cellFormat[key.(string)] = "%-" + d + "s"
		return true
	})

	// 打印表头
	for _, header := range headersOrder {
		sb.WriteString(
			fmt.Sprintf(cellFormat[header], header),
		)
	}
	sb.WriteString("\n")

	// 打印分隔线
	for _, header := range headersOrder {
		sb.WriteString(
			fmt.Sprintf(cellFormat[header], strings.Repeat("-", len(header))),
		)
	}
	sb.WriteString("\n")

	var end = true
	for rowIndex := 0; end; rowIndex++ {
		for _, header := range headersOrder {
			items := t.items[header]
			sb.WriteString(
				fmt.Sprintf(cellFormat[header], items[rowIndex]),
			)
			if rowIndex == len(items)-1 {
				end = false
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
