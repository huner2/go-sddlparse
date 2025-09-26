// Code generated from conditionalexpression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package conditionalparser // conditionalexpression
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type conditionalexpressionParser struct {
	*antlr.BaseParser
}

var ConditionalexpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func conditionalexpressionParserInit() {
	staticData := &ConditionalexpressionParserStaticData
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "", "'Member_of'", "'Not_Member_of'", "'Member_of_Any'",
		"'Not_Member_of_Any'", "'Device_Member_of'", "'Device_Member_of_Any'",
		"'Not_Device_Member_of'", "'Not_Device_Member_of_Any'", "'Exists'",
		"'Not_exists'", "'Contains'", "'Not_Contains'", "'Any_of'", "'Not_Any_of'",
		"'&&'", "'||'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
		"'('", "')'", "'{'", "'}'", "','", "'.'", "'@user.'", "'@device.'",
		"'@resource.'", "", "'SID('",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "WHITESPACE", "MEMBER_OF", "NOT_MEMBER_OF", "MEMBER_OF_ANY",
		"NOT_MEMBER_OF_ANY", "DEVICE_MEMBER_OF", "DEVICE_MEMBER_OF_ANY", "NOT_DEVICE_MEMBER_OF",
		"NOT_DEVICE_MEMBER_OF_ANY", "EXISTS", "NOT_EXISTS", "CONTAINS", "NOT_CONTAINS",
		"ANY_OF", "NOT_ANY_OF", "AND", "OR", "NOT", "EQ", "NEQ", "LT", "LTE",
		"GT", "GTE", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "COMMA", "DOT",
		"USER_PREFIX", "DEVICE_PREFIX", "RESOURCE_PREFIX", "SID_TOKEN", "SID_START",
		"HEX_INTEGER", "OCTAL_INTEGER", "DECIMAL_INTEGER", "STRING_LITERAL",
		"OCTET_STRING", "SID_FORMAT", "IDENTIFIER", "LIT_CHAR", "ANY_CHAR",
	}
	staticData.RuleNames = []string{
		"condExpr", "expr", "superTerm", "factor", "term", "memberofOp", "existsOp",
		"relOp", "relOp2", "containsOp", "anyofOp", "attrName", "simpleAttrName",
		"prefixedAttrName", "extendedIdentifier", "sidArray", "literalSID",
		"valueArray", "value", "integerValue", "sidString",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 167, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 49, 8, 1, 10, 1, 12, 1, 52, 9, 1,
		1, 2, 1, 2, 1, 2, 5, 2, 57, 8, 2, 10, 2, 12, 2, 60, 9, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 69, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 78, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 90, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 96, 8, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 102, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 108, 8, 10, 1, 11, 1, 11, 3, 11, 112, 8, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 3, 13, 119, 8, 13, 1, 14, 4, 14, 122, 8, 14, 11, 14, 12, 14,
		123, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 130, 8, 15, 10, 15, 12, 15, 133,
		9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 5, 17, 145, 8, 17, 10, 17, 12, 17, 148, 9, 17, 1, 17, 1, 17, 1, 17,
		3, 17, 153, 8, 17, 1, 18, 1, 18, 1, 18, 3, 18, 158, 8, 18, 1, 19, 3, 19,
		161, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 11,
		1, 0, 4, 11, 1, 0, 12, 13, 1, 0, 23, 26, 1, 0, 21, 22, 1, 0, 14, 15, 1,
		0, 16, 17, 1, 0, 33, 35, 1, 0, 44, 45, 1, 0, 1, 2, 1, 0, 38, 40, 2, 0,
		36, 36, 43, 43, 168, 0, 42, 1, 0, 0, 0, 2, 45, 1, 0, 0, 0, 4, 53, 1, 0,
		0, 0, 6, 68, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10, 79, 1, 0, 0, 0, 12, 82,
		1, 0, 0, 0, 14, 85, 1, 0, 0, 0, 16, 91, 1, 0, 0, 0, 18, 97, 1, 0, 0, 0,
		20, 103, 1, 0, 0, 0, 22, 111, 1, 0, 0, 0, 24, 113, 1, 0, 0, 0, 26, 115,
		1, 0, 0, 0, 28, 121, 1, 0, 0, 0, 30, 125, 1, 0, 0, 0, 32, 136, 1, 0, 0,
		0, 34, 152, 1, 0, 0, 0, 36, 157, 1, 0, 0, 0, 38, 160, 1, 0, 0, 0, 40, 164,
		1, 0, 0, 0, 42, 43, 3, 2, 1, 0, 43, 44, 5, 0, 0, 1, 44, 1, 1, 0, 0, 0,
		45, 50, 3, 4, 2, 0, 46, 47, 5, 19, 0, 0, 47, 49, 3, 4, 2, 0, 48, 46, 1,
		0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51,
		3, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 58, 3, 6, 3, 0, 54, 55, 5, 18, 0,
		0, 55, 57, 3, 6, 3, 0, 56, 54, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56,
		1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 5, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0,
		61, 69, 3, 8, 4, 0, 62, 63, 5, 27, 0, 0, 63, 64, 3, 2, 1, 0, 64, 65, 5,
		28, 0, 0, 65, 69, 1, 0, 0, 0, 66, 67, 5, 20, 0, 0, 67, 69, 3, 6, 3, 0,
		68, 61, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 7, 1, 0,
		0, 0, 70, 78, 3, 10, 5, 0, 71, 78, 3, 12, 6, 0, 72, 78, 3, 14, 7, 0, 73,
		78, 3, 16, 8, 0, 74, 78, 3, 18, 9, 0, 75, 78, 3, 20, 10, 0, 76, 78, 3,
		22, 11, 0, 77, 70, 1, 0, 0, 0, 77, 71, 1, 0, 0, 0, 77, 72, 1, 0, 0, 0,
		77, 73, 1, 0, 0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 76, 1,
		0, 0, 0, 78, 9, 1, 0, 0, 0, 79, 80, 7, 0, 0, 0, 80, 81, 3, 30, 15, 0, 81,
		11, 1, 0, 0, 0, 82, 83, 7, 1, 0, 0, 83, 84, 3, 22, 11, 0, 84, 13, 1, 0,
		0, 0, 85, 86, 3, 22, 11, 0, 86, 89, 7, 2, 0, 0, 87, 90, 3, 22, 11, 0, 88,
		90, 3, 36, 18, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 15, 1, 0,
		0, 0, 91, 92, 3, 22, 11, 0, 92, 95, 7, 3, 0, 0, 93, 96, 3, 22, 11, 0, 94,
		96, 3, 34, 17, 0, 95, 93, 1, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 17, 1, 0,
		0, 0, 97, 98, 3, 22, 11, 0, 98, 101, 7, 4, 0, 0, 99, 102, 3, 22, 11, 0,
		100, 102, 3, 34, 17, 0, 101, 99, 1, 0, 0, 0, 101, 100, 1, 0, 0, 0, 102,
		19, 1, 0, 0, 0, 103, 104, 3, 22, 11, 0, 104, 107, 7, 5, 0, 0, 105, 108,
		3, 22, 11, 0, 106, 108, 3, 34, 17, 0, 107, 105, 1, 0, 0, 0, 107, 106, 1,
		0, 0, 0, 108, 21, 1, 0, 0, 0, 109, 112, 3, 26, 13, 0, 110, 112, 3, 24,
		12, 0, 111, 109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 23, 1, 0, 0, 0,
		113, 114, 5, 44, 0, 0, 114, 25, 1, 0, 0, 0, 115, 118, 7, 6, 0, 0, 116,
		119, 5, 44, 0, 0, 117, 119, 3, 28, 14, 0, 118, 116, 1, 0, 0, 0, 118, 117,
		1, 0, 0, 0, 119, 27, 1, 0, 0, 0, 120, 122, 7, 7, 0, 0, 121, 120, 1, 0,
		0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0,
		124, 29, 1, 0, 0, 0, 125, 126, 5, 29, 0, 0, 126, 131, 3, 32, 16, 0, 127,
		128, 5, 31, 0, 0, 128, 130, 3, 32, 16, 0, 129, 127, 1, 0, 0, 0, 130, 133,
		1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0,
		0, 0, 133, 131, 1, 0, 0, 0, 134, 135, 5, 30, 0, 0, 135, 31, 1, 0, 0, 0,
		136, 137, 5, 37, 0, 0, 137, 138, 3, 40, 20, 0, 138, 139, 5, 28, 0, 0, 139,
		33, 1, 0, 0, 0, 140, 141, 5, 29, 0, 0, 141, 146, 3, 36, 18, 0, 142, 143,
		5, 31, 0, 0, 143, 145, 3, 36, 18, 0, 144, 142, 1, 0, 0, 0, 145, 148, 1,
		0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0,
		0, 148, 146, 1, 0, 0, 0, 149, 150, 5, 30, 0, 0, 150, 153, 1, 0, 0, 0, 151,
		153, 3, 36, 18, 0, 152, 140, 1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 35,
		1, 0, 0, 0, 154, 158, 3, 38, 19, 0, 155, 158, 5, 41, 0, 0, 156, 158, 5,
		42, 0, 0, 157, 154, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 156, 1, 0, 0,
		0, 158, 37, 1, 0, 0, 0, 159, 161, 7, 8, 0, 0, 160, 159, 1, 0, 0, 0, 160,
		161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 7, 9, 0, 0, 163, 39, 1,
		0, 0, 0, 164, 165, 7, 10, 0, 0, 165, 41, 1, 0, 0, 0, 16, 50, 58, 68, 77,
		89, 95, 101, 107, 111, 118, 123, 131, 146, 152, 157, 160,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// conditionalexpressionParserInit initializes any static state used to implement conditionalexpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewconditionalexpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConditionalexpressionParserInit() {
	staticData := &ConditionalexpressionParserStaticData
	staticData.once.Do(conditionalexpressionParserInit)
}

// NewconditionalexpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewconditionalexpressionParser(input antlr.TokenStream) *conditionalexpressionParser {
	ConditionalexpressionParserInit()
	this := new(conditionalexpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ConditionalexpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "conditionalexpression.g4"

	return this
}

// conditionalexpressionParser tokens.
const (
	conditionalexpressionParserEOF                      = antlr.TokenEOF
	conditionalexpressionParserT__0                     = 1
	conditionalexpressionParserT__1                     = 2
	conditionalexpressionParserWHITESPACE               = 3
	conditionalexpressionParserMEMBER_OF                = 4
	conditionalexpressionParserNOT_MEMBER_OF            = 5
	conditionalexpressionParserMEMBER_OF_ANY            = 6
	conditionalexpressionParserNOT_MEMBER_OF_ANY        = 7
	conditionalexpressionParserDEVICE_MEMBER_OF         = 8
	conditionalexpressionParserDEVICE_MEMBER_OF_ANY     = 9
	conditionalexpressionParserNOT_DEVICE_MEMBER_OF     = 10
	conditionalexpressionParserNOT_DEVICE_MEMBER_OF_ANY = 11
	conditionalexpressionParserEXISTS                   = 12
	conditionalexpressionParserNOT_EXISTS               = 13
	conditionalexpressionParserCONTAINS                 = 14
	conditionalexpressionParserNOT_CONTAINS             = 15
	conditionalexpressionParserANY_OF                   = 16
	conditionalexpressionParserNOT_ANY_OF               = 17
	conditionalexpressionParserAND                      = 18
	conditionalexpressionParserOR                       = 19
	conditionalexpressionParserNOT                      = 20
	conditionalexpressionParserEQ                       = 21
	conditionalexpressionParserNEQ                      = 22
	conditionalexpressionParserLT                       = 23
	conditionalexpressionParserLTE                      = 24
	conditionalexpressionParserGT                       = 25
	conditionalexpressionParserGTE                      = 26
	conditionalexpressionParserLPAREN                   = 27
	conditionalexpressionParserRPAREN                   = 28
	conditionalexpressionParserLBRACE                   = 29
	conditionalexpressionParserRBRACE                   = 30
	conditionalexpressionParserCOMMA                    = 31
	conditionalexpressionParserDOT                      = 32
	conditionalexpressionParserUSER_PREFIX              = 33
	conditionalexpressionParserDEVICE_PREFIX            = 34
	conditionalexpressionParserRESOURCE_PREFIX          = 35
	conditionalexpressionParserSID_TOKEN                = 36
	conditionalexpressionParserSID_START                = 37
	conditionalexpressionParserHEX_INTEGER              = 38
	conditionalexpressionParserOCTAL_INTEGER            = 39
	conditionalexpressionParserDECIMAL_INTEGER          = 40
	conditionalexpressionParserSTRING_LITERAL           = 41
	conditionalexpressionParserOCTET_STRING             = 42
	conditionalexpressionParserSID_FORMAT               = 43
	conditionalexpressionParserIDENTIFIER               = 44
	conditionalexpressionParserLIT_CHAR                 = 45
	conditionalexpressionParserANY_CHAR                 = 46
)

