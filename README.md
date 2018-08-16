# twitter-auth

Generate access token from CLI

## Installation

```bash
$ go get github.com/k0kubun/twitter-auth
```

## Usage

### Commandline

```bash
# choose your favorite consumer key
$ twitter-auth -k CONSUMER_KEY -s CONSUMER_SECRET

# show help
$ twitter-auth -h
```

### Library

```go
import "github.com/k0kubun/twitter-auth/auth"

func main() {
  credential := &auth.Credential{
    ConsumerKey:    "CONSUMER_KEY",
    ConsumerSecret: "CONSUMER_SECRET",
  }
  accessToken := auth.Authenticate(credential)

  println("AccessToken:", accessToken.Token)
  println("AccessTokenSecret:", accessToken.Secret)
}
```

## License

MIT Lisence
