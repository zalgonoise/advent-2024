package part01

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrInvalidNumChildren   = errors.New("invalid number of children")
	ErrInvalidTreeStructure = errors.New("invalid tree structure")
	ErrNilTokenOnHeadNumber = errors.New("nil token on head number")
	ErrNilTokenOnTailNumber = errors.New("nil token on tail number")
	ErrInvalidNumIntegers   = errors.New("invalid number of integers")
)

type Mul struct {
	a, b int
}

type Token struct {
	typ token
	val byte
	idx int
}

type AbstractSyntaxTree struct {
	parent   *Token
	children []*AbstractSyntaxTree
}

type token int

const (
	tokenMul token = iota
	tokenMulM
	tokenMulU
	tokenMulL
	tokenOpenParan
	tokenCloseParan
	tokenNum
	tokenComma
	tokenUnsupported
)

const (
	minTokens = 8
	maxLen    = 3
)

func Parse(input string) []Mul {
	tokens := lex(input)
	trees := parse(tokens)

	return convert(trees)
}

func Sum(muls []Mul) int {
	var n int

	for i := range muls {
		n += muls[i].a * muls[i].b
	}

	return n
}

func lex(input string) []Token {
	tokens := make([]Token, 0, len(input))

	for i := range input {
		switch input[i] {
		case 'm':
			tokens = append(tokens, Token{typ: tokenMulM, val: input[i], idx: i})
		case 'u':
			tokens = append(tokens, Token{typ: tokenMulU, val: input[i], idx: i})
		case 'l':
			tokens = append(tokens, Token{typ: tokenMulL, val: input[i], idx: i})
		case '(':
			tokens = append(tokens, Token{typ: tokenOpenParan, val: input[i], idx: i})
		case ')':
			tokens = append(tokens, Token{typ: tokenCloseParan, val: input[i], idx: i})
		case ',':
			tokens = append(tokens, Token{typ: tokenComma, val: input[i], idx: i})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			tokens = append(tokens, Token{typ: tokenNum, val: input[i], idx: i})
		default:
			tokens = append(tokens, Token{typ: tokenUnsupported, val: input[i], idx: i})
		}
	}

	return tokens
}

func parse(tokens []Token) []*AbstractSyntaxTree {
	trees := make([]*AbstractSyntaxTree, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		ast, idx := extractItem(tokens[i:])
		i += idx

		if ast != nil {
			trees = append(trees, ast)
		}
	}

	return trees
}

func convert(trees []*AbstractSyntaxTree) []Mul {
	muls := make([]Mul, 0, len(trees))

	for i := range trees {
		mul, err := convertTree(trees[i])
		if err != nil {
			// TODO: logging if failed
			continue
		}

		muls = append(muls, mul)
	}

	return muls
}

