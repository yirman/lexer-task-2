// Generated from ../Interfaces.g4 by ANTLR 4.13.2
// jshint ignore: start
import antlr4 from 'antlr4';
import InterfacesListener from './InterfacesListener.js';
const serializedATN = [4,1,20,87,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,
2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,1,0,4,0,22,8,0,11,0,12,0,23,1,0,
1,0,1,1,1,1,1,1,3,1,31,8,1,1,2,1,2,4,2,35,8,2,11,2,12,2,36,1,3,1,3,1,3,1,
3,1,3,5,3,44,8,3,10,3,12,3,47,9,3,1,4,1,4,1,4,1,4,5,4,53,8,4,10,4,12,4,56,
9,4,1,5,1,5,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,1,6,4,6,72,8,6,11,
6,12,6,73,3,6,76,8,6,1,7,1,7,1,7,1,8,1,8,1,8,1,8,1,9,1,9,1,9,0,0,10,0,2,
4,6,8,10,12,14,16,18,0,2,1,0,4,6,2,0,4,6,17,18,88,0,21,1,0,0,0,2,30,1,0,
0,0,4,32,1,0,0,0,6,38,1,0,0,0,8,48,1,0,0,0,10,57,1,0,0,0,12,75,1,0,0,0,14,
77,1,0,0,0,16,80,1,0,0,0,18,84,1,0,0,0,20,22,3,2,1,0,21,20,1,0,0,0,22,23,
1,0,0,0,23,21,1,0,0,0,23,24,1,0,0,0,24,25,1,0,0,0,25,26,5,0,0,1,26,1,1,0,
0,0,27,31,3,4,2,0,28,31,3,6,3,0,29,31,3,8,4,0,30,27,1,0,0,0,30,28,1,0,0,
0,30,29,1,0,0,0,31,3,1,0,0,0,32,34,5,1,0,0,33,35,5,17,0,0,34,33,1,0,0,0,
35,36,1,0,0,0,36,34,1,0,0,0,36,37,1,0,0,0,37,5,1,0,0,0,38,39,5,2,0,0,39,
40,5,17,0,0,40,41,5,3,0,0,41,45,3,10,5,0,42,44,3,12,6,0,43,42,1,0,0,0,44,
47,1,0,0,0,45,43,1,0,0,0,45,46,1,0,0,0,46,7,1,0,0,0,47,45,1,0,0,0,48,49,
5,13,0,0,49,50,5,17,0,0,50,54,3,14,7,0,51,53,3,16,8,0,52,51,1,0,0,0,53,56,
1,0,0,0,54,52,1,0,0,0,54,55,1,0,0,0,55,9,1,0,0,0,56,54,1,0,0,0,57,58,7,0,
0,0,58,11,1,0,0,0,59,60,5,7,0,0,60,76,5,16,0,0,61,62,5,8,0,0,62,76,5,16,
0,0,63,64,5,9,0,0,64,76,5,16,0,0,65,66,5,10,0,0,66,76,5,16,0,0,67,68,5,11,
0,0,68,76,5,16,0,0,69,71,5,12,0,0,70,72,5,16,0,0,71,70,1,0,0,0,72,73,1,0,
0,0,73,71,1,0,0,0,73,74,1,0,0,0,74,76,1,0,0,0,75,59,1,0,0,0,75,61,1,0,0,
0,75,63,1,0,0,0,75,65,1,0,0,0,75,67,1,0,0,0,75,69,1,0,0,0,76,13,1,0,0,0,
77,78,5,14,0,0,78,79,3,18,9,0,79,15,1,0,0,0,80,81,5,15,0,0,81,82,3,18,9,
0,82,83,3,18,9,0,83,17,1,0,0,0,84,85,7,1,0,0,85,19,1,0,0,0,7,23,30,36,45,
54,73,75];


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.atn.PredictionContextCache();

export default class InterfacesParser extends antlr4.Parser {

