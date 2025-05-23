# Plant Tracker API Documentation

## Базовая информация

- Base URL: `http://localhost:3000`
- Content-Type: `application/json`
- Кодировка: UTF-8

## Аутентификация

В текущей версии аутентификация не требуется.

## API Endpoints

### Plants

#### 1. Получение списка растений

```http
GET /api/plants
```

**Response** (200 OK)
```typescript
interface Plant[] {
  id: number;
  name: string;
  species: string;
  watering_frequency: number; // в днях
  last_watered: string; // ISO 8601 datetime
  created_at: string; // ISO 8601 datetime
  updated_at: string; // ISO 8601 datetime
}
```

**Пример ответа:**
```json
[
  {
    "id": 1,
    "name": "Монстера",
    "species": "Monstera Deliciosa",
    "watering_frequency": 7,
    "last_watered": "2024-03-20T10:00:00Z",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
  }
]
```

#### 2. Создание растения

```http
POST /api/plants
```

**Request Body**
```typescript
interface CreatePlantRequest {
  name: string;
  species: string;
  watering_frequency: number;
  last_watered: string; // ISO 8601 datetime
}
```

**Response** (201 Created)
```typescript
interface Plant {
  id: number;
  name: string;
  species: string;
  watering_frequency: number;
  last_watered: string;
  created_at: string;
  updated_at: string;
}
```

#### 3. Обновление растения

```http
PUT /api/plants/{plant_id}
```

**Path Parameters**
- `plant_id`: number (required)

**Request Body**
```typescript
interface UpdatePlantRequest {
  name: string;
  species: string;
  watering_frequency: number;
  last_watered: string; // ISO 8601 datetime
}
```

**Response** (200 OK)
```typescript
interface Plant {
  id: number;
  name: string;
  species: string;
  watering_frequency: number;
  last_watered: string;
  created_at: string;
  updated_at: string;
}
```

### Health Check

#### Проверка статуса сервиса

```http
GET /api/health
```

**Response** (200 OK)
```typescript
interface HealthResponse {
  status: "healthy";
}
```

## Коды ошибок

### HTTP Status Codes

- 200: Успешный запрос
- 201: Ресурс успешно создан
- 400: Некорректный запрос
- 404: Ресурс не найден
- 500: Внутренняя ошибка сервера

### Error Response Format

```typescript
interface ErrorResponse {
  error: string;
}
```

**Примеры ошибок:**

400 Bad Request
```json
{
  "error": "Key: 'CreatePlantRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

404 Not Found
```json
{
  "error": "plant not found"
}
```

500 Internal Server Error
```json
{
  "error": "internal server error"
}
```

## Рекомендации для iOS-разработки

### Работа с датами
- Все даты передаются в формате ISO 8601
- Рекомендуется использовать `ISO8601DateFormatter` для парсинга дат
- Временная зона: UTC

### Networking
- Рекомендуется использовать URLSession или Alamofire
- Implement proper error handling using Result type
- Используйте Codable для сериализации/десериализации JSON

### Пример Swift-моделей

```swift
struct Plant: Codable {
    let id: Int
    let name: String
    let species: String
    let wateringFrequency: Int
    let lastWatered: Date
    let createdAt: Date
    let updatedAt: Date
    
    enum CodingKeys: String, CodingKey {
        case id
        case name
        case species
        case wateringFrequency = "watering_frequency"
        case lastWatered = "last_watered"
        case createdAt = "created_at"
        case updatedAt = "updated_at"
    }
}

struct CreatePlantRequest: Codable {
    let name: String
    let species: String
    let wateringFrequency: Int
    let lastWatered: Date
    
    enum CodingKeys: String, CodingKey {
        case name
        case species
        case wateringFrequency = "watering_frequency"
        case lastWatered = "last_watered"
    }
}
```

### Пример базового сетевого слоя

```swift
enum APIError: Error {
    case invalidURL
    case networkError(Error)
    case decodingError(Error)
    case serverError(String)
}

class PlantAPI {
    private let baseURL = "http://localhost:3000"
    
    func getPlants() async throws -> [Plant] {
        guard let url = URL(string: "\(baseURL)/api/plants") else {
            throw APIError.invalidURL
        }
        
        let (data, response) = try await URLSession.shared.data(from: url)
        
        guard let httpResponse = response as? HTTPURLResponse,
              httpResponse.statusCode == 200 else {
            throw APIError.serverError("Invalid response")
        }
        
        return try JSONDecoder().decode([Plant].self, from: data)
    }
    
    // Добавьте остальные методы API
}
```

## Запуск сервера

```bash
# Запуск всех сервисов
docker-compose up -d

# Остановка сервисов
docker-compose down
```

## Переменные окружения

Создайте файл `.env` в корневой директории:

```env
POSTGRES_DB=plant_tracker
POSTGRES_USER=plant_user
POSTGRES_PASSWORD=your_secure_password
DATABASE_URL=postgres://plant_user:your_secure_password@db:5432/plant_tracker?sslmode=disable
PORT=3000
FLASK_ENV=development
```

## Мониторинг и логи

```bash
# API logs
docker logs plant_tracker_api

# Database logs
docker logs plant_tracker_db
```

## Версионирование

Текущая версия API: 0.1.0

## Поддержка

При возникновении проблем создавайте issue в репозитории проекта.
