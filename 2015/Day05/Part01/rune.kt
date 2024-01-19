import java.io.File

// Проверка, является ли символ гласной
fun isVowel(ch: Char): Boolean {
    return ch in "aeiou"
}

// Проверка, является ли строка "nice" 
fun isNice(s: String): Boolean {
    var vowelCount = 0
    var doubleLetter = false
    var containsBadStr = false

    // Список недопустимых строк
    val badStrs = listOf("ab", "cd", "pq", "xy")
    for (bad in badStrs) {
        if (bad in s) {
            containsBadStr = true
            break
        }
    }

    // Проверка наличия гласных и повторяющихся букв
    for (i in s.indices) {
        if (isVowel(s[i])) {
            vowelCount++
        }
        if (i > 0 && s[i] == s[i - 1]) {
            doubleLetter = true
        }
    }

    // Возвращается true, если строка соответствует всем критериям
    return vowelCount >= 3 && doubleLetter && !containsBadStr
}

fun main(args: Array<String>) {
    val inputFile = File("input.txt")
    var niceCount = 0

    // Чтение каждой строки из файла и подсчет "nice" строк
    inputFile.forEachLine { line ->
        if (isNice(line)) {
            niceCount++
        }
    }

    // Вывод количества "nice" строк
    println("Количество хороших строк: $niceCount")
}
