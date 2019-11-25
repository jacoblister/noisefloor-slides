package main // Obligatory Todo application
import (
	"strconv"

	"github.com/jacoblister/noisefloor/pkg/vdom"
)

//Render renders the Clicker component
func (t *Todo) Render() vdom.Element {
	items := vdom.MakeElement("div")

	for i := 0; i < len(t.items); i++ {
		element := t.renderItem(&t.items[i], i)
		items.AppendChild(element)
	}

	result := vdom.MakeElement("div",
		vdom.MakeElement("input",
			"id", "addItem",
			"placeholder", "add TODO item",
			vdom.MakeEventHandler(vdom.Change, func(element *vdom.Element, event *vdom.Event) {
				t.addItem(event.Data["Value"].(string))
			},
			),
		),
		items,
		vdom.MakeElement("br"),
		vdom.MakeElement("div",
			// vdom.MakeTextElement("Total items: "+strconv.Itoa(len(t.items))),
			vdom.MakeTextElement("Total items: "+strconv.Itoa(len(t.items))+
				", Checked items: "+strconv.Itoa(t.checkedItemCount())),
		),
	)
	return result
}

func main() {
	var todo Todo

	todo.items = append(todo.items, TodoItem{Name: "Implement VDOM", Completed: true})
	todo.items = append(todo.items, TodoItem{Name: "Implement Components", Completed: false})

	vdom.RenderComponentToDom(&todo)
	vdom.ListenAndServe()
}
