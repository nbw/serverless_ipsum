package main

import (
  "html/template"
	"bytes"
  "strconv"

	"github.com/nbw/serverless_ipsum/ipsum"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Page struct {
    Title string
    Body  []byte
}

type Response events.APIGatewayProxyResponse

var DefaultNum = 200
var HtmlTemplate = `
<html>
  <head>
    <link href="https://fonts.googleapis.com/css?family=Roboto:300" rel="stylesheet">
    <style>
      body {
        font-family: 'Roboto', sans-serif;
        text-rendering: optimizeLegibility;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        max-width: 1100px; 
        display: block;
        margin: auto;
        padding-top: 10px;
      }
    </style>
  </head>
  <body>
    <h1>{{ .Title }}</h1>
    <p>{{printf "%s" .Body}}</p>
  </body>
</html>`

func Handler(req events.APIGatewayProxyRequest) (Response, error) {
    numStr := req.QueryStringParameters["num"]

    num, err := strconv.Atoi(numStr)
    if err != nil {
      num = DefaultNum
    }

    resp := Response{
    StatusCode:      200,
    Body: RenderTemplate(HtmlTemplate, "Animal Ipsum Generator", ipsum.RandomIpsum(num)).String(),
    Headers: map[string]string{
      "Content-Type":           "text/html",
      "X-MyCompany-Func-Reply": "ipsum-handler",
    },
  }

  return resp, nil
}

func RenderTemplate(htmlTemplate, title, body string) (*bytes.Buffer) {
  var tpl bytes.Buffer

  p := &Page{Title: title, Body: []byte(body)}

  t, err := template.New("").Parse(htmlTemplate)
  if err != nil {
    tpl.Write([]byte("Template Parse Error."))
    return &tpl
  }

  err = t.Execute(&tpl, p)
  if err != nil {
    tpl.Write([]byte("Template Execution error."))
    return &tpl
  }

  return &tpl
}

func main() {
  lambda.Start(Handler)
}