    static grammarFileName = "Interfaces.g4";
    static literalNames = [ null, "'auto'", "'iface'", null, "'static'", 
                            "'dhcp'", "'loopback'", "'address'", "'netmask'", 
                            "'gateway'", "'network'", "'broadcast'", "'dns-nameservers'", 
                            "'mapping'", "'script'", "'map'" ];
    static symbolicNames = [ null, "AUTO", "IFACE", "INET", "STATIC", "DHCP", 
                             "LOOPBACK", "ADDRESS", "NETMASK", "GATEWAY", 
                             "NETWORK", "BROADCAST", "DNS_NAMESERVERS", 
                             "MAPPING", "SCRIPT", "MAP", "IP", "ID", "STRING", 
                             "WS", "COMMENT" ];
    static ruleNames = [ "file", "statement", "auto_stmt", "iface_stmt", 
                         "mapping_stmt", "method", "option", "script_def", 
                         "map_def", "value" ];

    constructor(input) {
        super(input);
        this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
        this.ruleNames = InterfacesParser.ruleNames;
        this.literalNames = InterfacesParser.literalNames;
        this.symbolicNames = InterfacesParser.symbolicNames;
    }



	file() {
	    let localctx = new FileContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 0, InterfacesParser.RULE_file);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 21; 
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        do {
	            this.state = 20;
	            this.statement();
	            this.state = 23; 
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        } while((((_la) & ~0x1f) === 0 && ((1 << _la) & 8198) !== 0));
	        this.state = 25;
	        this.match(InterfacesParser.EOF);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	statement() {
	    let localctx = new StatementContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 2, InterfacesParser.RULE_statement);
	    try {
	        this.state = 30;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 27;
	            this.auto_stmt();
	            break;
	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 28;
	            this.iface_stmt();
	            break;
	        case 13:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 29;
	            this.mapping_stmt();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	auto_stmt() {
	    let localctx = new Auto_stmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 4, InterfacesParser.RULE_auto_stmt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 32;
	        this.match(InterfacesParser.AUTO);
	        this.state = 34; 
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        do {
	            this.state = 33;
	            this.match(InterfacesParser.ID);
	            this.state = 36; 
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        } while(_la===17);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	iface_stmt() {
	    let localctx = new Iface_stmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 6, InterfacesParser.RULE_iface_stmt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 38;
	        this.match(InterfacesParser.IFACE);
	        this.state = 39;
	        this.match(InterfacesParser.ID);
	        this.state = 40;
	        this.match(InterfacesParser.INET);
	        this.state = 41;
	        this.method();
	        this.state = 45;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 8064) !== 0)) {
	            this.state = 42;
	            this.option();
	            this.state = 47;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	mapping_stmt() {
	    let localctx = new Mapping_stmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 8, InterfacesParser.RULE_mapping_stmt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 48;
	        this.match(InterfacesParser.MAPPING);
	        this.state = 49;
	        this.match(InterfacesParser.ID);
	        this.state = 50;
	        this.script_def();
	        this.state = 54;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===15) {
	            this.state = 51;
	            this.map_def();
	            this.state = 56;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	method() {
	    let localctx = new MethodContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 10, InterfacesParser.RULE_method);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 57;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 112) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	option() {
	    let localctx = new OptionContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 12, InterfacesParser.RULE_option);
	    var _la = 0;
	    try {
	        this.state = 75;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 7:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 59;
	            this.match(InterfacesParser.ADDRESS);
	            this.state = 60;
	            this.match(InterfacesParser.IP);
	            break;
	        case 8:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 61;
	            this.match(InterfacesParser.NETMASK);
	            this.state = 62;
	            this.match(InterfacesParser.IP);
	            break;
	        case 9:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 63;
	            this.match(InterfacesParser.GATEWAY);
	            this.state = 64;
	            this.match(InterfacesParser.IP);
	            break;
	        case 10:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 65;
	            this.match(InterfacesParser.NETWORK);
	            this.state = 66;
	            this.match(InterfacesParser.IP);
	            break;
	        case 11:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 67;
	            this.match(InterfacesParser.BROADCAST);
	            this.state = 68;
	            this.match(InterfacesParser.IP);
	            break;
	        case 12:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 69;
	            this.match(InterfacesParser.DNS_NAMESERVERS);
	            this.state = 71; 
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            do {
	                this.state = 70;
	                this.match(InterfacesParser.IP);
	                this.state = 73; 
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            } while(_la===16);
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	script_def() {
	    let localctx = new Script_defContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 14, InterfacesParser.RULE_script_def);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 77;
	        this.match(InterfacesParser.SCRIPT);
	        this.state = 78;
	        this.value();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	map_def() {
	    let localctx = new Map_defContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 16, InterfacesParser.RULE_map_def);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 80;
	        this.match(InterfacesParser.MAP);
	        this.state = 81;
	        this.value();
	        this.state = 82;
	        this.value();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	value() {
	    let localctx = new ValueContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 18, InterfacesParser.RULE_value);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 84;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 393328) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}


}

