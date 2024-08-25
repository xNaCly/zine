all: main

non-silent:
	pdflatex -shell-escape -jobname=hashmap main 

main:
	pdflatex -interaction=batchmode -shell-escape -jobname=hashmap main 

clean:
	rm -fr \
		*.toc \
		**/**.aux \
		*.aux \
		*.nav \
		*.out \
		*.snm \
		*.toc \
		*.log \
		*.bbl \
		*.blg \
		*.bcf \
		*.run.xml
