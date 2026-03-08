// Code generated from Interfaces.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Interfaces
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

type InterfacesParser struct {
	*antlr.BaseParser
}

var InterfacesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func interfacesParserInit() {
	staticData := &InterfacesParserStaticData
	staticData.LiteralNames = []string{
		"", "'auto'", "'iface'", "", "'static'", "'dhcp'", "'loopback'", "'address'",
		"'netmask'", "'gateway'", "'network'", "'broadcast'", "'dns-nameservers'",
		"'mapping'", "'script'", "'map'",
	}
	staticData.SymbolicNames = []string{
		"", "AUTO", "IFACE", "INET", "STATIC", "DHCP", "LOOPBACK", "ADDRESS",
		"NETMASK", "GATEWAY", "NETWORK", "BROADCAST", "DNS_NAMESERVERS", "MAPPING",
		"SCRIPT", "MAP", "IP", "ID", "STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"file", "statement", "auto_stmt", "iface_stmt", "mapping_stmt", "method",
		"option", "script_def", "map_def", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 20, 87, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 4,
		0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 31,
		8, 1, 1, 2, 1, 2, 4, 2, 35, 8, 2, 11, 2, 12, 2, 36, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 44, 8, 3, 10, 3, 12, 3, 47, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 53, 8, 4, 10, 4, 12, 4, 56, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 72, 8, 6, 11,
		6, 12, 6, 73, 3, 6, 76, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 2, 1,
		0, 4, 6, 2, 0, 4, 6, 17, 18, 88, 0, 21, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0,
		4, 32, 1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 57, 1, 0,
		0, 0, 12, 75, 1, 0, 0, 0, 14, 77, 1, 0, 0, 0, 16, 80, 1, 0, 0, 0, 18, 84,
		1, 0, 0, 0, 20, 22, 3, 2, 1, 0, 21, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0,
		23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 5,
		0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 31, 3, 4, 2, 0, 28, 31, 3, 6, 3, 0, 29,
		31, 3, 8, 4, 0, 30, 27, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 29, 1, 0, 0,
		0, 31, 3, 1, 0, 0, 0, 32, 34, 5, 1, 0, 0, 33, 35, 5, 17, 0, 0, 34, 33,
		1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0,
		37, 5, 1, 0, 0, 0, 38, 39, 5, 2, 0, 0, 39, 40, 5, 17, 0, 0, 40, 41, 5,
		3, 0, 0, 41, 45, 3, 10, 5, 0, 42, 44, 3, 12, 6, 0, 43, 42, 1, 0, 0, 0,
		44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 7, 1, 0,
		0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 5, 13, 0, 0, 49, 50, 5, 17, 0, 0, 50,
		54, 3, 14, 7, 0, 51, 53, 3, 16, 8, 0, 52, 51, 1, 0, 0, 0, 53, 56, 1, 0,
		0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 54,
		1, 0, 0, 0, 57, 58, 7, 0, 0, 0, 58, 11, 1, 0, 0, 0, 59, 60, 5, 7, 0, 0,
		60, 76, 5, 16, 0, 0, 61, 62, 5, 8, 0, 0, 62, 76, 5, 16, 0, 0, 63, 64, 5,
		9, 0, 0, 64, 76, 5, 16, 0, 0, 65, 66, 5, 10, 0, 0, 66, 76, 5, 16, 0, 0,
		67, 68, 5, 11, 0, 0, 68, 76, 5, 16, 0, 0, 69, 71, 5, 12, 0, 0, 70, 72,
		5, 16, 0, 0, 71, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 59, 1, 0, 0, 0, 75, 61, 1,
		0, 0, 0, 75, 63, 1, 0, 0, 0, 75, 65, 1, 0, 0, 0, 75, 67, 1, 0, 0, 0, 75,
		69, 1, 0, 0, 0, 76, 13, 1, 0, 0, 0, 77, 78, 5, 14, 0, 0, 78, 79, 3, 18,
		9, 0, 79, 15, 1, 0, 0, 0, 80, 81, 5, 15, 0, 0, 81, 82, 3, 18, 9, 0, 82,
		83, 3, 18, 9, 0, 83, 17, 1, 0, 0, 0, 84, 85, 7, 1, 0, 0, 85, 19, 1, 0,
		0, 0, 7, 23, 30, 36, 45, 54, 73, 75,
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

// InterfacesParserInit initializes any static state used to implement InterfacesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInterfacesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InterfacesParserInit() {
	staticData := &InterfacesParserStaticData
	staticData.once.Do(interfacesParserInit)
}

// NewInterfacesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInterfacesParser(input antlr.TokenStream) *InterfacesParser {
	InterfacesParserInit()
	this := new(InterfacesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &InterfacesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Interfaces.g4"

	return this
}

// InterfacesParser tokens.
const (
	InterfacesParserEOF             = antlr.TokenEOF
	InterfacesParserAUTO            = 1
	InterfacesParserIFACE           = 2
	InterfacesParserINET            = 3
	InterfacesParserSTATIC          = 4
	InterfacesParserDHCP            = 5
	InterfacesParserLOOPBACK        = 6
	InterfacesParserADDRESS         = 7
	InterfacesParserNETMASK         = 8
	InterfacesParserGATEWAY         = 9
	InterfacesParserNETWORK         = 10
	InterfacesParserBROADCAST       = 11
	InterfacesParserDNS_NAMESERVERS = 12
	InterfacesParserMAPPING         = 13
	InterfacesParserSCRIPT          = 14
	InterfacesParserMAP             = 15
	InterfacesParserIP              = 16
	InterfacesParserID              = 17
	InterfacesParserSTRING          = 18
	InterfacesParserWS              = 19
	InterfacesParserCOMMENT         = 20
)

// InterfacesParser rules.
const (
	InterfacesParserRULE_file         = 0
	InterfacesParserRULE_statement    = 1
	InterfacesParserRULE_auto_stmt    = 2
	InterfacesParserRULE_iface_stmt   = 3
	InterfacesParserRULE_mapping_stmt = 4
	InterfacesParserRULE_method       = 5
	InterfacesParserRULE_option       = 6
	InterfacesParserRULE_script_def   = 7
	InterfacesParserRULE_map_def      = 8
	InterfacesParserRULE_value        = 9
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(InterfacesParserEOF, 0)
}

func (s *FileContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *InterfacesParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InterfacesParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8198) != 0) {
		{
			p.SetState(20)
			p.Statement()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(InterfacesParserEOF)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Auto_stmt() IAuto_stmtContext
	Iface_stmt() IIface_stmtContext
	Mapping_stmt() IMapping_stmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Auto_stmt() IAuto_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuto_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuto_stmtContext)
}

