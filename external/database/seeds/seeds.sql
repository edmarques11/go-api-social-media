INSERT INTO tb_user (name, nick, email, password)
  VALUES
    ("User 1", "user_1", "user1@email.com", "$2a$10$85rkexO4PIIuwSzr0Vx/Vu.ADq8Dh/iYf0WvfatTivHz41zn1DQVO"),
    ("User 2", "user_2", "user2@email.com", "$2a$10$85rkexO4PIIuwSzr0Vx/Vu.ADq8Dh/iYf0WvfatTivHz41zn1DQVO"),
    ("User 3", "user_3", "user3@email.com", "$2a$10$85rkexO4PIIuwSzr0Vx/Vu.ADq8Dh/iYf0WvfatTivHz41zn1DQVO");

INSERT INTO tb_follower (user_id, follower_id)
  VALUES
    (4, 5),
    (6, 4),
    (4, 6);