func convertTree(tree *AbstractSyntaxTree) (Mul, error) {
	// root
	//  |- m
	//  |- u
	//  |- l
	//  +- (
	//  |  |- 4
	//  |  +- 3
	//  +- ,
	//  |  |- 1
	//  |  +- 2
	//  +- )

	// verify
	if len(tree.children) != 6 {
		return Mul{}, ErrInvalidNumChildren
	}

	if tree.children[0].parent == nil || tree.children[0].parent.typ != tokenMulM {
		return Mul{}, fmt.Errorf("%w: index 0", ErrInvalidTreeStructure)
	}
	if tree.children[1].parent == nil || tree.children[1].parent.typ != tokenMulU {
		return Mul{}, fmt.Errorf("%w: index 1", ErrInvalidTreeStructure)
	}
	if tree.children[2].parent == nil || tree.children[2].parent.typ != tokenMulL {
		return Mul{}, fmt.Errorf("%w: index 2", ErrInvalidTreeStructure)
	}
	if tree.children[3].parent == nil || tree.children[3].parent.typ != tokenOpenParan ||
		len(tree.children[3].children) == 0 {
		return Mul{}, fmt.Errorf("%w: index 3", ErrInvalidTreeStructure)
	}
	if tree.children[4].parent == nil || tree.children[4].parent.typ != tokenComma ||
		len(tree.children[4].children) == 0 {
		return Mul{}, fmt.Errorf("%w: index 4", ErrInvalidTreeStructure)
	}
	if tree.children[5].parent == nil || tree.children[5].parent.typ != tokenCloseParan {
		return Mul{}, fmt.Errorf("%w: index 5", ErrInvalidTreeStructure)
	}

	var (
		buf   = &bytes.Buffer{}
		left  int
		right int
		err   error
	)

	for i := range tree.children[3].children {
		if tree.children[3].children[i].parent == nil {
			return Mul{}, ErrNilTokenOnHeadNumber
		}

		if buf.Len() > maxLen {
			return Mul{}, ErrInvalidNumIntegers
		}

		buf.WriteByte(tree.children[3].children[i].parent.val)
	}

	if left, err = strconv.Atoi(buf.String()); err != nil {
		return Mul{}, err
	}

	buf.Reset()

	for i := range tree.children[4].children {
		if tree.children[4].children[i].parent == nil {
			return Mul{}, ErrNilTokenOnTailNumber
		}

		buf.WriteByte(tree.children[4].children[i].parent.val)
	}

	if right, err = strconv.Atoi(buf.String()); err != nil {
		return Mul{}, err
	}

	return Mul{left, right}, nil
}

func extractItem(tokens []Token) (*AbstractSyntaxTree, int) {
	i := 0

	for ; i < len(tokens); i++ {
		switch tokens[i].typ {
		case tokenMulM:
			if len(tokens[i:]) < minTokens {
				i += len(tokens[i:])

				continue
			}

			start := i

			i++
			if tokens[i].typ != tokenMulU {
				continue
			}

			i++
			if tokens[i].typ != tokenMulL {
				continue
			}

			i++
			if tokens[i].typ != tokenOpenParan {
				continue
			}

			ast, newIdx := extractMul(tokens[start:])
			i += newIdx

			if ast != nil {
				return ast, newIdx - 1
			}
		default:
			i--

			if i < 0 {
				i = 0
			}

			return nil, i
		}
	}

	i--

	if i < 0 {
		i = 0
	}

	return nil, i - 1
}

func extractMul(tokens []Token) (*AbstractSyntaxTree, int) {
	ast := &AbstractSyntaxTree{
		parent: &Token{
			typ: tokenMul,
			val: 0,
			idx: tokens[0].idx,
		},
		children: make([]*AbstractSyntaxTree, 0, 12),
	}

	openParen := &AbstractSyntaxTree{parent: &tokens[3], children: make([]*AbstractSyntaxTree, 0, 64)}

	ast.children = append(ast.children,
		&AbstractSyntaxTree{parent: &tokens[0]},
		&AbstractSyntaxTree{parent: &tokens[1]},
		&AbstractSyntaxTree{parent: &tokens[2]},
		openParen,
	)

	var (
		commaTree  = &AbstractSyntaxTree{children: make([]*AbstractSyntaxTree, 0, 64)}
		afterComma bool
		i          = 4
	)

	for ; i < len(tokens); i++ {
		switch tokens[i].typ {
		case tokenNum:
			if !afterComma {
				openParen.children = append(openParen.children, &AbstractSyntaxTree{parent: &tokens[i]})

				continue
			}

			commaTree.children = append(commaTree.children, &AbstractSyntaxTree{parent: &tokens[i]})

		case tokenComma:
			if len(openParen.children) == 0 {
				return nil, i
			}

			commaTree.parent = &tokens[i]
			ast.children = append(ast.children, commaTree)
			afterComma = true

		case tokenCloseParan:
			if !afterComma || len(commaTree.children) == 0 {
				return nil, i
			}

			ast.children = append(ast.children, &AbstractSyntaxTree{parent: &tokens[i]})

			return ast, i

		default:
			return nil, i
		}
	}

	return ast, i
}
