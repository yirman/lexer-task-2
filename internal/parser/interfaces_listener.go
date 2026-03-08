// Code generated from Interfaces.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Interfaces
import "github.com/antlr4-go/antlr/v4"

// InterfacesListener is a complete listener for a parse tree produced by InterfacesParser.
type InterfacesListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAuto_stmt is called when entering the auto_stmt production.
	EnterAuto_stmt(c *Auto_stmtContext)

	// EnterIface_stmt is called when entering the iface_stmt production.
	EnterIface_stmt(c *Iface_stmtContext)

	// EnterMapping_stmt is called when entering the mapping_stmt production.
	EnterMapping_stmt(c *Mapping_stmtContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterScript_def is called when entering the script_def production.
	EnterScript_def(c *Script_defContext)

	// EnterMap_def is called when entering the map_def production.
	EnterMap_def(c *Map_defContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAuto_stmt is called when exiting the auto_stmt production.
	ExitAuto_stmt(c *Auto_stmtContext)

	// ExitIface_stmt is called when exiting the iface_stmt production.
	ExitIface_stmt(c *Iface_stmtContext)

	// ExitMapping_stmt is called when exiting the mapping_stmt production.
	ExitMapping_stmt(c *Mapping_stmtContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitScript_def is called when exiting the script_def production.
	ExitScript_def(c *Script_defContext)

	// ExitMap_def is called when exiting the map_def production.
	ExitMap_def(c *Map_defContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
