programa {
  funcao inicio() {
    inteiro n1, n2 ,n3, pt

    escreva("Digite o número: ")
    leia(n1)
    escreva("Digite o número: ")
    leia(n2)
    escreva("Digite o número: ")
    leia(n3)
      se (n1 >= 10 ou n1 == 0 ou n2 >= 10 ou n2 == 0 ou n3 >= 10 ou n3 == 0){
        escreva("DIGITO INVÁLIDO")
      } senao {
        escreva(100 * n1 + 10 * n2 + n3)

        pt = 100 * n1 + 10 * n2 + n3
        escreva(" ")
        escreva(pt * pt)
      }
  }
}
