package main

// import "code.google.com/p/go-tour/pic"
// import "math/rand"

func Pic(dx, dy int) [][]uint8 {
    res := make([][]uint8, dy)
    for i := range res {
        value := make([]uint8, dx)
        for j := range value {
            // value[j] =  uint8(rand.Intn(256))
            value[j] =  uint8(i^j)
        }
        res[i] = value
    }
    return res
}

func main() {
    Show(Pic)
}
