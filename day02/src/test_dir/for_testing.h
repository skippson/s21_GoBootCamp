/**
 * @file serv.h
 * @brief Файл содержит класс Sgame, описывающий модель игры
 */
#ifndef SERV_H
#define SERV_H

#include <stdbool.h>

#include <array>
#include <iostream>

#include "apple.h"
#include "body.h"
#include "map.h"
#include "snakes.h"

/**
 * @namespace Пространство имен s21
 */
namespace s21 {
/**
 * @class Sgame
 */
class Sgame {
 public:
  /**
   * @brief Конструктор
   */
  Sgame();

  /**
   * @brief Деструктор
   */
  ~Sgame();

  /**
   * @brief Конструктор копирования
   */
  Sgame(const Sgame &other);

  /**
   * @brief Оператор присваивания
   */
  Sgame &operator=(const Sgame &rhs);

  /**
   * @return Игровое поле
   */
  int *arr();

  /**
   * @return Уровень игрока
   */
  int lvl();

  /**
   * @return Возвращает количество очков
   */
  int score();

  /**
   * @return Рекорд
   */
  int highScore();

  /**
   * @return true - ошибка игра | false - игра продолжается
   */
  bool getFail();

  /**
   * @brief Метод для начала новой игры
   * @param status true - новая игра | false - завершение игры
   */
  void again(bool status);

  /**
   * @return true - игра окончена | false - игра не окончена
   */
  bool getEnd();

  /**
   * @return true - победа игрока | false - игра не окончена
   */
  bool getWin();

  /**
   * @brief Метод для соверщения одного такта игры (тика)
   */
  void tick();

  /**
   * @brief Применение действия игрока
   * @param move Следующие движение змеи
   */
  void action(Direction move);

  /**
   * @return Значение скорости
   */
  int getSpeed();

  /**
   * @brief Убирает яблоко с игрового поля
   */
  void rmApple();

 private:
  /**
   * @brief Помещает змею на игровое поле
   */
  void putSnake();

  /**
   * @brief Проверяет победу игрока
   */
  void win();

  /**
   * @brief Помещает яблоко на игровое поле
   */
  void putApple();

  /**
   * @brief Убирает змею с игрового поля
   */
  void rmSnake();

  /**
   * @brief Проверяет съедено ли яблоко
   * @return true - яблоко съедено | false - яблоко на съедено
   */
  bool eatApple();

  /**
   * @brief Обновление статистики
   */
  void stat();

  /**
   * @brief Проверка на ошибку игрока
   */
  void fail();

  /**
   * @brief Сохраняет значение очков игрока(рекорд)
   */
  void saveScore();

  /**
   * @return Рекордное значение из файла
   */
  int getHighScore();

  /**
   * @brief Array контейнер со значениями скорости
   */
  std::array<int, 10> speed_;

  /**
   * @brief Флаг победы игрока
   */
  bool win_;

  /**
   * @brief Флаг ошибки игрока
   */
  bool fail_;

  /**
   * @brief Флаг окончания игры
   */
  bool end_;

  /**
   * @brief Флаг скушанного яблока
   */
  bool eatenApple_;

  /**
   * @brief Очки игрока
   */
  int points_;

  /**
   * @brief Рекордное значение
   */
  int highScore_;

  /**
   * @brief Уровень игрока
   */
  int level_;

  /**
   * @brief Класс змейки
   */
  Snake *snake_;

  /**
   * @brief Класс игрового поля
   */
  Map *map_;

  /**
   * @brief Класс яблока
   */
  Apple *apple_;
};
}  // namespace s21

#endif