# 🔗 Bootstrap Web Scraper

[![Go Report Card](https://goreportcard.com/badge/github.com/anbellouzi/bootstrap-scrape)](https://goreportcard.com/report/github.com/anbellouzi/bootstrap-scrape)

[Proposal](https://github.com/anbellouzi/bootstrap-scrape/blob/master/proposal.md) & [Slides](https://docs.google.com/presentation/d/1wEdtHDYqgtV5jkB8GsICcVsf4w5X6eG8RrrZt4qhrAY/edit?usp=sharing)

## Project Description

Golang Boostrap web scraper that retrives all boostrap names and html components and saves them into an API or file. 

![Terminal Output](https://i.imgur.com/jUpKKul.png)

## Project Structure

```bash
📂 bootstrap-scrape
├──proposal.md
├── components.txt
└── scrape.go
```


## Instructions

```bash
git clone https://github.com/anbellouzi/bootstrap-scrape.git
cd bootstrap-scrape
go build scrape.go
./scrape
    -flags: 
        -filename= 'string var:(optional) output file to store results. Default: output.json'
	    -toApi= 'bool var:(optional) to post results to API. Defailt: false'
	    -toFile= 'bool var:(optional) to write results to filename. Default: true'
	    -data= 'string var:(optional) file that contains html selectors. default: components.txt'
	    -remove= 'bool var:(optional) remove all components from API. default: false'
	    -readAPI= 'bool var:(optional) get all components from API. default: false'
```



