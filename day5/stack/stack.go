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
	for i := 0; i < itemCount; i++ {
		var top int = 0
		fromStackLen := len(fromStack.items)

		if fromStackLen > 0 {
			top = fromStackLen - 1
		}

		item := (*fromStack).items[top]
		(*toStack).items = append((*toStack).items, item)
		(*fromStack).items = (*fromStack).items[:top]
	}
}

func Reverse(toStack *Stack) {
	stk := (*toStack).items

	for i, j := 0, len(stk)-1; i < j; i, j = i+1, j-1 {
		stk[i], stk[j] = stk[j], stk[i]
	}
}
