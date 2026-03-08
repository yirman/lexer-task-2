grammar Interfaces;

// --- PARSER (Reglas Sintácticas) ---
file: statement+ EOF;

statement: auto_stmt | iface_stmt | mapping_stmt;

auto_stmt: AUTO ID+;

iface_stmt: IFACE ID INET method option*;

mapping_stmt: MAPPING ID script_def map_def*;

method: STATIC | DHCP | LOOPBACK;

option: ADDRESS IP
      | NETMASK IP
      | GATEWAY IP
      | NETWORK IP
      | BROADCAST IP
      | DNS_NAMESERVERS IP+;

script_def: SCRIPT value;
map_def: MAP value value;

value: STRING | ID | STATIC | DHCP | LOOPBACK;

// --- LEXER (Tokens / Expresiones Regulares) ---
AUTO: 'auto';
IFACE: 'iface';
INET: 'inet' | 'inet6';
STATIC: 'static';
DHCP: 'dhcp';
LOOPBACK: 'loopback';
ADDRESS: 'address';
NETMASK: 'netmask';
GATEWAY: 'gateway';
NETWORK: 'network';
BROADCAST: 'broadcast';
DNS_NAMESERVERS: 'dns-nameservers';
MAPPING: 'mapping';
SCRIPT: 'script';
MAP: 'map';

IP: [0-9]+ '.' [0-9]+ '.' [0-9]+ '.' [0-9]+;
ID: [a-zA-Z0-9_:-]+;
STRING: [a-zA-Z0-9_./:-]+;

WS: [ \t\r\n]+ -> skip; // Ignorar espacios en blanco
COMMENT: '#' ~[\r\n]* -> skip; // Ignorar comentarios