### **Что такое GOMAXPROCS?**

**`GOMAXPROCS`** — это настройка в Go, которая определяет количество операционных системных потоков (ядер), которые Go Runtime может использовать для выполнения горутин одновременно.

---

### **Зачем нужен GOMAXPROCS?**

1. **Контроль параллелизма**:
   - В Go горутины работают в пользовательском пространстве (user space) поверх операционных системных потоков.
   - `GOMAXPROCS` указывает, сколько потоков ядра процессора могут быть одновременно задействованы для выполнения горутин.

2. **Оптимизация производительности**:
   - Увеличивая или уменьшая `GOMAXPROCS`, можно настроить использование CPU в зависимости от задач:
     - **Большое значение** — больше параллелизма, подходит для вычислительных задач.
     - **Малое значение** — экономия ресурсов, полезно для задач ввода-вывода.

3. **Ограничение использования CPU**:
   - В некоторых случаях нужно ограничить количество используемых ядер (например, на сервере с ограничениями или при тестировании).

---

### **Как настроить GOMAXPROCS?**

1. **С помощью переменной окружения**:
   ```bash
   GOMAXPROCS=4 go run main.go
   ```

2. **Программно через `runtime.GOMAXPROCS`**:
   ```go
   package main

   import (
       "fmt"
       "runtime"
   )

   func main() {
       // Устанавливаем GOMAXPROCS
       old := runtime.GOMAXPROCS(4)
       fmt.Printf("Previous GOMAXPROCS: %d\n", old)

       // Вывод текущего значения
       fmt.Printf("Current GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
   }
   ```

---

### **Может ли приложение с GOMAXPROCS=4 использовать больше 4 ядер?**

Да, это возможно в следующих случаях:

1. **Блокирующие системные вызовы (syscalls)**:
   - Если горутина выполняет блокирующий системный вызов (например, `read`, `write`, `connect`), Go Runtime запускает дополнительный поток операционной системы для обработки других горутин.
   - В эти моменты количество активных потоков CPU может превысить значение `GOMAXPROCS`.

2. **Внешние библиотеки**:
   - Если Go-программа вызывает код на C через `cgo`, этот код может запускаться в системных потоках, не контролируемых `GOMAXPROCS`.

3. **Пул потоков Go Runtime**:
   - Go Runtime может создавать дополнительные потоки для обработки блокирующих операций, таких как garbage collection (сборка мусора).

---

### **Пример: GOMAXPROCS и блокирующие вызовы**

```go
package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) // Устанавливаем 2 ядра для выполнения горутин

	// Горутина для блокирующего вызова
	go func() {
		_, _ = http.Get("https://example.com") // Блокирующий вызов
		fmt.Println("HTTP call completed")
	}()

	// Горутина для вычислений
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Compute:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Ждём завершения
	time.Sleep(3 * time.Second)
}
```

**Что произойдёт:**
- Go Runtime выделит дополнительные потоки для обработки HTTP-запроса, если он блокирует выполнение.
- В результате, временно может быть использовано больше потоков, чем указано в `GOMAXPROCS`.

---

### **Зачем Go Runtime запускает дополнительные потоки?**

Это необходимо, чтобы предотвратить **"зависание"** горутин:
- Если один поток занят системным вызовом, другие горутины могут продолжать выполняться на дополнительных потоках.
- Go Runtime динамически управляет потоками для поддержания высокой производительности.

---

### **Итог: основные моменты про GOMAXPROCS**

1. **GOMAXPROCS** управляет числом потоков, которые Go Runtime может использовать для выполнения горутин одновременно.
2. Значение по умолчанию равно числу логических ядер процессора (`runtime.NumCPU()`).
3. Программа может использовать больше ядер, чем `GOMAXPROCS`, в следующих случаях:
   - При выполнении блокирующих системных вызовов.
   - При использовании `cgo` для вызова внешнего кода.
   - Для служебных задач, выполняемых Go Runtime (например, сборка мусора).

Настройка `GOMAXPROCS` позволяет балансировать производительность и использование ресурсов.