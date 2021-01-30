package email

import (
    "bytes"
    "fmt"
    "html/template"

    "github.com/pkg/errors"

    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
)

// nameWithHtmlExtension returns the template name with .html extension.
//
// Args:
//   name: the template name
// Return:
//   the template name with .html extension
func nameWithHtmlExtension(name string) string {
    return name + constants.DotHtml
}

// dict functions is used while parsing the templates. It maps the key to
// resource struct key in templates.
func dict(values ...interface{}) (map[string]interface{}, error) {
    if len(values)%2 != 0 {
        return nil, errors.New("Invalid key value mapping")
    }
    dict := make(map[string]interface{}, len(values)/2)
    for i := 0; i < len(values); i += 2 {
        key, ok := values[i].(string)
        if !ok {
            return nil, errors.New("Dict keys must be strings")
        }
        dict[key] = values[i+1]
    }
    return dict, nil
}

// newHtmlTemplate generates the new HTML template.
//
// Args:
//   templateName: the template name
// Return:
//   the new HTML template
func newHtmlTemplate(templateName string) *template.Template {
    t := template.New(nameWithHtmlExtension(templateName))
    keyFuncMap := template.FuncMap{}
    keyFuncMap[constants.FuncMapKey] = dict
    t.Funcs(keyFuncMap)
    return t
}

// ParseHtmlTemplateFiles parses the HTML template files.
//
// Args:
//   template: the template
//   filenames: the file names to be parsed
// Return:
//   template: the template with parsed files
func parseHtmlFiles(template *template.Template, filenames ...string) (*template.Template, error) {
    t, err := template.ParseFiles(filenames...)
    if err != nil {
        return nil, err
    }
    return t, nil
}

// executeTemplate gets the email body.
//
// Args:
//   t: the app template
// Return:
//   emailBody: the email body
//   error: if failed to generate the email body
func (e *email) executeTemplate(t *template.Template) (*bytes.Buffer, error) {
    var emailBody bytes.Buffer
    if err := t.ExecuteTemplate(&emailBody, nameWithHtmlExtension(constants.IntroduceTemplateName), e); err != nil {
        return nil, err
    }
    return &emailBody, nil
}

func (e *email) parseTemplates() (string, error) {
    t := newHtmlTemplate(constants.IntroduceTemplateName)

    t, err := parseHtmlFiles(t, fmt.Sprintf("%s/%s", e.User.IntroTemplate, nameWithHtmlExtension(constants.IntroduceTemplateName)))
    if err != nil {
        getLogger().Error("Error in parsing the HTML files", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return "", err
    }

    emailBody, err := e.executeTemplate(t)
    if err != nil {
        getLogger().Error("Error in replacing the dynamic data in the HTML files", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return "", err
    }
    return emailBody.String(), nil
}
