package component

import (
	"github.com/charmbracelet/bubbles/list"
)

type ListItem struct {
	title        string
	description  string
	actionSelect Action
	actionBack   Action
}

var _ list.DefaultItem = (*ListItem)(nil)

func NewListItem(title, description string, actionSelect, actionBack Action) ListItem {
	return ListItem{
		title:        title,
		description:  description,
		actionSelect: actionSelect,
		actionBack:   actionBack,
	}
}

func (m ListItem) Title() string {
	return m.title
}

func (m ListItem) Description() string {
	return m.description
}

func (m ListItem) FilterValue() string {
	return m.Title()
}
