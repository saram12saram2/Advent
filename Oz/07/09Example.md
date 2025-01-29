### **Похожий Task:**

У нас есть 3 сущности: **пользователь**, **проект** и **задача**.

- У пользователя есть имя и дата регистрации.
- У проекта есть название, описание и дата создания.
- У задачи есть текст, автор, статус (например, `TODO`, `IN_PROGRESS`, `DONE`), приоритет и дата создания.
- Пользователь может участвовать в нескольких проектах.
- Задача обязательно принадлежит одному проекту, но может быть создана только пользователем, который участвует в проекте.

**Задание:**
1. Опишите предметную область в виде таблиц.
2. Напишите SQL-запрос для получения всех задач пользователя с именем "Анна" в формате `(task_id, task_text, project_name)`.
3. Объясните разницу между `JOIN` и `LEFT JOIN` для этого запроса.
4. Что такое индекс, сколько и какие индексы будут созданы по этой схеме?
5. Как ускорить работу запросов при увеличении размера таблиц? Что показывает `EXPLAIN ANALYZE`?
6. Как заполнять поле `id` для каждой таблицы: использовать автоинкремент или UUID? Какой подход лучше для масштабируемости?
7. Если число задач быстро растет, а места на диске не хватает, что делать? Опишите подход к шардированию таблицы `tasks`.

---

### **Ответ**

#### **1. Описание предметной области в виде таблиц**
```sql
-- Таблица пользователей
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    registered_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица проектов
CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица задач
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    author_id INT NOT NULL REFERENCES users(id), -- Ссылка на автора задачи
    project_id INT NOT NULL REFERENCES projects(id), -- Ссылка на проект
    status TEXT NOT NULL CHECK (status IN ('TODO', 'IN_PROGRESS', 'DONE')),
    priority INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица связи пользователей и проектов (многие ко многим)
CREATE TABLE user_projects (
    user_id INT NOT NULL REFERENCES users(id),
    project_id INT NOT NULL REFERENCES projects(id),
    PRIMARY KEY (user_id, project_id)
);
```

---

#### **2. SQL-запрос для получения задач пользователя**
```sql
SELECT 
    t.id AS task_id, 
    t.text AS task_text, 
    p.name AS project_name
FROM users u
JOIN user_projects up ON up.user_id = u.id
JOIN projects p ON p.id = up.project_id
JOIN tasks t ON t.project_id = p.id AND t.author_id = u.id
WHERE u.name = 'Анна';
```

---

#### **3. Разница между `JOIN` и `LEFT JOIN`**
- **`JOIN`**:
  - Возвращает только совпадающие записи из всех таблиц.
  - Если пользователь "Анна" не участвует ни в одном проекте или не создал ни одной задачи, запрос вернет **пустой результат**.

- **`LEFT JOIN`**:
  - Возвращает все записи из левой таблицы (`users` в данном случае), даже если нет совпадений в других таблицах.
  - Если пользователь "Анна" существует, но не создал задачи, будут возвращены строки с NULL в колонках задач.

Пример с использованием `LEFT JOIN`:
```sql
SELECT 
    t.id AS task_id, 
    t.text AS task_text, 
    p.name AS project_name
FROM users u
LEFT JOIN user_projects up ON up.user_id = u.id
LEFT JOIN projects p ON p.id = up.project_id
LEFT JOIN tasks t ON t.project_id = p.id AND t.author_id = u.id
WHERE u.name = 'Анна';
```

---

#### **4. Индексы**
- **Что такое индекс?**
  Индекс — это вспомогательная структура данных, которая ускоряет поиск данных в таблице. 
  - В SQL индекс создается автоматически для полей с `PRIMARY KEY` и `UNIQUE`.
  - Дополнительные индексы можно добавлять вручную.

- **Индексы, создаваемые автоматически:**
  - `users(id)`
  - `projects(id)`
  - `tasks(id)`
  - `user_projects(user_id, project_id)`

- **Ручной индекс для ускорения поиска пользователя по имени:**
  ```sql
  CREATE INDEX idx_user_name ON users(name);
  ```

- **Индекс для ускорения поиска задач по проекту и автору:**
  ```sql
  CREATE INDEX idx_tasks_project_author ON tasks(project_id, author_id);
  ```

---

#### **5. Ускорение запросов и `EXPLAIN ANALYZE`**
- **Почему запросы замедляются?**
  - При увеличении размеров таблиц становится сложнее искать данные, так как происходит полный скан таблицы (Full Table Scan).
  - Это можно увидеть в плане запроса:
    ```sql
    EXPLAIN ANALYZE 
    SELECT 
        t.id AS task_id, 
        t.text AS task_text, 
        p.name AS project_name
    FROM users u
    JOIN user_projects up ON up.user_id = u.id
    JOIN projects p ON p.id = up.project_id
    JOIN tasks t ON t.project_id = p.id AND t.author_id = u.id
    WHERE u.name = 'Анна';
    ```
  - Узел `Seq Scan` (последовательный скан) указывает, что таблица просматривается целиком.

- **Как ускорить?**
  - Добавить индексы на поля, используемые в фильтрах и соединениях (`u.name`, `up.user_id`, `p.id`, `t.project_id`, `t.author_id`).

---

#### **6. Автоинкремент или UUID?**
- **Автоинкремент (`SERIAL`)**:
  - Удобен для небольших систем.
  - Менее уникален, сложен для шардирования.

- **UUID**:
  - Глобально уникален.
  - Подходит для масштабируемых распределенных систем.
  - Пример использования:
    ```sql
    CREATE TABLE tasks (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        text TEXT NOT NULL,
        author_id INT NOT NULL REFERENCES users(id),
        project_id INT NOT NULL REFERENCES projects(id),
        status TEXT NOT NULL CHECK (status IN ('TODO', 'IN_PROGRESS', 'DONE')),
        priority INT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );
    ```

---

#### **7. Шардирование таблицы `tasks`**
- Шардировать по `project_id`:
  - Разделить задачи по проектам, чтобы один проект хранился на одном шарде.
- Пример:
  - Проект с `id % 3 == 0` → Шард 1.
  - Проект с `id % 3 == 1` → Шард 2.
  - Проект с `id % 3 == 2` → Шард 3.

- **Обновленный запрос с учетом шардирования:**
  Выбор задач выполняется только на нужном шарде:
  ```sql
  SELECT 
      t.id AS task_id, 
      t.text AS task_text, 
      p.name AS project_name
  FROM shard_1.tasks t
  JOIN shard_1.projects p ON t.project_id = p.id
  WHERE t.author_id = (SELECT id FROM shard_1.users WHERE name = 'Анна');
  ```

---

Этот похожий сценарий иллюстрирует концепции нормализации данных, индексации и масштабируемости базы данных.