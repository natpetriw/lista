package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	LISTA_VACIA = "La lista esta vacia"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, LISTA_VACIA, func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, LISTA_VACIA, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, LISTA_VACIA, func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}
func TestLista1Elemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	lista.InsertarUltimo(3)
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestIteradorInternoListaConElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	elementos := []int{}
	lista.Iterar(func(e int) bool {
		elementos = append(elementos, e)
		return true
	})

	require.Equal(t, []int{1, 2, 3}, elementos, "Deberían recorrerse todos los elementos")
}

func TestIteradorInternoCortarRecorrido(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	elementos := []int{}
	lista.Iterar(func(e int) bool {
		elementos = append(elementos, e)
		return e != 2
	})

	require.EqualValues(t, []int{1, 2}, elementos, "El recorrido debería cortarse cuando visitar devuelve false")
}
func TestInsertarPrincipioVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 50; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	for i := 0; i < 50; i++ {
		require.EqualValues(t, i, iterador.VerActual())
		iterador.Siguiente()
	}
}
func TestInsertarFinalConIrerador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(44)
	require.EqualValues(t, 44, lista.VerUltimo())

}
func TestInsertarMedioIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 5; i++ {
		lista.InsertarUltimo(i)
	}
	iterador := lista.Iterador()
	for i := 0; i < 2; i++ {
		iterador.HaySiguiente()
	}
	iterador.Insertar(11)

	iterador2 := lista.Iterador()

	for iterador2.HaySiguiente() {
		valor := iterador.VerActual()
		iterador2.Siguiente()
		if valor == 2 {
			require.EqualValues(t, valor, iterador2.VerActual())
		}
	}
}
func TestIteradorBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.Borrar())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, iterador.VerActual())
}

func TestIteradorInternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	elementos := []int{}
	lista.Iterar(func(e int) bool {
		elementos = append(elementos, e)
		return true
	})

	require.Empty(t, elementos, "No debería recorrerse ningún elemento en una lista vacía")
}

func TestIteradorInternoVisitarSiempreFalse(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	elementos := []int{}
	lista.Iterar(func(e int) bool {
		elementos = append(elementos, e)
		return false
	})

	require.Len(t, elementos, 1, "Debería recorrerse sólo un elemento si visitar devuelve false al principio")
	require.EqualValues(t, 1, elementos[0], "El primer elemento debería ser el único recorrido")
}

func TestInsertarIteradorListaVacia(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	iterdor := lista.iterador()

	iterador.Insertar(1)

	require.EqualValues(t, 1, lista.VerPrimero(), "El primer elemento deberia ser el 1")
	require.EqualValues(t, 1, lista.VerUltimo(), "El ultimo elemento deberia ser el 1")
	require.EqualValues(t, 1, iterdaor.VerActual(), "El iterador deberia apuntar al nuevo elemento")
}

func TestVolumenLista(t *testing.T){
	lista :=TDALista.CrearListaEnlazada[int]()
	N := 1000

	for i:= 0, i < N; i++{
		lista.InsertarUltimo(i)
	}

	require.EqualValues(t, 0, lista.VerPrimero(), "El primer elemento deberia ser el 0")
	require.EqualValues(t, N-1, lista.VerUltimo())

	for i:= 0; i < N; i++{
		require.True(t, i, lista.BorrarPrimero(), "El eemnto borrado no coincide con el esperaod")
	}

	require.True(t, lista.EstaVacia(), "La lista debria estar vacia")
}

func TestVolumenIteradorExterno(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	N := 1000

	for i:= 0; i < N; i++{
		litsa.InsertarUltimo(i)
	}

	iter := lista.Iterador()
	contador := 0
	for iter.HaySiguiente(){
		valor := iterador.Siguiente()
		require.EqualValues(t, contador, valor, "El elemento es incorrecto")
		contador++
	}
	require.EqualValues(t, N, contador, "El iterador no recorrio tdos los elemtos")
}

func TestIteradorExternoListaVacia(t *testing){
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	require.False(t, iterador.HaySiguiente(), "No deberia tener siguiente ya que a lista esta vacia.")
	require.EqualValues(t, 0, iterador.VerActual(),"Deberia mostar 0 ya que la lista esta vacia")
}
