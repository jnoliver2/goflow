grammar Excellent;

// rebuild with % antlr4 -Dlanguage=Go excellent/gen/Excellent.g4 -visitor -package gen

import LexUnicode;

COMMA      : ',';
LPAREN     : '(';
RPAREN     : ')';
LBRACK     : '[';
RBRACK     : ']';

DOT        : '.';

PLUS       : '+';
MINUS      : '-';
TIMES      : '*';
DIVIDE     : '/';
EXPONENT   : '^';

EQ         : '=';
NEQ        : '!=';

LTE        : '<=';
LT         : '<';
GTE        : '>=';
GT         : '>';

AMPERSAND  : '&';

DECIMAL    : [0-9]+('.'[0-9]+)?;
STRING     : '"' (~["] | '""')* '"';

TRUE       : [Tt][Rr][Uu][Ee];
FALSE      : [Ff][Aa][Ll][Ss][Ee];

NAME       : UnicodeLetter+ (UnicodeLetter | UnicodeDigit | '_')*;    // variable names, e.g. contact.name or function names, e.g. SUM

WS         : [ \t\n\r]+ -> skip;        // ignore whitespace

ERROR      : . ;

parse      : expression EOF;

atom       : fnname LPAREN parameters? RPAREN             # functionCall
           | atom DOT atom                                # dotLookup
           | atom LBRACK expression RBRACK                # arrayLookup
           | NAME                                         # contextReference
           | STRING                                       # stringLiteral
           | DECIMAL                                      # decimalLiteral
           | TRUE                                         # true
           | FALSE                                        # false
           ;

expression : atom                                            # atomReference
           | MINUS expression                                # negation
           | expression EXPONENT expression                  # exponent
           | expression op=(TIMES | DIVIDE) expression       # multiplicationOrDivision
           | expression op=(PLUS | MINUS) expression         # additionOrSubtraction
           | expression op=(LTE | LT | GTE | GT) expression  # comparison
           | expression op=(EQ | NEQ) expression             # equality
           | expression AMPERSAND expression                 # concatenation
           | LPAREN expression RPAREN                        # parentheses
           ;

fnname     : NAME
           | TRUE
           | FALSE
           ;

parameters : expression (COMMA expression)*               # functionParameters
           ;