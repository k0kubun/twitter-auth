# twitter-auth

Generate access token from CLI

## Installation

```bash
$ go get github.com/k0kubun/twitter-auth
```

## Usage

### Commandline

```bash
# generate access token by default credential (own your risk)
$ twitter-auth

# choose your favorite consumer key
$ twitter-auth -k CONSUMER_KEY -s CONSUMER_SECRET

# generate access token by public client credential (own your risk)
$ twitter-auth -n "Twitter for iPhone"

# show help
$ twitter-auth -h
```

### Library

```go
import "github.com/k0kubun/twitter-auth/auth"

func main() {
  credential := &auth.Credential{
    ConsumerKey:    options.ConsumerKey,
    ConsumerSecret: options.ConsumerSecret,
  }
  accessToken := auth.Authenticate(credential)

  println("AccessToken:", accessToken.Token)
  println("AccessTokenSecret:", accessToken.Secret)
}
```

## License

MIT Lisence
