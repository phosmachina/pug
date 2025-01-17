// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	"bytes"

	pool "github.com/valyala/bytebufferpool"
)

func Jade_mixins(buffer *pool.ByteBuffer) {

	{
		buffer.WriteString(`<ul><li>foo</li><li>bar</li><li>baz</li></ul>`)

	}

	{
		buffer.WriteString(`<ul><li>foo</li><li>bar</li><li>baz</li></ul>`)

	}

	buffer.WriteString(`<ul>`)
	{
		var (
			name = "cat"
		)

		buffer.WriteString(`<li class="pet">`)
		WriteEscString(name, buffer)
		buffer.WriteString(`</li>`)
	}

	{
		var (
			name = "dog"
		)

		buffer.WriteString(`<li class="pet">`)
		WriteEscString(name, buffer)
		buffer.WriteString(`</li>`)
	}

	{
		var (
			name = "pig"
		)

		buffer.WriteString(`<li class="pet">`)
		WriteEscString(name, buffer)
		buffer.WriteString(`</li>`)
	}

	buffer.WriteString(`</ul>`)
	{
		var (
			title = "Hello world"
		)
		var block []byte
		buffer.WriteString(`<div class="article"><div class="article-wrapper"><h1>`)
		WriteEscString(title, buffer)
		buffer.WriteString(`</h1>`)
		if len(block) > 0 {
			buffer.Write(block)
		} else {
			buffer.WriteString(`<p>No content provided</p>`)

		}
		buffer.WriteString(`</div></div>`)
	}

	{
		var (
			title = "Hello world"
		)
		var block []byte
		{
			buffer := new(bytes.Buffer)
			buffer.WriteString(`<p>This is my</p><p>Amazing article</p>`)

			block = buffer.Bytes()
		}

		buffer.WriteString(`<div class="article"><div class="article-wrapper"><h1>`)
		WriteEscString(title, buffer)
		buffer.WriteString(`</h1>`)
		if len(block) > 0 {
			buffer.Write(block)
		} else {
			buffer.WriteString(`<p>No content provided</p>`)

		}
		buffer.WriteString(`</div></div>`)
	}

	{
		var (
			href = "/foo"
			name = "foo"
		)

		attributes := struct{ class string }{class: "btn"}
		buffer.WriteString(`<a class="`)
		WriteEscString(attributes.class, buffer)
		buffer.WriteString(`" href="`)
		WriteEscString(href, buffer)
		buffer.WriteString(`">`)
		WriteEscString(name, buffer)
		buffer.WriteString(`</a>`)
	}

	{
		var (
			href = fn("/foo", "bar", "baz")
			name = "foo"
		)

		attributes := struct{ class string }{class: "btn"}
		buffer.WriteString(`<a class="`)
		WriteEscString(attributes.class, buffer)
		buffer.WriteString(`" href="`)
		WriteAll(href, true, buffer)
		buffer.WriteString(`">`)
		WriteEscString(name, buffer)
		buffer.WriteString(`</a>`)
	}

	{
		var (
			href = "/foo"
			name = "foo"
		)

		buffer.WriteString(`<a href="`)
		WriteEscString(href, buffer)
		buffer.WriteString(`">`)
		WriteEscString(name, buffer)
		buffer.WriteString(`</a>`)
	}

	{
		var (
			title = "Default Title"
		)

		buffer.WriteString(`<div class="article"><div class="article-wrapper"><h1>`)
		WriteEscString(title, buffer)
		buffer.WriteString(`</h1></div></div>`)

	}

	{
		var (
			title = "Hello world"
		)

		buffer.WriteString(`<div class="article"><div class="article-wrapper"><h1>`)
		WriteEscString(title, buffer)
		buffer.WriteString(`</h1></div></div>`)

	}

	buffer.WriteString(`<!--  TODO for string -->`)
	{
		var (
			items = []string{"\"string\"", "2", "3.5", "4"}
			id    = fn("my-list")
		)

		buffer.WriteString(`<ul id="`)
		WriteAll(id, true, buffer)
		buffer.WriteString(`">`)
		for _, item := range items {
			buffer.WriteString(`<li>`)
			WriteEscString(item, buffer)
			buffer.WriteString(`</li>`)
		}
		buffer.WriteString(`</ul>`)
	}

	{
		var (
			foo = "My inner paragraph"
		)

		{
			var (
				bar = foo
			)

			buffer.WriteString(`<p>`)
			WriteEscString(bar, buffer)
			buffer.WriteString(`</p>`)
		}
	}

}
