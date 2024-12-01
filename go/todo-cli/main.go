package main

func main() {
	todos := Todos{}
    todos.add("Buy Milk")
    todos.add("Buy Bread")
	todos.print()
    todos.toggle(0)
    todos.print()

}