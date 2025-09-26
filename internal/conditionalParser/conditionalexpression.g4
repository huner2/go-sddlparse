grammar conditionalexpression;

// Lexer rules (must come first and be uppercase)

// Skip whitespace
WHITESPACE: [ \t\r\n]+ -> skip;

// Keywords and operators (specific tokens first to avoid conflicts)
MEMBER_OF: 'Member_of';
NOT_MEMBER_OF: 'Not_Member_of';
MEMBER_OF_ANY: 'Member_of_Any';
NOT_MEMBER_OF_ANY: 'Not_Member_of_Any';
DEVICE_MEMBER_OF: 'Device_Member_of';
DEVICE_MEMBER_OF_ANY: 'Device_Member_of_Any';
NOT_DEVICE_MEMBER_OF: 'Not_Device_Member_of';
NOT_DEVICE_MEMBER_OF_ANY: 'Not_Device_Member_of_Any';
EXISTS: 'Exists';
NOT_EXISTS: 'Not_exists';
CONTAINS: 'Contains';
NOT_CONTAINS: 'Not_Contains';
ANY_OF: 'Any_of';
NOT_ANY_OF: 'Not_Any_of';

// Logical operators
AND: '&&';
OR: '||';
NOT: '!';

// Comparison operators
EQ: '==';
NEQ: '!=';
LT: '<';
LTE: '<=';
GT: '>';
GTE: '>=';

// Symbols
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
COMMA: ',';
DOT: '.';

// Attribute prefixes
USER_PREFIX: '@user.';
DEVICE_PREFIX: '@device.';
RESOURCE_PREFIX: '@resource.';

// SID tokens
SID_TOKEN: 'DA'| 'DG' | 'DU' | 'ED' | 'DD' | 'DC' | 'BA' | 'BG' | 'BU' | 'LA' | 'LG' | 'AO'
         | 'BO' | 'PO' | 'SO' | 'AU' | 'PS' | 'CO' | 'CG' | 'SY' | 'PU' | 'WD' | 'RE' | 'IU' | 'NU'
         | 'SU' | 'RC' | 'WR' | 'AN' | 'SA' | 'CA' | 'RS' | 'EA' | 'PA' | 'RU' | 'LS' | 'NS' | 'RD'
         | 'NO' | 'MU' | 'LU' | 'IS' | 'CY' | 'OW' | 'ER' | 'RO' | 'CD' | 'AC' | 'RA' | 'ES' | 'MS'
         | 'UD' | 'HA' | 'CN' | 'AA' | 'RM' | 'LW' | 'ME' |'MP' | 'HI' | 'SI';

// SID literal wrapper
SID_START: 'SID(';

// Numeric values (specific patterns first)
HEX_INTEGER: '0x' [0-9a-fA-F]+;
OCTAL_INTEGER: '0' [0-7]+;
DECIMAL_INTEGER: [0-9]+;

// String literals
STRING_LITERAL: '"' (~["\r\n])* '"';

// Octet string
OCTET_STRING: '#' [0-9a-fA-F]*;

// SID format: S-1-IdentifierAuthority-SubAuthority1-SubAuthority2-...
SID_FORMAT: 'S-1-' ([0-9]+ | ('0x' [0-9a-fA-F]+)) ('-' [0-9]+)+;

// Identifiers (attribute names)
IDENTIFIER: [a-zA-Z] [a-zA-Z0-9_:./@]*;

// Special literal characters for extended identifiers
LIT_CHAR: ([#$'*+\-./:;?@[\\\]^_`{}~]) | [\u0080-\uFFFF] | ('%' [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F]);

// Catch-all for any remaining characters
ANY_CHAR: .;

// Parser rules (lowercase)

// Entry point
condExpr: expr EOF;

expr: superTerm (OR superTerm)*;

superTerm: factor (AND factor)*;

factor: term
      | LPAREN expr RPAREN
      | NOT factor
      ;

term: memberofOp | existsOp | relOp | relOp2 | containsOp | anyofOp | attrName;

// Membership operations
memberofOp: (MEMBER_OF | NOT_MEMBER_OF | MEMBER_OF_ANY | NOT_MEMBER_OF_ANY |
            DEVICE_MEMBER_OF | DEVICE_MEMBER_OF_ANY | NOT_DEVICE_MEMBER_OF |
            NOT_DEVICE_MEMBER_OF_ANY) sidArray;

// Existence operations
existsOp: (EXISTS | NOT_EXISTS) attrName;

// Relational operations (scalars only)
relOp: attrName (LT | LTE | GT | GTE) (attrName | value);

// Equality operations (scalar or list)
relOp2: attrName (EQ | NEQ) (attrName | valueArray);

// Contains operations
containsOp: attrName (CONTAINS | NOT_CONTAINS) (attrName | valueArray);

// Any of operations
anyofOp: attrName (ANY_OF | NOT_ANY_OF) (attrName | valueArray);

// Attribute names
attrName: prefixedAttrName | simpleAttrName;

// Simple attribute name (local attributes)
simpleAttrName: IDENTIFIER;

// Prefixed attribute names
prefixedAttrName: (USER_PREFIX | DEVICE_PREFIX | RESOURCE_PREFIX) (IDENTIFIER | extendedIdentifier);

// Extended identifier for special characters in attribute names
extendedIdentifier: (IDENTIFIER | LIT_CHAR)+;

// SID array
sidArray: LBRACE literalSID (COMMA literalSID)* RBRACE;
literalSID: SID_START sidString RPAREN;

// Value array
valueArray: LBRACE value (COMMA value)* RBRACE | value;

// Values
value: integerValue | STRING_LITERAL | OCTET_STRING;

// Integer values with optional sign
integerValue: ('+' | '-')? (HEX_INTEGER | OCTAL_INTEGER | DECIMAL_INTEGER);

// SID string format
sidString: SID_TOKEN | SID_FORMAT;