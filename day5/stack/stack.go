package stack

type Stack struct {
	items []string
}

func NewStack() *Stack {
	return &Stack{}
}

func Add(toStack *Stack, item string) {
	(*toStack).items = append((*toStack).items, item)
}

func Move(itemCount int, fromStack *Stack, toStack *Stack) {
	cutoff := len((*fromStack).items) - itemCount
	(*toStack).items = append((*toStack).items, (*fromStack).items[cutoff:]...)
	(*fromStack).items = (*fromStack).items[:cutoff]
}

func Reverse(toStack *Stack) {
	stk := (*toStack).items

	for i, j := 0, len(stk)-1; i < j; i, j = i+1, j-1 {
		stk[i], stk[j] = stk[j], stk[i]
	}
}
