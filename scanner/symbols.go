package scanner

// This is all lifter from https://raw.githubusercontent.com/asciimath/asciimathml/master/ASCIIMathML.js

type Symbol struct {
	input            string
	tag              string
	output           string
	ttype            int
	isFunc           bool
	rewriteLeftRight [2]string
	codes            int
	tex              string
	invisible        bool
	acc              bool
	atname           string
	atval            string
	notexcopy        bool
}

// token types
const (
	CONST          = 0
	UNARY          = 1
	BINARY         = 2
	INFIX          = 3
	LEFTBRACKET    = 4
	RIGHTBRACKET   = 5
	SPACE          = 6
	UNDEROVER      = 7
	DEFINITION     = 8
	LEFTRIGHT      = 9
	TEXT           = 10
	BIG            = 11
	LONG           = 12
	STRETCHY       = 13
	MATRIX         = 14
	UNARYUNDEROVER = 15
)

const (
	AMbbb = 0
	AMcal = 1
	AMfrk = 2
)

var AMquote = Symbol{input: "\"", tag: "mtext", output: "mbox", ttype: TEXT}

var AMsymbols = []Symbol{
	//some greek symbols
	{input: "alpha", tag: "mi", output: "\u03B1", ttype: CONST},
	{input: "beta", tag: "mi", output: "\u03B2", ttype: CONST},
	{input: "chi", tag: "mi", output: "\u03C7", ttype: CONST},
	{input: "delta", tag: "mi", output: "\u03B4", ttype: CONST},
	{input: "Delta", tag: "mo", output: "\u0394", ttype: CONST},
	{input: "epsi", tag: "mi", output: "\u03B5", tex: "epsilon", ttype: CONST},
	{input: "varepsilon", tag: "mi", output: "\u025B", ttype: CONST},
	{input: "eta", tag: "mi", output: "\u03B7", ttype: CONST},
	{input: "gamma", tag: "mi", output: "\u03B3", ttype: CONST},
	{input: "Gamma", tag: "mo", output: "\u0393", ttype: CONST},
	{input: "iota", tag: "mi", output: "\u03B9", ttype: CONST},
	{input: "kappa", tag: "mi", output: "\u03BA", ttype: CONST},
	{input: "lambda", tag: "mi", output: "\u03BB", ttype: CONST},
	{input: "Lambda", tag: "mo", output: "\u039B", ttype: CONST},
	{input: "lamda", tag: "mi", output: "\u03BB", ttype: CONST},
	{input: "Lamda", tag: "mo", output: "\u039B", ttype: CONST},
	{input: "mu", tag: "mi", output: "\u03BC", ttype: CONST},
	{input: "nu", tag: "mi", output: "\u03BD", ttype: CONST},
	{input: "omega", tag: "mi", output: "\u03C9", ttype: CONST},
	{input: "Omega", tag: "mo", output: "\u03A9", ttype: CONST},
	{input: "phi", tag: "mi", output: "\u03D5", ttype: CONST},
	{input: "varphi", tag: "mi", output: "\u03C6", ttype: CONST},
	{input: "Phi", tag: "mo", output: "\u03A6", ttype: CONST},
	{input: "pi", tag: "mi", output: "\u03C0", ttype: CONST},
	{input: "Pi", tag: "mo", output: "\u03A0", ttype: CONST},
	{input: "psi", tag: "mi", output: "\u03C8", ttype: CONST},
	{input: "Psi", tag: "mi", output: "\u03A8", ttype: CONST},
	{input: "rho", tag: "mi", output: "\u03C1", ttype: CONST},
	{input: "sigma", tag: "mi", output: "\u03C3", ttype: CONST},
	{input: "Sigma", tag: "mo", output: "\u03A3", ttype: CONST},
	{input: "tau", tag: "mi", output: "\u03C4", ttype: CONST},
	{input: "theta", tag: "mi", output: "\u03B8", ttype: CONST},
	{input: "vartheta", tag: "mi", output: "\u03D1", ttype: CONST},
	{input: "Theta", tag: "mo", output: "\u0398", ttype: CONST},
	{input: "upsilon", tag: "mi", output: "\u03C5", ttype: CONST},
	{input: "xi", tag: "mi", output: "\u03BE", ttype: CONST},
	{input: "Xi", tag: "mo", output: "\u039E", ttype: CONST},
	{input: "zeta", tag: "mi", output: "\u03B6", ttype: CONST},

	//binary operation symbols
	//{input:"-",  tag:"mo", output:"\u0096", tex:null, ttype:CONST},
	{input: "*", tag: "mo", output: "\u22C5", tex: "cdot", ttype: CONST},
	{input: "**", tag: "mo", output: "\u2217", tex: "ast", ttype: CONST},
	{input: "***", tag: "mo", output: "\u22C6", tex: "star", ttype: CONST},
	{input: "//", tag: "mo", output: "/", ttype: CONST},
	{input: "\\\\", tag: "mo", output: "\\", tex: "backslash", ttype: CONST},
	{input: "setminus", tag: "mo", output: "\\", ttype: CONST},
	{input: "xx", tag: "mo", output: "\u00D7", tex: "times", ttype: CONST},
	{input: "|><", tag: "mo", output: "\u22C9", tex: "ltimes", ttype: CONST},
	{input: "><|", tag: "mo", output: "\u22CA", tex: "rtimes", ttype: CONST},
	{input: "|><|", tag: "mo", output: "\u22C8", tex: "bowtie", ttype: CONST},
	{input: "-:", tag: "mo", output: "\u00F7", tex: "div", ttype: CONST},
	{input: "divide", tag: "mo", output: "-:", ttype: DEFINITION},
	{input: "@", tag: "mo", output: "\u2218", tex: "circ", ttype: CONST},
	{input: "o+", tag: "mo", output: "\u2295", tex: "oplus", ttype: CONST},
	{input: "ox", tag: "mo", output: "\u2297", tex: "otimes", ttype: CONST},
	{input: "o.", tag: "mo", output: "\u2299", tex: "odot", ttype: CONST},
	{input: "sum", tag: "mo", output: "\u2211", ttype: UNDEROVER},
	{input: "prod", tag: "mo", output: "\u220F", ttype: UNDEROVER},
	{input: "^^", tag: "mo", output: "\u2227", tex: "wedge", ttype: CONST},
	{input: "^^^", tag: "mo", output: "\u22C0", tex: "bigwedge", ttype: UNDEROVER},
	{input: "vv", tag: "mo", output: "\u2228", tex: "vee", ttype: CONST},
	{input: "vvv", tag: "mo", output: "\u22C1", tex: "bigvee", ttype: UNDEROVER},
	{input: "nn", tag: "mo", output: "\u2229", tex: "cap", ttype: CONST},
	{input: "nnn", tag: "mo", output: "\u22C2", tex: "bigcap", ttype: UNDEROVER},
	{input: "uu", tag: "mo", output: "\u222A", tex: "cup", ttype: CONST},
	{input: "uuu", tag: "mo", output: "\u22C3", tex: "bigcup", ttype: UNDEROVER},

	//binary relation symbols
	{input: "!=", tag: "mo", output: "\u2260", tex: "ne", ttype: CONST},
	{input: ":=", tag: "mo", output: ":=", ttype: CONST},
	{input: "lt", tag: "mo", output: "<", ttype: CONST},
	{input: "<=", tag: "mo", output: "\u2264", tex: "le", ttype: CONST},
	{input: "lt=", tag: "mo", output: "\u2264", tex: "leq", ttype: CONST},
	{input: "gt", tag: "mo", output: ">", ttype: CONST},
	{input: ">=", tag: "mo", output: "\u2265", tex: "ge", ttype: CONST},
	{input: "gt=", tag: "mo", output: "\u2265", tex: "geq", ttype: CONST},
	{input: "-<", tag: "mo", output: "\u227A", tex: "prec", ttype: CONST},
	{input: "-lt", tag: "mo", output: "\u227A", ttype: CONST},
	{input: ">-", tag: "mo", output: "\u227B", tex: "succ", ttype: CONST},
	{input: "-<=", tag: "mo", output: "\u2AAF", tex: "preceq", ttype: CONST},
	{input: ">-=", tag: "mo", output: "\u2AB0", tex: "succeq", ttype: CONST},
	{input: "in", tag: "mo", output: "\u2208", ttype: CONST},
	{input: "!in", tag: "mo", output: "\u2209", tex: "notin", ttype: CONST},
	{input: "sub", tag: "mo", output: "\u2282", tex: "subset", ttype: CONST},
	{input: "sup", tag: "mo", output: "\u2283", tex: "supset", ttype: CONST},
	{input: "sube", tag: "mo", output: "\u2286", tex: "subseteq", ttype: CONST},
	{input: "supe", tag: "mo", output: "\u2287", tex: "supseteq", ttype: CONST},
	{input: "-=", tag: "mo", output: "\u2261", tex: "equiv", ttype: CONST},
	{input: "~=", tag: "mo", output: "\u2245", tex: "cong", ttype: CONST},
	{input: "~~", tag: "mo", output: "\u2248", tex: "approx", ttype: CONST},
	{input: "prop", tag: "mo", output: "\u221D", tex: "propto", ttype: CONST},

	//logical symbols
	{input: "and", tag: "mtext", output: "and", ttype: SPACE},
	{input: "or", tag: "mtext", output: "or", ttype: SPACE},
	{input: "not", tag: "mo", output: "\u00AC", tex: "neg", ttype: CONST},
	{input: "=>", tag: "mo", output: "\u21D2", tex: "implies", ttype: CONST},
	{input: "if", tag: "mo", output: "if", ttype: SPACE},
	{input: "<=>", tag: "mo", output: "\u21D4", tex: "iff", ttype: CONST},
	{input: "AA", tag: "mo", output: "\u2200", tex: "forall", ttype: CONST},
	{input: "EE", tag: "mo", output: "\u2203", tex: "exists", ttype: CONST},
	{input: "_|_", tag: "mo", output: "\u22A5", tex: "bot", ttype: CONST},
	{input: "TT", tag: "mo", output: "\u22A4", tex: "top", ttype: CONST},
	{input: "|--", tag: "mo", output: "\u22A2", tex: "vdash", ttype: CONST},
	{input: "|==", tag: "mo", output: "\u22A8", tex: "models", ttype: CONST},

	//grouping brackets
	{input: "(", tag: "mo", output: "(", tex: "left(", ttype: LEFTBRACKET},
	{input: ")", tag: "mo", output: ")", tex: "right)", ttype: RIGHTBRACKET},
	{input: "[", tag: "mo", output: "[", tex: "left[", ttype: LEFTBRACKET},
	{input: "]", tag: "mo", output: "]", tex: "right]", ttype: RIGHTBRACKET},
	{input: "{", tag: "mo", output: "{", ttype: LEFTBRACKET},
	{input: "}", tag: "mo", output: "}", ttype: RIGHTBRACKET},
	{input: "|", tag: "mo", output: "|", ttype: LEFTRIGHT},
	{input: ":|:", tag: "mo", output: "|", ttype: CONST},
	{input: "|:", tag: "mo", output: "|", ttype: LEFTBRACKET},
	{input: ":|", tag: "mo", output: "|", ttype: RIGHTBRACKET},
	//{input:"||", tag:"mo", output:"||", tex:null, ttype:LEFTRIGHT},
	{input: "(:", tag: "mo", output: "\u2329", tex: "langle", ttype: LEFTBRACKET},
	{input: ":)", tag: "mo", output: "\u232A", tex: "rangle", ttype: RIGHTBRACKET},
	{input: "<<", tag: "mo", output: "\u2329", ttype: LEFTBRACKET},
	{input: ">>", tag: "mo", output: "\u232A", ttype: RIGHTBRACKET},
	{input: "{:", tag: "mo", output: "{:", ttype: LEFTBRACKET, invisible: true},
	{input: ":}", tag: "mo", output: ":}", ttype: RIGHTBRACKET, invisible: true},

	//miscellaneous symbols
	{input: "int", tag: "mo", output: "\u222B", ttype: CONST},
	{input: "dx", tag: "mi", output: "{:d x:}", ttype: DEFINITION},
	{input: "dy", tag: "mi", output: "{:d y:}", ttype: DEFINITION},
	{input: "dz", tag: "mi", output: "{:d z:}", ttype: DEFINITION},
	{input: "dt", tag: "mi", output: "{:d t:}", ttype: DEFINITION},
	{input: "oint", tag: "mo", output: "\u222E", ttype: CONST},
	{input: "del", tag: "mo", output: "\u2202", tex: "partial", ttype: CONST},
	{input: "grad", tag: "mo", output: "\u2207", tex: "nabla", ttype: CONST},
	{input: "+-", tag: "mo", output: "\u00B1", tex: "pm", ttype: CONST},
	{input: "O/", tag: "mo", output: "\u2205", tex: "emptyset", ttype: CONST},
	{input: "oo", tag: "mo", output: "\u221E", tex: "infty", ttype: CONST},
	{input: "aleph", tag: "mo", output: "\u2135", ttype: CONST},
	{input: "...", tag: "mo", output: "...", tex: "ldots", ttype: CONST},
	{input: ":.", tag: "mo", output: "\u2234", tex: "therefore", ttype: CONST},
	{input: ":'", tag: "mo", output: "\u2235", tex: "because", ttype: CONST},
	{input: "/_", tag: "mo", output: "\u2220", tex: "angle", ttype: CONST},
	{input: "/_\\", tag: "mo", output: "\u25B3", tex: "triangle", ttype: CONST},
	{input: "'", tag: "mo", output: "\u2032", tex: "prime", ttype: CONST},
	{input: "tilde", tag: "mover", output: "~", ttype: UNARY, acc: true},
	{input: "\\ ", tag: "mo", output: "\u00A0", ttype: CONST},
	{input: "frown", tag: "mo", output: "\u2322", ttype: CONST},
	{input: "quad", tag: "mo", output: "\u00A0\u00A0", ttype: CONST},
	{input: "qquad", tag: "mo", output: "\u00A0\u00A0\u00A0\u00A0", ttype: CONST},
	{input: "cdots", tag: "mo", output: "\u22EF", ttype: CONST},
	{input: "vdots", tag: "mo", output: "\u22EE", ttype: CONST},
	{input: "ddots", tag: "mo", output: "\u22F1", ttype: CONST},
	{input: "diamond", tag: "mo", output: "\u22C4", ttype: CONST},
	{input: "square", tag: "mo", output: "\u25A1", ttype: CONST},
	{input: "|__", tag: "mo", output: "\u230A", tex: "lfloor", ttype: CONST},
	{input: "__|", tag: "mo", output: "\u230B", tex: "rfloor", ttype: CONST},
	{input: "|~", tag: "mo", output: "\u2308", tex: "lceiling", ttype: CONST},
	{input: "~|", tag: "mo", output: "\u2309", tex: "rceiling", ttype: CONST},
	{input: "CC", tag: "mo", output: "\u2102", ttype: CONST},
	{input: "NN", tag: "mo", output: "\u2115", ttype: CONST},
	{input: "QQ", tag: "mo", output: "\u211A", ttype: CONST},
	{input: "RR", tag: "mo", output: "\u211D", ttype: CONST},
	{input: "ZZ", tag: "mo", output: "\u2124", ttype: CONST},
	{input: "f", tag: "mi", output: "f", ttype: UNARY, isFunc: true},
	{input: "g", tag: "mi", output: "g", ttype: UNARY, isFunc: true},

	//standard functions
	{input: "lim", tag: "mo", output: "lim", ttype: UNDEROVER},
	{input: "Lim", tag: "mo", output: "Lim", ttype: UNDEROVER},
	{input: "sin", tag: "mo", output: "sin", ttype: UNARY, isFunc: true},
	{input: "cos", tag: "mo", output: "cos", ttype: UNARY, isFunc: true},
	{input: "tan", tag: "mo", output: "tan", ttype: UNARY, isFunc: true},
	{input: "sinh", tag: "mo", output: "sinh", ttype: UNARY, isFunc: true},
	{input: "cosh", tag: "mo", output: "cosh", ttype: UNARY, isFunc: true},
	{input: "tanh", tag: "mo", output: "tanh", ttype: UNARY, isFunc: true},
	{input: "cot", tag: "mo", output: "cot", ttype: UNARY, isFunc: true},
	{input: "sec", tag: "mo", output: "sec", ttype: UNARY, isFunc: true},
	{input: "csc", tag: "mo", output: "csc", ttype: UNARY, isFunc: true},
	{input: "arcsin", tag: "mo", output: "arcsin", ttype: UNARY, isFunc: true},
	{input: "arccos", tag: "mo", output: "arccos", ttype: UNARY, isFunc: true},
	{input: "arctan", tag: "mo", output: "arctan", ttype: UNARY, isFunc: true},
	{input: "coth", tag: "mo", output: "coth", ttype: UNARY, isFunc: true},
	{input: "sech", tag: "mo", output: "sech", ttype: UNARY, isFunc: true},
	{input: "csch", tag: "mo", output: "csch", ttype: UNARY, isFunc: true},
	{input: "exp", tag: "mo", output: "exp", ttype: UNARY, isFunc: true},
	{input: "abs", tag: "mo", output: "abs", ttype: UNARY, rewriteLeftRight: [2]string{"|", "|"}},
	{input: "norm", tag: "mo", output: "norm", ttype: UNARY, rewriteLeftRight: [2]string{"\u2225", "\u2225"}},
	{input: "floor", tag: "mo", output: "floor", ttype: UNARY, rewriteLeftRight: [2]string{"\u230A", "\u230B"}},
	{input: "ceil", tag: "mo", output: "ceil", ttype: UNARY, rewriteLeftRight: [2]string{"\u2308", "\u2309"}},
	{input: "log", tag: "mo", output: "log", ttype: UNARY, isFunc: true},
	{input: "ln", tag: "mo", output: "ln", ttype: UNARY, isFunc: true},
	{input: "det", tag: "mo", output: "det", ttype: UNARY, isFunc: true},
	{input: "dim", tag: "mo", output: "dim", ttype: CONST},
	{input: "mod", tag: "mo", output: "mod", ttype: CONST},
	{input: "gcd", tag: "mo", output: "gcd", ttype: UNARY, isFunc: true},
	{input: "lcm", tag: "mo", output: "lcm", ttype: UNARY, isFunc: true},
	{input: "lub", tag: "mo", output: "lub", ttype: CONST},
	{input: "glb", tag: "mo", output: "glb", ttype: CONST},
	{input: "min", tag: "mo", output: "min", ttype: UNDEROVER},
	{input: "max", tag: "mo", output: "max", ttype: UNDEROVER},
	{input: "Sin", tag: "mo", output: "Sin", ttype: UNARY, isFunc: true},
	{input: "Cos", tag: "mo", output: "Cos", ttype: UNARY, isFunc: true},
	{input: "Tan", tag: "mo", output: "Tan", ttype: UNARY, isFunc: true},
	{input: "Arcsin", tag: "mo", output: "Arcsin", ttype: UNARY, isFunc: true},
	{input: "Arccos", tag: "mo", output: "Arccos", ttype: UNARY, isFunc: true},
	{input: "Arctan", tag: "mo", output: "Arctan", ttype: UNARY, isFunc: true},
	{input: "Sinh", tag: "mo", output: "Sinh", ttype: UNARY, isFunc: true},
	{input: "Cosh", tag: "mo", output: "Cosh", ttype: UNARY, isFunc: true},
	{input: "Tanh", tag: "mo", output: "Tanh", ttype: UNARY, isFunc: true},
	{input: "Cot", tag: "mo", output: "Cot", ttype: UNARY, isFunc: true},
	{input: "Sec", tag: "mo", output: "Sec", ttype: UNARY, isFunc: true},
	{input: "Csc", tag: "mo", output: "Csc", ttype: UNARY, isFunc: true},
	{input: "Log", tag: "mo", output: "Log", ttype: UNARY, isFunc: true},
	{input: "Ln", tag: "mo", output: "Ln", ttype: UNARY, isFunc: true},
	{input: "Abs", tag: "mo", output: "abs", ttype: UNARY, notexcopy: true, rewriteLeftRight: [2]string{"|", "|"}},

	//arrows
	{input: "uarr", tag: "mo", output: "\u2191", tex: "uparrow", ttype: CONST},
	{input: "darr", tag: "mo", output: "\u2193", tex: "downarrow", ttype: CONST},
	{input: "rarr", tag: "mo", output: "\u2192", tex: "rightarrow", ttype: CONST},
	{input: "->", tag: "mo", output: "\u2192", tex: "to", ttype: CONST},
	{input: ">->", tag: "mo", output: "\u21A3", tex: "rightarrowtail", ttype: CONST},
	{input: "->>", tag: "mo", output: "\u21A0", tex: "twoheadrightarrow", ttype: CONST},
	{input: ">->>", tag: "mo", output: "\u2916", tex: "twoheadrightarrowtail", ttype: CONST},
	{input: "|->", tag: "mo", output: "\u21A6", tex: "mapsto", ttype: CONST},
	{input: "larr", tag: "mo", output: "\u2190", tex: "leftarrow", ttype: CONST},
	{input: "harr", tag: "mo", output: "\u2194", tex: "leftrightarrow", ttype: CONST},
	{input: "rArr", tag: "mo", output: "\u21D2", tex: "Rightarrow", ttype: CONST},
	{input: "lArr", tag: "mo", output: "\u21D0", tex: "Leftarrow", ttype: CONST},
	{input: "hArr", tag: "mo", output: "\u21D4", tex: "Leftrightarrow", ttype: CONST},
	//commands with argument
	{input: "sqrt", tag: "msqrt", output: "sqrt", ttype: UNARY},
	{input: "root", tag: "mroot", output: "root", ttype: BINARY},
	{input: "frac", tag: "mfrac", output: "/", ttype: BINARY},
	{input: "/", tag: "mfrac", output: "/", ttype: INFIX},
	{input: "stackrel", tag: "mover", output: "stackrel", ttype: BINARY},
	{input: "overset", tag: "mover", output: "stackrel", ttype: BINARY},
	{input: "underset", tag: "munder", output: "stackrel", ttype: BINARY},
	{input: "_", tag: "msub", output: "_", ttype: INFIX},
	{input: "^", tag: "msup", output: "^", ttype: INFIX},
	{input: "hat", tag: "mover", output: "\u005E", ttype: UNARY, acc: true},
	{input: "bar", tag: "mover", output: "\u00AF", tex: "overline", ttype: UNARY, acc: true},
	{input: "vec", tag: "mover", output: "\u2192", ttype: UNARY, acc: true},
	{input: "dot", tag: "mover", output: ".", ttype: UNARY, acc: true},
	{input: "ddot", tag: "mover", output: "..", ttype: UNARY, acc: true},
	{input: "overarc", tag: "mover", output: "\u23DC", tex: "overparen", ttype: UNARY, acc: true},
	{input: "ul", tag: "munder", output: "\u0332", tex: "underline", ttype: UNARY, acc: true},
	{input: "ubrace", tag: "munder", output: "\u23DF", tex: "underbrace", ttype: UNARYUNDEROVER, acc: true},
	{input: "obrace", tag: "mover", output: "\u23DE", tex: "overbrace", ttype: UNARYUNDEROVER, acc: true},
	{input: "text", tag: "mtext", output: "text", ttype: TEXT},
	{input: "mbox", tag: "mtext", output: "mbox", ttype: TEXT},
	{input: "color", tag: "mstyle", ttype: BINARY},
	{input: "id", tag: "mrow", ttype: BINARY},
	{input: "class", tag: "mrow", ttype: BINARY},
	{input: "cancel", tag: "menclose", output: "cancel", ttype: UNARY},
	AMquote,
	{input: "bb", tag: "mstyle", atname: "mathvariant", atval: "bold", output: "bb", ttype: UNARY},
	{input: "mathbf", tag: "mstyle", atname: "mathvariant", atval: "bold", output: "mathbf", ttype: UNARY},
	{input: "sf", tag: "mstyle", atname: "mathvariant", atval: "sans-serif", output: "sf", ttype: UNARY},
	{input: "mathsf", tag: "mstyle", atname: "mathvariant", atval: "sans-serif", output: "mathsf", ttype: UNARY},
	{input: "bbb", tag: "mstyle", atname: "mathvariant", atval: "double-struck", output: "bbb", ttype: UNARY, codes: AMbbb},
	{input: "mathbb", tag: "mstyle", atname: "mathvariant", atval: "double-struck", output: "mathbb", ttype: UNARY, codes: AMbbb},
	{input: "cc", tag: "mstyle", atname: "mathvariant", atval: "script", output: "cc", ttype: UNARY, codes: AMcal},
	{input: "mathcal", tag: "mstyle", atname: "mathvariant", atval: "script", output: "mathcal", ttype: UNARY, codes: AMcal},
	{input: "tt", tag: "mstyle", atname: "mathvariant", atval: "monospace", output: "tt", ttype: UNARY},
	{input: "mathtt", tag: "mstyle", atname: "mathvariant", atval: "monospace", output: "mathtt", ttype: UNARY},
	{input: "fr", tag: "mstyle", atname: "mathvariant", atval: "fraktur", output: "fr", ttype: UNARY, codes: AMfrk},
	{input: "mathfrak", tag: "mstyle", atname: "mathvariant", atval: "fraktur", output: "mathfrak", ttype: UNARY, codes: AMfrk},
}
