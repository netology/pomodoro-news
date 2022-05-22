package status

type Pomodoro int

const (
	PomodoroStarted   Pomodoro = Pomodoro(1)
	PomodoroPaused             = Pomodoro(2)
	PomodoroCancelled          = Pomodoro(3)
	PomodoroFinished           = Pomodoro(4)
)
