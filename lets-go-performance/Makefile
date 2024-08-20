all: biber

biber: main
	biber performance

non-silent:
	pdflatex -shell-escape -jobname=performance main 

main:
	pdflatex -interaction=batchmode -shell-escape -jobname=performance main 

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
