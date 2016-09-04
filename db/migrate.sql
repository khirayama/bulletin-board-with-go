create table users(
  id integer primary key autoincrement,
  provider text,
  uid text,
  nickname text,
  image_url text
);
create table messages(
  id integer primary key autoincrement,
  user_id text,
  message text
);
