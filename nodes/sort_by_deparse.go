// Auto-generated - DO NOT EDIT

package pg_query

func (node SortBy) Deparse(_ ...string) string {
	output := "ORDER BY "
	output += node.Node.Deparse()
    if node.SortbyDir == SORTBY_ASC {
        output += " ASC"
    } else if node.SortbyDir == SORTBY_DESC {
        output += " DESC"
    }
	return output
}
