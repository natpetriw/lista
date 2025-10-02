package lista

type Lista[T any] interface {
	// Devuelve True si la lista no tiene elementos.
	EstaVacia() bool

	//Inserta el elemento en el inicio de la lista
	InsertarPrimero(T)

	//Inserta el elemento en el final de la lista
	InsertarUltimo(T)

	//Borra el primer elemento de la lista y devuelve su valor, si esta vacia entra en panico con un mensaje.
	BorrarPrimero() T

	//Devuelve el valor del primer elemento de la lista, si esta vacia entra en panico con un mensaje.
	VerPrimero() T

	//Devuelve el valor del ultimo elemento de la lista, si esta vacia entra en panico con un mensaje.
	VerUltimo() T

	//Devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	//Aplica la funcion que se pase por parametro a todos los elementos de la lista. Termina cuando no queden mas elementos o cuando la funcion visitar devuelva false.
	Iterar(visitar func(T) bool)

	//Crea un iterador.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	VerActual() T

	//Devuelve true si hay mas elementos para seguir iterando--.
	HaySiguiente() bool

	//Itera al siguiente elemento de la lista, entra en panico si no hay mas elementos.
	Siguiente()

	//Inserta un elemento en la posicion que se encuentra el iterador.
	Insertar(T)

	//Borra el elemento en la posicion que se encuentra el iterador y devuelve el valor, entra en panico si no es valido.
	Borrar() T
}
