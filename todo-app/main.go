package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("Buy milk")
	todos.add("Buy bread")
	todos.add("Buy beans")
	todos.toggle(0)
	todos.print()
	todos.delete(1)
	todos.print()
	storage.save(todos)
}