// conditionalexpressionParser rules.
const (
	conditionalexpressionParserRULE_condExpr           = 0
	conditionalexpressionParserRULE_expr               = 1
	conditionalexpressionParserRULE_superTerm          = 2
	conditionalexpressionParserRULE_factor             = 3
	conditionalexpressionParserRULE_term               = 4
	conditionalexpressionParserRULE_memberofOp         = 5
	conditionalexpressionParserRULE_existsOp           = 6
	conditionalexpressionParserRULE_relOp              = 7
	conditionalexpressionParserRULE_relOp2             = 8
	conditionalexpressionParserRULE_containsOp         = 9
	conditionalexpressionParserRULE_anyofOp            = 10
	conditionalexpressionParserRULE_attrName           = 11
	conditionalexpressionParserRULE_simpleAttrName     = 12
	conditionalexpressionParserRULE_prefixedAttrName   = 13
	conditionalexpressionParserRULE_extendedIdentifier = 14
	conditionalexpressionParserRULE_sidArray           = 15
	conditionalexpressionParserRULE_literalSID         = 16
	conditionalexpressionParserRULE_valueArray         = 17
	conditionalexpressionParserRULE_value              = 18
	conditionalexpressionParserRULE_integerValue       = 19
	conditionalexpressionParserRULE_sidString          = 20
)

// ICondExprContext is an interface to support dynamic dispatch.
type ICondExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsCondExprContext differentiates from other interfaces.
	IsCondExprContext()
}

type CondExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondExprContext() *CondExprContext {
	var p = new(CondExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_condExpr
	return p
}

func InitEmptyCondExprContext(p *CondExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_condExpr
}

func (*CondExprContext) IsCondExprContext() {}

func NewCondExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondExprContext {
	var p = new(CondExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_condExpr

	return p
}

func (s *CondExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CondExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CondExprContext) EOF() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserEOF, 0)
}

func (s *CondExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterCondExpr(s)
	}
}

func (s *CondExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitCondExpr(s)
	}
}

func (p *conditionalexpressionParser) CondExpr() (localctx ICondExprContext) {
	localctx = NewCondExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, conditionalexpressionParserRULE_condExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Expr()
	}
	{
		p.SetState(43)
		p.Match(conditionalexpressionParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSuperTerm() []ISuperTermContext
	SuperTerm(i int) ISuperTermContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllSuperTerm() []ISuperTermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISuperTermContext); ok {
			len++
		}
	}

	tst := make([]ISuperTermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISuperTermContext); ok {
			tst[i] = t.(ISuperTermContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) SuperTerm(i int) ISuperTermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuperTermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuperTermContext)
}

func (s *ExprContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(conditionalexpressionParserOR)
}

func (s *ExprContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserOR, i)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *conditionalexpressionParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, conditionalexpressionParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.SuperTerm()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == conditionalexpressionParserOR {
		{
			p.SetState(46)
			p.Match(conditionalexpressionParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.SuperTerm()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISuperTermContext is an interface to support dynamic dispatch.
type ISuperTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsSuperTermContext differentiates from other interfaces.
	IsSuperTermContext()
}

type SuperTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuperTermContext() *SuperTermContext {
	var p = new(SuperTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_superTerm
	return p
}

func InitEmptySuperTermContext(p *SuperTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_superTerm
}

func (*SuperTermContext) IsSuperTermContext() {}

func NewSuperTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuperTermContext {
	var p = new(SuperTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_superTerm

	return p
}

func (s *SuperTermContext) GetParser() antlr.Parser { return s.parser }

func (s *SuperTermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *SuperTermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *SuperTermContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(conditionalexpressionParserAND)
}

func (s *SuperTermContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserAND, i)
}

func (s *SuperTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuperTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterSuperTerm(s)
	}
}

func (s *SuperTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitSuperTerm(s)
	}
}

func (p *conditionalexpressionParser) SuperTerm() (localctx ISuperTermContext) {
	localctx = NewSuperTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, conditionalexpressionParserRULE_superTerm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Factor()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == conditionalexpressionParserAND {
		{
			p.SetState(54)
			p.Match(conditionalexpressionParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Factor()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Term() ITermContext
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	NOT() antlr.TerminalNode
	Factor() IFactorContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserLPAREN, 0)
}

func (s *FactorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserRPAREN, 0)
}

func (s *FactorContext) NOT() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT, 0)
}

func (s *FactorContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *conditionalexpressionParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, conditionalexpressionParserRULE_factor)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserMEMBER_OF, conditionalexpressionParserNOT_MEMBER_OF, conditionalexpressionParserMEMBER_OF_ANY, conditionalexpressionParserNOT_MEMBER_OF_ANY, conditionalexpressionParserDEVICE_MEMBER_OF, conditionalexpressionParserDEVICE_MEMBER_OF_ANY, conditionalexpressionParserNOT_DEVICE_MEMBER_OF, conditionalexpressionParserNOT_DEVICE_MEMBER_OF_ANY, conditionalexpressionParserEXISTS, conditionalexpressionParserNOT_EXISTS, conditionalexpressionParserUSER_PREFIX, conditionalexpressionParserDEVICE_PREFIX, conditionalexpressionParserRESOURCE_PREFIX, conditionalexpressionParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Term()
		}

	case conditionalexpressionParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(conditionalexpressionParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Expr()
		}
		{
			p.SetState(64)
			p.Match(conditionalexpressionParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case conditionalexpressionParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(conditionalexpressionParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Factor()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MemberofOp() IMemberofOpContext
	ExistsOp() IExistsOpContext
	RelOp() IRelOpContext
	RelOp2() IRelOp2Context
	ContainsOp() IContainsOpContext
	AnyofOp() IAnyofOpContext
	AttrName() IAttrNameContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) MemberofOp() IMemberofOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberofOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberofOpContext)
}

