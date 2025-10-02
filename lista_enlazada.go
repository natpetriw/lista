package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type iteradorListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}


func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}


func crearNodo[T any](nuevo_nodo T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = nuevo_nodo
	nuevoNodo.siguiente = nil
	return nuevoNodo
}

