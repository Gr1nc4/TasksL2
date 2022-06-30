package main

/*
		Строитель (англ. Builder) — порождающий шаблон проектирования предоставляет способ создания составного объекта.
Отделяет конструирование сложного объекта от его представления так, что в результате одного и того
же процесса конструирования могут получаться разные представления.
Проблема:
    Инициализация очень сложного, большого объекта со множеством параметров инициализации. Использовать один конструктор
    с множеством параметров - плохо (телескопический конструктор - анти-паттерн)
Решение:
    Паттерн Строитель предлагает вынести конструирование объекта за пределы его собственного класса, поручив это дело
    отдельным объектам, называемым строителями.
плюсы:
    +Позволяет создавать продукты пошагово.
 +Позволяет использовать один и тот же код для создания различных продуктов.
 +Изолирует сложный код сборки продукта от его основной бизнес-логики.
минусы:
    -Усложняет код программы из-за введения дополнительных классов.
 -Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

// Pizza Cложносостовной объект пицца, который разные представление которого
//могут включать разные элементы (нужные поля - True)
type Pizza struct {
	isChesse  bool
	isMeat    bool
	isOlivies bool
	isPepper  bool
}

// NewPizza Общий конструктор для всех элементов
func NewPizza(ch bool, mt bool, ol bool, pep bool) *Pizza {
	return &Pizza{isChesse: ch, isMeat: mt, isOlivies: ol, isPepper: pep}
}

// Builder Каждый Builder пиццы должны реализовывать этот интерфейс
//(даже если возвращаемые результаты сборки потом будут разных классов -
// у нас может потом быть не пицца, а например бумажное описание пиццы PapperPizza)
type Builder interface {
	setChesse()
	setMeat()
	setOlivies()
	setPepper()
}

// PizzaBuilder Реализуем Builder с набором полей как у исходного объекта
type PizzaBuilder struct {
	isChesse  bool
	isMeat    bool
	isOlivies bool
	isPepper  bool
}

// NewPizzaBuilder Конструктор без параметром, создающий дефолтный объект билдера
func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

//Сеттеры для каждого поля
func (p *PizzaBuilder) setChesse() *PizzaBuilder {
	p.isChesse = true
	return p
}

func (p *PizzaBuilder) setMeat() *PizzaBuilder {
	p.isMeat = true
	return p
}

func (p *PizzaBuilder) setOlivies() *PizzaBuilder {
	p.isOlivies = true
	return p
}

func (p *PizzaBuilder) setPepper() *PizzaBuilder {
	p.isPepper = true
	return p
}

//Метод создания итогового объекта - итоговый результат построения
//вызывается общий конструктор Пиццы
func (p *PizzaBuilder) build() *Pizza {
	return NewPizza(p.isChesse, p.isMeat, p.isOlivies, p.isPepper)
}

// func main() {
// 	pizza := NewPizzaBuilder(). //создаем объект билдер
// 					setChesse(). //добавляем нужные части объекта (нужные - становятся True)
// 					setPepper().
// 					build() //создаем возращаем объект Pizza
// 	fmt.Println(pizza)
// }
