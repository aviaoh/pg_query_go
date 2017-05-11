package pg_query

func (node SelectStmt) Deparse(_ ...string) string {
    output := ""

    if node.Op == SETOP_UNION {
        output += node.Larg.Deparse()
        output += "UNION "
        if node.All {
            output += "ALL "
        }
        output += node.Rarg.Deparse() + " "
    }

    if node.WithClause != nil {
        output += node.WithClause.Deparse() + " "
    }

    if len(node.TargetList.Items) > 0 {
        output += "SELECT " + node.TargetList.Deparse(":select") + " "
    }

    if len(node.FromClause.Items) > 0 {
        output += "FROM " + node.FromClause.Deparse() + " "
    }

    if node.WhereClause != nil {
        output += "WHERE " + node.WhereClause.Deparse() + " "
    }

    if len(node.ValuesLists) > 0 {
        output += "VALUES "
        for valuesIdx, values := range node.ValuesLists {
            if valuesIdx > 0 {
                output += ", "
            }
            output += "("
            for idx, value := range values {
                if idx > 0 {
                    output += ", "
                }
                output += value.Deparse()
            }
            output += ") "
        }
    }

    if len(node.GroupClause.Items) > 0 {
        output += "GROUP BY " + node.GroupClause.Deparse() + " "
    }

    if node.HavingClause != nil {
        output += "HAVING " + node.HavingClause.Deparse() + " "
    }

    if len(node.SortClause.Items) > 0 {
        output += node.SortClause.Deparse() + " "
    }

    if node.LimitCount != nil {
        output += "LIMIT " + node.LimitCount.Deparse() + " "
    }

    if node.LimitOffset != nil {
        output += "OFFSET " + node.LimitOffset.Deparse() + " "
    }

    if len(node.LockingClause.Items) > 0 {
        output += node.LockingClause.Deparse() + " "
    }

    return output[:len(output)-1]
}
