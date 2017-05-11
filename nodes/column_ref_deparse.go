package pg_query

func (node ColumnRef) Deparse(_ ...string) string {
	return node.Fields.Deparse()
    // TODO end.join('.')
}
