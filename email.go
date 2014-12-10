package ws

import (
	"github.com/mattbaird/gochimp"
	"os"
)

func EmailTemplate(templateName string, vars map[string]string) string {
	mandrill, _ := gochimp.NewMandrill(os.Getenv("MANDRILL_KEY"))
	gVars := []gochimp.Var{}

	for k, v := range vars {
		gVar := gochimp.Var{k, v}
		gVars = append(gVars, gVar)
	}

	t, _ := mandrill.TemplateRender(templateName, nil, gVars)

	return t
}

func EmailSend(to string, subject string, html string) error {
	recipients := []gochimp.Recipient{
		gochimp.Recipient{Email: to},
	}
	message := gochimp.Message{
		Html:      html,
		Subject:   subject,
		FromEmail: "no-reply@taxagdl.com",
		FromName:  "TaxaGDL",
		To:        recipients,
	}
	mandrill, _ := gochimp.NewMandrill(os.Getenv("MANDRILL_KEY"))
	_, err := mandrill.MessageSend(message, false)

	return err
}
