package jade

import (
	"bytes"
	"fmt"
	"strings"
)



// NodeType identifies the type of a parse tree node.
type NodeType int
const (
	NodeList    NodeType = iota
	NodeText
	NodeTag
	NodeAttr
	NodeDoctype
)



type NestNode struct {
	NodeType
	Pos
	tr    *Tree
	Nodes []Node

	typ 	itemType
	Tag     string
	Indent  int
	Nesting int

	id 		string
	class 	[]string
}

func (t *Tree) newNest(pos Pos, tag string, tp itemType, idt, nst int) *NestNode {
	return &NestNode{tr: t, NodeType: NodeTag, Pos: pos, Tag: tag, typ: tp, Indent: idt, Nesting: nst}
}

func (nn *NestNode) append(n Node) {
	nn.Nodes = append(nn.Nodes, n)
}
func (nn *NestNode) tree() *Tree {
	return nn.tr
}
func (nn *NestNode) tp() itemType {
	return nn.typ
}

func (nn *NestNode) String() string {
	// fmt.Printf("%s\t%s\n", itemToStr[nn.typ], nn.Tag)
	b   := new(bytes.Buffer)
	idt := new(bytes.Buffer)

	bgnFormat := "<%s"
	endFormat := "</%s>"

	if nn.typ != itemInlineTag { idt.WriteByte('\n') }

	if nestIndent {
		for i := 0; i < nn.Nesting; i++ {
			idt.WriteString(outputIndent)
		}
	} else {
		for i := 0; i < nn.Indent; i++ {
			idt.WriteByte(' ')
		}
	}

	switch nn.typ {
	case itemDiv:
		nn.Tag = "div"
	case itemComment:
		nn.Tag = "--"
		bgnFormat = "<!%s "
		endFormat = " %s>"
	case itemAction:
		bgnFormat = "{{ %s }}"
	}

	fmt.Fprintf(b, idt.String()+bgnFormat, nn.Tag)

	if len(nn.id) > 0 {
		fmt.Fprintf(b, " id=\"%s\"", nn.id)
	}
	if len(nn.class) > 0 {
		fmt.Fprintf(b, " class=\"%s\"", strings.Join(nn.class, " "))
	}
	var (
		endFmt string
		endFlag bool
	)

	for _, n := range nn.Nodes {
		   tp := n.tp()
		if tp == itemInlineText || tp == itemInlineAction {endFlag = false} else {endFlag = true}
		if tp != itemBlank { fmt.Fprint(b, n) }
	}

	if !endFlag { endFmt = endFormat } else { endFmt = idt.String()+endFormat }
	if nn.typ < itemVoidTag { fmt.Fprintf(b, endFmt, nn.Tag) }

	return b.String()
}

func (nn *NestNode) CopyNest() *NestNode {
	if nn == nil {
		return nn
	}
	n := nn.tr.newNest(nn.Pos, nn.Tag, nn.typ, nn.Indent, nn.Nesting)
	for _, elem := range nn.Nodes {
		n.append(elem.Copy())
	}
	return n
}
func (nn *NestNode) Copy() Node {
	return nn.CopyNest()
}



var doctype = map[string]string {
	"xml" 			: `<?xml version="1.0" encoding="utf-8" ?>`,
	"html" 			: `<!DOCTYPE html>`,
	"5" 			: `<!DOCTYPE html>`,
	"1.1" 			: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
	"xhtml" 		: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
	"basic" 		: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN" "http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">`,
	"mobile" 		: `<!DOCTYPE html PUBLIC "-//WAPFORUM//DTD XHTML Mobile 1.2//EN" "http://www.openmobilealliance.org/tech/DTD/xhtml-mobile12.dtd">`,
	"strict" 		: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`,
	"frameset" 		: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">`,
	"transitional" 	: `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`,
	"4" 			: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
	"4strict" 		: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`,
	"4frameset" 	: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">`,
	"4transitional" : `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN" "http://www.w3.org/TR/html4/frameset.dtd"> `,
}