InterfacesParser.EOF = antlr4.Token.EOF;
InterfacesParser.AUTO = 1;
InterfacesParser.IFACE = 2;
InterfacesParser.INET = 3;
InterfacesParser.STATIC = 4;
InterfacesParser.DHCP = 5;
InterfacesParser.LOOPBACK = 6;
InterfacesParser.ADDRESS = 7;
InterfacesParser.NETMASK = 8;
InterfacesParser.GATEWAY = 9;
InterfacesParser.NETWORK = 10;
InterfacesParser.BROADCAST = 11;
InterfacesParser.DNS_NAMESERVERS = 12;
InterfacesParser.MAPPING = 13;
InterfacesParser.SCRIPT = 14;
InterfacesParser.MAP = 15;
InterfacesParser.IP = 16;
InterfacesParser.ID = 17;
InterfacesParser.STRING = 18;
InterfacesParser.WS = 19;
InterfacesParser.COMMENT = 20;

InterfacesParser.RULE_file = 0;
InterfacesParser.RULE_statement = 1;
InterfacesParser.RULE_auto_stmt = 2;
InterfacesParser.RULE_iface_stmt = 3;
InterfacesParser.RULE_mapping_stmt = 4;
InterfacesParser.RULE_method = 5;
InterfacesParser.RULE_option = 6;
InterfacesParser.RULE_script_def = 7;
InterfacesParser.RULE_map_def = 8;
InterfacesParser.RULE_value = 9;

class FileContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_file;
    }

	EOF() {
	    return this.getToken(InterfacesParser.EOF, 0);
	};

	statement = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(StatementContext);
	    } else {
	        return this.getTypedRuleContext(StatementContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterFile(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitFile(this);
		}
	}


}



class StatementContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_statement;
    }

	auto_stmt() {
	    return this.getTypedRuleContext(Auto_stmtContext,0);
	};

	iface_stmt() {
	    return this.getTypedRuleContext(Iface_stmtContext,0);
	};

	mapping_stmt() {
	    return this.getTypedRuleContext(Mapping_stmtContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterStatement(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitStatement(this);
		}
	}


}



class Auto_stmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_auto_stmt;
    }

	AUTO() {
	    return this.getToken(InterfacesParser.AUTO, 0);
	};

	ID = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(InterfacesParser.ID);
	    } else {
	        return this.getToken(InterfacesParser.ID, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterAuto_stmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitAuto_stmt(this);
		}
	}


}



