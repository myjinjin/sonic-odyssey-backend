package email

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

//go:embed templates/*.html
var embeddedTemplates embed.FS

const (
	TemplateWelcome       = "welcome.html"
	TemplatePasswordReset = "password_reset.html"
)

type WelcomeData struct {
	Name string
}

type PasswordResetData struct {
	Name      string
	ResetLink string
}

type EmailSender interface {
	SendEmail(to, templateName string, data interface{}) error
}

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	renderer := &TemplateRenderer{
		templates: make(map[string]*template.Template),
	}

	err := renderer.loadTemplates()
	if err != nil {
		return nil, err
	}

	return renderer, nil
}

func (r *TemplateRenderer) loadTemplates() error {
	templates, err := template.ParseFS(embeddedTemplates, "templates/*.html")
	if err != nil {
		return err
	}

	for _, t := range templates.Templates() {
		name := filepath.Base(t.Name())
		r.templates[name] = t
	}

	return nil
}

func (r *TemplateRenderer) RenderTemplate(name string, data interface{}) (string, error) {
	t, ok := r.templates[name]
	if !ok {
		return "", fmt.Errorf("template %s not found", name)
	}

	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

type SMTPEmailSender struct {
	smtpHost         string
	smtpPort         string
	username         string
	password         string
	fromAddress      string
	templateRenderer *TemplateRenderer
}

func NewSMTPEmailSender(host, port, username, password, fromAddress string) (*SMTPEmailSender, error) {
	renderer, err := NewTemplateRenderer()
	if err != nil {
		return nil, err
	}

	return &SMTPEmailSender{
		smtpHost:         host,
		smtpPort:         port,
		username:         username,
		password:         password,
		fromAddress:      fromAddress,
		templateRenderer: renderer,
	}, nil
}

func (s *SMTPEmailSender) SendEmail(to, templateName string, data interface{}) error {
	body, err := s.templateRenderer.RenderTemplate(templateName, data)
	if err != nil {
		return err
	}

	subject, err := extractTitle(body)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	message := []byte(fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("From: %s\r\n", s.fromAddress) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		body)

	err = smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.fromAddress, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}

func extractTitle(body string) (string, error) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return "", err
	}

	var title string
	var traverse func(n *html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				title = n.FirstChild.Data
			}
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return title, nil
}
