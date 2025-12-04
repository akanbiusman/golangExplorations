package main

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy bread")
	todos.add("Buy beans")
	todos.toggle(0)
	todos.print()
	todos.delete(1)
	todos.print()
}
