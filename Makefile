build:
	mkdir -p functions
	go get ./...
	go build -o functions/hello-lambda ./...
	# mkdir -p functions/h2
	mkdir -p functions/example.bleve
	# cp functions/hello-lambda functions/h2/h2
	
	cp -r example.bleve functions/example.bleve 
	# echo "raw data" > functions/h2/data.txt
	# zip -rj functions/h2.zip functions/h2
	zip -rj functions/example.bleve.zip functions/example.bleve
	unzip -l functions/example.bleve.zip
	rm -r functions/h2
	ls functions
	ls functions/example.bleve
	#ls functions/example.bleve/example.bleve