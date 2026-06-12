# Tokens

## Keywords

| Token | Lexeme | Description |
|---|---|---|
| `LET` | `let` | Variable declaration |
| `FUNC` | `func` | Function definition |
| `CLASS` | `class` | Class definition |
| `INTERFACE` | `interface` | Interface definition |
| `ENUM` | `enum` | Enum definition |
| `STRUCT` | `struct` | Struct definition |
| `TYPE` | `type` | Type alias definition |
| `IMPLEMENTS` | `implements` | Class interface list |
| `SELF` | `Self` | Reference to current instance |
| `IF` | `if` | Conditional |
| `ELSE` | `else` | Conditional alternative |
| `SWITCH` | `switch` | Pattern matching |
| `IN` | `in` | Membership check (for-each) |
| `FOR` | `for` | For-each loop |
| `WHILE` | `while` | While loop |
| `DO` | `do` | Do-while loop |
| `SEQ` | `seq` | Sequential block |
| `PAR` | `par` | Parallel block |
| `RETURN` | `return` | Return statement |
| `BREAK` | `break` | Break out of loop |
| `CONTINUE` | `continue` | Skip to next iteration |
| `PASS` | `pass` | No-op statement |
| `GOTO` | `goto` | Jump to label |
| `AND` | `and` | Logical AND |
| `OR` | `or` | Logical OR |
| `TRUE` | `true` | Boolean true |
| `FALSE` | `false` | Boolean false |
| `PRINT` | `print` | Built-in print |
| `INPUT` | `input` | Built-in input |
| `S_CHANNEL` | `s_channel` | Send on channel |
| `C_CHANNEL` | `c_channel` | Receive on channel |

## Primitive Types

| Token         | Lexeme   |
| ------------- | -------- |
| `TYPE_I8`     | `i8`     |
| `TYPE_I16`    | `i16`    |
| `TYPE_I32`    | `i32`    |
| `TYPE_I64`    | `i64`    |
| `TYPE_U8`     | `u8`     |
| `TYPE_U16`    | `u16`    |
| `TYPE_U32`    | `u32`    |
| `TYPE_U64`    | `u64`    |
| `TYPE_F16`    | `f16`    |
| `TYPE_F32`    | `f32`    |
| `TYPE_F64`    | `f64`    |
| `TYPE_CHAR`   | `char`   |
| `TYPE_STRING` | `string` |
| `TYPE_BOOL`   | `bool`   |
| `TYPE_ANY`    | `any`    |
| `TYPE_VOID`   | `void`   |
| `TYPE_CHAN`   | `chan`   |

## Operators

| Token          | Lexeme | Description                  |
| -------------- | ------ | ---------------------------- |
| `ASSIGN`       | `=`    | Assignment                   |
| `PLUS`         | `+`    | Addition                     |
| `MINUS`        | `-`    | Subtraction / unary negation |
| `STAR`         | `*`    | Multiplication               |
| `SLASH`        | `/`    | Division                     |
| `PERCENT`      | `%`    | Modulo                       |
| `BANG`         | `!`    | Logical NOT                  |
| `EQ`           | `==`   | Equality                     |
| `NEQ`          | `!=`   | Inequality                   |
| `LT`           | `<`    | Less than                    |
| `GT`           | `>`    | Greater than                 |
| `LEQ`          | `<=`   | Less than or equal           |
| `GEQ`          | `>=`   | Greater than or equal        |
| `ARROW`        | `->`   | Function return type         |
| `FAT_ARROW`    | `=>`   | Case clause separator        |
| `PLUS_ASSIGN`  | `+=`   | Add and assign               |
| `MINUS_ASSIGN` | `-=`   | Subtract and assign          |
| `STAR_ASSIGN`  | `*=`   | Multiply and assign          |
| `SLASH_ASSIGN` | `/=`   | Divide and assign            |

## Delimiters

| Token | Lexeme | Description |
|---|---|---|
| `LPAREN` | `(` | Left parenthesis |
| `RPAREN` | `)` | Right parenthesis |
| `LBRACE` | `{` | Left brace |
| `RBRACE` | `}` | Right brace |
| `LBRACKET` | `[` | Left bracket |
| `RBRACKET` | `]` | Right bracket |
| `COMMA` | `,` | Argument / element separator |
| `SEMICOLON` | `;` | Statement terminator |
| `COLON` | `:` | Type annotation separator |
| `DOT` | `.` | Member access |

## Literals

| Token | Pattern | Example |
|---|---|---|
| `INT` | `0 \| [1-9][0-9]*` | `42` |
| `FLOAT` | `[0-9]+ "." [0-9]+` | `3.14` |
| `STRING` | `"\"" [^"]* "\""` | `"hello"` |
| `CHAR` | `"'" <char> "'"` | `'a'` |

## Identifier

| Token | Pattern | Example |
|---|---|---|
| `IDENT` | `[a-zA-Z_][a-zA-Z0-9_]*` | `foo`, `_bar`, `MyClass` |

## Comments (skipped by lexer)

| Kind | Pattern |
|---|---|
| Line comment | `#` until end of line |
| Block comment | `/* ... */` |
