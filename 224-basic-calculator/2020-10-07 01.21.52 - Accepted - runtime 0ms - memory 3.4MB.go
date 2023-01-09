func calculate(s string) int {
    sum := 0
    sign := 1
    number := 0
    stack := New()
    for _, v := range s {
        if v == '+' {
            sum += number * sign
            number = 0
            sign = 1
        }else if v == '-' {
            sum += number * sign
            number = 0
            sign = -1
        }else if v >= '0' && v <= '9' {
            number = number * 10 + int(v - '0') 
        }else if v == '(' {
            stack.Push(sign)
            stack.Push(sum)
            sum = 0
            sign = 1
        }else if v == ')' {
            sum += number * sign
            number = 0
            sum = stack.Pop().(int) + sum * stack.Pop().(int)
        }
    }
    sum += number * sign
    return sum
}

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}	
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}