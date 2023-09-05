# Finder

"Crawling the provided URL, I unearth the shattered remnants of social media links, ripe for the taking by malicious forces. These broken pathways serve as a sinister gateway for attackers, enabling them to orchestrate devastating phishing campaigns that can leave a trail of destruction, not only tarnishing a company's reputation but potentially inflicting severe financial losses.

The vulnerability of these fractured social media connections is so alarming that they are eagerly sought after in the darkest corners of the internet. Cybercriminals eagerly exploit these openings, and the consequences can be nothing short of cataclysmic.

Fear not, for these broken social media hijack issues are often deemed worthy of significant rewards in the shadowy realm of bug bounty programs. But remember, in the world of cyber warfare, time is of the essence, and failure to act swiftly may lead to dire consequences."

```go

    _______           __         
   / ____(_)___  ____/ /__  _____
  / /_  / / __ \/ __  / _ \/ ___/
 / __/ / / / / / /_/ /  __/ /    
/_/   /_/_/ /_/\__,_/\___/_/     
                                 
                                 
```

We are Looking for You

## Instalation

[Install Go on your system](https://go.dev/doc/install)
Run: (go install github.com/ayoubzulfiqar/finder)

## Usage

Finder Required Two Parameters

1. -f : Path of the text file that contains URLs line by line.
2. path : D:\url.txt
3. -w : The number of workers to run (e.g -w 10). The default value is 5. You can increase or decrease this by testing out the capability of your system.

url.txt will contain all the possible URL that needed to scanned.
You can create a local txt file and put the path and scan it.

```go
go run . -f url.txt
```

```go
./finder.exe - f D:/url.txt
```

## Commands

```go
// it will run the code
make run
```

```go
// It will build the exe file
make build
```

Currently, it supports Twitter, Facebook, Instagram and Tiktok without any API keys.
Will Support other Socila Media In Future..
