package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

type Config struct {
	Smtp_server string
	User        string
	Passwd      string
	Hostname    string
}

func main() {
	conf := parseConfig()
	sendMail(conf)

}

func sendMail(conf *Config) {
	client, err := smtp.Dial(conf.Smtp_server)
	if err != nil {
		log.Panic(err)
	}
	a := smtp.PlainAuth("", conf.User, conf.Passwd, conf.Hostname)
	var tlsConf tls.Config
	tlsConf.InsecureSkipVerify = true
	client.StartTLS(&tlsConf)
	err = client.Auth(a)
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	if err = client.Mail(conf.User); err != nil {
		log.Panic(err)
	}

	var recipient string
	fmt.Printf("To: ")
	fmt.Scanf("%s", &recipient)

	if err = client.Rcpt(recipient); err != nil {
		log.Panic(err)
	}

	data, err := client.Data()
	if err != nil {
		log.Panic(err)
	}
	payload := "From:"
	payload += conf.User
	payload += "\r\n"
	var subject string
	fmt.Printf("Subject: ")
	fmt.Scanln(&subject)
	payload += "Subject:"
	payload += subject
	fmt.Printf("Message: ")
	message, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	payload += string(message)
	data.Write([]byte(payload))
	data.Close()
	defer client.Quit()
	fmt.Println()
	fmt.Println("Sent email to " + recipient)

}

func parseConfig() *Config {
	var conf Config
	f, err := os.Open("config.json")
	defer f.Close()
	if err != nil {
		log.Panic(err)
	}
	dec := json.NewDecoder(bufio.NewReader(f))
	err = dec.Decode(&conf)
	if err != nil {
		log.Panic(err)
	}
	return &conf
}
