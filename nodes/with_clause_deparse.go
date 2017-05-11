// Auto-generated - DO NOT EDIT

package pg_query

func (node WithClause) Deparse(_ ...string) string {
	output := "WITH "
    if node.Recursive {
		output += "RECURSIVE "
	}
    return output + node.Ctes.Deparse()
}
