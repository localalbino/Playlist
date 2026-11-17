package views

type ViewState int

const (
	HomeView ViewState = iota
	AddView
	RemoveView
	UpdateView
	AuthView
)
