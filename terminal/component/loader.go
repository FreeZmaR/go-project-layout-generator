package component

import (
	"context"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"sync"
)

type Loader struct {
	ctx              context.Context
	cancelFN         context.CancelFunc
	spinner          spinner.Model
	text             string
	inProgress       bool
	isDone           bool
	process          LoaderProcess
	onCompleteAction Action
	onErrorAction    Action
	locker           sync.RWMutex
}

type LoaderProcess func(ctx context.Context) error

var _ Component = (*Loader)(nil)

func NewLoader(text string, process LoaderProcess, onComplete, onError Action) *Loader {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	s.Spinner = spinner.Dot

	ctx, cancel := context.WithCancel(context.Background())

	return &Loader{
		ctx:              ctx,
		cancelFN:         cancel,
		spinner:          s,
		text:             text,
		process:          process,
		onErrorAction:    onError,
		onCompleteAction: onComplete,
		inProgress:       false,
		isDone:           false,
	}
}

func (l *Loader) Init() tea.Cmd {
	return l.spinner.Tick
}

func (l *Loader) Update(msg tea.Msg) (Component, tea.Cmd) {
	if l.ctx.Err() != nil {
		return l, tea.Quit
	}

	l.locker.Lock()
	if !l.inProgress {
		l.inProgress = true

		go l.runProcess()
	}
	l.locker.Unlock()

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

func (l *Loader) View() string {
	l.locker.Lock()
	defer l.locker.Unlock()

	if l.ctx.Err() != nil {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render("Done")
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		l.spinner.View(),
		lipgloss.NewStyle().
			PaddingLeft(1).
			Foreground(lipgloss.Color("205")).
			Render(l.text),
	)
}

func (l *Loader) IsDone() bool {
	return l.isDone
}

func (l *Loader) SetSize(_, _ int) {}

func (l *Loader) runProcess() {
	err := l.process(l.ctx)

	l.locker.Lock()
	l.isDone = true
	l.cancelFN()
	l.locker.Unlock()

	if err != nil {
		_ = l.onErrorAction(err)

		return
	}

	_ = l.onCompleteAction(nil)
}
