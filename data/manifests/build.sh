#!/bin/bash
echo -e "# Total count\n" > total.md
for x in *.gz; do echo $x && echo $(gunzip -c $x | wc -l) && echo -e "\n"; done >> total.md
echo -e "total" >> total.md
gunzip -c *.gz | wc -l >> total.md
