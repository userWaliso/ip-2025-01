programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    inteiro c
    real Temp, A
    escreva("Digite o número de temperaturas a serem calculadas: ")
    leia(A)
    real vetor[A]
    para (c = 0; c < A; c = c + 1)
  {   escreva("Digite a temperatura em fahrenheit: ")
    leia(Temp)
    vetor[c] = (5 / 9) * (Temp - 32)}
    para (c = 0; c < A; c = c + 1) {
      escreva("TEMPERATURA EQUIVALE A: ", mat.arredondar(vetor[c],2), " ", "CELSIUS", "\n")
    }
    } 
  }