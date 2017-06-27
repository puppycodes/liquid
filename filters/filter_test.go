package filters

import (
	"fmt"
	"testing"

	"github.com/osteele/liquid/expressions"
	"github.com/stretchr/testify/require"
)

func init() {
	DefineStandardFilters()
}

var filterTests = []struct{ in, expected string }{
	// Jekyll extensions
	{`obj | inspect`, `{"a":1}`},

	// filters
	// product_price | default: 2.99 }}

	// list filters
	// site.pages | map: 'category' | compact | join "," %}
	// {% assign my_array = "apples, oranges, peaches, plums" | split: ", " %}my_array.first }}
	{`"John, Paul, George, Ringo" | split: ", " | join: " and "`, "John and Paul and George and Ringo"},
	{`animals | sort | join: ", "`, "Sally Snake, giraffe, octopus, zebra"},
	{`sort_prop | sort: "weight" | inspect`, `[{"weight":null},{"weight":1},{"weight":3},{"weight":5}]`},

	// last, map, slice, sort_natural, reverse, size, uniq

	// string filters
	// "/my/fancy/url" | append: ".html"
	// {% assign filename = "/index.html" %}"website.com" | append: filename

	// "title" | capitalize
	// "my great title" | capitalize

	// "Parker Moore" | downcase

	// "Have you read 'James & the Giant Peach'?" | escape
	// "1 < 2 & 3" | escape_once
	// "1 &lt; 2 &amp; 3" | escape_once

	// lstrip, newline_to_br, prepend, remove, remove_first, replace, replace_first
	// rstrip, split, strip, strip_html, strip_newlines, truncate, truncatewords, upcase
	// url_decode, url_encode

	// number filters
	// -17 | abs
	// 4 | abs
	// "-19.86" | abs

	// 1.2 | ceil
	// 2.0 | ceil
	// 183.357 | ceil
	// "3.5" | ceil

	// 16 | divided_by: 4
	// 5 | divided_by: 3
	// 20 | divided_by: 7.0

	// 1.2 | floor
	// 2.0 | floor
	// 183.357 | floor
	// minus, modulo, plus, round,times

	// date filters
	// article.published_at | date: "%a, %b %d, %y"
	// article.published_at | date: "%Y"
	// "March 14, 2016" | date: "%b %d, %y"
	// "now" | date: "%Y-%m-%d %H:%M" }
}

var filterTestContext = expressions.NewContext(map[string]interface{}{
	"x": 123,
	"obj": map[string]interface{}{
		"a": 1,
	},
	"animals": []string{"zebra", "octopus", "giraffe", "Sally Snake"},
	"pages": []map[string]interface{}{
		{"category": "business"},
		{"category": "celebrities"},
		{},
		{"category": "lifestyle"},
		{"category": "sports"},
		{},
		{"category": "technology"},
	},
	"sort_prop": []map[string]interface{}{
		{"weight": 1},
		{"weight": 5},
		{"weight": 3},
		{"weight": nil},
	},
	"ar": []string{"first", "second", "third"},
	"page": map[string]interface{}{
		"title": "Introduction",
	},
})

func TestFilters(t *testing.T) {
	for i, test := range filterTests {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			val, err := expressions.EvaluateExpr(test.in, filterTestContext)
			require.NoErrorf(t, err, test.in)
			actual := fmt.Sprintf("%s", val)
			require.Equalf(t, test.expected, actual, test.in)
		})
	}
}