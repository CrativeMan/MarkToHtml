package main

func parseToStandalone(markdown []string) ([]string, error) {
	html := generateBlankHtml()
	html = append(html, markdown...)
	html = append(html, "</body>", "</html>")

	return html, nil
}

func generateBlankHtml() []string {
	return []string{"<!DOCTYPE html>", "<html>", "<head>", "<title>Markdown to HTML</title>", "</head>", "<body>"}
}
