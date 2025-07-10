package main

import "fmt"

/***
Мы описываем квартиру
У каждой квартиры есть адрес, площадь, этаж, и стоимость
Каждая квартира может менять свою стоимость
Адрес не должен быть пустой, и не может изменяться
Площадь должна быть больше 10 квадратных метров и не может изменяться
Этаж должен быть больше нуля и меньше либо равен ста
Стоимость должна быть больше нуля и может изменяться
Необходимо описать соответсвующую структуру с конструктором, полями и методами
То, что поля на текущем этапе изучение мы сможем изменить напрямую и сделать их невалидными мы во внимание не берём
***/

type Flat struct {
	address string
	square  float64
	floor   int
	cost    int
}

func NewFlat(
	address string,
	square float64,
	floor int,
	cost int,
) Flat {
	fmt.Println("- Проверка адреса на непустоту")
	if address == "" {
		fmt.Println("	>>> WARN: проверка не пройдена")
		return Flat{}
	}
	fmt.Println("- Проверка площади квартиры, не меньше 10 кв. м.")
	if square < 10 {
		fmt.Println("	>>> WARN: проверка не пройдена")
		return Flat{}
	}
	fmt.Println("- Проверка этажа, не меньше 0 и не больше 100")
	if floor < 0 || floor > 100 {
		fmt.Println("	>>> WARN: проверка не пройдена")
		return Flat{}
	}
	fmt.Println("- Проверка стоимости, не меньше 0")
	if cost < 0 {
		fmt.Println("	>>> WARN: проверка не пройдена")
		return Flat{}
	}
	return Flat{
		address: address,
		square:  square,
		floor:   floor,
		cost:    cost,
	}
}

func (mf *Flat) changeCostFlat(newCost int) {
	if newCost > 0 {
		mf.cost = newCost
		fmt.Println("\nСтомость квартиры по адресу", mf.address, "изменена на:", mf.cost)
	} else {
		fmt.Println("\n	>>> WARN: Изменение стоимости квартиры не произведено, т.к. не может быть меньше 0")
	}
}

func main() {
	flat_1 := NewFlat(
		"Москва, Подольских курсантов, д 12 к 1 кв 11",
		44.0,
		4,
		15000000,
	)
	fmt.Println("\nИсходные данные по квартире:\n",
		"		Адрес: ", flat_1.address,
		"\n		Площадь: ", flat_1.square,
		"\n		Этаж: ", flat_1.floor,
		"\n		Стоимость: ", flat_1.cost)

	flat_2 := NewFlat(
		"Москва, Крутицкая набережная д 15 кв 14",
		39.0,
		2,
		18000000,
	)
	fmt.Println("\nИсходные данные по квартире:\n",
		"		Адрес: ", flat_2.address,
		"\n		Площадь: ", flat_2.square,
		"\n		Этаж: ", flat_2.floor,
		"\n		Стоимость: ", flat_2.cost)

	flat_1.changeCostFlat(16000000)
	flat_2.changeCostFlat(25000000)
	// можем изменить напрямую , при этом проверки не будет
	// my_flat.cost = -100

	fmt.Println("\nОбновлённые данные по квартире:\n",
		"		Адрес: ", flat_1.address,
		"\n		Площадь: ", flat_1.square,
		"\n		Этаж: ", flat_1.floor,
		"\n		Стоимость: ", flat_1.cost)
}
