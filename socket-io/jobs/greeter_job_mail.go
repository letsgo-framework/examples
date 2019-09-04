package jobs

import (
	"github.com/letsgo-framework/examples/socket-io/mail"
)

func GreetingMail()  {
	mail.SendMail([]string{"greet@letsgo.com"}, "Greetings", "greeter-template.html", struct {
		App string
	}{
		App: "LetsGO",
	})
}
