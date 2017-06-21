package counters

type alertCounter int

func New(val int) alertCounter {
	return alertCounter(val)
}
