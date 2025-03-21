Давай разберём практическое применение горутин, `defer`, `panic` и `recover` в **реальных сценариях**.  

---

## **📌 Где применяются горутины?**
Горутины полезны, когда надо выполнять несколько задач **параллельно или конкурентно**.  

### **1. Обработка множества запросов одновременно (сервер)**
Представь, что у нас есть веб-сервер, который обрабатывает входящие HTTP-запросы. Если сервер обрабатывает запросы **последовательно**, пользователи будут ждать, пока завершится предыдущий запрос.

С **горутинами** можно запускать обработку **каждого запроса в отдельной горутине**, что ускоряет работу.

**Пример веб-сервера с горутинами**:
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	go processRequest(r.URL.Path) // Запускаем обработку в горутине
	fmt.Fprintln(w, "Запрос принят, обрабатывается...")
}

func processRequest(path string) {
	time.Sleep(2 * time.Second) // Имитация сложной обработки
	fmt.Println("Обработан запрос:", path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```
✅ Теперь сервер **не зависает** на каждом запросе, а может обрабатывать их параллельно.

---

### **2. Параллельная загрузка данных**
Допустим, у нас есть несколько источников данных (API, базы данных, файлы), и мы хотим загружать данные **параллельно**, а не по очереди.

**Пример загрузки данных из нескольких источников одновременно**:
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(source string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счётчик после завершения
	time.Sleep(time.Second * 2) // Имитация задержки
	fmt.Println("Данные загружены из:", source)
}

func main() {
	var wg sync.WaitGroup

	sources := []string{"БД", "API", "Файл"}
	wg.Add(len(sources)) // Устанавливаем счётчик горутин

	for _, source := range sources {
		go fetchData(source, &wg) // Запускаем горутину
	}

	wg.Wait() // Ждём завершения всех горутин
	fmt.Println("Все данные загружены")
}
```
**Вывод:**
```
Данные загружены из: БД
Данные загружены из: API
Данные загружены из: Файл
Все данные загружены
```
✅ Без горутин программа загружала бы данные **по очереди** (6 секунд), а с горутинами – **всего 2 секунды**.

---

## **📌 Где применяются `defer`?**
`defer` удобен, когда нужно выполнить **уборочные операции** (закрытие файлов, соединений, освобождение ресурсов).

### **1. Гарантированное закрытие файла**
При работе с файлами **нужно закрывать файл после использования**. `defer` помогает сделать это автоматически.

**Пример с `defer` при работе с файлами**:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close() // Гарантированное закрытие

	fmt.Println("Файл открыт, читаем данные...")
}
```
✅ Даже если возникнет ошибка или программа выйдет раньше, `file.Close()` **всё равно выполнится**.

---

### **2. Закрытие соединения с базой данных**
```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=mydb sslmode=disable")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer db.Close() // Закрываем соединение при выходе из функции
}
```
✅ Не забудем закрыть соединение!

---

## **📌 Где применяются `panic` и `recover`?**
### **1. `panic` при критических ошибках**
Когда **программа не может продолжать работу** (например, повреждение данных, невозможность выделить память), лучше вызвать `panic`.

**Пример выхода за границы массива (автоматическая паника):**
```go
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[5]) // Паника: выход за границы массива
}
```
Вывод:
```
panic: runtime error: index out of range
```
✅ `panic` помогает **сразу обнаружить ошибку**.

---

### **2. `recover` для перехвата `panic` и продолжения работы**
Вместо полного краха программы можно **поймать `panic` и обработать ошибку**.

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Перехвачена паника:", r)
		}
	}()

	fmt.Println("Начало программы")
	panic("Что-то пошло не так!") // Вызов паники
	fmt.Println("Эта строка не выполнится")
}
```
Вывод:
```
Начало программы
Перехвачена паника: Что-то пошло не так!
```
✅ Программа **не крашится**, а продолжает работу.

---

### **3. Использование `panic` в HTTP-сервере**
Если веб-сервер обработал ошибку **с `panic`**, он не должен полностью останавливаться.

```go
package main

import (
	"fmt"
	"net/http"
)

func safeHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка обработана:", r)
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		}
	}()

	panic("Что-то пошло не так!") // Искусственная паника
}

func main() {
	http.HandleFunc("/", safeHandler)
	http.ListenAndServe(":8080", nil)
}
```
✅ Теперь сервер **не падает**, даже если в обработчике произошла ошибка.

---

## **📌 Итог**
| **Концепция**  | **Применение** |
|--------------|---------------|
| **Горутины (`go`)** | Обработка множества задач параллельно (серверы, загрузка данных) |
| **`sync.WaitGroup`** | Ожидание завершения всех горутин |
| **`defer`** | Автоматическое закрытие файлов, соединений, освобождение ресурсов |
| **`panic`** | Критические ошибки (нельзя продолжить работу) |
| **`recover`** | Перехват `panic`, чтобы не падало всё приложение |

🔥 **Если тебе нужно:**  
✅ **Обрабатывать запросы одновременно** → используй **горутины**  
✅ **Дождаться выполнения всех горутин** → используй **`sync.WaitGroup`**  
✅ **Гарантированно закрывать файлы и соединения** → используй **`defer`**  
✅ **Остановить программу при критической ошибке** → используй **`panic`**  
✅ **Не дать программе упасть из-за `panic`** → используй **`recover`**  

**Теперь ты знаешь, как и где это применять в реальных задачах!** 🚀