type DoctypeNode struct {
	NodeType
	Pos
	tr    *Tree
	Doctype string
}

func (t *Tree) newDoctype(pos Pos, dt string) *DoctypeNode {
	return &DoctypeNode{tr: t, NodeType: NodeDoctype, Pos: pos, Doctype: dt}
}

func (d *DoctypeNode) String() string {
	if dt, ok := doctype[d.Doctype]; ok {
		return fmt.Sprintf("%s", dt)
	}
	return fmt.Sprintf("<!DOCTYPE html>")
}

func (d *DoctypeNode) tp() itemType {
	return itemDoctype
}
func (d *DoctypeNode) tree() *Tree {
	return d.tr
}
func (d *DoctypeNode) Copy() Node {
	return &DoctypeNode{tr: d.tr, NodeType: NodeDoctype, Pos: d.Pos, Doctype: d.Doctype}
}



// LineNode holds plain text.
type LineNode struct {
	NodeType
	Pos
	tr   *Tree

	Text []byte // The text; may span newlines.
	typ 	itemType
	Indent  int
	Nesting int
}

func (t *Tree) newLine(pos Pos, text string, tp itemType, idt, nst int) *LineNode {
	return &LineNode{tr: t, NodeType: NodeText, Pos: pos, Text: []byte(text), typ: tp, Indent: idt, Nesting: nst}
}

func (tx *LineNode) String() string {
	// fmt.Printf("%s\t%s\n", itemToStr[tx.typ], tx.Text)
	idt := new(bytes.Buffer)

	lnFormat := "%s"
	if lineIndent {
		for i := 0; i < tx.Nesting; i++ {
			idt.WriteString(outputIndent)
		}
	} else {
		for i := 0; i < tx.Indent; i++ {
			idt.WriteByte(' ')
		}
	}

	switch tx.typ {
	case itemText:
		lnFormat = "\n"+idt.String()+lnFormat
	case itemHtmlTag:
		lnFormat = "\n"+idt.String()+lnFormat
	case itemInlineText:
	case itemInlineAction:
		lnFormat = "{{%s }}"
	}

	return fmt.Sprintf( lnFormat, tx.Text )
}

func (tx *LineNode) tp() itemType {
	return tx.typ
}
func (tx *LineNode) tree() *Tree {
	return tx.tr
}
func (tx *LineNode) Copy() Node {
	return &LineNode{tr: tx.tr, NodeType: NodeText, Pos: tx.Pos, Text: append([]byte{}, tx.Text...), typ: tx.typ, Indent: tx.Indent, Nesting: tx.Nesting}
}


type AttrNode struct {
	NodeType
	Pos
	tr    *Tree
	Attr  string
	typ   itemType
}

func (t *Tree) newAttr(pos Pos, attr string, tp itemType) *AttrNode {
	return &AttrNode{tr: t, NodeType: NodeAttr, Pos: pos, Attr: attr, typ: tp}
}

func (a *AttrNode) String() string {
	switch a.typ {
	case itemEndAttr:
		return fmt.Sprintf( "%s", a.Attr )
	case itemAttr:
		return fmt.Sprintf( "=%s", a.Attr )
	case itemAttrN:
		return fmt.Sprintf( "=\"%s\"", a.Attr )
	case itemAttrVoid:
		return fmt.Sprintf( " %s=\"%s\"", a.Attr, a.Attr )
	default:
		return fmt.Sprintf( " %s", a.Attr )
	}
}

func (a *AttrNode) tp() itemType {
	return a.typ
}
func (a *AttrNode) tree() *Tree {
	return a.tr
}
func (a *AttrNode) Copy() Node {
	return &AttrNode{tr: a.tr, NodeType: NodeAttr, Pos: a.Pos, Attr: a.Attr, typ: a.typ}
}