# wolf
Learning how to build a programming language

## Stage 0
Learn the basics of building your own programing language following, How to build an interpreter in go, then optimize for the needs on the scripting language to be built.

## Stage 1
This language aims to be simple and easy to learn, it should be easy to install and update compiler.

### Phase 1
- [ ] Create a new rust package to initialise the Wolf scripting language
- [ ] Create initial documentation showing a draft of the language capabilities
- [ ] Create the initial draft of grammar/syntax of the language
- [ ] Use pest to create the language grammar and tokens
- [ ] Create initial cli to help with development and testing

### Phase 2
- [ ] Build interpreter
- [ ] Build compiler

At the end  of this phase you should have a working programming language with the bare minimun functionality, no standart library.

### Phase 3
- [ ] Build garbage collector
- [ ] Build JIT

we should be able to execute the script in the command line using the JIT

### Phase 4
- [ ] Build AOT

we should be able to create executables of the script

### Phase 5
- [ ] Create standart library for string utilities
- [ ] Create standart library for number utilities
- [ ] Create standart library for array utilities
- [ ] Create standart library for hashmap utilities

## Stage 2
This stage is to create your own "browser" like application that will have an input that will accept a link to a script in our language and draw ui.



Wolf language Draft:
this syntax is in early stages, it will change

##### Variable Declaration
```
var number: number = 1; // number is an alias for i32 or f64
var string: string = "string";
```

##### types can be auto detected
```
var number = 2;
var string = "hello cecile";
```

##### Array & Builtin function
```
var arr = [1, 2, 3, 4]; // same as var arr: number[] = [1, 2, 3, 4];
arr.push(5);

println(arr); // Out: [1, 2, 3, 4, 5]
```

##### Hashmap & Builtin function
```
var map = { 
  "a": 1, 
  "b": 2, 
  "c": 3, 
  "d": 4 
};

// Or
var map: string[number] = { 
  "a": 1, 
  "b": 2, 
  "c": 3, 
  "d": 4 
};
```

##### Function Declaration
```
fn hello() -> String ::
  return "hello";
end

println(hello()); // Out: "hello"

// Or
var greet = fn() -> String ::
  return "hello";
end

println(greet()); // Out: "hello"
```

##### If Declaration
```
if str == "hello" ::
  println(str);
end

if str and msg ::
  println(foo);
end

if str or msg ::
  println(foo);
end

if not str ::
  println(foo);
end
```

##### Case Declaration
```
case str == "hello" ::
  true => println(str);
  false => println("");
end

var msg = case str == "hello" ::
  "hello" => str + " world!";
  _ => str;
end

println(msg) // Out: "hello world!" or "hello"
```

##### For Declaration
```
var arr = [1, 2, 3, 4];

for value in arr ::
 println(value);
end

for index, value in arr ::
 println(index);
 println(value);
end

var map = {
  "a": 1, 
  "b": 2, 
  "c": 3, 
  "d": 4 
};
for key, value in map ::
 println(key);
 println(value);
end
```

##### Type Declaration
```
type Area = number;
type Message = string;
type List = string[];
type Items = string[string];
```

##### Struct Declaration
```
struct Point ::
  x: int,
  y: int,
end

impl Point ::
  fn new(x: int, y: int) ::
    self.x = x;
    self.y = y;
  end
end

// OR

fn Point.new(x: int, y: int) ::
  self.x = x;
  self.y = y;
end

var point: Point = Point();
```

### GRAMMAR Helpers
https://github.com/abhimanyu003/qubit ,  https://github.com/lyledean1/cyclang

PROGRAM                <-  STATEMENTS

STATEMENTS             <-  (STATEMENT ';'?)*
STATEMENT              <-  ASSIGNMENT / RETURN / EXPRESSION_STATEMENT

ASSIGNMENT             <-  'let' IDENTIFIER '=' EXPRESSION
RETURN                 <-  'return' EXPRESSION
EXPRESSION_STATEMENT   <-  EXPRESSION

EXPRESSION             <-  INFIX_EXPR(PREFIX_EXPR, INFIX_OPE)
INFIX_EXPR(ATOM, OPE)  <-  ATOM (OPE ATOM)* {
                             precedence
                               L == !=
                               L < >
                               L + -
                               L * /
                           }

