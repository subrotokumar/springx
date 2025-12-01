package listview

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/subrotokumar/springx/internal/spring"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type ItemType = item

type item struct {
	spring.DependencyDetail
	title, desc string
}

func NewItem(param spring.DependencyDetail) item {
	return item{
		title:            param.Name,
		desc:             param.Description,
		DependencyDetail: param,
	}
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func New(options []ItemType) model {
	inputOptions := []list.Item{}

	for _, v := range options {
		inputOptions = append(inputOptions, v)
	}
	m := model{list: list.New(inputOptions, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Dependencies"
	return m
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *model) View() string {
	return docStyle.Render(m.list.View())
}

func (m model) Run() {
	p := tea.NewProgram(&m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
