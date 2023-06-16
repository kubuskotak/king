-- Create "hellos" table
CREATE TABLE `hellos` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `title` text NOT NULL DEFAULT '', `body` text NOT NULL DEFAULT '', `description` text NOT NULL DEFAULT '', `slug` text NOT NULL, `user_id` integer NULL);
-- Create index "hellos_slug_key" to table: "hellos"
CREATE UNIQUE INDEX `hellos_slug_key` ON `hellos` (`slug`);
-- Create "ymirs" table
CREATE TABLE `ymirs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `version` text NOT NULL DEFAULT 'alpha-test-dev1');
