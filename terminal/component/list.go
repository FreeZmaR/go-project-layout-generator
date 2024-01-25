package component

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	instance list.Model
	isDone   bool
}

var _ Component = (*List)(nil)

func NewList(title string, items ...ListItem) *List {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = title

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("<-"),
				key.WithHelp("<-", "back"),
			),
		}
	}

	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter", "<-"),
				key.WithHelp("enter", "select"),
				key.WithHelp("<-", "back"),
			),
		}
	}

	l.KeyMap.Filter.Unbind()
	l.KeyMap.GoToEnd.Unbind()
	l.KeyMap.GoToStart.Unbind()
	l.KeyMap.Quit.Unbind()

	for i, item := range items {
		l.InsertItem(i, item)
	}

	return &List{
		instance: l,
		isDone:   false,
	}
}

func (l *List) Init() tea.Cmd {
	return nil
}

func (l *List) Update(msg tea.Msg) (Component, tea.Cmd) {
	var cmd tea.Cmd

	switch t := msg.(type) {
	case tea.KeyMsg:
		l.handelKeyMsg(t)
	}

	l.instance, cmd = l.instance.Update(msg)

	return l, cmd
}

func (l *List) SetSize(width, height int) {
	l.instance.SetSize(width, height)
}

func (l *List) View() string {
	return l.instance.View()
}

func (l *List) IsDone() bool {
	return l.isDone
}

func (l *List) handelKeyMsg(msg tea.KeyMsg) (Component, tea.Cmd) {
	switch true {
	case EnterKeyUp(msg):
		l.handleEnterKey()
	case BackspaceKeyUp(msg):
		l.handleBackKey()
	case CtrlCKeyUp(msg):
		return l, tea.Quit
	}

	return l, nil
}

func (l *List) handleEnterKey() {
	if l.instance.SelectedItem() == nil {
		return
	}

	item, ok := l.instance.SelectedItem().(ListItem)
	if !ok {
		return
	}

	item.actionSelect(l)
	l.isDone = true
}

func (l *List) handleBackKey() {
	if len(l.instance.Items()) == 0 {
		return
	}

	item, ok := l.instance.Items()[0].(ListItem)
	if !ok {
		return
	}

	if nil == item.actionBack {
		return
	}

	item.actionBack(l)
	l.isDone = true
}

func (l *List) SetIsNotDone() {
	l.isDone = false
}