func (s *TermContext) ExistsOp() IExistsOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExistsOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExistsOpContext)
}

func (s *TermContext) RelOp() IRelOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelOpContext)
}

func (s *TermContext) RelOp2() IRelOp2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelOp2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelOp2Context)
}

func (s *TermContext) ContainsOp() IContainsOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainsOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainsOpContext)
}

func (s *TermContext) AnyofOp() IAnyofOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyofOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyofOpContext)
}

func (s *TermContext) AttrName() IAttrNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *conditionalexpressionParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, conditionalexpressionParserRULE_term)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.MemberofOp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.ExistsOp()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.RelOp()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.RelOp2()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(74)
			p.ContainsOp()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)
			p.AnyofOp()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(76)
			p.AttrName()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMemberofOpContext is an interface to support dynamic dispatch.
type IMemberofOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SidArray() ISidArrayContext
	MEMBER_OF() antlr.TerminalNode
	NOT_MEMBER_OF() antlr.TerminalNode
	MEMBER_OF_ANY() antlr.TerminalNode
	NOT_MEMBER_OF_ANY() antlr.TerminalNode
	DEVICE_MEMBER_OF() antlr.TerminalNode
	DEVICE_MEMBER_OF_ANY() antlr.TerminalNode
	NOT_DEVICE_MEMBER_OF() antlr.TerminalNode
	NOT_DEVICE_MEMBER_OF_ANY() antlr.TerminalNode

	// IsMemberofOpContext differentiates from other interfaces.
	IsMemberofOpContext()
}

type MemberofOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberofOpContext() *MemberofOpContext {
	var p = new(MemberofOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_memberofOp
	return p
}

func InitEmptyMemberofOpContext(p *MemberofOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_memberofOp
}

func (*MemberofOpContext) IsMemberofOpContext() {}

func NewMemberofOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberofOpContext {
	var p = new(MemberofOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_memberofOp

	return p
}

func (s *MemberofOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberofOpContext) SidArray() ISidArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISidArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISidArrayContext)
}

func (s *MemberofOpContext) MEMBER_OF() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserMEMBER_OF, 0)
}

func (s *MemberofOpContext) NOT_MEMBER_OF() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT_MEMBER_OF, 0)
}

func (s *MemberofOpContext) MEMBER_OF_ANY() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserMEMBER_OF_ANY, 0)
}

func (s *MemberofOpContext) NOT_MEMBER_OF_ANY() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT_MEMBER_OF_ANY, 0)
}

func (s *MemberofOpContext) DEVICE_MEMBER_OF() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserDEVICE_MEMBER_OF, 0)
}

func (s *MemberofOpContext) DEVICE_MEMBER_OF_ANY() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserDEVICE_MEMBER_OF_ANY, 0)
}

func (s *MemberofOpContext) NOT_DEVICE_MEMBER_OF() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT_DEVICE_MEMBER_OF, 0)
}

func (s *MemberofOpContext) NOT_DEVICE_MEMBER_OF_ANY() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT_DEVICE_MEMBER_OF_ANY, 0)
}

func (s *MemberofOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberofOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberofOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterMemberofOp(s)
	}
}

func (s *MemberofOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitMemberofOp(s)
	}
}

func (p *conditionalexpressionParser) MemberofOp() (localctx IMemberofOpContext) {
	localctx = NewMemberofOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, conditionalexpressionParserRULE_memberofOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4080) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(80)
		p.SidArray()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExistsOpContext is an interface to support dynamic dispatch.
type IExistsOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrName() IAttrNameContext
	EXISTS() antlr.TerminalNode
	NOT_EXISTS() antlr.TerminalNode

	// IsExistsOpContext differentiates from other interfaces.
	IsExistsOpContext()
}

type ExistsOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExistsOpContext() *ExistsOpContext {
	var p = new(ExistsOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_existsOp
	return p
}

func InitEmptyExistsOpContext(p *ExistsOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_existsOp
}

func (*ExistsOpContext) IsExistsOpContext() {}

func NewExistsOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExistsOpContext {
	var p = new(ExistsOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_existsOp

	return p
}

func (s *ExistsOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExistsOpContext) AttrName() IAttrNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *ExistsOpContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserEXISTS, 0)
}

func (s *ExistsOpContext) NOT_EXISTS() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT_EXISTS, 0)
}

func (s *ExistsOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExistsOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterExistsOp(s)
	}
}

func (s *ExistsOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitExistsOp(s)
	}
}

func (p *conditionalexpressionParser) ExistsOp() (localctx IExistsOpContext) {
	localctx = NewExistsOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, conditionalexpressionParserRULE_existsOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(_la == conditionalexpressionParserEXISTS || _la == conditionalexpressionParserNOT_EXISTS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(83)
		p.AttrName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelOpContext is an interface to support dynamic dispatch.
type IRelOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttrName() []IAttrNameContext
	AttrName(i int) IAttrNameContext
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	Value() IValueContext

	// IsRelOpContext differentiates from other interfaces.
	IsRelOpContext()
}

type RelOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelOpContext() *RelOpContext {
	var p = new(RelOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_relOp
	return p
}

func InitEmptyRelOpContext(p *RelOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_relOp
}

func (*RelOpContext) IsRelOpContext() {}

func NewRelOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelOpContext {
	var p = new(RelOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_relOp

	return p
}

func (s *RelOpContext) GetParser() antlr.Parser { return s.parser }

func (s *RelOpContext) AllAttrName() []IAttrNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrNameContext); ok {
			len++
		}
	}

	tst := make([]IAttrNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrNameContext); ok {
			tst[i] = t.(IAttrNameContext)
			i++
		}
	}

	return tst
}

