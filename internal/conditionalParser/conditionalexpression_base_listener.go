// Code generated from conditionalexpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package conditionalparser // conditionalexpression
import "github.com/antlr4-go/antlr/v4"

// BaseconditionalexpressionListener is a complete listener for a parse tree produced by conditionalexpressionParser.
type BaseconditionalexpressionListener struct{}

var _ conditionalexpressionListener = &BaseconditionalexpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseconditionalexpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseconditionalexpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseconditionalexpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseconditionalexpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCondExpr is called when production condExpr is entered.
func (s *BaseconditionalexpressionListener) EnterCondExpr(ctx *CondExprContext) {}

// ExitCondExpr is called when production condExpr is exited.
func (s *BaseconditionalexpressionListener) ExitCondExpr(ctx *CondExprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseconditionalexpressionListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseconditionalexpressionListener) ExitExpr(ctx *ExprContext) {}

// EnterSuperTerm is called when production superTerm is entered.
func (s *BaseconditionalexpressionListener) EnterSuperTerm(ctx *SuperTermContext) {}

// ExitSuperTerm is called when production superTerm is exited.
func (s *BaseconditionalexpressionListener) ExitSuperTerm(ctx *SuperTermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseconditionalexpressionListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseconditionalexpressionListener) ExitFactor(ctx *FactorContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseconditionalexpressionListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseconditionalexpressionListener) ExitTerm(ctx *TermContext) {}

// EnterMemberofOp is called when production memberofOp is entered.
func (s *BaseconditionalexpressionListener) EnterMemberofOp(ctx *MemberofOpContext) {}

// ExitMemberofOp is called when production memberofOp is exited.
func (s *BaseconditionalexpressionListener) ExitMemberofOp(ctx *MemberofOpContext) {}

// EnterExistsOp is called when production existsOp is entered.
func (s *BaseconditionalexpressionListener) EnterExistsOp(ctx *ExistsOpContext) {}

// ExitExistsOp is called when production existsOp is exited.
func (s *BaseconditionalexpressionListener) ExitExistsOp(ctx *ExistsOpContext) {}

// EnterRelOp is called when production relOp is entered.
func (s *BaseconditionalexpressionListener) EnterRelOp(ctx *RelOpContext) {}

// ExitRelOp is called when production relOp is exited.
func (s *BaseconditionalexpressionListener) ExitRelOp(ctx *RelOpContext) {}

// EnterRelOp2 is called when production relOp2 is entered.
func (s *BaseconditionalexpressionListener) EnterRelOp2(ctx *RelOp2Context) {}

// ExitRelOp2 is called when production relOp2 is exited.
func (s *BaseconditionalexpressionListener) ExitRelOp2(ctx *RelOp2Context) {}

// EnterContainsOp is called when production containsOp is entered.
func (s *BaseconditionalexpressionListener) EnterContainsOp(ctx *ContainsOpContext) {}

// ExitContainsOp is called when production containsOp is exited.
func (s *BaseconditionalexpressionListener) ExitContainsOp(ctx *ContainsOpContext) {}

// EnterAnyofOp is called when production anyofOp is entered.
func (s *BaseconditionalexpressionListener) EnterAnyofOp(ctx *AnyofOpContext) {}

// ExitAnyofOp is called when production anyofOp is exited.
func (s *BaseconditionalexpressionListener) ExitAnyofOp(ctx *AnyofOpContext) {}

// EnterAttrName is called when production attrName is entered.
func (s *BaseconditionalexpressionListener) EnterAttrName(ctx *AttrNameContext) {}

// ExitAttrName is called when production attrName is exited.
func (s *BaseconditionalexpressionListener) ExitAttrName(ctx *AttrNameContext) {}

// EnterSimpleAttrName is called when production simpleAttrName is entered.
func (s *BaseconditionalexpressionListener) EnterSimpleAttrName(ctx *SimpleAttrNameContext) {}

// ExitSimpleAttrName is called when production simpleAttrName is exited.
func (s *BaseconditionalexpressionListener) ExitSimpleAttrName(ctx *SimpleAttrNameContext) {}

// EnterPrefixedAttrName is called when production prefixedAttrName is entered.
func (s *BaseconditionalexpressionListener) EnterPrefixedAttrName(ctx *PrefixedAttrNameContext) {}

// ExitPrefixedAttrName is called when production prefixedAttrName is exited.
func (s *BaseconditionalexpressionListener) ExitPrefixedAttrName(ctx *PrefixedAttrNameContext) {}

// EnterExtendedIdentifier is called when production extendedIdentifier is entered.
func (s *BaseconditionalexpressionListener) EnterExtendedIdentifier(ctx *ExtendedIdentifierContext) {}

// ExitExtendedIdentifier is called when production extendedIdentifier is exited.
func (s *BaseconditionalexpressionListener) ExitExtendedIdentifier(ctx *ExtendedIdentifierContext) {}

// EnterSidArray is called when production sidArray is entered.
func (s *BaseconditionalexpressionListener) EnterSidArray(ctx *SidArrayContext) {}

// ExitSidArray is called when production sidArray is exited.
func (s *BaseconditionalexpressionListener) ExitSidArray(ctx *SidArrayContext) {}

// EnterLiteralSID is called when production literalSID is entered.
func (s *BaseconditionalexpressionListener) EnterLiteralSID(ctx *LiteralSIDContext) {}

// ExitLiteralSID is called when production literalSID is exited.
func (s *BaseconditionalexpressionListener) ExitLiteralSID(ctx *LiteralSIDContext) {}

// EnterValueArray is called when production valueArray is entered.
func (s *BaseconditionalexpressionListener) EnterValueArray(ctx *ValueArrayContext) {}

// ExitValueArray is called when production valueArray is exited.
func (s *BaseconditionalexpressionListener) ExitValueArray(ctx *ValueArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseconditionalexpressionListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseconditionalexpressionListener) ExitValue(ctx *ValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseconditionalexpressionListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseconditionalexpressionListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterSidString is called when production sidString is entered.
func (s *BaseconditionalexpressionListener) EnterSidString(ctx *SidStringContext) {}

// ExitSidString is called when production sidString is exited.
func (s *BaseconditionalexpressionListener) ExitSidString(ctx *SidStringContext) {}
