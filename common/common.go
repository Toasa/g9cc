package common


// Vector
// 異なるデータ型*Token, *IRなどをスライスとして扱うための構造体(wrapperみたいなものか?)
type Vector struct {
    Data []interface{}
    Capacity int
    Len int
}

type Map struct {
    Keys *Vector
    Vals *Vector
}

type StringBuilder struct {
    Data string
    Capacity int
    Len int
}

type Type struct {
    Ty int
    Size int
    Align int

    // Pointer
    Ptr_to *Type

    // Array
    Ary_of *Type // Array
    Len int

    // Struct
    Members *Vector
    Offset int
}

// token.go

const (
    TK_NUM = iota + 256 // Number Literal
    TK_STR // String literal
    TK_IDENT // Identifier
    TK_ARROW // ->
    TK_EXTERN // "extern"
    TK_TYPEDEF // "typedef"
    TK_INT // "int"
    TK_CHAR // "char"
    TK_VOID // "void"
    TK_STRUCT // "struct"
    TK_IF // "if"
    TK_ELSE // "else"
    TK_FOR // "for"
    TK_DO // "do"
    TK_WHILE // "while"
    TK_BREAK // "break"
    TK_EQ // ==
    TK_NE // !=
    TK_LE // <=
    TK_GE // >=
    TK_LOGOR // ||
    TK_LOGAND // &&
    TK_SHL // <<
    TK_SHR // >>
    TK_INC // ++
    TK_DEC // --
    TK_MUL_EQ // *=
    TK_DIV_EQ // /=
    TK_MOD_EQ // %=
    TK_ADD_EQ // +=
    TK_SUB_EQ // -=
    TK_SHL_EQ // <<=
    TK_SHR_EQ // >>=
    TK_BITAND_EQ // &=
    TK_XOR_EQ // ^=
    TK_BITOR_EQ // |=
    TK_RETURN // "return"
    TK_SIZEOF // "sizeof"
    TK_ALIGNOF // "_Alignof"
    TK_EOF
)

// Tokenizer
type Token struct {
    Ty int // token type
    Val int // number literal
    Name string // identifier
    Input string // token string

    // String literal
    Str string
    Len int
}


// parse.go

const (
    ND_NUM = iota + 256 // number literal
    ND_STR // String literal
    ND_IDENT // identifier
    ND_STRUCT // Struct
    ND_VARDEF // Variable definition
    ND_LVAR // Local variable reference
    ND_GVAR // Glocal variable reference
    ND_IF // "if"
    ND_FOR // "for"
    ND_DO_WHILE // do ... while
    ND_BREAK // break
    ND_ADDR // address operator ("&")
    ND_DEREF // pointer dereference ("*")
    ND_DOT // Struct member access
    ND_EQ // ==
    ND_NE // !=
    ND_LE // <=
    ND_LOGAND // &&
    ND_LOGOR // ||
    ND_SHL // <<
    ND_SHR // >>
    ND_MOD // %
    ND_NEG // -
    ND_POST_INC // post ++
    ND_POST_DEC // post --
    ND_MUL_EQ // *=
    ND_DIV_EQ // /=
    ND_MOD_EQ // %=
    ND_ADD_EQ // +=
    ND_SUB_EQ // -=
    ND_SHL_EQ // <<=
    ND_SHR_EQ // >>=
    ND_BITAND_EQ // &=
    ND_XOR_EQ // ^=
    ND_BITOR_EQ // |=
    ND_RETURN // "return"
    ND_SIZEOF // "sizeof"
    ND_ALIGNOF // "_Alignof"
    ND_CALL // Function call
    ND_FUNC // Function definition
    ND_COMP_STMT // Compound statement
    ND_EXPR_STMT // Expression statement
    ND_STMT_EXPR // Statement expression (GNU extn)
    ND_NULL // Null statement
)

const (
    INT = iota
    CHAR
    VOID
    PTR
    ARY
    STRUCT
)

type Node struct {
    Op int // Node type
    Ty *Type // C type
    Lhs *Node // left-hand side
    Rhs *Node // right-hand side
    Val int // number literal
    Expr *Node // "return" or Expression statement
    Stmts *Vector // Compound statement

    Name string // Identifier

    // Global variable
    Is_extern bool
    Data string
    Len int

    // // Struct
    // Members *Vector
    //
    // // Struct access
    // Member string

    // "if" (cond) then "else" els
    // "for" (init; cond; inc) body
    Cond *Node // condtion in IF stmt
    Then *Node
    Els *Node

    Init *Node
    Inc *Node
    Body *Node

    // Function definition
    Stacksize int
    Globals *Vector

    // Offset from BP or begining of a struct
    Offset int

    // Function call
    Args *Vector
}

// sema.go

type Var struct {
    Ty *Type
    Is_local bool

    // local
    Offset int

    // global
    Name string
    Is_extern bool
    Data string
    Len int
}

// ir_dump.go
type IRInfo struct {
    Name string
    Ty int
}

// gen_ir.go

const (
    IR_ADD = iota
    IR_SUB
    IR_MUL
    IR_DIV
    IR_IMM
    IR_BPREL
    IR_MOV
    IR_RETURN
    IR_CALL
    IR_LABEL
    IR_LABEL_ADDR
    IR_EQ
    IR_NE
    IR_LE
    IR_LT
    IR_AND
    IR_OR
    IR_XOR
    IR_SHL
    IR_SHR
    IR_MOD
    IR_NEG
    IR_JMP
    IR_IF
    IR_UNLESS
    IR_LOAD
    IR_STORE
    IR_STORE_ARG
    IR_KILL
    IR_NOP
)

type IR struct {
    Op int
    Lhs int
    Rhs int

    // load/store size in bytes
    Size int

    // For binary operator.
    // If true, rhs is an immediate value.
    Is_imm bool

    // Function call
    Name string
    Nargs int
    Args [6]int
}

const (
    IR_TY_NOARG = iota
    IR_TY_BINARY
    IR_TY_REG
    IR_TY_IMM
    IR_TY_MEM
    IR_TY_JMP
    IR_TY_LABEL
    IR_TY_LABEL_ADDR
    IR_TY_REG_REG
    IR_TY_REG_IMM
    IR_TY_STORE_ARG
    IR_TY_REG_LABEL
    IR_TY_CALL
)

type Function struct {
    Name string
    Stacksize int
    Globals *Vector
    Ir *Vector
}
