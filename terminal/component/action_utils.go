package component

import tea "github.com/charmbracelet/bubbletea"

type Action func(data any) error

func EnterKeyUp(key tea.KeyMsg) bool {
	return key.String() == "enter"
}

func BackspaceKeyUp(key tea.KeyMsg) bool {
	return key.String() == "backspace" || key.Type == tea.KeyCtrlQuestionMark
}

func CtrlCKeyUp(key tea.KeyMsg) bool {
	return key.String() == "ctrl+c"
}
