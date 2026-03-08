// Code generated from Interfaces.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Interfaces
import "github.com/antlr4-go/antlr/v4"

// BaseInterfacesListener is a complete listener for a parse tree produced by InterfacesParser.
type BaseInterfacesListener struct{}

var _ InterfacesListener = &BaseInterfacesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInterfacesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInterfacesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInterfacesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInterfacesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseInterfacesListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseInterfacesListener) ExitFile(ctx *FileContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseInterfacesListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseInterfacesListener) ExitStatement(ctx *StatementContext) {}

// EnterAuto_stmt is called when production auto_stmt is entered.
func (s *BaseInterfacesListener) EnterAuto_stmt(ctx *Auto_stmtContext) {}

// ExitAuto_stmt is called when production auto_stmt is exited.
func (s *BaseInterfacesListener) ExitAuto_stmt(ctx *Auto_stmtContext) {}

// EnterIface_stmt is called when production iface_stmt is entered.
func (s *BaseInterfacesListener) EnterIface_stmt(ctx *Iface_stmtContext) {}

// ExitIface_stmt is called when production iface_stmt is exited.
func (s *BaseInterfacesListener) ExitIface_stmt(ctx *Iface_stmtContext) {}

// EnterMapping_stmt is called when production mapping_stmt is entered.
func (s *BaseInterfacesListener) EnterMapping_stmt(ctx *Mapping_stmtContext) {}

// ExitMapping_stmt is called when production mapping_stmt is exited.
func (s *BaseInterfacesListener) ExitMapping_stmt(ctx *Mapping_stmtContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseInterfacesListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseInterfacesListener) ExitMethod(ctx *MethodContext) {}

// EnterOption is called when production option is entered.
func (s *BaseInterfacesListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseInterfacesListener) ExitOption(ctx *OptionContext) {}

// EnterScript_def is called when production script_def is entered.
func (s *BaseInterfacesListener) EnterScript_def(ctx *Script_defContext) {}

// ExitScript_def is called when production script_def is exited.
func (s *BaseInterfacesListener) ExitScript_def(ctx *Script_defContext) {}

// EnterMap_def is called when production map_def is entered.
func (s *BaseInterfacesListener) EnterMap_def(ctx *Map_defContext) {}

// ExitMap_def is called when production map_def is exited.
func (s *BaseInterfacesListener) ExitMap_def(ctx *Map_defContext) {}

// EnterValue is called when production value is entered.
func (s *BaseInterfacesListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseInterfacesListener) ExitValue(ctx *ValueContext) {}
