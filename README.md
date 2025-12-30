
# DarkDarkGo

DarkDarkGo is a concurrent Tor network scraper built with Go. It is specifically designed for Cyber Threat Intelligence (CTI) operations to automate the collection of data from multiple .onion addresses anonymously and efficiently.


## Key Features


- #### Anonymous Routing:
    Securely routes all traffic through a local Tor SOCKS5 proxy to ensure complete anonymity and prevent IP leakage.



- #### High Concurrency:
    Leverages Go's goroutines to    scan hundreds of onion sites simultaneously, significantly reducing collection time.


- #### Resilient Error Handling:
    Sophisticated error management that logs failed or timed-out requests without interrupting the overall scanning process.


- #### Automated Reporting:

    Generates a structured scan_report.log summarizing the status (Success/Failed) of every target.


- #### Structured Storage: 
    Automatically organizes fetched HTML content into separate directories named after the target URL.



- #### YAML-Based Input:
    Easily manage target lists with a clean and simple YAML configuration

  
## Prerequisites

- Go: Version 1.20 or higher installed.

- Tor Service: A Tor background service running on your machine.

- Default expected port: 9050


  
## Installation & Usage



```bash 
  git clone https://github.com/Eneskalin/Dark-Dark-Go.git
```
    
```bash 
  cd DarkDarkGo
  go run main.go

```
### Configure Targets
```bash 
  cd config
       |_ targets.yaml
```

```bash 
  http://check.torproject.org

```



    
