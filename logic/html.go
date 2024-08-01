package logic

import (
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/cbroglie/mustache"
)

type FormElement struct {
	Type        string
	Name        string
	Label       string
	Placeholder string
	Required    bool
	LabelStyle  string
	Hint        string
	Rows        int
	Multiple    bool
	Options     []Option
	Checks      []Check
	Items       []Item
	TotalCount  string
}

type Option struct {
	Value string
	Label string
}

type Check struct {
	Value string
	Label string
}

type Item struct {
	Value string
	Label string
	Count string
}

type Form struct {
	Header   string
	Tagline  string
	Action   string
	Method   string
	Elements []FormElement
}

type Backend struct {
	Model string
}

type JSONData struct {
	Backend Backend
	Form    Form
}

func GenerateHTML(jsonData JSONData) (string, error) {
	tmpl, err := mustache.ParseFile(
		path.Join("templates", "template.mustache"),
	)
	if err != nil {
		return "", err
	}

	renderedTemplate, err := tmpl.Render(map[string]interface{}{
		"form":       jsonData.Form,
		"formHolder": generateHTMLForm(jsonData),
	})
	if err != nil {
		return "", err
	}

	return renderedTemplate, nil
}

func generateHTMLForm(jsonData JSONData) string {
	formHTML := fmt.Sprintf(`<form id="mainForm" name="mainForm" action="%s" method="%s">`, jsonData.Form.Action, jsonData.Form.Method)
	formHTML += `<div class="card-body">`
	for _, element := range jsonData.Form.Elements {
		formHTML += generateHTMLElement(element)
	}
	formHTML += `</div>`
	formHTML += `
    <div class="card-footer">
        <button class="btn btn-primary" type="submit">Submit</button>
    </div>
    `
	formHTML += `</form>`
	return formHTML
}

func generateHTMLElement(element FormElement) string {
	html := ""
	labelStyle := ""
	if element.LabelStyle != "" {
		labelStyle = element.LabelStyle
	}
	switch element.Type {
	case "text", "password", "email", "tel", "number", "date", "url", "file", "color":
		html += fmt.Sprintf(`<div class="form-group">
            <label class="form-label %s" for="%s">%s</label>
            <input class="form-input" type="%s" name="%s" id="%s" autocomplete="true"`, labelStyle, element.Name, element.Label, element.Type, element.Name, element.Name)
		if element.Placeholder != "" {
			html += fmt.Sprintf(` placeholder="%s"`, element.Placeholder)
		}
		if element.Required {
			html += ` required`
		}
		if element.Type == "date" {
			elementValue := time.Now().Format("2006-01-02")
			html += fmt.Sprintf(` value="%s"`, elementValue)
		}
		html += ` ></input>`
		if element.Hint != "" {
			html += fmt.Sprintf(`<p class="form-input-hint">%s</p>`, element.Hint)
		}
		html += `</div>`

	case "textarea":
		html += `<div class="form-group">`
		html += fmt.Sprintf(`<label class="form-label %s">%s</label>`, labelStyle, element.Label)
		html += fmt.Sprintf(`<textarea id="%s" name="%s" class="form-input"`, element.Name, element.Name)
		if element.Rows > 0 {
			html += fmt.Sprintf(` rows="%d"`, element.Rows)
		}
		if element.Required {
			html += ` required`
		}
		html += ` ></textarea></div>`

	case "select":
		html += `<div class="form-group">`
		html += fmt.Sprintf(`<label class="form-label %s">%s</label>`, labelStyle, element.Label)
		html += fmt.Sprintf(`<select name="%s" id="%s" class="form-select"`, element.Name, element.Name)
		if element.Multiple {
			html += ` multiple`
		}
		html += `>`
		for _, option := range element.Options {
			html += fmt.Sprintf(`<option value="%s">%s</option>`, option.Value, option.Label)
		}
		html += `</select>`
		html += `</div>`

	case "radio":
		html += `<div class="form-group">`
		html += fmt.Sprintf(`<label class="form-label %s">%s</label>`, labelStyle, element.Label)
		for _, option := range element.Options {
			html += `<label class="form-radio">`
			html += fmt.Sprintf(`<input type="radio" id="%s" name="%s" value="%s" />`, element.Name, element.Name, option.Value)
			html += `<i class="form-icon"></i> ` + option.Label + ``
			html += `</label>`
		}
		html += `</div>`

	case "checkbox":
		html += `<div class="form-group">`
		html += `<label class="form-switch">`
		html += fmt.Sprintf(`<input type="checkbox" id="%s" name="%s" >`, element.Name, element.Name)
		html += `<i class="form-icon"></i> ` + element.Label + ``
		html += `</label>`
		html += `</div>`

	case "checkbox-tile":
		html += `<div class="form-group">`
		html += `<label class="form-label ` + labelStyle + `">Sample</label>`
		html += `</div>`
		html += `<div class="tile-checkbox-container mb-2">`
		for _, check := range element.Checks {
			html += `<div class="tile-checkbox mb-2">`
			html += fmt.Sprintf(`<input type="checkbox" id="%s" name="%s">`, check.Value, check.Value)
			html += fmt.Sprintf(`<label for="%s">%s</label>`, check.Value, check.Label)
			html += `</div>`
		}
		html += `</div>`

	case "radio-tile":
		html += `<div class="form-group">`
		html += fmt.Sprintf(`<label class="form-label %s">%s</label>`, labelStyle, element.Label)
		html += `<div class="tile-radio-container mb-2">`
		for _, option := range element.Options {
			html += `<div class="tile-radio mb-2 mt-2">`
			html += fmt.Sprintf(`<input type="radio" id="%s" name="%s" value="%s">`, option.Value, element.Name, option.Value)
			html += fmt.Sprintf(`<label for="%s">%s</label>`, option.Value, option.Label)
			html += `</div>`
		}
		html += `</div>`
		html += `</div>`

	case "list-tile":
		html += `<div class="form-group">`
		html += fmt.Sprintf(`<label class="form-label %s">%s</label>`, labelStyle, element.Label)
		html += `<div class="tile-list-container">`
		for _, item := range element.Items {
			html += `<div class="tile-list-item">` + item.Label + `</div>`
		}
		html += `</div>`
		html += `</div>`

	case "poll":
		html += `<div class="form-group">`
		html += fmt.Sprintf(`<label class="form-label %s">%s</label>`, labelStyle, element.Label)
		html += `<div class="tile-list-container">`
		for _, item := range element.Items {
			count, _ := strconv.Atoi(item.Count)
			totalCount, _ := strconv.Atoi(element.TotalCount)
			percentage := 100 * float64(count) / float64(totalCount)
			strPercentage := fmt.Sprintf("%.2f%%", percentage)
			html += `<div class="tile-list-item-sm">` + item.Label + `</div>`
			html += fmt.Sprintf(`<div class="bar bar-sm">
                <div class="bar-item" role="progressbar" style="width:%s;" aria-valuenow="%f" aria-valuemin="0" aria-valuemax="100"></div>
                </div>`, strPercentage, percentage)
		}
		html += `</div>`
		html += `</div>`
	}
	return html
}
