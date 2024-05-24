package email

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
)

//go:embed templates/*.html
var embeddedTemplates embed.FS

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
		name = name[:len(name)-len(filepath.Ext(name))]
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

type EmailSender interface {
	SendEmail(to, templateName string, data interface{}) error
}

type SMTPEmailSender struct {
	smtpHost         string
	smtpPort         string
	username         string
	password         string
	fromAddress      string
	templateRenderer *TemplateRenderer
}

func NewSMTPEmailSender(host, port, username, password, fromAddress, templateDir string) (*SMTPEmailSender, error) {
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

func (s *SMTPEmailSender) SendEmail(to, subject, templateName string, data interface{}) error {
	body, err := s.templateRenderer.RenderTemplate(templateName, data)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", s.username, s.password, s.smtpHost)
	message := fmt.Sprintf("To: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", to, body)
	if err := smtp.SendMail(s.smtpHost+":"+s.smtpPort, auth, s.fromAddress, []string{to}, []byte(message)); err != nil {
		return err
	}
	return nil
}

type WelcomeData struct {
	Name string
}