func (s *StatementContext) Iface_stmt() IIface_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIface_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIface_stmtContext)
}

func (s *StatementContext) Mapping_stmt() IMapping_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapping_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapping_stmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *InterfacesParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InterfacesParserRULE_statement)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InterfacesParserAUTO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Auto_stmt()
		}

	case InterfacesParserIFACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Iface_stmt()
		}

	case InterfacesParserMAPPING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
			p.Mapping_stmt()
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

// IAuto_stmtContext is an interface to support dynamic dispatch.
type IAuto_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AUTO() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsAuto_stmtContext differentiates from other interfaces.
	IsAuto_stmtContext()
}

type Auto_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuto_stmtContext() *Auto_stmtContext {
	var p = new(Auto_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_auto_stmt
	return p
}

func InitEmptyAuto_stmtContext(p *Auto_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_auto_stmt
}

func (*Auto_stmtContext) IsAuto_stmtContext() {}

func NewAuto_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Auto_stmtContext {
	var p = new(Auto_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_auto_stmt

	return p
}

func (s *Auto_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Auto_stmtContext) AUTO() antlr.TerminalNode {
	return s.GetToken(InterfacesParserAUTO, 0)
}

func (s *Auto_stmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(InterfacesParserID)
}

func (s *Auto_stmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(InterfacesParserID, i)
}

func (s *Auto_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Auto_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Auto_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterAuto_stmt(s)
	}
}

func (s *Auto_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitAuto_stmt(s)
	}
}

func (p *InterfacesParser) Auto_stmt() (localctx IAuto_stmtContext) {
	localctx = NewAuto_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InterfacesParserRULE_auto_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(InterfacesParserAUTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == InterfacesParserID {
		{
			p.SetState(33)
			p.Match(InterfacesParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(36)
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

// IIface_stmtContext is an interface to support dynamic dispatch.
type IIface_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IFACE() antlr.TerminalNode
	ID() antlr.TerminalNode
	INET() antlr.TerminalNode
	Method() IMethodContext
	AllOption() []IOptionContext
	Option(i int) IOptionContext

	// IsIface_stmtContext differentiates from other interfaces.
	IsIface_stmtContext()
}

type Iface_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIface_stmtContext() *Iface_stmtContext {
	var p = new(Iface_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_iface_stmt
	return p
}

func InitEmptyIface_stmtContext(p *Iface_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_iface_stmt
}

func (*Iface_stmtContext) IsIface_stmtContext() {}

func NewIface_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iface_stmtContext {
	var p = new(Iface_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_iface_stmt

	return p
}

func (s *Iface_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Iface_stmtContext) IFACE() antlr.TerminalNode {
	return s.GetToken(InterfacesParserIFACE, 0)
}

func (s *Iface_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(InterfacesParserID, 0)
}

func (s *Iface_stmtContext) INET() antlr.TerminalNode {
	return s.GetToken(InterfacesParserINET, 0)
}

func (s *Iface_stmtContext) Method() IMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *Iface_stmtContext) AllOption() []IOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionContext); ok {
			len++
		}
	}

	tst := make([]IOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionContext); ok {
			tst[i] = t.(IOptionContext)
			i++
		}
	}

	return tst
}

func (s *Iface_stmtContext) Option(i int) IOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
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

	return t.(IOptionContext)
}

