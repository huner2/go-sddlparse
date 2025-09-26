// Code generated from conditionalexpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package conditionalparser // conditionalexpression
import "github.com/antlr4-go/antlr/v4"

// conditionalexpressionListener is a complete listener for a parse tree produced by conditionalexpressionParser.
type conditionalexpressionListener interface {
	antlr.ParseTreeListener

	// EnterCondExpr is called when entering the condExpr production.
	EnterCondExpr(c *CondExprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterSuperTerm is called when entering the superTerm production.
	EnterSuperTerm(c *SuperTermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMemberofOp is called when entering the memberofOp production.
	EnterMemberofOp(c *MemberofOpContext)

	// EnterExistsOp is called when entering the existsOp production.
	EnterExistsOp(c *ExistsOpContext)

	// EnterRelOp is called when entering the relOp production.
	EnterRelOp(c *RelOpContext)

	// EnterRelOp2 is called when entering the relOp2 production.
	EnterRelOp2(c *RelOp2Context)

	// EnterContainsOp is called when entering the containsOp production.
	EnterContainsOp(c *ContainsOpContext)

	// EnterAnyofOp is called when entering the anyofOp production.
	EnterAnyofOp(c *AnyofOpContext)

	// EnterAttrName is called when entering the attrName production.
	EnterAttrName(c *AttrNameContext)

	// EnterSimpleAttrName is called when entering the simpleAttrName production.
	EnterSimpleAttrName(c *SimpleAttrNameContext)

	// EnterPrefixedAttrName is called when entering the prefixedAttrName production.
	EnterPrefixedAttrName(c *PrefixedAttrNameContext)

	// EnterExtendedIdentifier is called when entering the extendedIdentifier production.
	EnterExtendedIdentifier(c *ExtendedIdentifierContext)

	// EnterSidArray is called when entering the sidArray production.
	EnterSidArray(c *SidArrayContext)

	// EnterLiteralSID is called when entering the literalSID production.
	EnterLiteralSID(c *LiteralSIDContext)

	// EnterValueArray is called when entering the valueArray production.
	EnterValueArray(c *ValueArrayContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterSidString is called when entering the sidString production.
	EnterSidString(c *SidStringContext)

	// ExitCondExpr is called when exiting the condExpr production.
	ExitCondExpr(c *CondExprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitSuperTerm is called when exiting the superTerm production.
	ExitSuperTerm(c *SuperTermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMemberofOp is called when exiting the memberofOp production.
	ExitMemberofOp(c *MemberofOpContext)

	// ExitExistsOp is called when exiting the existsOp production.
	ExitExistsOp(c *ExistsOpContext)

	// ExitRelOp is called when exiting the relOp production.
	ExitRelOp(c *RelOpContext)

	// ExitRelOp2 is called when exiting the relOp2 production.
	ExitRelOp2(c *RelOp2Context)

	// ExitContainsOp is called when exiting the containsOp production.
	ExitContainsOp(c *ContainsOpContext)

	// ExitAnyofOp is called when exiting the anyofOp production.
	ExitAnyofOp(c *AnyofOpContext)

	// ExitAttrName is called when exiting the attrName production.
	ExitAttrName(c *AttrNameContext)

	// ExitSimpleAttrName is called when exiting the simpleAttrName production.
	ExitSimpleAttrName(c *SimpleAttrNameContext)

	// ExitPrefixedAttrName is called when exiting the prefixedAttrName production.
	ExitPrefixedAttrName(c *PrefixedAttrNameContext)

	// ExitExtendedIdentifier is called when exiting the extendedIdentifier production.
	ExitExtendedIdentifier(c *ExtendedIdentifierContext)

	// ExitSidArray is called when exiting the sidArray production.
	ExitSidArray(c *SidArrayContext)

	// ExitLiteralSID is called when exiting the literalSID production.
	ExitLiteralSID(c *LiteralSIDContext)

	// ExitValueArray is called when exiting the valueArray production.
	ExitValueArray(c *ValueArrayContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitSidString is called when exiting the sidString production.
	ExitSidString(c *SidStringContext)
}
