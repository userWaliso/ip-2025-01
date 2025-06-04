package main
import f "fmt"

func main() {
    var n, m int
    f.Scan(&n, &m)

    trans := make([]int, n*m)

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            var v int
            f.Scan(&v)
            trans[j*n + i] = v
        }
    }

    for i := 0; i < m; i++ {
        base := i * n
        for j := 0; j < n; j++ {
            f.Printf("%d ", trans[base + j])
        }
        f.Println()
    }
}