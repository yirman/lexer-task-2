# Generated from ../Interfaces.g4 by ANTLR 4.13.2
# encoding: utf-8
from antlr4 import *
from io import StringIO
import sys
if sys.version_info[1] > 5:
	from typing import TextIO
else:
	from typing.io import TextIO

def serializedATN():
    return [
        4,1,20,87,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,2,8,7,8,2,9,7,9,1,0,4,0,22,8,0,11,0,12,0,23,1,0,1,0,1,
        1,1,1,1,1,3,1,31,8,1,1,2,1,2,4,2,35,8,2,11,2,12,2,36,1,3,1,3,1,3,
        1,3,1,3,5,3,44,8,3,10,3,12,3,47,9,3,1,4,1,4,1,4,1,4,5,4,53,8,4,10,
        4,12,4,56,9,4,1,5,1,5,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,
        6,1,6,4,6,72,8,6,11,6,12,6,73,3,6,76,8,6,1,7,1,7,1,7,1,8,1,8,1,8,
        1,8,1,9,1,9,1,9,0,0,10,0,2,4,6,8,10,12,14,16,18,0,2,1,0,4,6,2,0,
        4,6,17,18,88,0,21,1,0,0,0,2,30,1,0,0,0,4,32,1,0,0,0,6,38,1,0,0,0,
        8,48,1,0,0,0,10,57,1,0,0,0,12,75,1,0,0,0,14,77,1,0,0,0,16,80,1,0,
        0,0,18,84,1,0,0,0,20,22,3,2,1,0,21,20,1,0,0,0,22,23,1,0,0,0,23,21,
        1,0,0,0,23,24,1,0,0,0,24,25,1,0,0,0,25,26,5,0,0,1,26,1,1,0,0,0,27,
        31,3,4,2,0,28,31,3,6,3,0,29,31,3,8,4,0,30,27,1,0,0,0,30,28,1,0,0,
        0,30,29,1,0,0,0,31,3,1,0,0,0,32,34,5,1,0,0,33,35,5,17,0,0,34,33,
        1,0,0,0,35,36,1,0,0,0,36,34,1,0,0,0,36,37,1,0,0,0,37,5,1,0,0,0,38,
        39,5,2,0,0,39,40,5,17,0,0,40,41,5,3,0,0,41,45,3,10,5,0,42,44,3,12,
        6,0,43,42,1,0,0,0,44,47,1,0,0,0,45,43,1,0,0,0,45,46,1,0,0,0,46,7,
        1,0,0,0,47,45,1,0,0,0,48,49,5,13,0,0,49,50,5,17,0,0,50,54,3,14,7,
        0,51,53,3,16,8,0,52,51,1,0,0,0,53,56,1,0,0,0,54,52,1,0,0,0,54,55,
        1,0,0,0,55,9,1,0,0,0,56,54,1,0,0,0,57,58,7,0,0,0,58,11,1,0,0,0,59,
        60,5,7,0,0,60,76,5,16,0,0,61,62,5,8,0,0,62,76,5,16,0,0,63,64,5,9,
        0,0,64,76,5,16,0,0,65,66,5,10,0,0,66,76,5,16,0,0,67,68,5,11,0,0,
        68,76,5,16,0,0,69,71,5,12,0,0,70,72,5,16,0,0,71,70,1,0,0,0,72,73,
        1,0,0,0,73,71,1,0,0,0,73,74,1,0,0,0,74,76,1,0,0,0,75,59,1,0,0,0,
        75,61,1,0,0,0,75,63,1,0,0,0,75,65,1,0,0,0,75,67,1,0,0,0,75,69,1,
        0,0,0,76,13,1,0,0,0,77,78,5,14,0,0,78,79,3,18,9,0,79,15,1,0,0,0,
        80,81,5,15,0,0,81,82,3,18,9,0,82,83,3,18,9,0,83,17,1,0,0,0,84,85,
        7,1,0,0,85,19,1,0,0,0,7,23,30,36,45,54,73,75
    ]