func (s *RelOpContext) AttrName(i int) IAttrNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *RelOpContext) LT() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserLT, 0)
}

func (s *RelOpContext) LTE() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserLTE, 0)
}

func (s *RelOpContext) GT() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserGT, 0)
}

func (s *RelOpContext) GTE() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserGTE, 0)
}

func (s *RelOpContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *RelOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterRelOp(s)
	}
}

func (s *RelOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitRelOp(s)
	}
}

func (p *conditionalexpressionParser) RelOp() (localctx IRelOpContext) {
	localctx = NewRelOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, conditionalexpressionParserRULE_relOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.AttrName()
	}
	{
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&125829120) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserUSER_PREFIX, conditionalexpressionParserDEVICE_PREFIX, conditionalexpressionParserRESOURCE_PREFIX, conditionalexpressionParserIDENTIFIER:
		{
			p.SetState(87)
			p.AttrName()
		}

	case conditionalexpressionParserT__0, conditionalexpressionParserT__1, conditionalexpressionParserHEX_INTEGER, conditionalexpressionParserOCTAL_INTEGER, conditionalexpressionParserDECIMAL_INTEGER, conditionalexpressionParserSTRING_LITERAL, conditionalexpressionParserOCTET_STRING:
		{
			p.SetState(88)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelOp2Context is an interface to support dynamic dispatch.
type IRelOp2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttrName() []IAttrNameContext
	AttrName(i int) IAttrNameContext
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	ValueArray() IValueArrayContext

	// IsRelOp2Context differentiates from other interfaces.
	IsRelOp2Context()
}

type RelOp2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelOp2Context() *RelOp2Context {
	var p = new(RelOp2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_relOp2
	return p
}

func InitEmptyRelOp2Context(p *RelOp2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_relOp2
}

func (*RelOp2Context) IsRelOp2Context() {}

func NewRelOp2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelOp2Context {
	var p = new(RelOp2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_relOp2

	return p
}

func (s *RelOp2Context) GetParser() antlr.Parser { return s.parser }

func (s *RelOp2Context) AllAttrName() []IAttrNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrNameContext); ok {
			len++
		}
	}

	tst := make([]IAttrNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrNameContext); ok {
			tst[i] = t.(IAttrNameContext)
			i++
		}
	}

	return tst
}

func (s *RelOp2Context) AttrName(i int) IAttrNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *RelOp2Context) EQ() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserEQ, 0)
}

func (s *RelOp2Context) NEQ() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNEQ, 0)
}

func (s *RelOp2Context) ValueArray() IValueArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueArrayContext)
}

func (s *RelOp2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelOp2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelOp2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterRelOp2(s)
	}
}

func (s *RelOp2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitRelOp2(s)
	}
}

func (p *conditionalexpressionParser) RelOp2() (localctx IRelOp2Context) {
	localctx = NewRelOp2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, conditionalexpressionParserRULE_relOp2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.AttrName()
	}
	{
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !(_la == conditionalexpressionParserEQ || _la == conditionalexpressionParserNEQ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserUSER_PREFIX, conditionalexpressionParserDEVICE_PREFIX, conditionalexpressionParserRESOURCE_PREFIX, conditionalexpressionParserIDENTIFIER:
		{
			p.SetState(93)
			p.AttrName()
		}

	case conditionalexpressionParserT__0, conditionalexpressionParserT__1, conditionalexpressionParserLBRACE, conditionalexpressionParserHEX_INTEGER, conditionalexpressionParserOCTAL_INTEGER, conditionalexpressionParserDECIMAL_INTEGER, conditionalexpressionParserSTRING_LITERAL, conditionalexpressionParserOCTET_STRING:
		{
			p.SetState(94)
			p.ValueArray()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContainsOpContext is an interface to support dynamic dispatch.
type IContainsOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttrName() []IAttrNameContext
	AttrName(i int) IAttrNameContext
	CONTAINS() antlr.TerminalNode
	NOT_CONTAINS() antlr.TerminalNode
	ValueArray() IValueArrayContext

	// IsContainsOpContext differentiates from other interfaces.
	IsContainsOpContext()
}

type ContainsOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainsOpContext() *ContainsOpContext {
	var p = new(ContainsOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_containsOp
	return p
}

func InitEmptyContainsOpContext(p *ContainsOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_containsOp
}

func (*ContainsOpContext) IsContainsOpContext() {}

func NewContainsOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContainsOpContext {
	var p = new(ContainsOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_containsOp

	return p
}

func (s *ContainsOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ContainsOpContext) AllAttrName() []IAttrNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrNameContext); ok {
			len++
		}
	}

	tst := make([]IAttrNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrNameContext); ok {
			tst[i] = t.(IAttrNameContext)
			i++
		}
	}

	return tst
}

func (s *ContainsOpContext) AttrName(i int) IAttrNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *ContainsOpContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserCONTAINS, 0)
}

func (s *ContainsOpContext) NOT_CONTAINS() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT_CONTAINS, 0)
}

func (s *ContainsOpContext) ValueArray() IValueArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueArrayContext)
}

func (s *ContainsOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContainsOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterContainsOp(s)
	}
}

func (s *ContainsOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitContainsOp(s)
	}
}

