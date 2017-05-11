// Auto-generated - DO NOT EDIT

package pg_query

func (node RangeVar) Deparse(_ ...string) string {
	output := ""
	if node.InhOpt == INH_NO {
		output += "ONLY "
	}
    output += "\"" + *node.Relname + "\""
    if node.Alias != nil {
        output += " " + node.Alias.Deparse()
    }
    return output
}
