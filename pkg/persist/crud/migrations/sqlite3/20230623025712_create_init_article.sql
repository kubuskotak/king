-- Create "articles" table
CREATE TABLE `articles` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `title` text NOT NULL DEFAULT '', `body` text NOT NULL DEFAULT '', `description` text NOT NULL DEFAULT '', `slug` text NOT NULL, `user_id` integer NULL);
-- Create index "articles_slug_key" to table: "articles"
CREATE UNIQUE INDEX `articles_slug_key` ON `articles` (`slug`);
-- Create "ymirs" table
CREATE TABLE `ymirs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `version` text NOT NULL DEFAULT 'alpha-test-dev1');