func (p *conditionalexpressionParser) ContainsOp() (localctx IContainsOpContext) {
	localctx = NewContainsOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, conditionalexpressionParserRULE_containsOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.AttrName()
	}
	{
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !(_la == conditionalexpressionParserCONTAINS || _la == conditionalexpressionParserNOT_CONTAINS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserUSER_PREFIX, conditionalexpressionParserDEVICE_PREFIX, conditionalexpressionParserRESOURCE_PREFIX, conditionalexpressionParserIDENTIFIER:
		{
			p.SetState(99)
			p.AttrName()
		}

	case conditionalexpressionParserT__0, conditionalexpressionParserT__1, conditionalexpressionParserLBRACE, conditionalexpressionParserHEX_INTEGER, conditionalexpressionParserOCTAL_INTEGER, conditionalexpressionParserDECIMAL_INTEGER, conditionalexpressionParserSTRING_LITERAL, conditionalexpressionParserOCTET_STRING:
		{
			p.SetState(100)
			p.ValueArray()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnyofOpContext is an interface to support dynamic dispatch.
type IAnyofOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttrName() []IAttrNameContext
	AttrName(i int) IAttrNameContext
	ANY_OF() antlr.TerminalNode
	NOT_ANY_OF() antlr.TerminalNode
	ValueArray() IValueArrayContext

	// IsAnyofOpContext differentiates from other interfaces.
	IsAnyofOpContext()
}

type AnyofOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyofOpContext() *AnyofOpContext {
	var p = new(AnyofOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_anyofOp
	return p
}

func InitEmptyAnyofOpContext(p *AnyofOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_anyofOp
}

func (*AnyofOpContext) IsAnyofOpContext() {}

func NewAnyofOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyofOpContext {
	var p = new(AnyofOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_anyofOp

	return p
}

func (s *AnyofOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyofOpContext) AllAttrName() []IAttrNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrNameContext); ok {
			len++
		}
	}

	tst := make([]IAttrNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrNameContext); ok {
			tst[i] = t.(IAttrNameContext)
			i++
		}
	}

	return tst
}

func (s *AnyofOpContext) AttrName(i int) IAttrNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *AnyofOpContext) ANY_OF() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserANY_OF, 0)
}

func (s *AnyofOpContext) NOT_ANY_OF() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserNOT_ANY_OF, 0)
}

func (s *AnyofOpContext) ValueArray() IValueArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueArrayContext)
}

func (s *AnyofOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyofOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyofOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterAnyofOp(s)
	}
}

func (s *AnyofOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitAnyofOp(s)
	}
}

func (p *conditionalexpressionParser) AnyofOp() (localctx IAnyofOpContext) {
	localctx = NewAnyofOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, conditionalexpressionParserRULE_anyofOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.AttrName()
	}
	{
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !(_la == conditionalexpressionParserANY_OF || _la == conditionalexpressionParserNOT_ANY_OF) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserUSER_PREFIX, conditionalexpressionParserDEVICE_PREFIX, conditionalexpressionParserRESOURCE_PREFIX, conditionalexpressionParserIDENTIFIER:
		{
			p.SetState(105)
			p.AttrName()
		}

	case conditionalexpressionParserT__0, conditionalexpressionParserT__1, conditionalexpressionParserLBRACE, conditionalexpressionParserHEX_INTEGER, conditionalexpressionParserOCTAL_INTEGER, conditionalexpressionParserDECIMAL_INTEGER, conditionalexpressionParserSTRING_LITERAL, conditionalexpressionParserOCTET_STRING:
		{
			p.SetState(106)
			p.ValueArray()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrNameContext is an interface to support dynamic dispatch.
type IAttrNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrefixedAttrName() IPrefixedAttrNameContext
	SimpleAttrName() ISimpleAttrNameContext

	// IsAttrNameContext differentiates from other interfaces.
	IsAttrNameContext()
}

type AttrNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrNameContext() *AttrNameContext {
	var p = new(AttrNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_attrName
	return p
}

func InitEmptyAttrNameContext(p *AttrNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_attrName
}

func (*AttrNameContext) IsAttrNameContext() {}

func NewAttrNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrNameContext {
	var p = new(AttrNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_attrName

	return p
}

func (s *AttrNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrNameContext) PrefixedAttrName() IPrefixedAttrNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixedAttrNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixedAttrNameContext)
}

func (s *AttrNameContext) SimpleAttrName() ISimpleAttrNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleAttrNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleAttrNameContext)
}

func (s *AttrNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterAttrName(s)
	}
}

func (s *AttrNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitAttrName(s)
	}
}

func (p *conditionalexpressionParser) AttrName() (localctx IAttrNameContext) {
	localctx = NewAttrNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, conditionalexpressionParserRULE_attrName)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserUSER_PREFIX, conditionalexpressionParserDEVICE_PREFIX, conditionalexpressionParserRESOURCE_PREFIX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.PrefixedAttrName()
		}

	case conditionalexpressionParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.SimpleAttrName()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleAttrNameContext is an interface to support dynamic dispatch.
type ISimpleAttrNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsSimpleAttrNameContext differentiates from other interfaces.
	IsSimpleAttrNameContext()
}

type SimpleAttrNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleAttrNameContext() *SimpleAttrNameContext {
	var p = new(SimpleAttrNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_simpleAttrName
	return p
}

func InitEmptySimpleAttrNameContext(p *SimpleAttrNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_simpleAttrName
}

func (*SimpleAttrNameContext) IsSimpleAttrNameContext() {}

func NewSimpleAttrNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleAttrNameContext {
	var p = new(SimpleAttrNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_simpleAttrName

	return p
}

func (s *SimpleAttrNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleAttrNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserIDENTIFIER, 0)
}

func (s *SimpleAttrNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleAttrNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleAttrNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterSimpleAttrName(s)
	}
}

func (s *SimpleAttrNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitSimpleAttrName(s)
	}
}

func (p *conditionalexpressionParser) SimpleAttrName() (localctx ISimpleAttrNameContext) {
	localctx = NewSimpleAttrNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, conditionalexpressionParserRULE_simpleAttrName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(conditionalexpressionParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefixedAttrNameContext is an interface to support dynamic dispatch.
type IPrefixedAttrNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USER_PREFIX() antlr.TerminalNode
	DEVICE_PREFIX() antlr.TerminalNode
	RESOURCE_PREFIX() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ExtendedIdentifier() IExtendedIdentifierContext

	// IsPrefixedAttrNameContext differentiates from other interfaces.
	IsPrefixedAttrNameContext()
}

type PrefixedAttrNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixedAttrNameContext() *PrefixedAttrNameContext {
	var p = new(PrefixedAttrNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_prefixedAttrName
	return p
}

func InitEmptyPrefixedAttrNameContext(p *PrefixedAttrNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_prefixedAttrName
}

func (*PrefixedAttrNameContext) IsPrefixedAttrNameContext() {}

func NewPrefixedAttrNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixedAttrNameContext {
	var p = new(PrefixedAttrNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_prefixedAttrName

	return p
}

func (s *PrefixedAttrNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixedAttrNameContext) USER_PREFIX() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserUSER_PREFIX, 0)
}

func (s *PrefixedAttrNameContext) DEVICE_PREFIX() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserDEVICE_PREFIX, 0)
}

