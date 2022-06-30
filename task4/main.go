package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/
func main() {

	var arr = []string{"пятак", "пятка", "тяпка", "лситок", "слиток", "столик"}
	m := SearchAnagram(arr)

	fmt.Println(m)
}

func SearchAnagram(arr []string) map[string][]string {
	m := make(map[string][]string)
	//проходимся по нашему исходному массиву
	for _, item := range arr {
		//создаем массив, в который будем складывать анаграммы
		var anagrams []string
		//каждое слово приводим к нижнему регистру
		item = strings.ToLower(strings.Trim(item, " "))
		//так же проходимся по исходному массиву
		for _, val := range arr {
			val = strings.ToLower(strings.Trim(val, " "))
			//сравниваем колличество символов в словах
			if utf8.RuneCountInString(item) == utf8.RuneCountInString(val) {

				//создаем мапу, в которой будем хранить буквы и их колличества в слове
				hash := make(map[string]int)
				//проходимся по слову
				for _, r := range item {
					//создаем j переменную, в которую кдадем значение с ключем равной букве из слова
					j := hash[string(r)]
					//если j равен 0, значит в мапе нет значения с таким ключем, складываем туда нашу букву (так делаем для каждого след слова)
					if j == 0 {
						hash[string(r)] = 1
					} else {
						//иначе инкрментим значение
						hash[string(r)] = j + 1
					}
					//	fmt.Println(r)
				}
				//аналогично проходимся по слову
				for _, r := range val {
					j := hash[string(r)]

					if j == 0 {
						hash[string(r)] = 1
					} else {
						hash[string(r)] = j + 1
					}
					//	fmt.Println(string(r))
				}
				//проверка на анаграмму
				var isAnagram bool = true
				for _, value := range hash {
					if value%2 != 0 {
						isAnagram = false
					}
				}
				if isAnagram {
					anagrams = append(anagrams, val)
					if len(anagrams) >= 2 {
						m[anagrams[0]] = anagrams[1:]
					}
				}
			}
		}
	}
	for val := range m {
		sort.Strings(m[val])
	}
	return m
}
