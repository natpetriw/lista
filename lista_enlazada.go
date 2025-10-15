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

const(
	LISTA_VACIA = "La lista está vacía"
	ITERADOR = "El iterador termino de iterar"
)

func CrearListaEnlazada[T any]() *listaEnlazada[T] {
	return new(listaEnlazada[T])
}

func crearNodo[T any](nuevo_nodo T) *nodoLista[T] {
	nuevoNodo := new(nodoLista[T])
	nuevoNodo.dato = nuevo_nodo
	nuevoNodo.siguiente = nil
	return nuevoNodo
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := &nodoLista[T]{dato: dato, siguiente: lista.primero}
	
	if lista.primero == nil {
		lista.ultimo = nuevoNodo
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := &nodoLista[T]{dato: dato, siguiente: nil}

	if lista.ultimo != nil {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo

	if lista.primero == nil {
		lista.primero = nuevoNodo
	}
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(LISTA_VACIA)
	}
	elemento := lista.primero.dato
	lista.primero = lista.primero.siguiente

	if lista.primero == nil {
		lista.ultimo = nil
	}

	lista.largo--
	return elemento
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(LISTA_VACIA)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(LISTA_VACIA)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil; actual = actual.siguiente{
		if !visitar(actual.dato) {
			return
		}
		
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
}

func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic(ITERADOR)
	}
	return iterador.actual.dato
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic(ITERADOR)
	}

	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iteradorListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)

	if iterador.anterior == nil {
		iterador.lista.InsertarPrimero(dato)
		iterador.actual = iterador.lista.primero
	} else {
		nuevoNodo.siguiente = iterador.actual
		iterador.anterior.siguiente = nuevoNodo
		if iterador.actual == nil {
			iterador.lista.ultimo = nuevoNodo
		}
		iterador.lista.largo++
		iterador.actual = nuevoNodo
	}
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic(ITERADOR)
	}

	dato := iterador.actual.dato

	if iterador.anterior == nil {
		iterador.lista.BorrarPrimero()
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
		if iterador.actual == iterador.lista.ultimo {
			iterador.lista.ultimo = iterador.anterior
		}
		iterador.lista.largo--
	}

	iterador.actual = iterador.actual.siguiente

	return dato
}