IF                     <-  'if' '(' EXPRESSION ')' BLOCK ('else' BLOCK)?

FUNCTION               <-  'fn' '(' PARAMETERS ')' BLOCK
PARAMETERS             <-  LIST(IDENTIFIER, ',')

BLOCK                  <-  '{' STATEMENTS '}'

CALL                   <-  PRIMARY (ARGUMENTS / INDEX)*
ARGUMENTS              <-  '(' LIST(EXPRESSION, ',') ')'
INDEX                  <-   '[' EXPRESSION ']'

PREFIX_EXPR            <-  PREFIX_OPE* CALL
PRIMARY                <-  IF / FUNCTION / ARRAY / HASH / INTEGER / BOOLEAN / NULL / IDENTIFIER / STRING / '(' EXPRESSION ')'

ARRAY                  <-  '[' LIST(EXPRESSION, ',') ']'

HASH                   <-  '{' LIST(HASH_PAIR, ',') '}'
HASH_PAIR              <-  EXPRESSION ':' EXPRESSION

IDENTIFIER             <-  < [a-zA-Z]+ >
INTEGER                <-  < [0-9]+ >
STRING                 <-  < ["] < (!["] .)* > ["] >
BOOLEAN                <-  'true' / 'false'
NULL                   <-  'null'
PREFIX_OPE             <-  < [-!] >
INFIX_OPE              <-  < [-+/*<>] / '==' / '!=' >

KEYWORD                <-  'null' | 'true' | 'false' | 'let' | 'return' | 'if' | 'else' | 'fn'

LIST(ITEM, DELM)       <-  (ITEM (~DELM ITEM)*)?

LINE_COMMENT           <-  '//' (!LINE_END .)* &LINE_END
LINE_END               <-  '\r\n' / '\r' / '\n' / !.

%whitespace            <-  ([ \t\r\n]+ / LINE_COMMENT)*
%word                  <-  [a-zA-Z]+




expression_list = { SOI ~ ( stmt_inner | expression_list_inner ) ~ (WHITESPACE* ~ (stmt_inner | expression_list_inner )*) ~ EOI }
stmt_inner = _{ if_stmt | while_stmt| for_stmt | func_stmt | block_stmt }
expression_list_inner = _{((( expression |  index_stmt  |let_stmt  | print_stmt | call_stmt | grouping ) ~ (semicolon ~ WHITESPACE? ~ (binary | expression |index_stmt| let_stmt |  print_stmt | call_stmt | grouping))*) ~ semicolon)}
expression = _ { binary | literal }


// for loop 
initialization = { "let" ~ name ~ WHITESPACE? ~ "=" ~ WHITESPACE? ~ number }
iteration = { name ~ WHITESPACE? ~ ("++" | "--") }
condition = { name ~ WHITESPACE? ~ ("<" | "<=" | ">" | ">=" ) ~ WHITESPACE? ~ number }
for_stmt = { "for" ~ WHITESPACE? ~ "(" ~ initialization ~ ";" ~ condition ~ ";" ~ iteration ~ ")" ~ block_stmt }

// logical types
if_stmt = { "if" ~ WHITESPACE? ~ "(" ~ (expression | name ) ~ ")" ~ WHITESPACE? ~ block_stmt ~ (WHITESPACE? ~ "else" ~ block_stmt)? }
while_stmt = {"while" ~ WHITESPACE? ~ "(" ~ (expression | name) ~ ")" ~ WHITESPACE? ~ block_stmt}
block_stmt = { "{" ~ WHITESPACE? ~ (return_stmt | expression_list_inner | stmt_inner | WHITESPACE?) ~ (WHITESPACE? ~ (return_stmt | expression_list_inner | stmt_inner)*) ~ (WHITESPACE*)? ~ return_stmt? ~ WHITESPACE? ~ "}" }

// let statements and functions
let_stmt = { (((("let" ~ WHITESPACE?)? ~ name)) ~ WHITESPACE?) ~ (colon ~ type_name ~ WHITESPACE?)? ~ assignment_stmt}
index_stmt = {list_index ~ WHITESPACE?  ~ assignment_stmt  }
assignment_stmt = _{equal ~ WHITESPACE? ~ (list_index | call_stmt | expression | grouping | name)}
func_stmt = { "fn" ~ WHITESPACE? ~ name ~ "(" ~ func_arg* ~ ")" ~ (WHITESPACE? ~ arrow ~ WHITESPACE? ~ type_name)? ~ WHITESPACE? ~ block_stmt }
func_arg = { WHITESPACE? ~ type_name ~ WHITESPACE? ~ name ~ WHITESPACE? ~ comma? }
type_name = { base_type | list_type}
call_stmt = { name ~ "(" ~ (expression | name)? ~ (comma ~ (expression | name))* ~ ")" }
print_stmt = { "print(" ~ (call_stmt | list_index | expression | name ) ~ ")" }
string_type = {"string"}
i32_type = {"i32"}
i64_type = {"i64"}
bool_type = { "bool"}
base_type = _{bool_type | i32_type | i64_type | string_type}
list_type = {"List<" ~  (base_type | list_type )~ ">"}
// binary statemeents
binary = {  operand ~ WHITESPACE? ~ operator_sequence }
operand = _{ literal ~ WHITESPACE? | grouping | call_stmt | name  }
operator_sequence = _{ operator ~ WHITESPACE* ~ operand ~ (WHITESPACE* ~ operator_sequence)? }
operator = { "==" | "!=" | ">=" | "<=" | ">" | "<" | "+" | "-" | "*" | "/" | "^" }

grouping = { "(" ~ expression ~ ")" }
literal = { number | string | bool | nil | list  }

list = { lbracket ~ WHITESPACE? ~ literal ~ (WHITESPACE? ~ "," ~ WHITESPACE? ~ literal)* ~ rbracket }
list_index = {(call_stmt  |expression | name) ~ lbracket ~ (expression  |number | name | call_stmt) ~ rbracket}
name = { (alpha | "_") ~ (alpha | digits | "_")* }
number = { "-"? ~ digits }
digits = @{ ASCII_DIGIT+ }
alpha = { ASCII_ALPHA | "_" }
string = { "\"" ~ (!"\"" ~ ANY)* ~ "\"" }
nil = { "nil" }
bool = { "true" | "false" }
equal = { "=" }
semicolon = { ";" }
colon = { ":" }
arrow = { "->" }
lbracket  = {"["}
rbracket = {"]"}
return_keyword = _{ "return" }
return_stmt = { return_keyword ~ WHITE_SPACE? ~ ((binary | grouping | literal | name | call_stmt)? ~ WHITESPACE? ~ semicolon?)? }
comma = { WHITESPACE? ~ "," ~ WHITESPACE? }
WHITESPACE = _{ " " | "\t" | NEWLINE }








WHITESPACE = _{ " " | "\t" | NEWLINE }

// Identifier
ident = @{ ASCII_ALPHA ~ (ASCII_ALPHANUMERIC | "_")* }

// Numbers
int = @{ ("+" | "-")? ~ ASCII_DIGIT+ }
float = @{ int ~ "." ~ ASCII_DIGIT+ ~ (^"e" ~ int)? }
// float = @{ int ~ ("." ~ ASCII_DIGIT*)? ~ (^"e" ~ int)? }
number = _{ float | int }

// Strings
string = { "\"" ~ (!"\"" ~ ANY)* ~ "\"" }

null = { "nil" }

boolean = { "true" | "false" }

// Operations on numbers
operation = _{ subtract | add | multiply | divide | power | modulus }
add         = { "+" }
subtract    = { "-" }
multiply    = { "*" }
divide      = { "/" }
power       = { "^" }
modulus     = { "%" }

type = { base_type }
base_type = _{ "bool" | "int" | "float" | "string" }

parameters = { ( ident ~ ( "," ~ ident )* )? }

atom = _{ block_stmt | if_stmt | function_stmt | call_stmt | boolean | null | ident | number | string | "(" ~ expression ~ ")"  }
expression = { atom ~ (operation ~ atom)* }

// expression_stmt = { expression }
return_stmt = { "return" ~ expression }
let_stmt = { "let" ~ ident ~ ( ":" ~ type )? ~ assign_stmt }
assign_stmt = { "=" ~ expression }

if_stmt = { "if" ~ expression ~ block_stmt }
function_stmt = { "fn" ~ "(" ~ parameters ~ ")" ~ ( "->" ~ type )? ~ block_stmt }
block_stmt = { "::" ~ statements ~ "end" }
call_stmt = { ident ~ "(" ~ (expression | ident)? ~ ("," ~ (expression | ident))* ~ ")" }

statement = { let_stmt | return_stmt | expression }
statements = { (statement ~ ";")* }

program = _{ SOI ~ statements ~ EOI }



























program = { SOI ~ ( stmt_inner | expression_list_inner ) ~ (WHITESPACE* ~ (stmt_inner | expression_list_inner )*) ~ EOI }
stmt_inner = _{ if_stmt | func_stmt | block_stmt }
expression_list_inner = _{((( expression  |let_stmt  | print_stmt | call_stmt | grouping ) ~ (semicolon ~ WHITESPACE? ~ (binary | expression | let_stmt |  print_stmt | call_stmt | grouping))*) ~ semicolon)}
expression = _{ binary | literal }


// logical types
if_stmt = { "if" ~ WHITESPACE? ~ "(" ~ (expression | name ) ~ ")" ~ WHITESPACE? ~ block_stmt ~ (WHITESPACE? ~ "else" ~ block_stmt)? }
block_stmt = { "::" ~ WHITESPACE? ~ (return_stmt | expression_list_inner | stmt_inner | WHITESPACE?) ~ (WHITESPACE? ~ (return_stmt | expression_list_inner | stmt_inner)*) ~ (WHITESPACE*)? ~ return_stmt? ~ WHITESPACE? ~ "end" }

// let statements and functions
let_stmt = { (((("let" ~ WHITESPACE?)? ~ name)) ~ WHITESPACE?) ~ (colon ~ type_name ~ WHITESPACE?)? ~ assignment_stmt}
assignment_stmt = _{equal ~ WHITESPACE? ~ (call_stmt | expression | grouping | name)}
func_stmt = { "fn" ~ WHITESPACE? ~ name ~ "(" ~ func_arg* ~ ")" ~ (WHITESPACE? ~ arrow ~ WHITESPACE? ~ type_name)? ~ WHITESPACE? ~ block_stmt }
// func_arg = { WHITESPACE? ~ type_name ~ WHITESPACE? ~ name ~ WHITESPACE? ~ comma? }
func_arg = { WHITESPACE? ~ name ~ (colon ~ WHITESPACE? ~ type_name)? ~ WHITESPACE? ~ comma? }
type_name = _{ base_type }
call_stmt = { name ~ "(" ~ (expression | name)? ~ (comma ~ (expression | name))* ~ ")" }
print_stmt = { "print(" ~ (call_stmt | expression | name ) ~ ")" }

string_type = {"string"}
i32_type = {"i32"}
i64_type = {"i64"}
float_type = {"float"}
int_type = {"int"}
bool_type = { "bool"}

base_type = _{bool_type | float_type | int_type | i32_type | i64_type | string_type}

// binary statements
binary = {  operand ~ WHITESPACE? ~ operator_sequence }
operand = _{ literal ~ WHITESPACE? | grouping | call_stmt | name  }
operator_sequence = _{ operator ~ WHITESPACE* ~ operand ~ (WHITESPACE* ~ operator_sequence)? }
operator = { "==" | "!=" | ">=" | "<=" | ">" | "<" | "+" | "-" | "*" | "/" | "^" | "%" }

grouping = { "(" ~ expression ~ ")" }
literal = _{ number | string | bool | nil  }

name = @{ ASCII_ALPHA ~ (ASCII_ALPHANUMERIC | "_")* }
int = @{ ("+" | "-")? ~ ASCII_DIGIT+ }
float = @{ int ~ "." ~ ASCII_DIGIT+ ~ (^"e" ~ int)? }
number = _{ float | int  }
string = { "\"" ~ (!"\"" ~ ANY)* ~ "\"" }

nil = { "nil" }
bool = { "true" | "false" }
equal = { "=" }
semicolon = { ";" }
colon = { ":" }
arrow = { "->" }

return_stmt = { "return" ~ WHITE_SPACE? ~ ((binary | grouping | literal | name | call_stmt)? ~ WHITESPACE? ~ semicolon?)? }
comma = { WHITESPACE? ~ "," ~ WHITESPACE? }
WHITESPACE = _{ " " | "\t" | NEWLINE }