func (s *Iface_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iface_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iface_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterIface_stmt(s)
	}
}

func (s *Iface_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitIface_stmt(s)
	}
}

func (p *InterfacesParser) Iface_stmt() (localctx IIface_stmtContext) {
	localctx = NewIface_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InterfacesParserRULE_iface_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(InterfacesParserIFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(InterfacesParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(InterfacesParserINET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Method()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8064) != 0 {
		{
			p.SetState(42)
			p.Option()
		}

		p.SetState(47)
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

// IMapping_stmtContext is an interface to support dynamic dispatch.
type IMapping_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAPPING() antlr.TerminalNode
	ID() antlr.TerminalNode
	Script_def() IScript_defContext
	AllMap_def() []IMap_defContext
	Map_def(i int) IMap_defContext

	// IsMapping_stmtContext differentiates from other interfaces.
	IsMapping_stmtContext()
}

type Mapping_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapping_stmtContext() *Mapping_stmtContext {
	var p = new(Mapping_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_mapping_stmt
	return p
}

func InitEmptyMapping_stmtContext(p *Mapping_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_mapping_stmt
}

func (*Mapping_stmtContext) IsMapping_stmtContext() {}

func NewMapping_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mapping_stmtContext {
	var p = new(Mapping_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_mapping_stmt

	return p
}

func (s *Mapping_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Mapping_stmtContext) MAPPING() antlr.TerminalNode {
	return s.GetToken(InterfacesParserMAPPING, 0)
}

func (s *Mapping_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(InterfacesParserID, 0)
}

func (s *Mapping_stmtContext) Script_def() IScript_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScript_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScript_defContext)
}

func (s *Mapping_stmtContext) AllMap_def() []IMap_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMap_defContext); ok {
			len++
		}
	}

	tst := make([]IMap_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMap_defContext); ok {
			tst[i] = t.(IMap_defContext)
			i++
		}
	}

	return tst
}

func (s *Mapping_stmtContext) Map_def(i int) IMap_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMap_defContext); ok {
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

	return t.(IMap_defContext)
}

func (s *Mapping_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mapping_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mapping_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterMapping_stmt(s)
	}
}

func (s *Mapping_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitMapping_stmt(s)
	}
}

