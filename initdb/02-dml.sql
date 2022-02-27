USE banking;

LOCK TABLES `customers` WRITE;
/*!40000 ALTER TABLE `customers`
    DISABLE KEYS */;
INSERT INTO `customers`
VALUES (2000, 'mike', '1988-12-27', 'tokyo', '110075', 1),
       (2001, 'popcorn', '1999-03-17', 'osaka', '12550', 1),
       (2002, 'kendrick', '2000-04-19', 'paris', '07631', 1),
       (2003, 'west', '2002-01-01', 'sapporo', '03102', 0),
       (2004, 'kanye', '1986-07-12', 'kyoto', '48348', 1),
       (2005, 'lamar', '1968-10-09', 'melbourne', '20782', 0);

UNLOCK TABLES;


LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts`
    DISABLE KEYS */;
INSERT INTO `accounts`
VALUES (95470, 2000, '2020-08-22 10:20:06', 'saving', 6823.23, 1),
       (95471, 2002, '2020-08-09 10:27:22', 'checking', 3342.96, 1),
       (95472, 2001, '2020-08-09 10:35:22', 'saving', 7000, 1),
       (95473, 2001, '2020-08-09 10:38:22', 'saving', 5861.86, 1);
/*!40000 ALTER TABLE `accounts`
    ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users`
    DISABLE KEYS */;
INSERT INTO `users`
VALUES ('admin', 'abc123', 'admin', NULL, '2020-08-09 10:27:22'),
       ('2001', 'abc123', 'user', 2001, '2020-08-09 10:27:22'),
       ('2000', 'abc123', 'user', 2000, '2020-08-09 10:27:22');
/*!40000 ALTER TABLE `users`
    ENABLE KEYS */;

UNLOCK TABLES;
