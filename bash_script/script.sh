#!/bin/bash

echo "Script name >>> $0"

# Все аргументы записали в переменную как строку с пробелом
argument=$*
# Переменная общего подсчёта строк
countline=0
# Функция проверки файла: существует или нет и является текстом или нет
checkfile() {
[ -f $oneArg ]  && file $oneArg | grep -q "text"
}
# Функция подсчёта строк
checkcount() {
    line=$(wc -l < $oneArg)
    countline=$((countline + line + 1))
    echo "Файл $oneArg существует и он ТЕКСТ"
    echo "Количество перехода строк в $oneArg равно >>> $line"
}
# Используем цикл for для перебора переданных аргументов
for oneArg in $argument; do

# Условие на существование файла и является ли текстом с использованием двух функций
    if checkfile; then
    checkcount $oneArg
    else
    echo "      WARNING >>> Файл $oneArg НЕ существует или НЕ текст"
    fi
done
echo "__________________________________________"
echo "Общее количество строк в файлах: $countline"

sleep 10

