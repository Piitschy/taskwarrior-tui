package utils

var Suggestions = []string{"add", "annotate", "append", "denotate", "done", "duplicate", "edit", "export", "import", "log", "long", "modify", "prepend", "start", "stop", "undo", "unmodify", "untag", "add project:"}

func AddProjectSuggestions(suggestions, projects []string) []string {
	for _, project := range projects {
		suggestions = append(suggestions, "add project:"+project)
	}
	return suggestions
}

func ProjectSuggestions(projects []string) []string {
	suggestions := []string{""}
	for _, project := range projects {
		suggestions = append(suggestions, "project:"+project)
	}
	return suggestions
}
