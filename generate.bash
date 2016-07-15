#!/bin/bash

# Converts Go source code to annotated HTML

set -euo pipefail
IFS=$'\n\t'

for f in $(ls src); do
	out=${f}.html
	
	title=$(awk 'NR==2' src/$f)
	comments=$(awk 'NR==4' src/$f)

	# write header with title
	cat tmpl/head.html | sed "s%Edges of Go%$title%" > pages/$out

	# 1. Remove comment link number
	# 2. Pygmentize
	# 3. Mark up comments similar to GoDoc. This regex will totally
	# break if there are <> in the go code.
	#
	# Comment markup:
	# Headings are comments lines without punctuation
	# URLs are linked
	# Output example starts with //=
	# Commented code lacks the space after the //

	cat src/$f \
		| sed '4d' \
		| pygmentize -l go -f html \
		| sed 's%"cm">\([^:?!/=,.]*\)<%><h1>\1</h1><%' \
		| sed 's%"c1">\([^:?!=,.]*\)<%><h2>\1</h2><%' \
		| sed 's%>\(// \)%><i>\1</i>%' \
		| sed 's%>\(//= \)\(.*\)<%><i>\1</i><output>\2</output><%' \
		| sed 's%>\(//\)\(\w.*\)<%><b>\1</b><code>\2</code><%' \
		| sed 's%>\(/\*\)%><i>\1</i>%' \
		| sed 's%>\(\*/\)%><i>\1</i>%' \
		| sed 's%\(http.*\)<%<a href="\1">\1</a><%' \
		>> pages/$out
	# Did you catch that? I just manipulated HTML with regex, what terrible practice!
	 
	# write footer with comment link
	sed "s%https://github.com/JonahBraun/edges/issues/%$comments%" tmpl/foot.html >> pages/$out
done
