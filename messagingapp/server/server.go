package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"messagingapp/common"
	"net/http"
	"text/template"
)

// top of page template object
type TopOfPage struct {
	Title  string
	CSS    string
	NavBar []common.Pair
}

// top of page template object
type BottomOfPage struct {
	Title string
}

// create top of page object
var topOfPage TopOfPage = TopOfPage{
	Title: "Messaging App",
	CSS:   "resources/main.css",
	NavBar: []common.Pair{
		{Key: "Messages", Value: "messaging"},
	},
}

// create top of page object
var bottomOfPage BottomOfPage = BottomOfPage{
	Title: "Messaging App",
}

func createAndParseTemplate(obj interface{}, name, templateString string, templateFunctions ...template.FuncMap) string {
	newTemplate := template.New(name)

	for _, function := range templateFunctions {
		newTemplate = newTemplate.Funcs(function)
	}

	newTemplate, err := newTemplate.Parse(templateString)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	if err = newTemplate.Execute(&tpl, obj); err != nil {
		panic(err)
	}

	return tpl.String()
}

func CreateLoginPageHandler(tss *common.Database) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			renderLoginPage(rw, tss)
		case http.MethodPost:
			//newusername := r.FormValue("UserName")
			renderLoginPage(rw, tss)
		}
	}
}

func renderLoginPage(rw http.ResponseWriter, tss *common.Database) {
	// load top of page
	dat, err := ioutil.ReadFile("htmltemplates/top_of_page_template.html")
	common.Check(err)

	// create functions that will be used to parse template
	// nav bar function
	navBarFunc := template.FuncMap{
		"writeBannerElement": func(navBarPair common.Pair) string {
			if navBarPair.Key == "Login" {
				return fmt.Sprintf("<li class=\"active\"><a href=\"%v\">%v</a></li>", navBarPair.Value, navBarPair.Key)
			} else {
				return fmt.Sprintf("<li><a href=\"%v\">%v</a></li>", navBarPair.Value, navBarPair.Key)
			}
		},
	}

	// parse template
	pageString := createAndParseTemplate(topOfPage, "TopOfPage", string(dat), navBarFunc)

	// parse login template
	dat, err = ioutil.ReadFile("htmltemplates/login_app_template.html")
	common.Check(err)

	pageString = pageString + string(dat)

	// load bottom of page
	dat, err = ioutil.ReadFile("htmltemplates/bottom_of_page_template.html")
	common.Check(err)

	pageString = pageString + createAndParseTemplate(bottomOfPage, "BottomOfPage", string(dat))

	fmt.Fprintf(rw, "%s", pageString)
}

func CreateMessagePageHandler(tss *common.Database) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		renderMessagePage(rw, tss)
	}
}

func renderMessagePage(rw http.ResponseWriter, tss *common.Database) {
	// load top of page
	dat, err := ioutil.ReadFile("htmltemplates/top_of_page_template.html")
	common.Check(err)

	// create functions that will be used to parse template
	// nav bar function
	navBarFunc := template.FuncMap{
		"writeBannerElement": func(navBarPair common.Pair) string {
			if navBarPair.Key == "Messages" {
				return fmt.Sprintf("<li class=\"active\"><a href=\"%v\">%v</a></li>", navBarPair.Value, navBarPair.Key)
			} else {
				return fmt.Sprintf("<li><a href=\"%v\">%v</a></li>", navBarPair.Value, navBarPair.Key)
			}
		},
	}

	// parse template
	pageString := createAndParseTemplate(topOfPage, "TopOfPage", string(dat), navBarFunc)

	// parse message template
	dat, err = ioutil.ReadFile("htmltemplates/messaging_app_template.html")
	common.Check(err)

	// create functions that will be used to parse template
	messagingAppFunc := template.FuncMap{
		"writeMessage": func(message common.Message) string {
			return "<dt>" + message.Username + " @ " + message.TimeSent + "</dt><dd>" + message.Message + "<dd>"
		},
	}

	tss.Mut.RLock()
	pageString = pageString + createAndParseTemplate(tss.Messages, "MessagingPart", string(dat), messagingAppFunc)
	tss.Mut.RUnlock()

	// load bottom of page
	dat, err = ioutil.ReadFile("htmltemplates/bottom_of_page_template.html")
	common.Check(err)

	pageString = pageString + createAndParseTemplate(bottomOfPage, "BottomOfPage", string(dat))

	fmt.Fprintf(rw, "%s", pageString)
}
