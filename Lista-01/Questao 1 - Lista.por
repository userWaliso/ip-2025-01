programa {
  funcao inicio() {
    real n1, n2, n3, n4, media

    escreva("Digite a nota: ")
    leia(n1)
    escreva("Digite a nota: ")
    leia(n2)
    escreva("Digite a nota: ")
    leia(n3)

    n4 = n1 + n2 + n3
    media = n4 / 3
    escreva("Média = ")
    escreva(media)

    se (media >= 6)
    escreva("\nAPROVADO")
    senao
    escreva ("\nREPROVADO")

   }
  }
