#!/bin/bash
echo -e "# Total count\n" > total.md
for x in *.gz; do echo $x && echo $(gunzip -c $x | wc -l) && echo -e "\n"; done >> total.md