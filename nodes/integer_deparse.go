package pg_query

import "strconv"

const DECIMAL_BASE = 10

func (node Integer) Deparse(_ ...string) string {
	return strconv.FormatInt(node.Ival, DECIMAL_BASE)
}
