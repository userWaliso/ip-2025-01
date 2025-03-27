programa {
  inclua biblioteca Matematica --> mat
  real t, m, c, mm
  funcao inicio() {
    
    escreva("Digite a temperatura em °F: ")
    leia(t)
    escreva("Digite o valor em polegadas: ")
    leia(m)

    c = (5*t-160)/9
    mm = m * 25.4

    escreva("O VALOR EM CELSIUS: ", mat.arredondar(c,2), "\n")
    escreva("A QUANTIDADE DE CHUVA É: ", mat.arredondar(mm,2))
  }
}