func (s *PrefixedAttrNameContext) RESOURCE_PREFIX() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserRESOURCE_PREFIX, 0)
}

func (s *PrefixedAttrNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserIDENTIFIER, 0)
}

func (s *PrefixedAttrNameContext) ExtendedIdentifier() IExtendedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendedIdentifierContext)
}

func (s *PrefixedAttrNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixedAttrNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixedAttrNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterPrefixedAttrName(s)
	}
}

func (s *PrefixedAttrNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitPrefixedAttrName(s)
	}
}

func (p *conditionalexpressionParser) PrefixedAttrName() (localctx IPrefixedAttrNameContext) {
	localctx = NewPrefixedAttrNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, conditionalexpressionParserRULE_prefixedAttrName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60129542144) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(116)
			p.Match(conditionalexpressionParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(117)
			p.ExtendedIdentifier()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtendedIdentifierContext is an interface to support dynamic dispatch.
type IExtendedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllLIT_CHAR() []antlr.TerminalNode
	LIT_CHAR(i int) antlr.TerminalNode

	// IsExtendedIdentifierContext differentiates from other interfaces.
	IsExtendedIdentifierContext()
}

type ExtendedIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendedIdentifierContext() *ExtendedIdentifierContext {
	var p = new(ExtendedIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_extendedIdentifier
	return p
}

func InitEmptyExtendedIdentifierContext(p *ExtendedIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_extendedIdentifier
}

func (*ExtendedIdentifierContext) IsExtendedIdentifierContext() {}

func NewExtendedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendedIdentifierContext {
	var p = new(ExtendedIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_extendedIdentifier

	return p
}

func (s *ExtendedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendedIdentifierContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(conditionalexpressionParserIDENTIFIER)
}

func (s *ExtendedIdentifierContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserIDENTIFIER, i)
}

func (s *ExtendedIdentifierContext) AllLIT_CHAR() []antlr.TerminalNode {
	return s.GetTokens(conditionalexpressionParserLIT_CHAR)
}

func (s *ExtendedIdentifierContext) LIT_CHAR(i int) antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserLIT_CHAR, i)
}

func (s *ExtendedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterExtendedIdentifier(s)
	}
}

func (s *ExtendedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitExtendedIdentifier(s)
	}
}

func (p *conditionalexpressionParser) ExtendedIdentifier() (localctx IExtendedIdentifierContext) {
	localctx = NewExtendedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, conditionalexpressionParserRULE_extendedIdentifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == conditionalexpressionParserIDENTIFIER || _la == conditionalexpressionParserLIT_CHAR {
		{
			p.SetState(120)
			_la = p.GetTokenStream().LA(1)

			if !(_la == conditionalexpressionParserIDENTIFIER || _la == conditionalexpressionParserLIT_CHAR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISidArrayContext is an interface to support dynamic dispatch.
type ISidArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	AllLiteralSID() []ILiteralSIDContext
	LiteralSID(i int) ILiteralSIDContext
	RBRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSidArrayContext differentiates from other interfaces.
	IsSidArrayContext()
}

type SidArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySidArrayContext() *SidArrayContext {
	var p = new(SidArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_sidArray
	return p
}

func InitEmptySidArrayContext(p *SidArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_sidArray
}

func (*SidArrayContext) IsSidArrayContext() {}

func NewSidArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SidArrayContext {
	var p = new(SidArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_sidArray

	return p
}

func (s *SidArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *SidArrayContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserLBRACE, 0)
}

func (s *SidArrayContext) AllLiteralSID() []ILiteralSIDContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralSIDContext); ok {
			len++
		}
	}

	tst := make([]ILiteralSIDContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralSIDContext); ok {
			tst[i] = t.(ILiteralSIDContext)
			i++
		}
	}

	return tst
}

func (s *SidArrayContext) LiteralSID(i int) ILiteralSIDContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralSIDContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralSIDContext)
}

func (s *SidArrayContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserRBRACE, 0)
}

func (s *SidArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(conditionalexpressionParserCOMMA)
}

func (s *SidArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserCOMMA, i)
}

func (s *SidArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SidArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SidArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterSidArray(s)
	}
}

func (s *SidArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitSidArray(s)
	}
}

func (p *conditionalexpressionParser) SidArray() (localctx ISidArrayContext) {
	localctx = NewSidArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, conditionalexpressionParserRULE_sidArray)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(conditionalexpressionParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.LiteralSID()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == conditionalexpressionParserCOMMA {
		{
			p.SetState(127)
			p.Match(conditionalexpressionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.LiteralSID()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(conditionalexpressionParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralSIDContext is an interface to support dynamic dispatch.
type ILiteralSIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SID_START() antlr.TerminalNode
	SidString() ISidStringContext
	RPAREN() antlr.TerminalNode

	// IsLiteralSIDContext differentiates from other interfaces.
	IsLiteralSIDContext()
}

type LiteralSIDContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralSIDContext() *LiteralSIDContext {
	var p = new(LiteralSIDContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_literalSID
	return p
}

func InitEmptyLiteralSIDContext(p *LiteralSIDContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_literalSID
}

func (*LiteralSIDContext) IsLiteralSIDContext() {}

func NewLiteralSIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralSIDContext {
	var p = new(LiteralSIDContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_literalSID

	return p
}

func (s *LiteralSIDContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralSIDContext) SID_START() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserSID_START, 0)
}

func (s *LiteralSIDContext) SidString() ISidStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISidStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISidStringContext)
}

func (s *LiteralSIDContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserRPAREN, 0)
}

func (s *LiteralSIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralSIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterLiteralSID(s)
	}
}

