# Variablen, Datentypen und formatierte Ausgabe

Beim Programmieren in Go sind vorab einige wichtige Sachen zu berücksichtigen:

- Go ist eine stark typisierte Programmiersprache. Variablen haben einen
  bestimmen Datentyp, und können nur Werte dieses einen Typs annehmen.
- *Bezeichner* (von Variablen, Konstanten, Typen usw.) beginnen mit einem
  Buchstaben oder einem Unterstrich (`_`), worauf weitere Buchstaben, Ziffern
  und Underscores folgen können. Es wird zwischen Gross- und Kleinschreibung
  unterschieden. Man verwendet `CamelCase` (und nicht `snake_case`).
- Werden Variablen deklariert aber nirgends verwendet, kann das Programm nicht
  kompiliert werden. Damit soll sichergestellt werden, dass sich keine unnötigen
  Variablen in den Code einschleichen.

## Variablen

*Variablen* werden mit dem Schlüsselwort `var` deklariert:

```go
var [Bezeichner] [Datentyp]
```

Im Gegensatz zu vielen anderen Programmiersprachen wird zuerst der Bezeichner,
und erst dann der Datentyp angegeben!

Beispiele:

```go
var result int
var caption string
var isReady bool
```

Eine neu deklarierte Variable wird mit dem _Nullwert_ (_zero value_) des
jeweiligen Datentyps initialisiert, z.B. `0` bei Ganzzahlen, `false` bei
Wahrheitswerten, oder `""` bei Zeichenketten.

Es ist auch möglich, bei der Deklaration einen Wert anzugeben:

```go
var [Bezeichner] [Datentyp] = [Wert]
```

Beispiele:

```go
var result int = 5
var caption string = "Hello"
var isReady bool = true
```

Gibt man einen Wert an, kann der Compiler den Datentyp anhand dieses Werts
festlegen. Der Datentyp kann dann auch weggelassen werden:

```go
var result = 5
var caption = "Hello"
var isReady = true
```

Die Deklaration kann mit folgender Kurzschreibweise und dem Operator `:=` weiter
vereinfacht werden:

```go
result := 5
isReady := true
caption := "Hello"
```

Es ist auch möglich, mehrere Variablen in einer Zeile zu deklarieren und zu
initialisieren:

```go
var a, b, c = 1, "Hey", false
d, e, f := 2, "Ho", true
```

### Konstanten

*Konstanten* lassen sich mit dem Schlüsselwort `const` definieren:

```go
const [Bezeichner] [Datentyp] = [Wert]
```

Da sich Konstanten im weiteren Programmverlauf _nicht_ ändern können, muss der
Wert gleich bei der Deklaration angegeben werden. Der Datentyp ist wiederum
optional, da er vom Compiler erraten werden kann. Beispiele:

```go
const g float32 = 9.81
const pi float64 = 3.14159265359
const name = "Brandon"
```

Konstanten stehen nur für Ausdrücke zur Verfügung, die zur Kompilierzeit
ermittelt werden können.

## Datentypen

TODO

    int int8 int16 int32 int64
    uint uint8 uint16 uint32 uint64
    float32 float64
    bool
    string
    error
    byte -> uint8?
    rune -> uint32?
