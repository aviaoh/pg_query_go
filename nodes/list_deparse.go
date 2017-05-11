package pg_query

func (node List) Deparse(context ...string) string {
    contextStr := ""
    if len(context) > 0 {
        contextStr = context[0]
    }
    output := ""
    for idx, item := range node.Items {
        if idx > 0 {
            output += ", "
        }
        output += item.Deparse(contextStr)
    }
    return output
}
