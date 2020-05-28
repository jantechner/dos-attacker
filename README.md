# dos-attacker
Simple DOS attacker written in Golang. It floods the given host with a given number of requests per minute. If it receives 5 negative responses in a row the attack ends, considered as successful. If after a given timeout the attack does not break the host server, the attack ends with a failure.

## Setup 
Provide execution permission to `attacker` file.

`chmod +x attacker`

## Usage
To run this program simply run 
```golang
./attacker <host> <port> <requestsFrequency> <timeout>
```
`requestsFrequency` - how many requests you want to send per minute

`timeout` - attack duration in seconds

## Example 

```golang
./attacker 127.0.0.1 8080 120 30
```

This example show the attack for localhost on port 8080, where requests are sent every 0.5 second, and the attack ends after 30 seconds.