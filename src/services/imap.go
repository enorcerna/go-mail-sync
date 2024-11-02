package services

import (
	"go-mail-sync/src/constants"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func createIMAPClient() (*client.Client, error) {
	authMail := constants.AuthMail
	c, err := client.DialTLS(authMail.HostName+":993", nil)
	if err != nil {
		return nil, err
	}
	if err := c.Login(authMail.Email, authMail.Password); err != nil {
		c.Logout() // close sesion
		return nil, err
	}
	return c, nil
}

func GetInbox() ([]map[string]interface{}, error) {
	c, err := createIMAPClient()
	if err != nil {
		log.Println("Failed to create IMAP client:", err)
		return nil, err
	}
	defer c.Logout()

	if _, err := c.Select("INBOX", false); err != nil {
		log.Println("Failed to select INBOX:", err)
		return nil, err
	}

	segSet := new(imap.SeqSet)
	segSet.AddRange(1, 5)
	//fetch messages
	messages := make(chan *imap.Message, 5)
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope}
	go func() {
		if err := c.Fetch(segSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	var mails []map[string]interface{}
	for msg := range messages {
		if msg == nil { // Verifica si el mensaje es nil
			log.Println("Received nil message")
			continue
		}
		mails = append(mails, map[string]interface{}{
			"subject": msg.Envelope.Subject,
			"date":    msg.Envelope.Date,
			// "to":      msg.Envelope.To,
			// "cc":      msg.Envelope.Cc,

		})
	}

	return mails, nil
}
