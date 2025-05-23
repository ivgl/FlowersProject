#!/bin/bash

# Проверяем, установлен ли Newman
if ! command -v newman &> /dev/null; then
    echo "Newman не установлен. Устанавливаем..."
    npm install -g newman
fi

# Запускаем тесты
echo "Запуск тестов API..."
newman run plant-tracker-collection.json \
    -e plant-tracker-environment.json \
    --reporters cli,json \
    --reporter-json-export newman-report.json

# Проверяем результат
if [ $? -eq 0 ]; then
    echo "Все тесты успешно пройдены!"
else
    echo "Тесты завершились с ошибками"
    exit 1
fi 

