build:
	mkdir -p functions
	go get ./...
	go build -o functions/hello-lambda ./...
	# mkdir -p functions/h2
	# cp functions/hello-lambda functions/h2/h2
	cp -r example.bleve functions/example.bleve/example.bleve
	# echo "raw data" > functions/h2/data.txt
	# zip -rj functions/h2.zip functions/h2
	# unzip -l functions/h2.zip
	# rm -r functions/h2
	ls functions
	ls functions/example.bleve
	ls functions/example.bleve/example.bleve