func (s *LiteralSIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitLiteralSID(s)
	}
}

func (p *conditionalexpressionParser) LiteralSID() (localctx ILiteralSIDContext) {
	localctx = NewLiteralSIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, conditionalexpressionParserRULE_literalSID)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(conditionalexpressionParserSID_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.SidString()
	}
	{
		p.SetState(138)
		p.Match(conditionalexpressionParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueArrayContext is an interface to support dynamic dispatch.
type IValueArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	RBRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValueArrayContext differentiates from other interfaces.
	IsValueArrayContext()
}

type ValueArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueArrayContext() *ValueArrayContext {
	var p = new(ValueArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_valueArray
	return p
}

func InitEmptyValueArrayContext(p *ValueArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_valueArray
}

func (*ValueArrayContext) IsValueArrayContext() {}

func NewValueArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueArrayContext {
	var p = new(ValueArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_valueArray

	return p
}

func (s *ValueArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueArrayContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserLBRACE, 0)
}

func (s *ValueArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ValueArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueArrayContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserRBRACE, 0)
}

func (s *ValueArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(conditionalexpressionParserCOMMA)
}

func (s *ValueArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserCOMMA, i)
}

func (s *ValueArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterValueArray(s)
	}
}

func (s *ValueArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitValueArray(s)
	}
}

func (p *conditionalexpressionParser) ValueArray() (localctx IValueArrayContext) {
	localctx = NewValueArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, conditionalexpressionParserRULE_valueArray)
	var _la int

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(conditionalexpressionParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Value()
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == conditionalexpressionParserCOMMA {
			{
				p.SetState(142)
				p.Match(conditionalexpressionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.Value()
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(149)
			p.Match(conditionalexpressionParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case conditionalexpressionParserT__0, conditionalexpressionParserT__1, conditionalexpressionParserHEX_INTEGER, conditionalexpressionParserOCTAL_INTEGER, conditionalexpressionParserDECIMAL_INTEGER, conditionalexpressionParserSTRING_LITERAL, conditionalexpressionParserOCTET_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerValue() IIntegerValueContext
	STRING_LITERAL() antlr.TerminalNode
	OCTET_STRING() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) IntegerValue() IIntegerValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *ValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserSTRING_LITERAL, 0)
}

func (s *ValueContext) OCTET_STRING() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserOCTET_STRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *conditionalexpressionParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, conditionalexpressionParserRULE_value)
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case conditionalexpressionParserT__0, conditionalexpressionParserT__1, conditionalexpressionParserHEX_INTEGER, conditionalexpressionParserOCTAL_INTEGER, conditionalexpressionParserDECIMAL_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.IntegerValue()
		}

	case conditionalexpressionParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Match(conditionalexpressionParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case conditionalexpressionParserOCTET_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)
			p.Match(conditionalexpressionParserOCTET_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerValueContext is an interface to support dynamic dispatch.
type IIntegerValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HEX_INTEGER() antlr.TerminalNode
	OCTAL_INTEGER() antlr.TerminalNode
	DECIMAL_INTEGER() antlr.TerminalNode

	// IsIntegerValueContext differentiates from other interfaces.
	IsIntegerValueContext()
}

type IntegerValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerValueContext() *IntegerValueContext {
	var p = new(IntegerValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_integerValue
	return p
}

func InitEmptyIntegerValueContext(p *IntegerValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_integerValue
}

func (*IntegerValueContext) IsIntegerValueContext() {}

func NewIntegerValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerValueContext {
	var p = new(IntegerValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_integerValue

	return p
}

func (s *IntegerValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerValueContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserHEX_INTEGER, 0)
}

func (s *IntegerValueContext) OCTAL_INTEGER() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserOCTAL_INTEGER, 0)
}

func (s *IntegerValueContext) DECIMAL_INTEGER() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserDECIMAL_INTEGER, 0)
}

func (s *IntegerValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterIntegerValue(s)
	}
}

func (s *IntegerValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitIntegerValue(s)
	}
}

func (p *conditionalexpressionParser) IntegerValue() (localctx IIntegerValueContext) {
	localctx = NewIntegerValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, conditionalexpressionParserRULE_integerValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == conditionalexpressionParserT__0 || _la == conditionalexpressionParserT__1 {
		{
			p.SetState(159)
			_la = p.GetTokenStream().LA(1)

			if !(_la == conditionalexpressionParserT__0 || _la == conditionalexpressionParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(162)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISidStringContext is an interface to support dynamic dispatch.
type ISidStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SID_TOKEN() antlr.TerminalNode
	SID_FORMAT() antlr.TerminalNode

	// IsSidStringContext differentiates from other interfaces.
	IsSidStringContext()
}

type SidStringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySidStringContext() *SidStringContext {
	var p = new(SidStringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_sidString
	return p
}

func InitEmptySidStringContext(p *SidStringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = conditionalexpressionParserRULE_sidString
}

func (*SidStringContext) IsSidStringContext() {}

func NewSidStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SidStringContext {
	var p = new(SidStringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = conditionalexpressionParserRULE_sidString

	return p
}

func (s *SidStringContext) GetParser() antlr.Parser { return s.parser }

func (s *SidStringContext) SID_TOKEN() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserSID_TOKEN, 0)
}

func (s *SidStringContext) SID_FORMAT() antlr.TerminalNode {
	return s.GetToken(conditionalexpressionParserSID_FORMAT, 0)
}

func (s *SidStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SidStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SidStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.EnterSidString(s)
	}
}

func (s *SidStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(conditionalexpressionListener); ok {
		listenerT.ExitSidString(s)
	}
}

func (p *conditionalexpressionParser) SidString() (localctx ISidStringContext) {
	localctx = NewSidStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, conditionalexpressionParserRULE_sidString)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		_la = p.GetTokenStream().LA(1)

		if !(_la == conditionalexpressionParserSID_TOKEN || _la == conditionalexpressionParserSID_FORMAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
