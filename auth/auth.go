package auth

import (
	"bufio"
	"github.com/mrjones/oauth"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func Authenticate(credential *Credential) *oauth.AccessToken {
	consumer := NewTwitterConsumer(credential.ConsumerKey, credential.ConsumerSecret)
	requestToken, url, err := consumer.GetRequestTokenAndUrl("")
	if err != nil {
		log.Fatal(err)
	}

	// openUrl(url)
	log.Print(url)
	pin := ReadPin()

	accessToken, err := consumer.AuthorizeToken(requestToken, pin)
	if err != nil {
		log.Fatal(err)
	}
	return accessToken
}

func NewTwitterConsumer(consumerKey, consumerSecret string) *oauth.Consumer {
	return oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
}

func ReadPin() string {
	print("PIN: ")

	stdin := bufio.NewReader(os.Stdin)
	input, err := stdin.ReadBytes('\n')
	if err != nil {
		log.Fatal("canceled")
	}

	if input[len(input)-2] == '\r' {
		input = input[0 : len(input)-2]
	} else {
		input = input[0 : len(input)-1]
	}
	return string(input)
}

func openUrl(url string) {
	cmd := "xdg-open"
	args := []string{cmd, url}
	if runtime.GOOS == "windows" {
		cmd = "rundll32.exe"
		args = []string{cmd, "url.dll,FileProtocolHandler", url}
	} else if runtime.GOOS == "darwin" {
		cmd = "open"
		args = []string{cmd, url}
	} else if runtime.GOOS == "plan9" {
		cmd = "plumb"
	}

	cmd, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatal("command not found:", err)
	}

	p, err := os.StartProcess(cmd, args, &os.ProcAttr{Dir: "", Files: []*os.File{nil, nil, os.Stderr}})
	if err != nil {
		log.Fatal("failed to start command:", err)
	}
	defer p.Release()
}
