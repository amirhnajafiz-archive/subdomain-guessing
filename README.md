<h1 align="center">
Subdomain Guessing
</h1>

Guessing a targets subdomains and other DNS records. All you need to do
is creating a wordlist, finding a DNS address and find the source ip address
or domain name. 

## Parameters

### Wordlist

Program uses a wordlist to start searching for subdomains. You can create
a wordlist of any words that you like. Remember that if you use
more words in your wordlist, you have a better chance of finding subdomains
of your target. 

It's better to use the words related to domain company that you
are searching for its subdomains.

Create a new wordlist:

```shell
touch wordlist.txt
```

```shell
wc -l wordlist.txt
```

### DNS address

In order to find the subdomains of a domain, you need to choose
a DNS server. A good DNS server address is google DNS server ```8.8.8.8:53```.
For having a better chance in finding subdomains of domain, you can 
find DNS servers that have the server data in them.

### Domain name or IP

The most important parameter is domain information. You need to
give the domain name or domain IP address. 

### Number of workers (optional)

Program runs with workers to create a concurrency in processing
words in wordlist. You can also set the number of workers.

## Execute

Run the application with following command:

```shell
go run main.go -domain microsoft.com -wordlist wordlist.txt -server 8.8.8.8:53 -c 100
```
