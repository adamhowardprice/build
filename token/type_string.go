// generated by stringer -type Type; DO NOT EDIT

package token

import "fmt"

const _Type_name = "EOFErrorNewlineStringSpaceNumberFloatHexLeftCurlyRightCurlyLeftParenRightParenLeftBracRightBracQuoteEqualColonCommaSemicolonPeriodCommentPipeElipsisTrueFalseMultiLineStringTargetDeclFunc"

var _Type_index = [...]uint8{0, 3, 8, 15, 21, 26, 32, 37, 40, 49, 59, 68, 78, 86, 95, 100, 105, 110, 115, 124, 130, 137, 141, 148, 152, 157, 172, 182, 186}

func (i Type) String() string {
	if i < 0 || i+1 >= Type(len(_Type_index)) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
