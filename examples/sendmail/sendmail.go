//send email through smtp
package main

import(
  "bitbucket.org/ekanna/smtpmail"
  "fmt"
)

func main(){
	a := smtpmail.Auth{"me@example.com", "password", "smtp.example.com", "465"} //here we pass the values directly
 //a, _ := smtpmail.Config("./cfg.json")

  email := smtpmail.Email{
      To:      "receiver@example.com",
      Subject: "Hi ",
      Body:    "Hi , How are you?",
  }

  err := a.SendMail(email)
  if err != nil {
    fmt.Println(err)
  }
}

/*

//create a separate file in prj folder ,name it as cfg.json  and save a below code 

{
"UserName": "user@gmail.com",
"Password": "password",
"Host":"smtp.gmail.com",
"Port":465
}

*/
