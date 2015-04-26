gopath

[10]byte // array of 10 bytes

Public / private variables
stor forbokstav er public
liten - private

slice vs array
slices play the role of dynamically sized arrays
slices are part of an underlying array.
changing the underlaying array will change the slice covering that part


map iteration order is not specified

for key, value := range m {
	value = value + 1 //wont work (assigns to a temporary variable)
	m[key] = value + 1 // assign to the object inside the map
}

func Printf(format string, args ...int){ // ubestemt mengde input arg. blir lagret i array "args"
	if len(args)>0{

	}
}

type Point strukct {
		x, y int
}

func main(){
	p := point{2,3}
	fmt.Println(p.String())
	p.x = 4
	fmt.println(Point{3,5}.string())
}

----------------------------

go doc - kan vise documentation for funksjoner i en package
godoc.org
STOR forbokstav i en funksjon vil gjøre den public i en package
liten forbokstav vil gjøre den private

-----------------------------

const (
	Mon Weekday = iota // iota gir Mon verdien 0
	Tue   //alle under Mon blir av typen Weekday
	Wed
	Thu

	)

-------------------------------

INTERFACE :

interface{} // empty interface

interface {
	String( string)
}

interface{
	Len() int
	Swap(i, j int)
	Less(i, j, int) bool
}

type tor struct{
	name string
}

func (t tor) Len() int{
	len(t.name)
}

---------
les opp på interface

------------------------------

func( ch chan<- string) // setter så funksjonen bare leser fra kanalen

fikse concurrency

lag wait group for å synkroisere program
Package sync -- sjekk ut

"" - empty string

Panic // er en funksjon. stopper systemet
recover() // kan catche panics


sjekk ut meetupen til gutta


