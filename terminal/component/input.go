package component

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Input struct {
	textInput   textinput.Model
	label       string
	value       string
	isDone      bool
	isCancelled bool
}

var _ Component = (*Input)(nil)

func NewInput(label, placeholder string) *Input {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()

	return &Input{
		textInput: ti,
		label:     label,
		isDone:    false,
	}
}

func (i *Input) Init() tea.Cmd {
	return textinput.Blink
}

func (i *Input) Update(msg tea.Msg) (Component, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch true {
		case CtrlCKeyUp(msg):
			return i, tea.Quit
		case EnterKeyUp(msg):
			i.isDone = true
		}
	}

	var cmd tea.Cmd
	i.textInput, cmd = i.textInput.Update(msg)

	return i, cmd
}

func (i *Input) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#EDEDED")).Render(i.label),
		i.textInput.View(),
	)
}

func (i *Input) SetSize(_, _ int) {}

func (i *Input) IsDone() bool {
	return i.isDone
}

func (i *Input) IsCancelled() bool {
	return i.isCancelled
}

func (i *Input) Reset() {
	i.textInput.Reset()
}

func (i *Input) Value() string {
	return i.textInput.Value()
}
