package main

type Board struct {
	tokens []int
}

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = -1
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == -1 {
		return "x"
	}
	return ""
}

func (b *Board) judge() bool {
	// 縦列判定
	for i := 0; i < 3; i++ {
		sum := b.tokens[i] + b.tokens[i+3] + b.tokens[i+6]
		if sum == 3 || sum == -3 {
			return true
		}
	}
	return false
}

func main() {

}
