package main

// Реалізація S-блоку
func Sblock(bytes uint8) uint8 {
	// Ініціалізуємо масив констант
	Sconstant := [16]uint8{0b0000, 0b1101, 0b1011, 0b1100, 0b1001, 0b1111, 0b0001, 0b0101, 0b0010, 0b1000, 0b1110, 0b0110, 0b0111, 0b0100, 0b0011, 0b1010}

	// Розбиваємо вхідний байт на дві тетради h і l відповідно
	h := bytes >> 4     // для отримання старши 4 бітів просто зсуваємо вправо на 4 біти
	l := bytes & 0b1111 // для отримання молодших 4 бітів робимо операцію & з числом 15

	// Ініціалізуємо змінну яку будемо повертати
	var output uint8

	// Виконуємо підстановку
	output += (Sconstant[h] << 4) + Sconstant[l]

	// Повертаємо результат
	return output
}

// Реалізація S-блок для обрахунку оберненого значення
func SblockR(bytes uint8) uint8 {
	// Ініціалізуємо масив констант обернених до прямої функції
	Sconstant := [16]uint8{0b0000, 0b1101, 0b1011, 0b1100, 0b1001, 0b1111, 0b0001, 0b0101, 0b0010, 0b1000, 0b1110, 0b0110, 0b0111, 0b0100, 0b0011, 0b1010}
	SconstantR := SconstantReverso(Sconstant)

	// Розбиваємо вхідний байт на дві тетради h і l відповідно
	h := bytes >> 4     // для отримання старши 4 бітів просто зсуваємо вправо на 4 біти
	l := bytes & 0b1111 // для отримання молодших 4 бітів робимо операцію & з числом 15

	// Ініціалізуємо змінну яку будемо повертати
	var output uint8

	// Виконуємо підстановку
	output += (SconstantR[h] << 4) + SconstantR[l]

	// Повертаємо результат
	return output
}

// Функція для обрахування оберненої таблиці констант
func SconstantReverso(Sconstant [16]uint8) [16]uint8 {
	// Ініціалізуємо масив обернених констант
	var SconstantR [16]uint8

	// Знаходимо до констант обернені
	for i := 0; i < len(Sconstant); i++ {
		SconstantR[Sconstant[i]] = uint8(i)
	}

	return SconstantR
}

// Реалізація P-блоку
func Pblock(bytes uint8) uint8 {
	// Ініціалізуємо масив який зберігає нові позиції бітів (мапа перестановки)
	Ppositions := [8]uint8{7, 3, 0, 1, 6, 2, 5, 4}
	// Ініціалізуємо вихідний масив
	var output uint8
	// Виконуємо перестановку
	for i := 0; i < 8; i++ {
		output += ((bytes >> i) & 1) << Ppositions[i]
	}

	return output
}

// Реалізація P-блоку для обрахунку оберненого значення
func PblockR(bytes uint8) uint8 {
	// Ініціалізуємо масив який зберігає нові позиції бітів (мапа перестановки)
	Ppositions := [8]uint8{7, 3, 0, 1, 6, 2, 5, 4}
	// Ініціалізуємо вихідний масив
	var output uint8
	// Виконуємо зворотню перестановку
	for i := 0; i < 8; i++ {
		output += ((bytes >> Ppositions[i]) & 1) << i
	}

	return output
}

// Зауважимо що Pblock(PblockR(uint8)) == PblockR(Pblock(uint8)), тобто якщо одна функція буде служити прямою, то інша буде оберненою.
// Так само для S-блоків

func main() {
	// Це точка входу в програму
}