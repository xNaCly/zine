all: biber

biber: main
	biber interpreters

non-silent:
	pdflatex -shell-escape -jobname=interpreters main 

main:
	pdflatex -interaction=batchmode -shell-escape -jobname=interpreters main 

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
