# Generated from ../Interfaces.g4 by ANTLR 4.13.2
from antlr4 import *
if "." in __name__:
    from .InterfacesParser import InterfacesParser
else:
    from InterfacesParser import InterfacesParser

# This class defines a complete listener for a parse tree produced by InterfacesParser.
class InterfacesListener(ParseTreeListener):

    # Enter a parse tree produced by InterfacesParser#file.
    def enterFile(self, ctx:InterfacesParser.FileContext):
        pass

    # Exit a parse tree produced by InterfacesParser#file.
    def exitFile(self, ctx:InterfacesParser.FileContext):
        pass


    # Enter a parse tree produced by InterfacesParser#statement.
    def enterStatement(self, ctx:InterfacesParser.StatementContext):
        pass

    # Exit a parse tree produced by InterfacesParser#statement.
    def exitStatement(self, ctx:InterfacesParser.StatementContext):
        pass


    # Enter a parse tree produced by InterfacesParser#auto_stmt.
    def enterAuto_stmt(self, ctx:InterfacesParser.Auto_stmtContext):
        pass

    # Exit a parse tree produced by InterfacesParser#auto_stmt.
    def exitAuto_stmt(self, ctx:InterfacesParser.Auto_stmtContext):
        pass


    # Enter a parse tree produced by InterfacesParser#iface_stmt.
    def enterIface_stmt(self, ctx:InterfacesParser.Iface_stmtContext):
        pass

    # Exit a parse tree produced by InterfacesParser#iface_stmt.
    def exitIface_stmt(self, ctx:InterfacesParser.Iface_stmtContext):
        pass


    # Enter a parse tree produced by InterfacesParser#mapping_stmt.
    def enterMapping_stmt(self, ctx:InterfacesParser.Mapping_stmtContext):
        pass

    # Exit a parse tree produced by InterfacesParser#mapping_stmt.
    def exitMapping_stmt(self, ctx:InterfacesParser.Mapping_stmtContext):
        pass


    # Enter a parse tree produced by InterfacesParser#method.
    def enterMethod(self, ctx:InterfacesParser.MethodContext):
        pass

    # Exit a parse tree produced by InterfacesParser#method.
    def exitMethod(self, ctx:InterfacesParser.MethodContext):
        pass


    # Enter a parse tree produced by InterfacesParser#option.
    def enterOption(self, ctx:InterfacesParser.OptionContext):
        pass

    # Exit a parse tree produced by InterfacesParser#option.
    def exitOption(self, ctx:InterfacesParser.OptionContext):
        pass


    # Enter a parse tree produced by InterfacesParser#script_def.
    def enterScript_def(self, ctx:InterfacesParser.Script_defContext):
        pass

    # Exit a parse tree produced by InterfacesParser#script_def.
    def exitScript_def(self, ctx:InterfacesParser.Script_defContext):
        pass


    # Enter a parse tree produced by InterfacesParser#map_def.
    def enterMap_def(self, ctx:InterfacesParser.Map_defContext):
        pass

    # Exit a parse tree produced by InterfacesParser#map_def.
    def exitMap_def(self, ctx:InterfacesParser.Map_defContext):
        pass


    # Enter a parse tree produced by InterfacesParser#value.
    def enterValue(self, ctx:InterfacesParser.ValueContext):
        pass

    # Exit a parse tree produced by InterfacesParser#value.
    def exitValue(self, ctx:InterfacesParser.ValueContext):
        pass



del InterfacesParser