func (p *InterfacesParser) Mapping_stmt() (localctx IMapping_stmtContext) {
	localctx = NewMapping_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, InterfacesParserRULE_mapping_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(InterfacesParserMAPPING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(InterfacesParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Script_def()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InterfacesParserMAP {
		{
			p.SetState(51)
			p.Map_def()
		}

		p.SetState(56)
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

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STATIC() antlr.TerminalNode
	DHCP() antlr.TerminalNode
	LOOPBACK() antlr.TerminalNode

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_method
	return p
}

func InitEmptyMethodContext(p *MethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_method
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) STATIC() antlr.TerminalNode {
	return s.GetToken(InterfacesParserSTATIC, 0)
}

func (s *MethodContext) DHCP() antlr.TerminalNode {
	return s.GetToken(InterfacesParserDHCP, 0)
}

func (s *MethodContext) LOOPBACK() antlr.TerminalNode {
	return s.GetToken(InterfacesParserLOOPBACK, 0)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *InterfacesParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, InterfacesParserRULE_method)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&112) != 0) {
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

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADDRESS() antlr.TerminalNode
	AllIP() []antlr.TerminalNode
	IP(i int) antlr.TerminalNode
	NETMASK() antlr.TerminalNode
	GATEWAY() antlr.TerminalNode
	NETWORK() antlr.TerminalNode
	BROADCAST() antlr.TerminalNode
	DNS_NAMESERVERS() antlr.TerminalNode

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) ADDRESS() antlr.TerminalNode {
	return s.GetToken(InterfacesParserADDRESS, 0)
}

func (s *OptionContext) AllIP() []antlr.TerminalNode {
	return s.GetTokens(InterfacesParserIP)
}

func (s *OptionContext) IP(i int) antlr.TerminalNode {
	return s.GetToken(InterfacesParserIP, i)
}

func (s *OptionContext) NETMASK() antlr.TerminalNode {
	return s.GetToken(InterfacesParserNETMASK, 0)
}

func (s *OptionContext) GATEWAY() antlr.TerminalNode {
	return s.GetToken(InterfacesParserGATEWAY, 0)
}

func (s *OptionContext) NETWORK() antlr.TerminalNode {
	return s.GetToken(InterfacesParserNETWORK, 0)
}

func (s *OptionContext) BROADCAST() antlr.TerminalNode {
	return s.GetToken(InterfacesParserBROADCAST, 0)
}

func (s *OptionContext) DNS_NAMESERVERS() antlr.TerminalNode {
	return s.GetToken(InterfacesParserDNS_NAMESERVERS, 0)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *InterfacesParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, InterfacesParserRULE_option)
	var _la int

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InterfacesParserADDRESS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(InterfacesParserADDRESS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(InterfacesParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InterfacesParserNETMASK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(InterfacesParserNETMASK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.Match(InterfacesParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InterfacesParserGATEWAY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Match(InterfacesParserGATEWAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(InterfacesParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InterfacesParserNETWORK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Match(InterfacesParserNETWORK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(InterfacesParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InterfacesParserBROADCAST:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(67)
			p.Match(InterfacesParserBROADCAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(InterfacesParserIP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InterfacesParserDNS_NAMESERVERS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.Match(InterfacesParserDNS_NAMESERVERS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == InterfacesParserIP {
			{
				p.SetState(70)
				p.Match(InterfacesParserIP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IScript_defContext is an interface to support dynamic dispatch.
type IScript_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SCRIPT() antlr.TerminalNode
	Value() IValueContext

	// IsScript_defContext differentiates from other interfaces.
	IsScript_defContext()
}

type Script_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScript_defContext() *Script_defContext {
	var p = new(Script_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_script_def
	return p
}

func InitEmptyScript_defContext(p *Script_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_script_def
}

func (*Script_defContext) IsScript_defContext() {}

func NewScript_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Script_defContext {
	var p = new(Script_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_script_def

	return p
}

func (s *Script_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Script_defContext) SCRIPT() antlr.TerminalNode {
	return s.GetToken(InterfacesParserSCRIPT, 0)
}

func (s *Script_defContext) Value() IValueContext {
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

func (s *Script_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Script_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Script_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterScript_def(s)
	}
}

func (s *Script_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitScript_def(s)
	}
}

func (p *InterfacesParser) Script_def() (localctx IScript_defContext) {
	localctx = NewScript_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, InterfacesParserRULE_script_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(InterfacesParserSCRIPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Value()
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

// IMap_defContext is an interface to support dynamic dispatch.
type IMap_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsMap_defContext differentiates from other interfaces.
	IsMap_defContext()
}

type Map_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_defContext() *Map_defContext {
	var p = new(Map_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_map_def
	return p
}

func InitEmptyMap_defContext(p *Map_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_map_def
}

func (*Map_defContext) IsMap_defContext() {}

func NewMap_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_defContext {
	var p = new(Map_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_map_def

	return p
}

func (s *Map_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Map_defContext) MAP() antlr.TerminalNode {
	return s.GetToken(InterfacesParserMAP, 0)
}

func (s *Map_defContext) AllValue() []IValueContext {
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

func (s *Map_defContext) Value(i int) IValueContext {
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

func (s *Map_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterMap_def(s)
	}
}

func (s *Map_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitMap_def(s)
	}
}

func (p *InterfacesParser) Map_def() (localctx IMap_defContext) {
	localctx = NewMap_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, InterfacesParserRULE_map_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(InterfacesParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Value()
	}
	{
		p.SetState(82)
		p.Value()
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
	STRING() antlr.TerminalNode
	ID() antlr.TerminalNode
	STATIC() antlr.TerminalNode
	DHCP() antlr.TerminalNode
	LOOPBACK() antlr.TerminalNode

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
	p.RuleIndex = InterfacesParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InterfacesParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InterfacesParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(InterfacesParserSTRING, 0)
}

func (s *ValueContext) ID() antlr.TerminalNode {
	return s.GetToken(InterfacesParserID, 0)
}

func (s *ValueContext) STATIC() antlr.TerminalNode {
	return s.GetToken(InterfacesParserSTATIC, 0)
}

func (s *ValueContext) DHCP() antlr.TerminalNode {
	return s.GetToken(InterfacesParserDHCP, 0)
}

func (s *ValueContext) LOOPBACK() antlr.TerminalNode {
	return s.GetToken(InterfacesParserLOOPBACK, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InterfacesListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *InterfacesParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, InterfacesParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&393328) != 0) {
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
