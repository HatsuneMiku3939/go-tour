package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    Walker(t.Left, ch)
    ch <- t.Value
    Walker(t.Right, ch)

    close(ch)
}

func Walker(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }

    Walker(t.Left, ch)
    ch <- t.Value
    Walker(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    same := true

    for {
        val1, ok1 := <- ch1
        val2, ok2 := <- ch2

        if !ok1 || !ok2 {
            return ok1 == ok2 && same
        }

        fmt.Println(val1, val2)

        if val1 != val2 {
            same = false
        }
    }
    return true
}

func main() {
    fmt.Println("Same(tree.New(1), tree.New(1))")
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println("\n")

    fmt.Println("Same(tree.New(1), tree.New(2))")
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