class Iface_stmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_iface_stmt;
    }

	IFACE() {
	    return this.getToken(InterfacesParser.IFACE, 0);
	};

	ID() {
	    return this.getToken(InterfacesParser.ID, 0);
	};

	INET() {
	    return this.getToken(InterfacesParser.INET, 0);
	};

	method() {
	    return this.getTypedRuleContext(MethodContext,0);
	};

	option = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(OptionContext);
	    } else {
	        return this.getTypedRuleContext(OptionContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterIface_stmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitIface_stmt(this);
		}
	}


}



class Mapping_stmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_mapping_stmt;
    }

	MAPPING() {
	    return this.getToken(InterfacesParser.MAPPING, 0);
	};

	ID() {
	    return this.getToken(InterfacesParser.ID, 0);
	};

	script_def() {
	    return this.getTypedRuleContext(Script_defContext,0);
	};

	map_def = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(Map_defContext);
	    } else {
	        return this.getTypedRuleContext(Map_defContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterMapping_stmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitMapping_stmt(this);
		}
	}


}



class MethodContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_method;
    }

	STATIC() {
	    return this.getToken(InterfacesParser.STATIC, 0);
	};

	DHCP() {
	    return this.getToken(InterfacesParser.DHCP, 0);
	};

	LOOPBACK() {
	    return this.getToken(InterfacesParser.LOOPBACK, 0);
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterMethod(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitMethod(this);
		}
	}


}



class OptionContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_option;
    }

	ADDRESS() {
	    return this.getToken(InterfacesParser.ADDRESS, 0);
	};

	IP = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(InterfacesParser.IP);
	    } else {
	        return this.getToken(InterfacesParser.IP, i);
	    }
	};


	NETMASK() {
	    return this.getToken(InterfacesParser.NETMASK, 0);
	};

	GATEWAY() {
	    return this.getToken(InterfacesParser.GATEWAY, 0);
	};

	NETWORK() {
	    return this.getToken(InterfacesParser.NETWORK, 0);
	};

	BROADCAST() {
	    return this.getToken(InterfacesParser.BROADCAST, 0);
	};

	DNS_NAMESERVERS() {
	    return this.getToken(InterfacesParser.DNS_NAMESERVERS, 0);
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterOption(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitOption(this);
		}
	}


}



class Script_defContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_script_def;
    }

	SCRIPT() {
	    return this.getToken(InterfacesParser.SCRIPT, 0);
	};

	value() {
	    return this.getTypedRuleContext(ValueContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterScript_def(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitScript_def(this);
		}
	}


}



class Map_defContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_map_def;
    }

	MAP() {
	    return this.getToken(InterfacesParser.MAP, 0);
	};

	value = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(ValueContext);
	    } else {
	        return this.getTypedRuleContext(ValueContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterMap_def(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitMap_def(this);
		}
	}


}



class ValueContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = InterfacesParser.RULE_value;
    }

	STRING() {
	    return this.getToken(InterfacesParser.STRING, 0);
	};

	ID() {
	    return this.getToken(InterfacesParser.ID, 0);
	};

	STATIC() {
	    return this.getToken(InterfacesParser.STATIC, 0);
	};

	DHCP() {
	    return this.getToken(InterfacesParser.DHCP, 0);
	};

	LOOPBACK() {
	    return this.getToken(InterfacesParser.LOOPBACK, 0);
	};

	enterRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.enterValue(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof InterfacesListener ) {
	        listener.exitValue(this);
		}
	}


}




InterfacesParser.FileContext = FileContext; 
InterfacesParser.StatementContext = StatementContext; 
InterfacesParser.Auto_stmtContext = Auto_stmtContext; 
InterfacesParser.Iface_stmtContext = Iface_stmtContext; 
InterfacesParser.Mapping_stmtContext = Mapping_stmtContext; 
InterfacesParser.MethodContext = MethodContext; 
InterfacesParser.OptionContext = OptionContext; 
InterfacesParser.Script_defContext = Script_defContext; 
InterfacesParser.Map_defContext = Map_defContext; 
InterfacesParser.ValueContext = ValueContext; 
