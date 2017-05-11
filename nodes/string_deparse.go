// Auto-generated - DO NOT EDIT

package pg_query

func (node String) Deparse(_ ...string) string {
	     //  if context == A_CONST
        //  format("'%s'", node['str'].gsub("'", "''"))
        //elsif [FUNC_CALL, TYPE_NAME, :operator, :defname_as].include?(context)
        //  node['str']
        //else
        //  format('"%s"', node['str'].gsub('"', '""'))
        //end
	return node.Str //TODO
}
