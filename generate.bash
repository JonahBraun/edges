#!/bin/bash

# What, were you expecting a Go app?
# When hanging a poster, only simple tape is needed.

set -euo pipefail
IFS=$'\n\t'

cat tmpl/head.html tmpl/welcome.html > index.html
echo "<nav>" >> index.html
 
for f in $(ls src); do
	out=${f}.html
	title=$(awk 'NR==2' src/$f)
	echo "<a href='/pages/$out'>$title</a>" >> index.html

	cat tmpl/head.html > pages/$out
	pygmentize -f html src/$f \
		| sed 's%"cm">\([^?/!,.]*\)<%><h1>\1</h1><%' \
		| sed 's%"c1">\([^?/!,.]*\)<%><h2>\1</h2><%' \
		| sed 's%>\(// \)%><i>\1</i>%' \
		| sed 's%>\(//= \)\(.*\)<%><i>\1</i><output>\2</output><%' \
		| sed 's%>\(/\*\)%><i>\1</i>%' \
		| sed 's%>\(\*/\)%><i>\1</i>%' \
		| sed 's%\(http.*\)<%<a href="\1">\1</a><%' \
		>> pages/$out
	# Did you catch that? I just manipulated HTML with regex, what terrible practice!
	 
	cat tmpl/foot.html >> pages/$out
done

echo "</nav>" >> index.html
cat tmpl/foot.html >> index.html