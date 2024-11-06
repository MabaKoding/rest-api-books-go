-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               PostgreSQL 12.10, compiled by Visual C++ build 1914, 64-bit
-- Server OS:                    
-- HeidiSQL Version:             12.0.0.6468
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table public.books
CREATE TABLE IF NOT EXISTS "books" (
	"books_isbn" VARCHAR(60) NOT NULL,
	"books_title" VARCHAR(255) NOT NULL,
	"books_subtitle" VARCHAR(255) NOT NULL,
	"books_author" VARCHAR(100) NOT NULL,
	"books_description" TEXT NOT NULL,
	"books_published" DATE NOT NULL,
	"books_publisher" VARCHAR(100) NOT NULL,
	PRIMARY KEY ("books_isbn")
);

-- Dumping data for table public.books: 0 rows
/*!40000 ALTER TABLE "books" DISABLE KEYS */;
INSERT INTO "books" ("books_isbn", "books_title", "books_subtitle", "books_author", "books_description", "books_published", "books_publisher") VALUES
	('9781449365036', 'Speaking JavaScript', 'An In-Depth Guide for Programmers', 'Axel Rauschmayer', 'Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position.', '2014-04-08', 'OReilly Media'),
	('9781449365046', 'Speaking JavaScript', 'An In-Depth Guide for Programmers', 'Axel Rauschmayer', 'Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position.', '2014-04-08', 'OReilly Media'),
	('9781449365048', 'Speaking JavaScript', 'An In-Depth Guide for Programmers', 'Axel Rauschmayer', 'Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position.', '2014-04-08', 'OReilly Media'),
	('9781449365040', 'Speaking JavaScript', 'An In-Depth Guide for Programmers', 'Axel Rauschmayer', 'Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position.', '2014-04-08', 'OReilly Media'),
	('9781449365020', 'Speaking JavaScript', 'An In-Depth Guide for Programmers', 'Axel Rauschmayer', 'Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position.', '2014-04-08', 'OReilly Media'),
	('9781449365035', 'Speaking JavaScript', 'An In-Depth Guide for Programmers', 'Axel Rauschmayersss', 'Like it or not, JavaScript is everywhere these days -from browser to server to mobile- and now you, too, need to learn the language or dive deeper than you have. This concise book guides you into and through JavaScript, written by a veteran programmer who once found himself in the same position.', '2014-04-08', 'OReilly Media');
/*!40000 ALTER TABLE "books" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
