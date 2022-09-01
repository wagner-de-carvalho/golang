package formas

import(
	"testing"
	"math"
)

// Fatalf = ao gerar erro, para a execução de testes
// Errorf = executa todos os testes, mesmo após erros
func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Área do retângulo recebida = %f é diferente da área = %f esperada", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * (10 * 10))
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Área do círculo recebida = %f é diferente da área = %f esperada", areaRecebida, areaEsperada)
		}
	})

}
