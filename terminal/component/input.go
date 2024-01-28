package component

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Input struct {
	textInput   textinput.Model
	label       string
	value       string
	isDone      bool
	isCancelled bool
	help        help.Model
	width       int
	keyMap      inputHelpKeyMap
	infoValue   string
}

type inputInfoPool struct {
	inputs []*Input
	data   map[string]string
}

func newInputInfoPool(inputs ...*Input) *inputInfoPool {
	return &inputInfoPool{
		inputs: inputs,
	}
}

func (p *inputInfoPool) PutInfoValue(name, value string) *inputInfoPool {
	if nil == p.data {
		p.data = make(map[string]string)
	}

	if _, ok := p.data[name]; ok {
		return p
	}

	p.data[name] = value

	for _, input := range p.inputs {
		input.PutInfoValue(name, value)
	}

	return p
}

type inputHelpKeyMap struct {
	Enter     key.Binding
	Escape    key.Binding
	Quit      key.Binding
	Backspace key.Binding
	Left      key.Binding
	Right     key.Binding
	Help      key.Binding
}

func (k inputHelpKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Enter, k.Escape, k.Backspace, k.Left, k.Right}
}

func (k inputHelpKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Escape, k.Backspace, k.Left, k.Right}, // first column
		{k.Help, k.Quit}, // second column
	}
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
		help:      help.New(),
		keyMap:    initInputHelpKeyMap(),
	}
}

func (i *Input) Init() tea.Cmd {
	return textinput.Blink
}

func (i *Input) Update(msg tea.Msg) (Component, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case CtrlCKeyUp(msg):
			return i, tea.Quit
		case EnterKeyUp(msg):
			i.isDone = true
		case key.Matches(msg, i.keyMap.Help):
			i.help.ShowAll = !i.help.ShowAll
		}
	}

	var cmd tea.Cmd
	i.textInput, cmd = i.textInput.Update(msg)

	return i, cmd
}

func (i *Input) View() string {
	text := i.textInput.Value()
	helpText := i.help.View(i.keyMap)
	height := 8 - strings.Count(text, "\n") - strings.Count(helpText, "\n")

	if len(i.infoValue) == 0 {
		return lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.NewStyle().Foreground(lipgloss.Color("#EDEDED")).Render(i.label),
			lipgloss.JoinVertical(
				lipgloss.Left,
				i.textInput.View(),
				"\n"+strings.Repeat("\n", height)+helpText,
			),
		)
	}

	infoWrap := lipgloss.NewStyle().Foreground(lipgloss.Color("#EDEDED")).Render(i.infoValue)

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.NewStyle().Foreground(lipgloss.Color("#EDEDED")).Render(i.label),
			lipgloss.JoinVertical(
				lipgloss.Left,
				i.textInput.View(),
				"\n"+strings.Repeat("\n", height)+helpText,
			),
		),
		infoWrap,
	)
}

func (i *Input) SetSize(width, _ int) {
	i.width = width
	i.help.Width = width
}

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

func (i *Input) PutInfoValue(name, value string) *Input {
	i.infoValue += "| " + name + ": " + value + "\n"

	return i
}

func initInputHelpKeyMap() inputHelpKeyMap {
	return inputHelpKeyMap{
		Enter:     key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "Submit")),
		Escape:    key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "Cancel")),
		Quit:      key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "Quit")),
		Backspace: key.NewBinding(key.WithKeys("delete"), key.WithHelp("\u232B", "Delete")),
		Left:      key.NewBinding(key.WithKeys("right"), key.WithHelp("←", "Move cursor left")),
		Right:     key.NewBinding(key.WithKeys("left"), key.WithHelp("→", "Move cursor right")),
	}
}
