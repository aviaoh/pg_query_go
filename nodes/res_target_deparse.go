package pg_query

import "fmt"

func (node ResTarget) Deparse(context ...string) string {
    contextStr := ""
    if len(context) > 0 {
        contextStr = context[0]
    }

	if contextStr == ":select" {
        renaming := ""
        if node.Name != nil {
            renaming = " AS " + *node.Name
        }
        return node.Val.Deparse() +renaming
    } else if contextStr == ":update" {
        return *node.Name + " = " + node.Val.Deparse()
    } else if node.Val == nil {
        return *node.Name
    } else {
        panic(fmt.Errorf("Can't deparse %s in context %s", node, context))
    }
}
