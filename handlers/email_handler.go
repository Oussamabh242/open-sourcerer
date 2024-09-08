package handlers

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendEmails(body string, to []string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", os.Getenv("SMTPMAIL"), os.Getenv("SMTPPASS"), smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, os.Getenv("SMTPMAIL"), to, []byte(body))
	if err != nil {
		return err
	}
	return nil
}

func SendNewPostAdded(to []string, link, title string) {

	err := SendEmails(fmt.Sprintf(bodynewPost, title, link), to)
	if err != nil {
		log.Println(err)
	}
}

func SendConfirmSubscription(to string, link string) {
	var strto []string
	strto = append(strto, to)
	err := SendEmails(fmt.Sprintf(bodyConfirm, link), strto)
	if err != nil {
		log.Println(err)
	}
}

const bodynewPost = "Subject: New Blog Post Published\r\n" +
	"MIME-version: 1.0;\r\n" +
	"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
	`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
            color: #333;
        }
        .container {
            max-width: 600px;
            margin: 40px auto;
            background-color: #ffffff;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .header {
            background-color: #4CAF50;
            color: white;
            text-align: center;
            padding: 20px 0;
        }
        .header h1 {
            margin: 0;
            font-size: 24px;
        }
        .content {
            padding: 20px;
        }
        .content p {
            font-size: 16px;
            line-height: 1.5;
        }
        .content a.button {
            display: inline-block;
            margin: 20px 0;
            padding: 12px 24px;
            font-size: 16px;
            color: white;
            background-color: #4CAF50;
            text-decoration: none;
            border-radius: 5px;
        }
        .content a.button:hover {
            background-color: #45a049;
        }
        .footer {
            background-color: #f4f4f4;
            color: #888;
            text-align: center;
            padding: 10px;
            font-size: 14px;
        }
    </style>
</head>
<body>

    <div class="container">
        <div class="header">
            <h1>New Blog Post: %s</h1>
        </div>

        <div class="content">
            <p>Hello,</p>
            <p>I've just published a new post on our blog! Click the button below to read the full article:</p>
            <a href="%s" class="button">Read the Post</a>

        </div>

        <div class="footer">
            <p>Open Sourcerer - Oussama Ben Hassen</p>
        </div>
    </div>

</body>
</html>
`

const bodyConfirm = "Subject: Confirm Your Subscription\r\n" +
	"MIME-version: 1.0;\r\n" +
	"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
	`
			<!DOCTYPE html>
			<html lang="en">
			<head>
			    <meta charset="UTF-8">
			    <meta name="viewport" content="width=device-width, initial-scale=1.0">
			    <style>
			        body {
			            font-family: Arial, sans-serif;
			            background-color: #f4f4f4;
			            margin: 0;
			            padding: 0;
			            color: #333;
			        }
			        .container {
			            max-width: 600px;
			            margin: 40px auto;
			            background-color: #ffffff;
			            border-radius: 8px;
			            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
			            overflow: hidden;
			        }
			        .header {
			            background-color: #4CAF50;
			            color: white;
			            text-align: center;
			            padding: 20px 0;
			        }
			        .header h1 {
			            margin: 0;
			            font-size: 24px;
			        }
			        .content {
			            padding: 20px;
			        }
			        .content p {
			            font-size: 16px;
			            line-height: 1.5;
			        }
			        .content a.button {
			            display: inline-block;
			            margin: 20px 0;
			            padding: 12px 24px;
			            font-size: 16px;
			            color: white;
			            background-color: #4CAF50;
			            text-decoration: none;
			            border-radius: 5px;
			        }
			        .content a.button:hover {
			            background-color: #45a049;
			        }
			        .footer {
			            background-color: #f4f4f4;
			            color: #888;
			            text-align: center;
			            padding: 10px;
			            font-size: 14px;
			        }
			    </style>
			</head>
			<body>

			    <div class="container">
			        <div class="header">
			            <h1>Confirm Your Subscription</h1>
			        </div>

			        <div class="content">
			            <p>Hello,</p>
			            <p>Thank you for subscribing to my blog! To complete your subscription, please confirm your email address by clicking the button below.</p>
			            <a href="%s" class="button">Confirm Subscription</a>
			            <p>If you did not subscribe, please ignore this email.</p>
			        </div>

			        <div class="footer">
			            <p>Open Sourcerer - Oussama Ben Hassen</p>
			        </div>
			    </div>

			</body>
			</html>
      `
