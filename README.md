# 🔗 Bootstrap Web Scraper

[![Go Report Card](https://goreportcard.com/badge/github.com/anbellouzi/bootstrap-scrape)](https://goreportcard.com/report/github.com/anbellouzi/bootstrap-scrape)

[Proposal](https://github.com/anbellouzi/bootstrap-scrape/blob/master/proposal.md) & [Slides](https://docs.google.com/presentation/d/1wEdtHDYqgtV5jkB8GsICcVsf4w5X6eG8RrrZt4qhrAY/edit?usp=sharing)

## Project Description

Golang Boostrap web scraper that retrives all boostrap names and html components and saves them into an API or file. 


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
    - flags: 
        - filename: 'name of .json or .txt file you want to save components to'
	    - toApi: 'true or false do you want to save components to API?'
	    - toFile: 'true or false do you want to save components to a file?'
	    - data: '.txt file that contains html selectors you want to scrape'
	    - remove: 'true or false do you want to remove all components from API?'
	    - readAPI: 'true or false do you want to get all components from API?'
```




