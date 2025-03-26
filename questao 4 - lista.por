programa {
  funcao inicio() {
    real sal, kw, custo_kw, desconto, custo_consumo

    escreva("Digite o salário: ")
    leia(sal)
    escreva("Digite o gasto de energia: ")
    leia(kw)

  custo_kw = sal * 0.7
  custo_consumo = custo_kw * kw
  desconto = custo_consumo * 0.9

  escreva("Custo por Kw: ")
  escreva(custo_kw)
  escreva("\nCusto do consumo: ")
  escreva(custo_consumo)
  escreva("\nCusto do desconto: ")
  escreva(desconto)
   
  }
}
