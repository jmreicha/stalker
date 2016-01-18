package util

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	//"os"
	"strings"
)

var (
	ServerPort = 587
)

// EmailBodyTemplate is used to store variables to display in the email.
type EmailBodyTemplate struct {
	Repos []string
}

// Iterate over tags in email body
var emailBodyTemplate = `<h3>Recently updated repos</h3>
<ul>
{{range $repo := .Repos}}
	<li>{{$repo}}</li>
{{end}}
</ul>`

// extracUserName chops off the @domain portion of an eamil address so the user
// name can be extracted.
func extractUserName(emailaddress string) string {
	Split := strings.Index(emailaddress, "@")
	Username := emailaddress[:Split]
	return Username
}

// AlertNewProjectTag sends an email to a configured address when a project tag
// gets updated.
func AlertNewProjectTag(new_tags []string) {

	// Create email body
	var Body bytes.Buffer
	// Read in configuration to correctly set up emails
	configuration := ReadConfig()

	// Build variables from config to use in email
	EmailServer := configuration.Email.Server
	Username := extractUserName(configuration.Email.Address)
	Password := configuration.Email.Password
	FromAddress := configuration.Email.Address
	ToAddress := configuration.Email.Address
	Subject := "Github repos updated"
	// Create and render email template
	emailBody := EmailBodyTemplate{new_tags}
	t, err := template.New("emailbody").Parse(emailBodyTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(&Body, emailBody)
	if err != nil {
		panic(err)
	}

	// Construct email credentials
	d := gomail.NewPlainDialer(EmailServer, ServerPort, Username, Password)

	// Create the message
	m := gomail.NewMessage()
	m.SetHeader("From", FromAddress)
	m.SetHeader("To", ToAddress)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", Body.String())

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Println("Email sent successfully")
	}
}
