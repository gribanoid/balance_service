docker-compose up --build


Методы пользователя:

- создание пользователя - [GET] http://localhost:80/user/create?user_id={user_id}
- получение баланса - [GET] http://localhost:80/user/balance?user_id={user_id}
- пополнение баланса - [GET] http://localhost:80/user/deposit?user_id={user_id}&amount={amount}
- списание баланса - [GET] http://localhost:80/user/withdrawal?user_id={user_id}&amount={amount}
- перевод средств от одного пользователя другому - [GET] http://localhost:80/user/send?from={from_user_id}&to={to_user_id}&amount={amount}