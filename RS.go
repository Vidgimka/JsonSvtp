package main

func uniqRSLists() []string {
	var station []string

	for k, _ := range geodatList { // перебор массива по ключу, ищем уникальные базы
		ifExists := false // создаем счетчик встречаось или нет перебираемая БС, по умолчанию нет.

		// проверка на уникальность имени БС
		for _, s := range station {
			if geodatList[k].Station == s { // сравниваем  элементы geodatList с к-м пребранным элементом массива
				ifExists = true // если нашлось совпаниение - выход из условного оператора и конец итерации
				break
			}
		}
		// проверка на уникальность имени БС

		if ifExists == false { // если сетчик не переключился те не нашли совпадение
			station = append(station, geodatList[k].Station) // значит добавляем в слайс  строк station перебираемый элемент
		}

	}

	return station
}