class InterfacesParser ( Parser ):

    grammarFileName = "Interfaces.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "'auto'", "'iface'", "<INVALID>", "'static'", 
                     "'dhcp'", "'loopback'", "'address'", "'netmask'", "'gateway'", 
                     "'network'", "'broadcast'", "'dns-nameservers'", "'mapping'", 
                     "'script'", "'map'" ]

    symbolicNames = [ "<INVALID>", "AUTO", "IFACE", "INET", "STATIC", "DHCP", 
                      "LOOPBACK", "ADDRESS", "NETMASK", "GATEWAY", "NETWORK", 
                      "BROADCAST", "DNS_NAMESERVERS", "MAPPING", "SCRIPT", 
                      "MAP", "IP", "ID", "STRING", "WS", "COMMENT" ]

    RULE_file = 0
    RULE_statement = 1
    RULE_auto_stmt = 2
    RULE_iface_stmt = 3
    RULE_mapping_stmt = 4
    RULE_method = 5
    RULE_option = 6
    RULE_script_def = 7
    RULE_map_def = 8
    RULE_value = 9

    ruleNames =  [ "file", "statement", "auto_stmt", "iface_stmt", "mapping_stmt", 
                   "method", "option", "script_def", "map_def", "value" ]

    EOF = Token.EOF
    AUTO=1
    IFACE=2
    INET=3
    STATIC=4
    DHCP=5
    LOOPBACK=6
    ADDRESS=7
    NETMASK=8
    GATEWAY=9
    NETWORK=10
    BROADCAST=11
    DNS_NAMESERVERS=12
    MAPPING=13
    SCRIPT=14
    MAP=15
    IP=16
    ID=17
    STRING=18
    WS=19
    COMMENT=20

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.13.2")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class FileContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def EOF(self):
            return self.getToken(InterfacesParser.EOF, 0)

        def statement(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(InterfacesParser.StatementContext)
            else:
                return self.getTypedRuleContext(InterfacesParser.StatementContext,i)


        def getRuleIndex(self):
            return InterfacesParser.RULE_file

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFile" ):
                listener.enterFile(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFile" ):
                listener.exitFile(self)




    def file_(self):

        localctx = InterfacesParser.FileContext(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_file)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 21 
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while True:
                self.state = 20
                self.statement()
                self.state = 23 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if not ((((_la) & ~0x3f) == 0 and ((1 << _la) & 8198) != 0)):
                    break

            self.state = 25
            self.match(InterfacesParser.EOF)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def auto_stmt(self):
            return self.getTypedRuleContext(InterfacesParser.Auto_stmtContext,0)


        def iface_stmt(self):
            return self.getTypedRuleContext(InterfacesParser.Iface_stmtContext,0)


        def mapping_stmt(self):
            return self.getTypedRuleContext(InterfacesParser.Mapping_stmtContext,0)


        def getRuleIndex(self):
            return InterfacesParser.RULE_statement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterStatement" ):
                listener.enterStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitStatement" ):
                listener.exitStatement(self)




    def statement(self):

        localctx = InterfacesParser.StatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_statement)
        try:
            self.state = 30
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [1]:
                self.enterOuterAlt(localctx, 1)
                self.state = 27
                self.auto_stmt()
                pass
            elif token in [2]:
                self.enterOuterAlt(localctx, 2)
                self.state = 28
                self.iface_stmt()
                pass
            elif token in [13]:
                self.enterOuterAlt(localctx, 3)
                self.state = 29
                self.mapping_stmt()
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Auto_stmtContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def AUTO(self):
            return self.getToken(InterfacesParser.AUTO, 0)

        def ID(self, i:int=None):
            if i is None:
                return self.getTokens(InterfacesParser.ID)
            else:
                return self.getToken(InterfacesParser.ID, i)

        def getRuleIndex(self):
            return InterfacesParser.RULE_auto_stmt

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAuto_stmt" ):
                listener.enterAuto_stmt(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAuto_stmt" ):
                listener.exitAuto_stmt(self)




    def auto_stmt(self):

        localctx = InterfacesParser.Auto_stmtContext(self, self._ctx, self.state)
        self.enterRule(localctx, 4, self.RULE_auto_stmt)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 32
            self.match(InterfacesParser.AUTO)
            self.state = 34 
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while True:
                self.state = 33
                self.match(InterfacesParser.ID)
                self.state = 36 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                if not (_la==17):
                    break

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Iface_stmtContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def IFACE(self):
            return self.getToken(InterfacesParser.IFACE, 0)

        def ID(self):
            return self.getToken(InterfacesParser.ID, 0)

        def INET(self):
            return self.getToken(InterfacesParser.INET, 0)

        def method(self):
            return self.getTypedRuleContext(InterfacesParser.MethodContext,0)


        def option(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(InterfacesParser.OptionContext)
            else:
                return self.getTypedRuleContext(InterfacesParser.OptionContext,i)


        def getRuleIndex(self):
            return InterfacesParser.RULE_iface_stmt

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterIface_stmt" ):
                listener.enterIface_stmt(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitIface_stmt" ):
                listener.exitIface_stmt(self)




    def iface_stmt(self):

        localctx = InterfacesParser.Iface_stmtContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_iface_stmt)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 38
            self.match(InterfacesParser.IFACE)
            self.state = 39
            self.match(InterfacesParser.ID)
            self.state = 40
            self.match(InterfacesParser.INET)
            self.state = 41
            self.method()
            self.state = 45
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while (((_la) & ~0x3f) == 0 and ((1 << _la) & 8064) != 0):
                self.state = 42
                self.option()
                self.state = 47
                self._errHandler.sync(self)
                _la = self._input.LA(1)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Mapping_stmtContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def MAPPING(self):
            return self.getToken(InterfacesParser.MAPPING, 0)

        def ID(self):
            return self.getToken(InterfacesParser.ID, 0)

        def script_def(self):
            return self.getTypedRuleContext(InterfacesParser.Script_defContext,0)


        def map_def(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(InterfacesParser.Map_defContext)
            else:
                return self.getTypedRuleContext(InterfacesParser.Map_defContext,i)


        def getRuleIndex(self):
            return InterfacesParser.RULE_mapping_stmt

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMapping_stmt" ):
                listener.enterMapping_stmt(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMapping_stmt" ):
                listener.exitMapping_stmt(self)




    def mapping_stmt(self):

        localctx = InterfacesParser.Mapping_stmtContext(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_mapping_stmt)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 48
            self.match(InterfacesParser.MAPPING)
            self.state = 49
            self.match(InterfacesParser.ID)
            self.state = 50
            self.script_def()
            self.state = 54
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while _la==15:
                self.state = 51
                self.map_def()
                self.state = 56
                self._errHandler.sync(self)
                _la = self._input.LA(1)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class MethodContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def STATIC(self):
            return self.getToken(InterfacesParser.STATIC, 0)

        def DHCP(self):
            return self.getToken(InterfacesParser.DHCP, 0)

        def LOOPBACK(self):
            return self.getToken(InterfacesParser.LOOPBACK, 0)

        def getRuleIndex(self):
            return InterfacesParser.RULE_method

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMethod" ):
                listener.enterMethod(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMethod" ):
                listener.exitMethod(self)




    def method(self):

        localctx = InterfacesParser.MethodContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_method)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 57
            _la = self._input.LA(1)
            if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 112) != 0)):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class OptionContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def ADDRESS(self):
            return self.getToken(InterfacesParser.ADDRESS, 0)

        def IP(self, i:int=None):
            if i is None:
                return self.getTokens(InterfacesParser.IP)
            else:
                return self.getToken(InterfacesParser.IP, i)

        def NETMASK(self):
            return self.getToken(InterfacesParser.NETMASK, 0)

        def GATEWAY(self):
            return self.getToken(InterfacesParser.GATEWAY, 0)

        def NETWORK(self):
            return self.getToken(InterfacesParser.NETWORK, 0)

        def BROADCAST(self):
            return self.getToken(InterfacesParser.BROADCAST, 0)

        def DNS_NAMESERVERS(self):
            return self.getToken(InterfacesParser.DNS_NAMESERVERS, 0)

        def getRuleIndex(self):
            return InterfacesParser.RULE_option

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterOption" ):
                listener.enterOption(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitOption" ):
                listener.exitOption(self)




    def option(self):

        localctx = InterfacesParser.OptionContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_option)
        self._la = 0 # Token type
        try:
            self.state = 75
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [7]:
                self.enterOuterAlt(localctx, 1)
                self.state = 59
                self.match(InterfacesParser.ADDRESS)
                self.state = 60
                self.match(InterfacesParser.IP)
                pass
            elif token in [8]:
                self.enterOuterAlt(localctx, 2)
                self.state = 61
                self.match(InterfacesParser.NETMASK)
                self.state = 62
                self.match(InterfacesParser.IP)
                pass
            elif token in [9]:
                self.enterOuterAlt(localctx, 3)
                self.state = 63
                self.match(InterfacesParser.GATEWAY)
                self.state = 64
                self.match(InterfacesParser.IP)
                pass
            elif token in [10]:
                self.enterOuterAlt(localctx, 4)
                self.state = 65
                self.match(InterfacesParser.NETWORK)
                self.state = 66
                self.match(InterfacesParser.IP)
                pass
            elif token in [11]:
                self.enterOuterAlt(localctx, 5)
                self.state = 67
                self.match(InterfacesParser.BROADCAST)
                self.state = 68
                self.match(InterfacesParser.IP)
                pass
            elif token in [12]:
                self.enterOuterAlt(localctx, 6)
                self.state = 69
                self.match(InterfacesParser.DNS_NAMESERVERS)
                self.state = 71 
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while True:
                    self.state = 70
                    self.match(InterfacesParser.IP)
                    self.state = 73 
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)
                    if not (_la==16):
                        break

                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Script_defContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def SCRIPT(self):
            return self.getToken(InterfacesParser.SCRIPT, 0)

        def value(self):
            return self.getTypedRuleContext(InterfacesParser.ValueContext,0)


        def getRuleIndex(self):
            return InterfacesParser.RULE_script_def

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterScript_def" ):
                listener.enterScript_def(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitScript_def" ):
                listener.exitScript_def(self)




    def script_def(self):

        localctx = InterfacesParser.Script_defContext(self, self._ctx, self.state)
        self.enterRule(localctx, 14, self.RULE_script_def)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 77
            self.match(InterfacesParser.SCRIPT)
            self.state = 78
            self.value()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Map_defContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def MAP(self):
            return self.getToken(InterfacesParser.MAP, 0)

        def value(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(InterfacesParser.ValueContext)
            else:
                return self.getTypedRuleContext(InterfacesParser.ValueContext,i)


        def getRuleIndex(self):
            return InterfacesParser.RULE_map_def

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMap_def" ):
                listener.enterMap_def(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMap_def" ):
                listener.exitMap_def(self)




    def map_def(self):

        localctx = InterfacesParser.Map_defContext(self, self._ctx, self.state)
        self.enterRule(localctx, 16, self.RULE_map_def)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 80
            self.match(InterfacesParser.MAP)
            self.state = 81
            self.value()
            self.state = 82
            self.value()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ValueContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def STRING(self):
            return self.getToken(InterfacesParser.STRING, 0)

        def ID(self):
            return self.getToken(InterfacesParser.ID, 0)

        def STATIC(self):
            return self.getToken(InterfacesParser.STATIC, 0)

        def DHCP(self):
            return self.getToken(InterfacesParser.DHCP, 0)

        def LOOPBACK(self):
            return self.getToken(InterfacesParser.LOOPBACK, 0)

        def getRuleIndex(self):
            return InterfacesParser.RULE_value

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterValue" ):
                listener.enterValue(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitValue" ):
                listener.exitValue(self)




    def value(self):

        localctx = InterfacesParser.ValueContext(self, self._ctx, self.state)
        self.enterRule(localctx, 18, self.RULE_value)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 84
            _la = self._input.LA(1)
            if not((((_la) & ~0x3f) == 0 and ((1 << _la) & 393328) != 0)):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx





