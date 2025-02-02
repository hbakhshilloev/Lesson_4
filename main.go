package main

import "fmt"

const (
	salary              = 150
	coat_cost           = 350
	cleaning_count      = 100
	game_ball           = 20
	min_dirty_limit     = 90
	min_satiety_limit   = 0
	min_happyness_limit = 10
	coat_ball           = 60
	dirty_min           = 5
	eat_min             = 30
	any_action_point    = 10
)

type House struct {
	man     *Man
	woman   *Woman
	money   int
	product int
	dirty   int
}
type Man struct {
	house     *House
	name      string
	satiety   int
	happyness int
}

//Добавление дней на дом
func (h *House) add_day(day int) {
	//проверка на голодност и счастье
	h.dirty = h.dirty + day*dirty_min
	if h.man.satiety < min_satiety_limit {
		fmt.Println(h.man.name, " died from starvation")
	} else if h.woman.satiety < min_satiety_limit {
		fmt.Println(h.woman.name, " died from starvation")
	} else if h.man.happyness < min_happyness_limit {
		fmt.Println(h.man.name, " died from depression")
	} else if h.woman.happyness < min_happyness_limit {
		fmt.Println(h.woman.name, " died from depression")
	}
	// проверка дом на грязи
	if h.dirty > dirty_min {
		h.man.happyness = h.man.happyness - day*any_action_point
		h.woman.happyness = h.woman.happyness - day*any_action_point
	} else {
		fmt.Println("Your house is clean, Clean point is : ", h.dirty)
	}
}

// Кормление мужика
func (m *Man) eat(count int) { //max 30
	if count > eat_min {
		fmt.Println(" You are cannot eat so much in one time")
	} else if count <= 0 {
		fmt.Println("You are eat nothing")
	} else if m.house.product > count {
		m.house.product = m.house.product - count
		m.satiety = m.satiety + count
	} else if m.house.product < count {
		fmt.Println("You dont have eats in your house , Tell your wife to buy products. Products count : ", m.house.product)
	}

}

//Развлечение мужика
func (m *Man) play_game() {
	m.satiety = m.satiety - any_action_point
	m.happyness = m.happyness + game_ball
}

// Работа мужика
func (m *Man) work() {
	m.house.money = m.house.money + salary
	m.satiety = m.satiety - any_action_point
}

type Woman struct {
	house     *House
	name      string
	satiety   int
	happyness int
}

// Кормление жены
func (w *Woman) eat(count int) { //max 30
	if count > eat_min {
		fmt.Println(" You are cannot eat so much in one time")
	} else if count <= 0 {
		fmt.Println("You are eat nothing")
	} else if w.house.product > count {
		w.house.product = w.house.product - count
		w.satiety = w.satiety + count
	} else if w.house.product < count {
		fmt.Println("You dont have eats in your house , Please buy products. Products count : ", w.house.product)
	}
}

// Покупка продукты
func (w *Woman) buy_product(count int) {
	if w.house.money < count {
		fmt.Println("You dont have enough money. Tell your husband to work. You remain is :", w.house.money)
		w.satiety = w.satiety - any_action_point
	} else {
		w.house.money = w.house.money - count
		w.house.product = w.house.product + count
		w.satiety = w.satiety - any_action_point
	}
}

// Покупка шубы для жены
func (w *Woman) buy_coat() {
	if w.house.money < coat_cost {
		fmt.Println("You dont have enough money. Tell your husband to work. You remain is :", w.house.money)
		w.satiety = w.satiety - any_action_point
	} else {
		w.house.money = w.house.money - coat_cost
		w.happyness = w.happyness + coat_ball
		w.satiety = w.satiety - any_action_point
	}
}

//чистка дома
func (w *Woman) clean_home() {
	w.satiety = w.satiety - any_action_point
	if w.house.dirty > min_dirty_limit {
		w.house.dirty = w.house.dirty - cleaning_count
		w.satiety = w.satiety - any_action_point
	} else {
		fmt.Print("Your home is clean, Index cleaness is : ", w.house.dirty)
	}

}
func main() {
	house := House{
		money:   100,
		product: 50,
		dirty:   0,
	}
	man := Man{
		name:      "James",
		satiety:   20,
		happyness: 35,
	}
	woman := Woman{
		name:      "Anna",
		satiety:   25,
		happyness: 15,
	}
	man.work()
	man.play_game()
	man.eat(10)
	woman.eat(8)
	woman.buy_product(5)
	house.add_day(1)
}
