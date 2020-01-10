# Markdown parser:

## Proposition to implementation tactics:

1. parse(test string) (Tree, error)
2. render(tree Tree) (string, error)

Try to model Tree with types

## Task:

1. Implement any three rules from markdown:

* em: _Hello_ => <em>Hello</em>
* bold: **Hello** => <b>Hello</b>
* italics: *Hello* => <i>Hello</i>

