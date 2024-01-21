package terminal

import (
	"context"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Loader struct {
	ctx      context.Context
	cancelFN context.CancelFunc
	spinner  spinner.Model
	style    lipgloss.Style
	text     string
}

var _ tea.Model = (*Loader)(nil)

func NewLoader(ctx context.Context, cancelFN context.CancelFunc, style lipgloss.Style, text string) *Loader {
	s := spinner.New()
	//s.Style = style.Foreground(lipgloss.Color("205"))
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	s.Spinner = spinner.Dot

	return &Loader{
		ctx:      ctx,
		cancelFN: cancelFN,
		spinner:  s,
		style:    style,
		text:     text,
	}
}

func (l Loader) Init() tea.Cmd {
	return l.spinner.Tick
}

func (l Loader) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if l.ctx.Err() != nil {
		return l, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			l.cancelFN()

			return l, tea.Quit
		}

		return l, nil

	default:
		var cmd tea.Cmd
		l.spinner, cmd = l.spinner.Update(msg)

		return l, cmd
	}
}

func (l Loader) View() string {
	if l.ctx.Err() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render("Done")
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		l.spinner.View(),
		lipgloss.NewStyle().
			PaddingLeft(1).
			Foreground(lipgloss.Color("205")).
			Background(l.style.GetBackground()).
			Render(l.text),
	)
}
