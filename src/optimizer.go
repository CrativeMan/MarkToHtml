package main

import "fmt"

const (
	divider            = "<hr>"
	heading1           = "<h1>%s</h1>"
	heading2           = "<h2>%s</h2>"
	heading3           = "<h3>%s</h3>"
	heading4           = "<h4>%s</h4>"
	heading5           = "<h5>%s</h5>"
	heading6           = "<h6>%s</h6>"
	checkListUnchecked = "<input type=\"checkbox\" disabled> %s<br>"
	checkListChecked   = "<input type=\"checkbox\" disabled checked> %s<br>"
	listItem           = "<li>- %s</li>"
	lineBreak          = "<br>"
	paragraph          = "<p>%s</p>"
	emptyLine          = ""
)

func optimizeMarkdown(markdown []string) ([]string, error) {
	optimized := []string{}

	for _, line := range markdown {
		// Remove empty lines
		if len(line) == 0 {
			optimized = append(optimized, emptyLine)
			continue
		}

		// optimize dividers
		if line == "***" || line == "---" {
			optimized = append(optimized, divider)
			continue
		}

		// optimize todo lists
		if len(line) >= 5 && line[:5] == "- [ ]" {
			list := fmt.Sprintf(checkListUnchecked, line[6:])
			optimized = append(optimized, list)
			continue
		}

		// optimize todo lists
		if (len(line) >= 5 && line[:5] == "- [x]") || (len(line) >= 5 && line[:5] == "- [X]") {
			list := fmt.Sprintf(checkListChecked, line[6:])
			optimized = append(optimized, list)
			continue
		}

		// optimize headings
		if line[:1] == "#" {
			switch line[:2] {
			case "# ":
				heading := fmt.Sprintf(heading1, line[2:])
				optimized = append(optimized, heading)
			case "##":
				heading := fmt.Sprintf(heading2, line[3:])
				optimized = append(optimized, heading)
			case "###":
				heading := fmt.Sprintf(heading3, line[4:])
				optimized = append(optimized, heading)
			case "####":
				heading := fmt.Sprintf(heading4, line[5:])
				optimized = append(optimized, heading)
			case "#####":
				heading := fmt.Sprintf(heading5, line[6:])
				optimized = append(optimized, heading)
			case "######":
				heading := fmt.Sprintf(heading6, line[7:])
				optimized = append(optimized, heading)
			}
			continue
		}

		// default
		text := fmt.Sprintf(paragraph, line)
		optimized = append(optimized, text)
	}

	return optimized, nil
}
