package pg_query

import "fmt"

func (node A_Expr) Deparse(context ...string) string {
    var output string
    switch node.Kind {
    case AEXPR_OP:
        output = deparseAExpr(node, context)
    //case AEXPR_OP_ANY: //TODO!!
    //    deparse_aexpr_any(node)
    //case AEXPR_IN:
    //    deparse_aexpr_in(node)
    //case CONSTR_TYPE_FOREIGN:
    //    deparse_aexpr_like(node)
    //case AEXPR_BETWEEN, AEXPR_NOT_BETWEEN, AEXPR_BETWEEN_SYM, AEXPR_NOT_BETWEEN_SYM:
    //    deparse_aexpr_between(node)
    //case AEXPR_NULLIF:
    //    deparse_aexpr_nullif(node)
    default:
        panic(fmt.Errorf("Can't deparse expresion: %+v", node))

    }
    return output
}

func deparseAExpr(expr A_Expr, context []string) string {
    output := ""
    contextStr := "true" //TODO
    if len(context) > 0 {
        contextStr = context[0]
    }

    output += expr.Lexpr.Deparse(contextStr)
    output += " " + expr.Name.Items[0].Deparse(":operator") + " "
    output += expr.Rexpr.Deparse(contextStr)
    if len(context) > 0 {
        // This is a nested expression, add parentheses
        output = "(" + output + ")"
    }
    return output
}
